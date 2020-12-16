package api

import (
	"net/url"
)

func constructUrl(scheme, host, path string, query map[string]string) *url.URL {
	urlV := &url.URL{
		Scheme: scheme,
		Host:   host,
		Path:   path}

	values := make(url.Values)

	for k, v := range query {
		values.Set(k, v)
	}

	urlV.RawQuery = values.Encode()

	return urlV
}
