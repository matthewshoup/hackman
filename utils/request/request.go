package request

import "net/http"
import "io"
import "io/ioutil"

func Get(url, accessToken string) []byte {
	req, _ := http.NewRequest("GET", url, nil)

	authHeader := "token " + accessToken
	req.Header.Set("Authorization", authHeader)
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Accept", "application/json")

	client := &http.Client{}
	res, _ := client.Do(req)

	defer res.Body.Close()
	body, _ := ioutil.ReadAll(res.Body)

	return body
}

func Post(url string, payloadReader io.Reader) []byte {
	req, _ := http.NewRequest("POST", url, payloadReader)

	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Accept", "application/json")

	client := &http.Client{}
	res, _ := client.Do(req)

	defer res.Body.Close()
	body, _ := ioutil.ReadAll(res.Body)

	return body
}