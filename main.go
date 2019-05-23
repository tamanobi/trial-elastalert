package main

import (
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"os"
	"os/signal"
	"strings"
	"syscall"
	"time"
)

func ticker() {
	t := time.NewTicker(1 * time.Second)
	defer t.Stop()

	sig := make(chan os.Signal, 1)
	signal.Notify(sig,
		syscall.SIGHUP,
		syscall.SIGINT,
		syscall.SIGTERM,
		syscall.SIGQUIT,
	)
	defer signal.Stop(sig)

	for {
		select {
		case <-t.C:
			process()
		case s := <-sig:
			switch s {
			case syscall.SIGHUP, syscall.SIGINT, syscall.SIGTERM, syscall.SIGQUIT:
				return
			}
		}
	}
}

func process() {
	values := url.Values{}
	values.Set("json", "{\"action\":\"login\",\"user\":2}")

	req, err := http.NewRequest("POST", "http://fluentd:8888/test.cycle", strings.NewReader(values.Encode()))
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")
	if err != nil {
		panic(errors.New("failed request"))
	}

	client := new(http.Client)
	res, err := client.Do(req)
	if err != nil {
		panic(err)
	}
	defer res.Body.Close()

	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		panic(errors.New("failed reading response"))
	}
	fmt.Println(string(body))
}

func main() {
	ticker()
}
