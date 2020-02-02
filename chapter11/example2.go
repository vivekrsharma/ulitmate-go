package main

import (
	"context"
	"fmt"
	"net/http"
	"time"
)

func makeHttpRequest() {

	ctx,cancel := context.WithTimeout(context.Background(),150 * time.Millisecond)
	defer cancel()

	ch := make(chan string, 1)
	tr := &http.Transport{
		MaxIdleConns:       10,
		IdleConnTimeout:    30 * time.Second,
		DisableCompression: true,
	}
	client := &http.Client{Transport: tr}
	req,err := http.NewRequest("GET","https://www.google.com",nil)
	if err!=nil {
		panic(err)
		return
	}
	go func() {
		resp, err := client.Do(req)
		if err != nil {
			ch <- "Errored on the request"
			panic(err)
		}
		defer resp.Body.Close()
		ch <- resp.Proto

	}()

	select {
		case p:= <-ch:
			fmt.Printf("Client http call succeeded. Response %s",p)
		case td := <-ctx.Done():
			fmt.Printf("Client reqeust timeout. Details %s",td)
			tr.CancelRequest(req)
			fmt.Printf("%s",<-ch)
	}

}

func main() {
	makeHttpRequest()
}
