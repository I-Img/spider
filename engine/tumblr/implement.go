package tumblr

import (
	"os"
	"time"

	"github.com/BurntSushi/toml"
	"github.com/sirupsen/logrus"
	"github.com/tumblr/tumblrclient.go"
	"spider/common"
)

//调用tumblr API获取指定的博客图片

type Tumblr struct {
	cli  *tumblrclient.Client
	conf TumblrConfig
}

type TumblrPic struct {
	Content string
	Url     string
	Width   int
	Height  int
}

type TumblrConfig struct {
	//Key tumblr的oauth key
	Key string `toml:"key"`
	//Secert 与Oauth Key相对应的Secert Key
	Secert string `toml:"secert"`
	//PageCount 每次API要求返回的博客页面数量
	PageCount int `toml:"page_count"`
	//Blogs 博客详情
	Blogs []TumblrBlog `toml:"blogs"`
}

type TumblrBlog struct {
	URL   string `toml:"url"`
	Alias string `toml:"alias"`
}

func NewTumblr() (*Tumblr, error) {

	tc := TumblrConfig{}
	_, err := toml.DecodeFile(os.Getenv("TUMBLR_CONFIG"), &tc)
	if err != nil {
		return nil, err
	}

	return &Tumblr{
		cli:  tumblrclient.NewClient(tc.Key, tc.Secert),
		conf: tc,
	}, nil
}

func (t *Tumblr) EngineInit() (err error) {
	_t, err := NewTumblr()
	//if err != nil {
	//	return err
	//}

	t.cli = _t.cli
	t.conf = _t.conf
	return err
}

func (t *Tumblr) CrawlLatestContents() (pics []common.Content, err error) {

	logrus.Debugf("In Tumblr Conf: [%v]", t.conf)
	for _, b := range t.conf.Blogs {
		logrus.Debugf("Fetch %s In Tumblr", b.URL)
		pictures, err := t.GetPicFormBlog(b.URL, t.conf.PageCount)
		if err != nil {
			return nil, err
		}

		//t := time.Now().Unix()
		for _, p := range pictures {
			pics = append(pics, common.Content{
				Name:      p.Content,
				Auth:      b.Alias,
				CType:     common.PIC,
				CrawlTime: time.Now(),
				Content:   p.Url,
				Width:     p.Width,
				Height:    p.Height,
			})
		}
	}

	return
}

func (t *Tumblr) SaveContentHook(contents []common.Content) error {
	return nil
}
