package requestr

import (
	"crypto/tls"
	"log"
	"net/http"
	"net/http/cookiejar"
	"net/url"
	"os"
)

var Client *http.Client
var tr *http.Transport

func init() {
	// retrieve environment variable to check for presence of proxy
	proxyURL := os.Getenv("HTTP_PROXY")
	if proxyURL != "" {
		// disable TLS verification and set proxy URL
		proxyUrl, _ := url.Parse(proxyURL)
		tr = &http.Transport{
			TLSClientConfig:    &tls.Config{InsecureSkipVerify: true},
			DisableCompression: true, // to ensure that we can obtain Content-Length response header
			Proxy:              http.ProxyURL(proxyUrl),
		}
	} else {
		tr = &http.Transport{
			TLSClientConfig:    &tls.Config{InsecureSkipVerify: true},
			DisableCompression: true, // to ensure that we can obtain Content-Length response header
		}
	}

	// cookie jar to help us manage cookies
	jar, err := cookiejar.New(nil)
	if err != nil {
		log.Fatalln("Error while creating cookie jar: ", err)
	}

	// create our HTTP client using the above transport and cookie jar, then set the global variable
	Client = &http.Client{
		Transport: tr,
		Jar:       jar,
	}
}
