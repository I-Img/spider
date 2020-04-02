package tumblr

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/PuerkitoBio/goquery"
	"github.com/tumblr/tumblr.go"
	"github.com/tumblr/tumblrclient.go"
)

//调用tumblr API获取指定的博客图片

type Tumblr struct {
	cli *tumblrclient.Client
}

type TumblrPic struct {
	url    string
	width  int
	height int
}

func NewTumber(key, secert string) *Tumblr {
	return &Tumblr{
		cli: tumblrclient.NewClient(key, secert),
	}
}

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

		//获取图片地址和原始尺寸
		img := doc.Find("img")
		for _, g := range img.Nodes {
			fmt.Println(g.Attr)
			for _, a := range g.Attr {
				if a.Key == "src" {
					tp.url = a.Val
				}
				if a.Key == "data-orig-height" {
					tp.height, _ = strconv.Atoi(a.Val)
				}
				if a.Key == "data-orig-width" {
					tp.width, _ = strconv.Atoi(a.Val)
				}
			}
		}

		pics = append(pics, tp)
	}

	return
}
