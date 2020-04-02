package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

var flicker = "https://api.flickr.com/services/rest/"

func getRecentPic() error {
	req, err := http.NewRequest("GET", flicker, nil)
	if err != nil {
		return err
	}

	q := req.URL.Query()
	q.Add("api_key", key)
	q.Add("method", "flickr.photos.getRecent")
	req.URL.RawQuery = q.Encode()

	var resp *http.Response
	resp, err = http.DefaultClient.Do(req)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	data, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return err
	}

	ids, err := parse(data)

	for _, i := range ids {
		url, err := fetchPic(i)
		if err != nil {
			return err
		}

		for _, u := range url {
			fmt.Println(u)
		}
	}
	return err
}

func fetchPic(id string) (urls []string, err error) {
	req, err := http.NewRequest("GET", flicker, nil)
	if err != nil {
		return
	}

	q := req.URL.Query()
	q.Add("photo_id", id)
	q.Add("api_key", key)
	q.Add("method", "flickr.photos.getSizes")
	req.URL.RawQuery = q.Encode()

	var resp *http.Response
	resp, err = http.DefaultClient.Do(req)
	if err != nil {
		return
	}
	defer resp.Body.Close()

	data, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return
	}

	return parseSize(data)

}
