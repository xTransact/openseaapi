package openseaapi

import "time"

type options struct {
	apiKey     string
	privateKey string
	verbose    bool
	timeout    time.Duration
}

type OptionFn func(*options)

func WithApiKey(key string) OptionFn {
	return func(o *options) {
		o.apiKey = key
	}
}

func EnableVerbose() OptionFn {
	return func(o *options) {
		o.verbose = true
	}
}

func WithTimeout(timeout time.Duration) OptionFn {
	return func(o *options) {
		o.timeout = timeout
	}
}

type requestOptions struct {
	testnets bool
}

type RequestOptionFn func(*requestOptions)

func UseTestnets() RequestOptionFn {
	return func(o *requestOptions) {
		o.testnets = true
	}
}
