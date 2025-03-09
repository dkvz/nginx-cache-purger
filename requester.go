package nginxcachepurger

import (
	"log"
	"net/http"
	"strings"
	"time"
)

// TODO Provide a way for owner to close the channel
type Requester struct {
	Config  *Config
	reqChan chan string
}

// Creates the Requester but does not start it.
func NewRequester(config *Config) *Requester {
	return &Requester{
		Config:  config,
		reqChan: make(chan string, 5),
	}
}

func makePurgeRequest(fullPurgeUrl string) {
	// I could set a low timeout but we don't really care if
	// it's long.
	log.Printf("clearing: %v\n", fullPurgeUrl)
	resp, err := http.Get(fullPurgeUrl)
	if err != nil || resp.StatusCode > 210 {
		// Crash on error by design:
		log.Fatalln("error making the purge request, exiting")
	}
}

// Starts the requester:
func (r *Requester) Start() {
	go func() {
		for {
			url := <-r.reqChan
			url = strings.TrimSpace(url)
			if len(url) < 1 {
				continue
			}
			// Check if first character is a slash, or make it a slash:
			if url[0] != '/' {
				url = "/" + url
			}
			if url[len(url)-1] != '/' {
				url = url + "/"
			}
			fullPurgeUrl := r.Config.PurgeBaseUrl + url
			makePurgeRequest(fullPurgeUrl)
			time.Sleep(time.Duration(r.Config.RequestSleepInterval) * time.Second)
		}
	}()
}

func (r *Requester) PurgeUrl(url string) {
	r.reqChan <- url
}
