package main

import (
	"bytes"
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"os/signal"

	"xmh.go-eg/model"
	"xmh.go-eg/setting"
)

var server http.HandlerFunc = func(w http.ResponseWriter, r *http.Request) {
	token := r.Header.Get("Authorization")
	log.Println("auth-header:", token)
	var body = model.Identity{
		Token: token,
	}
	bs, err := json.Marshal(body)
	if err != nil {
		log.Println(err.Error())
		w.WriteHeader(http.StatusUnauthorized)
		return
	}

	resp, err := http.Post(setting.Config.Auth.Url, "application/json", bytes.NewReader(bs))
	if err != nil || resp.StatusCode > 200 {
		w.WriteHeader(http.StatusUnauthorized)
		return
	}

	bs, err = ioutil.ReadAll(r.Body)
	log.Println("received:", string(bs))
	if err != nil {
		w.Write([]byte("error: " + err.Error()))
		return
	}
	w.Write([]byte("read: " + string(bs)))
}

func main() {
	log.Println("starting repeater")
	defer log.Println("bye")
	defer log.Println("tear down")

	// start server
	chErr := make(chan error)
	go func() {
		log.Println("listening:", setting.Config.App.Port)
		chErr <- http.ListenAndServe("0.0.0.0:"+setting.Config.App.Port, server)
	}()

	// process kill and interrupt
	chInterrupt := make(chan os.Signal)
	signal.Notify(chInterrupt, os.Interrupt)
	signal.Notify(chInterrupt, os.Kill)

	select {
	case <-chInterrupt:
		return
	case err := <-chErr:
		log.Println(err.Error())
		return
	}
}
