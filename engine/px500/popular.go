package px500

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
	"os"
	"strconv"
	"time"

	"github.com/BurntSushi/toml"
	"github.com/sirupsen/logrus"
	"spider/common"
	"spider/storge/postgres"
)

type PX500 struct {
	url           string
	pc            PXConf
	lastTimestamp time.Time
	db            *postgres.Driver
}

type PXConf struct {
	API  string   `toml:"api"`
	User []string `toml:"user"`
}

func newPx500Cli() (*PX500, error) {
	pc := PXConf{}

	_, err := toml.DecodeFile(os.Getenv("PX500_CONFIG"), &pc)
	if err != nil {
		return nil, err
	}

	return &PX500{url: pc.API, pc: pc, db: postgres.GetDriver()}, nil
}

func (p *PX500) EngineInit() error {
	_p, err := newPx500Cli()
	if err != nil {
		return err
	}

	p.pc = _p.pc
	p.url = _p.url
	p.db = _p.db
	p.getLastTimestamp()

	return nil
}

func (p *PX500) CrawlLatestContents() ([]common.Content, error) {

	var cts []common.Content

	populars, err := p.fetchPopular()
	if err != nil {
		return nil, err
	}

	users, err := p.fetchUserPics()
	if err != nil {
		return nil, err
	}

	cts = append(cts, populars...)
	cts = append(cts, users...)
	//for _, c := range populars {
	//	if c.CrawlTime.After(p.lastTimestamp) {
	//		cts = append(cts, c)
	//	} else {
	//		break
	//	}
	//}

	return cts, nil
}

func (p *PX500) SaveContentHook([]common.Content) error {
	return nil
}

func (p *PX500) fetchPopular() (contents []common.Content, err error) {
	req, err := http.NewRequest(http.MethodGet, p.url, nil)
	if err != nil {
		return nil, err
	}

	q := req.URL.Query()
	q.Add("rpp", "100")
	q.Add("feature", "popular")
	q.Add("image_size%5B%5D", "640")
	q.Add("include_states", "false")
	q.Add("formats", "jpeg%2Clytro")
	q.Add("sort", "created_at")

	page := 1
	for {
		q.Add("page", strconv.Itoa(page))
		req.URL.RawQuery = q.Encode()

		logrus.Debugf("query: %s", req.URL.String())

		resp, err := http.DefaultClient.Do(req)
		if err != nil {
			return nil, err
		}

		defer resp.Body.Close()

		var r response
		data, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			return nil, err
		}

		err = json.Unmarshal(data, &r)
		if err != nil {
			return nil, err
		}

		logrus.Printf("photo [%d]", len(r.Photos))
		t := time.Now()
		for _, pic := range r.Photos {
			logrus.Debugf("create time: %d. now: %d", pic.CreatedAt.Unix(), t.Unix())
			if pic.CreatedAt.After(p.lastTimestamp) {
				contents = append(contents, common.Content{
					//ID:        pic.ID,
					Name:      pic.Name,
					Auth:      pic.User.Username,
					CrawlTime: t,
					Content:   pic.ImageURL[0],
					CType:     common.PX500,
					Width:     pic.Width,
					Height:    pic.Height,
					Remarks:   pic.Description,
				})
			}
		}

		logrus.Debugf("Page %d fetch %d pics", page, len(contents))
		page++
		if page == 10 {
			break
		}
	}

	return contents, p.updateLastTimestamp()
}

func (p *PX500) fetchUserPics() (contents []common.Content, err error) {
	return nil, nil
}

func (p *PX500) getLastTimestamp() {
	t, err := p.db.FetchSource("px500")
	if err != nil {
		logrus.Errorf("Fetch Source Error: %s. Use time.Now() as timestamp ", err.Error())
		p.lastTimestamp = time.Now()
	}

	p.lastTimestamp = t.LastTime
}

func (p *PX500) updateLastTimestamp() error {
	p.lastTimestamp = time.Now()

	return p.db.UpdateSource(common.Source{
		Name:     "px500",
		LastTime: p.lastTimestamp,
	})
}
