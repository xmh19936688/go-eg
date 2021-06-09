package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"os/signal"
)

func main() {
	fmt.Println("starting repeater")
	defer fmt.Println("bye")
	defer fmt.Println("tear down")

	// create server
	var server http.HandlerFunc = func(w http.ResponseWriter, r *http.Request) {
		bs, err := ioutil.ReadAll(r.Body)
		if err != nil {
			w.Write([]byte("error: " + err.Error()))
			return
		}
		w.Write([]byte("read: " + string(bs)))
	}

	// start server
	chErr := make(chan error)
	go func() {
		chErr <- http.ListenAndServe("0.0.0.0:9000", server)
	}()

	// process kill and interrupt
	chInterrupt := make(chan os.Signal)
	signal.Notify(chInterrupt, os.Interrupt)
	signal.Notify(chInterrupt, os.Kill)

	select {
	case <-chInterrupt:
		return
	case err := <-chErr:
		fmt.Println(err.Error())
		return
	}
}
