package main

import (
	"net/http"

	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	"github.com/sirupsen/logrus"
)

func main() {
	startAPIServer()
}

var externalUrls = []string{
	"https://httpstat.us/200",
	"https://httpstat.us/503",
}

func startAPIServer() {
	for _, url := range externalUrls {
		req := newRequestCollector(url)
		prometheus.MustRegister(req)
	}

	logrus.Printf("setting up handlers (docs, status, /)")
	http.Handle("/metrics", promhttp.Handler())
	srv := &http.Server{
		Addr: "localhost:8080",
	}
	logrus.Printf("listening on 8080")
	func() {
		logrus.Printf("inside go goroutine")
		if err := srv.ListenAndServe(); err != http.ErrServerClosed {
			logrus.Infof("listening stopped: %v", err)
		}
	}()

}
