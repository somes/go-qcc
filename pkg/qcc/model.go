package qcc

import "github.com/go-resty/resty/v2"

type Client struct {
	Client *resty.Client
	tid    string
}
