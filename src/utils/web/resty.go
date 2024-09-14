package web

import "github.com/go-resty/resty/v2"

// RestyClient 复用Resty
var RestyClient = resty.New()
