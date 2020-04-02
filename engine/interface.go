package engine

import "spider/common"

type EngineInterface interface {
	//EngineInit 初始化抓取引擎
	EngineInit() error
	//CrawlLatestContents 抓取最新内容
	CrawlLatestContents() ([]common.Content, error)
	//SaveContent 保存抓取到的内容
	SaveContentHook([]common.Content) error
}
