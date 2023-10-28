package main

import "log"

func main() {
	endpoints := getEndpoints()

	for _, endpoint := range endpoints {
		switch endpoint.Target.Type {
		case "webpage":
			err := pageWatcher(endpoint.Name, endpoint.Target.Endpoint, endpoint.Target.Interval)
			if err != nil {
				log.Println(err)
			}
		case "ip":
			err := ipWatcher(endpoint.Name, endpoint.Target.Endpoint, endpoint.Target.Interval)
			if err != nil {
				log.Println(err)
			}
		case "api":
			err := apiWatcher(endpoint.Name, endpoint.Target.Endpoint, endpoint.Target.Interval)
			if err != nil {
				log.Println(err)
			}
		}
	}
}
