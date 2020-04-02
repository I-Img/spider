package tumblr

import (
	"strconv"
	"strings"

	"github.com/PuerkitoBio/goquery"
	"github.com/tumblr/tumblr.go"
)

func (t *Tumblr) getBlogInfo(blogName string) (*tumblr.Blog, error) {
	return tumblr.GetBlogInfo(t.cli, blogName)
}

//GetBlogAvatar 获取博客缩略头像图
func (t *Tumblr) GetBlogAvatar(blogName string) (avatar string, err error) {
	blog, err := t.getBlogInfo(blogName)
	if err != nil {
		return
	}

	return blog.GetAvatar()
}

func (t *Tumblr) getBlogNames() ([]string, error) {
	return nil, nil
}

//GetPicFormBlog 获取博客中的图片和宽高尺寸
func (t *Tumblr) GetPicFormBlog(blogName string, count int) (pics []TumblrPic, err error) {
	blog, err := t.getBlogInfo(blogName)
	if err != nil {
		return
	}

	values := make(map[string][]string)
	values["type"] = []string{"photo"}
	values["limit"] = []string{strconv.Itoa(count)}

	posts, err := blog.GetPosts(values)
	if err != nil {
		return
	}

	allPosts, err := posts.All()
	if err != nil {
		return
	}

	for _, p := range allPosts {

		buffer := strings.NewReader(p.GetSelf().Body)
		doc, err := goquery.NewDocumentFromReader(buffer)
		if err != nil {
			continue
		}

		var tp TumblrPic

		add := false
		tp.Content = strings.TrimSpace(doc.Find("p").Text())

		if len(tp.Content) > 255 {
			title := []rune(tp.Content)

			tp.Content = string(title[:200])
		}

		//获取图片地址和原始尺寸
		img := doc.Find("img")
		for _, g := range img.Nodes {
			for _, a := range g.Attr {
				if a.Key == "src" {
					if strings.TrimSpace(a.Val) == "" {
						break
					}
					tp.Url = a.Val
				}
				if a.Key == "data-orig-height" {
					tp.Height, _ = strconv.Atoi(a.Val)
				}
				if a.Key == "data-orig-width" {
					tp.Width, _ = strconv.Atoi(a.Val)
				}
			}

			add = true
		}

		if add {
			pics = append(pics, tp)
		}

	}

	return
}
