package main

import (
	"flag"
	"fmt"
	"net"
	"net/http"

	"github.com/go-kit/kit/log"
	"github.com/go-kit/kit/metrics"
	"github.com/go-kit/kit/metrics/prometheus"
	stdprometheus "github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	"google.golang.org/grpc"
	service2 "spider/endpoints/service"
	rpc_v1 "spider/grpc/golang"

	"os"
	"text/tabwriter"

	kitgrpc "github.com/go-kit/kit/transport/grpc"
	"github.com/oklog/oklog/pkg/group"
	"github.com/sirupsen/logrus"
	"spider/engine"
	tumblr2 "spider/engine/tumblr"
	"spider/storge/postgres"
)

var tumblr *tumblr2.Tumblr
var d *postgres.Driver

func usageFor(fs *flag.FlagSet) func() {
	short := os.Args[0] + " [flags]"
	return func() {
		fmt.Fprintf(os.Stderr, "USAGE\n")
		fmt.Fprintf(os.Stderr, "  %s\n", short)
		fmt.Fprintf(os.Stderr, "\n")
		fmt.Fprintf(os.Stderr, "FLAGS\n")
		w := tabwriter.NewWriter(os.Stderr, 0, 2, 2, ' ', 0)
		fs.VisitAll(func(f *flag.Flag) {
			fmt.Fprintf(w, "\t-%s %s\t%s\n", f.Name, f.DefValue, f.Usage)
		})
		w.Flush()
		fmt.Fprintf(os.Stderr, "\n")
	}
}

func main() {
	logrus.SetLevel(logrus.DebugLevel)

	fs := flag.NewFlagSet("spider", flag.ExitOnError)

	var grpcAddr = fs.String("grpc-addr", ":80", "gRPC Listen Address")
	var httpAddr = fs.String("http-addr",":2000","Http Listen Address")

	fs.Usage = usageFor(fs)
	fs.Parse(os.Args[1:])

	var logger log.Logger
	{
		logger = log.NewLogfmtLogger(os.Stderr)
		logger = log.With(logger, "ts", log.DefaultTimestampUTC)
		logger = log.With(logger, "caller", log.DefaultCaller)
	}

	var ints metrics.Counter
	{
		// Business-level metrics.
		ints = prometheus.NewCounterFrom(stdprometheus.CounterOpts{
			Namespace: "i_img",
			Subsystem: "Fetch",
			Name:      "Pictures",
			Help:      "Total count of fetch API invokes.",
		}, []string{})

	}
	var duration metrics.Histogram
	{
		// Endpoint-level metrics.
		duration = prometheus.NewSummaryFrom(stdprometheus.SummaryOpts{
			Namespace: "i_img",
			Subsystem: "Fetch",
			Name:      "FetchPictures",
			Help:      "FetchPictures Request duration in seconds.",
		}, []string{"method", "success"})
	}

	httpServer := http.DefaultServeMux
	httpServer.Handle("/metrics", promhttp.Handler())

	var (
		service    = service2.New(logger, ints)
		endpoints  = service2.NewSet(service, logger, duration)
		grpcServer = service2.NewGRPCServer(endpoints, logger)
	)

	//init db
	err := initDB()
	if err != nil {
		panic(err)
	}

	var g group.Group
	{
		grpcListener, err := net.Listen("tcp", *grpcAddr)
		if err != nil {
			logger.Log("transport", "gRPC", "during", "Listen", "err", err)
			os.Exit(-1)
		}

		g.Add(func() error {
			logger.Log("transport", "gRPC", "addr", *grpcAddr)
			baseServer := grpc.NewServer(grpc.UnaryInterceptor(kitgrpc.Interceptor))
			rpc_v1.RegisterPictureServiceServer(baseServer, grpcServer)
			return baseServer.Serve(grpcListener)
		}, func(err error) {
			grpcListener.Close()
		})
	}
	{
		httpListener, err := net.Listen("tcp",*httpAddr)
		if err != nil{
			logger.Log("transport","http","during","Listen","err",err)
			os.Exit(-1)
		}

		g.Add(func() error {
			logger.Log("transport","Http","addr",*httpAddr)
			return http.Serve(httpListener, httpServer)
		}, func(err error) {
			httpListener.Close()
		})
	}
	logger.Log("exit",g.Run())


	//
	//spiders := initSpider()
	//logrus.Debugf("Has [%d] spiders", len(spiders))
	//for _, s := range spiders {
	//	preys, err := s.CrawlLatestContents()
	//	if err != nil {
	//		logrus.Errorf("%s", err.Error())
	//		continue
	//	}
	//
	//	s.SaveContentHook(preys)
	//
	//	logrus.Debugf("Has [%d] Preys", len(preys))
	//	if err := d.SaveContents(preys); err != nil {
	//		logrus.Errorf("Save Preys Error: %s", err.Error())
	//	}
	//}

}

func initSpider() []engine.EngineInterface {
	var eei []engine.EngineInterface

	//tumblr
	t := new(tumblr2.Tumblr)
	err := t.EngineInit()
	if err != nil {
		logrus.Errorf("Tumblr Init Error: [%s]", err.Error())
	} else {
		eei = append(eei, t)
	}

	return eei
}

func initDB() (err error) {

	d, err = postgres.NewDriver()
	return err
}
