package http

import (
	"io/ioutil"
	"net/http"
	"time"
)

type Client struct {
	httpClient http.Client
	Response   *http.Response
}

func NewClient() (client *Client, error error) {
	client = new(Client)
	timeout := time.Duration(5 * time.Second)
	client.httpClient = http.Client{
		Timeout: timeout,
	}
	return
}

func (client *Client) Get(url string) (body string, error error) {
	client.Response, error = client.httpClient.Get(url)
	if error != nil {
		return
	}
	bodyBytes, error := ioutil.ReadAll(client.Response.Body)
	defer client.Response.Body.Close()
	if error == nil {
		body = string(bodyBytes)
	}
	return
}
