package network

import (
	"fmt"
	"net/url"
	"net/http"
	"clitut/file"
)

func Fetch() (*http.Response, error) {

	user, err := file.ReadFile()
	if (err == nil) {

		//&sort=id DESC
		apiUrl := "http://128.199.89.41:9001/data?user="+user.Id+"&limit=1&sort=id%20DESC"

		u, _ := url.ParseRequestURI(apiUrl)
		urlStr := fmt.Sprintf("%v", u)

		client := &http.Client{}
		r, _ := http.NewRequest("GET", urlStr, nil) // <-- URL-encoded payload
		r.Header.Add("x-access-token", user.Token)
		r.Header.Add("Content-Type", "application/x-www-form-urlencoded")


		// resp:=client.Do(req)
		return client.Do(r)
		defer r.Body.Close()
	} else {
		return nil,err

	}

	return nil,nil


}
