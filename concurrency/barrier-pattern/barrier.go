package barrier_pattern

import (
	"fmt"
	"io"
	"net/http"
	"time"
)

var timeoutMilliseconds time.Duration = 5000

type barrierResp struct {
	Body string
	Err  error
}

func barrier(endpoints ...string) {
	requestNumber := len(endpoints)

	in := make(chan barrierResp, requestNumber)
	defer close(in)

	responses := make([]barrierResp, requestNumber)

	for _, endpoint := range endpoints {
		go makeRequest(in, endpoint)
	}

	var hasError bool
	for i := 0; i < requestNumber; i++ {
		resp := <-in
		if resp.Err != nil {
			fmt.Println("ERROR:", resp.Err)
			hasError = true
		}
		responses[i] = resp
	}

	if !hasError {
		for _, resp := range responses {
			fmt.Println(resp.Body)
		}
	}
}

func makeRequest(out chan<- barrierResp, url string) {
	var resp barrierResp

	client := http.Client{Timeout: timeoutMilliseconds * time.Millisecond}

	httpResp, err := client.Get(url)
	if err != nil {
		resp.Err = err
		out <- resp
		return
	}

	respBody, err := io.ReadAll(httpResp.Body)
	if err != nil {
		resp.Err = err
		out <- resp
		return
	}

	resp.Body = string(respBody)
	out <- resp
}
