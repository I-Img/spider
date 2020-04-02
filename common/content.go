package common

import "time"

type ContentType int

const (
	HTML = iota
	PIC
)

//Content 抓取的内容
type Content struct {
	//Name 内容ID或者名称
	Name string `json:"name"`
	//Auth 内容作者或者内容主人公
	Auth string `json:"auth"`
	//CrawlTime 抓取时间戳
	CrawlTime time.Time `json:"crawl_time"`
	//Content 具体内容. 可能是地址，也可能是html代码, 通过后面的ContentType来区分
	Content string `json:"content"`
	//CType Content类型值
	CType ContentType `json:"c_type"`
	//Width 如果是图片，则表示图片宽
	Width float32 `json:"width"`
	//Height 如果是图片，则表示图片高
	Height float32 `json:"height"`
	//Remarks 备注信息
	Remarks string `json:"remarks"`
}
