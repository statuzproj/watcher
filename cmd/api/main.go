package main

import (
	"github.com/prometheus/client_golang/prometheus/promhttp"
	"github.com/statuzproj/watcher/utils/healthz"
	"log"
	"net/http"
	"sync"
)

func main() {
	endpoints := getEndpoints()
	var wg sync.WaitGroup

	wg.Add(1)
	go func() {
		defer wg.Done()
		http.HandleFunc("/healthz", healthz.HealthCheck)
		http.Handle("/metrics", promhttp.Handler())
		err := http.ListenAndServe(":8081", nil)
		if err != nil {
			log.Fatalf("HTTP server error: %v", err)
		}
	}()

	for _, endpoint := range endpoints {
		wg.Add(1)

		go func(endpoint Endpoint) {
			defer wg.Done()
			for {
				switch endpoint.Target.Type {
				case "webpage":
					err := pageWatcher(endpoint.Name, endpoint.Target.Endpoint, endpoint.Target.Interval)
					if err != nil {
						log.Println(err)
					}
				}
			}
		}(endpoint)
	}
	wg.Wait()
}
