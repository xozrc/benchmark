package client

type Client interface {
	Set(key string, val string) error
}
