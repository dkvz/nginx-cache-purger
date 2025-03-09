package nginxcachepurger

import "log"

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

// Starts the requester:
func (r *Requester) Start() {
	go func() {
		for {
			url := <-r.reqChan
			log.Printf("processing %v\n", url)
		}
	}()
}

func (r *Requester) PurgeUrl(url string) {
	r.reqChan <- url
}
