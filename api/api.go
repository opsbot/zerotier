package api

import (
	log "github.com/sirupsen/logrus"
)

// RandomToken - GET  /randomToken - This generates a random token suitable for use as an API token server-side using a secure random source.
func RandomToken() Response {
	log.Trace("RandomToken")
	url := getEndpoint("randomToken")
	resp := Request("GET", url, []byte{})

	showFullResponseBody(resp)

	if resp.StatusCode != 200 {
		log.Fatalf("RequestError: %v", resp.Status)
	}
	return resp
}

// Status - GET  /status - Obtain information about this server and/or useful to the Central web UI.
func Status() Response {
	log.Trace("Status")
	url := getEndpoint("status")
	resp := Request("GET", url, []byte{})

	showFullResponseBody(resp)

	if resp.StatusCode != 200 {
		log.Fatalf("RequestError: %v", resp.Status)
	}
	return resp
}
