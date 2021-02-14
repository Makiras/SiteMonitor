package utils

import (
	"io/ioutil"
	"log"
	"net/http"
)

// CheckHTTPPage provides http server accessibility check
func CheckHTTPPage(method string, params string, urlpath string) error {
	client := &http.Client{}
	req, err := http.NewRequest(method, urlpath, nil)
	if err != nil {
		log.Print(err)
		return err
	}
	resp, err := client.Do(req)
	if err != nil {
		log.Print(err)
		return err
	}
	data, _ := ioutil.ReadAll(resp.Body)
	resp.Body.Close()
	if data != nil {
		// todo: re match
	}
	return nil
}
