package util

import (
	"io"
	"io/ioutil"
	"net/http"
)

func PostWithJson(url string, body io.Reader) ([]byte, error) {
	client := &http.Client{}
	req, _ := http.NewRequest("POST", url, body)
	req.Header.Set("Content-Type", "application/json")
	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	if resp.Body == nil {
		return nil, nil
	}
	defer resp.Body.Close()
	r, err := ioutil.ReadAll(resp.Body)
	return r, err
}

func HttpGet(url string) ([]byte, error) {
	client := &http.Client{}
	req, _ := http.NewRequest("Get", url, nil)
	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	if resp.Body == nil {
		return nil, nil
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	return body, err
}
