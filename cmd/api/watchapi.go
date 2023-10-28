package main

import "log"

func apiWatcher(name string, endpoint string, interval string) error {
	log.Printf(" watching %s, %s, %s", name, endpoint, interval)
	return nil
}
