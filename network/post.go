package network

import (
	"bytes"
	"fmt"
	"net/http"
	"net/url"
	"strconv"
)

func Login(email string, password string) (*http.Response, error) {
	apiUrl := "http://128.199.89.41:9001"
	resource := "/user/login"
	data := url.Values{}
	data.Set("email", email)
	data.Add("password", password)

	u, _ := url.ParseRequestURI(apiUrl)
	u.Path = resource
	urlStr := fmt.Sprintf("%v", u) // "https://api.com/user/"

	client := &http.Client{}
	r, _ := http.NewRequest("POST", urlStr, bytes.NewBufferString(data.Encode())) // <-- URL-encoded payload
	r.Header.Add("Authorization", "auth_token=\"XXXXXXX\"")
	r.Header.Add("Content-Type", "application/x-www-form-urlencoded")
	r.Header.Add("Content-Length", strconv.Itoa(len(data.Encode())))
	defer r.Body.Close()

	// resp:=client.Do(req)
	return client.Do(r)

}
