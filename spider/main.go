package main

import (
	"github.com/sirupsen/logrus"
	"spider/engine"
	"spider/engine/px500"
	"spider/storge/postgres"
)

var d *postgres.Driver
var p *px500.PX500

func main() {
	logrus.SetLevel(logrus.DebugLevel)

	err := initDB()
	if err != nil {
		panic(err)
	}

	spiders := initSpider()
	logrus.Debugf("Has [%d] spiders", len(spiders))
	for _, s := range spiders {
		preys, err := s.CrawlLatestContents()
		if err != nil {
			logrus.Errorf("%s", err.Error())
			continue
		}

		s.SaveContentHook(preys)

		logrus.Debugf("Has [%d] Preys", len(preys))
		if err := d.SaveContents(preys); err != nil {
			logrus.Errorf("Save Preys Error: %s", err.Error())
		}
	}
}

func initDB() (err error) {

	d, err = postgres.NewDriver()
	return err
}

func initSpider() []engine.EngineInterface {

	var eei []engine.EngineInterface

	// px500
	p = new(px500.PX500)
	err := p.EngineInit()
	if err != nil {
		logrus.Errorf("PX500 Init Error: [%s]", err.Error())
	} else {
		eei = append(eei, p)
	}

	//tumblr
	//t := new(tumblr2.Tumblr)
	//err := t.EngineInit()
	//if err != nil {
	//	logrus.Errorf("Tumblr Init Error: [%s]", err.Error())
	//} else {
	//	eei = append(eei, t)
	//}

	return eei
}
