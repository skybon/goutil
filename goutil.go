package goutil

import (
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
)

var ErrMismatch = errors.New("Mismatch")
var ErrPanic = errors.New("Procedure panic")

func ErrorOut(err error, expectation, result interface{}) string {
	return fmt.Sprintf("Error: %s\nExpected %#v\nReceived %#v\n", err.Error(), expectation, result)
}

func ErrorOutJSON(err error, expectation, result interface{}) string {
	expJSON, _ := json.Marshal(expectation)
	resJSON, _ := json.Marshal(result)
	return fmt.Sprintf("Error: %s\nExpected %s\nReceived %s\n", err.Error(), expJSON, resJSON)
}

func SprintfCompare(expectation, result interface{}) bool {
	return fmt.Sprintf("%#v", expectation) == fmt.Sprintf("%#v", result)
}

func JSONcompare(expectation, result interface{}) bool {
	var e, _ = json.Marshal(expectation)
	var r, _ = json.Marshal(result)

	if string(e) != string(r) {
		return false
	}

	return true
}

func DownloadURL(urlData url.URL) (data string, err error) {
	return Download(urlData.String())
}

// Download retrieves data from the specified HTTP address.
func Download(url string) (data string, err error) {
	var resp *http.Response
	resp, err = http.Get(url)
	if err == nil {
		defer resp.Body.Close()
		readallContents, _ := ioutil.ReadAll(resp.Body)
		data = string(readallContents)
	}
	return data, err
}
