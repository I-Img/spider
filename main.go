package main

import (
	"github.com/sirupsen/logrus"
	"spider/engine"
	tumblr2 "spider/engine/tumblr"
	"spider/storge/postgres"
)

var tumblr *tumblr2.Tumblr
var d *postgres.Driver

func main() {
	logrus.SetLevel(logrus.DebugLevel)

	err := initDB()
	if err != nil {
		panic(err)
	}

	startGRPC(":80")
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
