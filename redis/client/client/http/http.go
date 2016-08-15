package http

import (
	"bytes"
	"fmt"
	"net/http"
)

type HttpClient struct {
	addr string
}

func (hc *HttpClient) Set(key string, val string) error {
	//url := fmt.Sprintf("%s/%s", hc.addr, key)
	url := hc.addr

	resp, err := http.Post(url, "", bytes.NewBufferString(val))
	if err != nil {
		return err
	}
	if resp.StatusCode != http.StatusOK {
		return fmt.Errorf("error status code %d", resp.StatusCode)
	}
	return nil
}

func NewHttpClient(addr string) *HttpClient {
	return &HttpClient{
		addr: addr,
	}
}
