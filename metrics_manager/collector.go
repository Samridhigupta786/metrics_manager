package main

import (
	"net/http"
	"strings"
	"time"

	"github.com/prometheus/client_golang/prometheus"
	"github.com/sirupsen/logrus"
)

type response struct {
	Url           string
	ResponseTime  float64
	ExternalUrlUp float64
}

//Define a struct for you collector that contains pointers
//to prometheus descriptors for each metric you wish to expose.
//Note you can also include fields of other types if they provide utility
//but we just won't be exposing them as metrics.
type RequestCollector struct {
	ExternalUrl  *prometheus.Desc
	ResponseTime *prometheus.Desc
}

//You must create a constructor for you collector that
//initializes every descriptor and returns a pointer to the collector
func newRequestCollector(url string) *RequestCollector {
	return &RequestCollector{
		ExternalUrl: prometheus.NewDesc("sample_external_url_up",
			"shows if url is up",
			nil, prometheus.Labels{
				"url": url},
		),
		ResponseTime: prometheus.NewDesc("sample_external_url_response_ms",
			"response time for url",
			nil, prometheus.Labels{
				"url": url},
		),
	}
}

//Each and every collector must implement the Describe function.
//It essentially writes all descriptors to the prometheus desc channel.
func (collector *RequestCollector) Describe(ch chan<- *prometheus.Desc) {

	//Update this section with the each metric you create for a given collector
	// ch <- collector.ExternalUrl503
	ch <- collector.ExternalUrl
	// ch <- collector.ResponseTime503
	ch <- collector.ResponseTime
}

//Collect implements required collect function for all promehteus collectors
func (collector *RequestCollector) Collect(ch chan<- prometheus.Metric) {

	//Implement logic here to determine proper metric value to return to prometheus
	//for each descriptor or call other functions that do so.
	resp := GetMetrics(collector.ExternalUrl.String())

	//Write latest value for each metric in the prometheus metric channel.
	//Note that you can pass CounterValue, GaugeValue, or UntypedValue types here.
	ch <- prometheus.MustNewConstMetric(collector.ExternalUrl, prometheus.CounterValue, resp.ExternalUrlUp)
	ch <- prometheus.MustNewConstMetric(collector.ResponseTime, prometheus.CounterValue, resp.ResponseTime)

}

func GetMetrics(url string) response {
	logrus.Printf("reached to metrics")
	externalUrls := map[string]string{
		"200": "https://httpstat.us/200",
		"503": "https://httpstat.us/503",
	}
	logrus.Printf("reached here.... %v", url)
	var resp response
	if strings.Contains(url, externalUrls["200"]) {
		logrus.Printf("inside 200 url if condition.... %v", url)
		resp = handleHTTPRequest(externalUrls["200"])
	} else {
		logrus.Printf("inside 503 url if condition.... %v", url)
		resp = handleHTTPRequest(externalUrls["503"])
	}

	return resp

}
func handleHTTPRequest(url string) response {

	start := time.Now()
	resp, err := http.Get(url)

	diff := time.Now().Sub(start).Milliseconds()
	if err != nil {
		logrus.Fatal(err)
	}
	logrus.Printf("time durations : %v, %v", diff, float64(diff))
	result := response{
		Url:          url,
		ResponseTime: float64(diff),
	}
	if resp.StatusCode == 200 {
		result.ExternalUrlUp = 1
	}
	return result
}
