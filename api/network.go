package api

import (
	"encoding/json"
	"fmt"

	log "github.com/sirupsen/logrus"
)

// NetworkCreate - POST /network/ - Creates a new network
func NetworkCreate(network Network) Response {
	log.Trace("NetworkCreate")
	marshalled, _ := json.Marshal(network)
	url := getEndpoint("network/")
	resp := Request("POST", url, []byte(marshalled))

	showFullResponseBody(resp)

	if resp.StatusCode != 201 {
		log.Fatalf("RequestError: %v", resp.Status)
	}
	return resp
}

// NetworkDelete - DELETE /network/{id} - Deletes a network
func NetworkDelete(id string) Response {
	log.Trace("NetworkDelete")
	url := getEndpoint(fmt.Sprintf("network/%v", id))
	resp := Request("DELETE", url, []byte{})

	showFullResponseBody(resp)

	if resp.StatusCode != 200 {
		log.Fatalf("RequestError: %v", resp.Status)
	}
	return resp
}

// NetworkGet - GET /network/{id} - Retrieves a specific network
func NetworkGet(id string) Response {
	log.Trace("NetworkGet")
	url := getEndpoint(fmt.Sprintf("network/%v", id))
	resp := Request("GET", url, []byte{})

	showFullResponseBody(resp)

	if resp.StatusCode != 200 {
		log.Fatalf("RequestError: %v", resp.Status)
	}
	return resp
}

// NetworkList - GET /network/ - Retrieves all members
func NetworkList() Response {
	log.Trace("NetworkList")
	url := getEndpoint("network")
	resp := Request("GET", url, []byte{})

	showFullResponseBody(resp)

	if resp.StatusCode != 200 {
		log.Fatalf("RequestError: %v", resp.Status)
	}
	return resp
}

// NetworkPut - PUT /network/{id} - Creates or updates a network
func NetworkPut(network Network) Response {
	log.Trace("NetworkPut")
	marshalled, _ := json.Marshal(network)
	url := getEndpoint(fmt.Sprintf("network/%v", network.ID))
	resp := Request("PUT", url, []byte(marshalled))

	showFullResponseBody(resp)

	if resp.StatusCode > 201 {
		log.Fatalf("RequestError: %v", resp.Status)
	}
	return resp
}
