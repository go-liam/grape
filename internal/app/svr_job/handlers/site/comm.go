package site

import (
	"encoding/json"
	"errors"
	"io/ioutil"
	"net/http"
)

const httpOk = 200

func GetHttp(url string) (*Model, error) {
	o := new(Model)
	req, _ := http.NewRequest("GET", url, nil)
	req.Header.Add("accept", "application/json")
	req.Header.Add("content-type", "application/json")
	res, err := (&http.Client{}).Do(req)
	if err != nil {
		return o, err
	}
	defer func() { _ = res.Body.Close() }()
	if res.StatusCode != httpOk {
		return o, errors.New("httpStatusFail")
	}
	got, err2 := ioutil.ReadAll(res.Body)
	json.Unmarshal(got, &o)
	return o, err2
}
