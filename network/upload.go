package network

import (
	"net/http"

	"fmt"
	//"bytes"
	//"strconv"
	"clitut/file"
	"net/url"
	"bytes"
	"strconv"
)

func Upload(text string) (*http.Response, error) {

	fmt.Println("Going to Upload Data Bro")
	apiUrl := "http://128.199.89.41:9001"
	resource := "/data"

	user, err := file.ReadFile()
	if (err == nil) {

		//fmt.Println(user.Token)
		data := url.Values{}
		data.Set("user", user.Id)
		data.Add("text", text)

		u, _ := url.ParseRequestURI(apiUrl)
		u.Path = resource
		urlStr := fmt.Sprintf("%v", u) // "https://api.com/user/"

		client := &http.Client{}
		r, _ := http.NewRequest("POST", urlStr, bytes.NewBufferString(data.Encode())) // <-- URL-encoded payload
		r.Header.Add("x-access-token", user.Token)
		r.Header.Add("Content-Type", "application/x-www-form-urlencoded")
		r.Header.Add("Content-Length", strconv.Itoa(len(data.Encode())))
		defer r.Body.Close()

		// resp:=client.Do(req)
		return client.Do(r)
	} else {
		return nil,err

	}

}
