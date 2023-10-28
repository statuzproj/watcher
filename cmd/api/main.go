package main

import (
	"log"
	"sync"
)

func main() {
	endpoints := getEndpoints()
	var wg sync.WaitGroup

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
