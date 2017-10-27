package main

import (
	"io/ioutil"
	"net/http"
)

func main() {
	s := &http.Server{
		Addr: ":9933",
		Handler: http.HandlerFunc(func(rw http.ResponseWriter, req *http.Request) {
			f, err := ioutil.ReadFile("/etc/cf-assets/envoy_config/listeners.json")
			if err != nil {
				panic(err)
			}
			rw.Write(f)
		}),
	}
	s.ListenAndServe()
	panic("shouldn't be here")
}
