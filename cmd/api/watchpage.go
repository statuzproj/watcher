package main

import (
	"github.com/statuzproj/watcher/utils/prometheus"
	"log"
	"net/http"
	"strconv"
	"time"
)

func pageWatcher(name string, url string, intervalInSeconds string) error {
	log.Printf("starting watcher for %s\n", name)

	interval, err := strconv.ParseInt(intervalInSeconds, 10, 64)
	if err != nil {
		log.Println(err)
		return err
	}

	ticker := time.NewTicker(time.Duration(interval) * time.Second)
	defer ticker.Stop()

	for range ticker.C {
		pageChecker(name, url)
	}

	return nil
}

func pageChecker(name string, url string) {
	resp, err := http.Get(url)
	if err != nil {
		log.Printf("Error checking URL %s: %v\n", url, err)
		return
	}
	defer resp.Body.Close()

	statusText := http.StatusText(resp.StatusCode)

	if resp.StatusCode == http.StatusOK {
		prometheus.SetWebpagePrometheusMetric(name, url, http.StatusOK)
		log.Printf("URL %s is returning %s\n", url, statusText)
	} else {
		prometheus.SetWebpagePrometheusMetric(name, url, int64(resp.StatusCode))
		log.Printf("URL %s is returning %s\n", url, statusText)
	}
}
