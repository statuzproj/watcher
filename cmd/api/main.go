package main

import "log"

func main() {
	endpoints := getEndpoints()

	for _, endpoint := range endpoints {
		switch endpoint.Target.Type {
		case "webpage":
			log.Println("webpage")
		case "ip":
			log.Println("ip")
		case "api":
			log.Println("api")
		}
	}
}
