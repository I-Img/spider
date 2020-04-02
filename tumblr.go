package main

import (
	"fmt"

	"github.com/sirupsen/logrus"
	tumblr2 "spider/engine/tumblr"
	"spider/layout/flutter"
)

func startTumblr() {
	blogs := []string{"mmqilysm.tumblr.com"}

	for _, b := range blogs {
		pics, err := tumblr.GetPicFormBlog(b, 20)
		if err != nil {
			logrus.Errorf("Get Tumblr Pic [%s] Error: %s", b, err.Error())
			continue
		}

		for _, p := range pics {
			fmt.Println(p.Url)
		}

		outputFlutterLayout(pics)
	}

}

func outputFlutterLayout(pics []tumblr2.TumblrPic) {
	var sz []flutter.Size

	for _, p := range pics {
		sz = append(sz, flutter.Size{Width: p.Width, Height: p.Height})
	}

	size := flutter.Calc(sz)

	for _, s := range size {
		logrus.Println(s)
	}
}
