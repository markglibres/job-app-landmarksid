package urlHelper

import (
	"net/url"
)

func Authority(urlToParse string) string {
	urlObject, err := url.Parse(urlToParse)
	if err != nil {
		panic(err.Error())
	}

	var port string
	if port = urlObject.Port(); port != "" {
		port = ":" + port
	}
	return urlObject.Scheme + "://" + urlObject.Host + port
}
