package main

import (
	"encoding/json"
	"io"
	"log"
	"net/http"
)

type Endpoint struct {
	Name   string `json:"Name"`
	Target struct {
		Type     string `json:"Type"`
		Endpoint string `json:"Endpoint"`
		Method   string `json:"Method"`
	} `json:"Target"`
}

func getEndpoints() (result []Endpoint) {
	getEndpoints, err := http.Get("http://genie:8080/endpoints")
	if err != nil {
		log.Println(err)
	}
	defer getEndpoints.Body.Close()

	body, err := io.ReadAll(getEndpoints.Body)
	if err != nil {
		log.Println(err)
	}

	var endpoints []Endpoint
	err = json.Unmarshal(body, &endpoints)
	if err != nil {
		log.Println(err)
	}
	return endpoints
}
