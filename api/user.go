package api

import (
	"encoding/json"
	"fmt"

	log "github.com/sirupsen/logrus"
)

// UserCreate - POST /{db}/user/ - Creates a new user
func UserCreate(user User) Response {
	log.Trace("zerotier UserCreate")
	marshalled, _ := json.Marshal(user)
	url := getEndpoint("user/")
	resp := Request("POST", url, []byte(marshalled))

	showFullResponseBody(resp)

	if resp.StatusCode != 201 {
		log.Fatalf("RequestError: %v", resp.Status)
	}
	return resp
}

// UserDelete - DELETE /{db}/user/{name} - Deletes a user
func UserDelete(username string) Response {
	log.Trace("zerotier UserDelete")
	url := getEndpoint(fmt.Sprintf("user/%v", username))
	resp := Request("DELETE", url, []byte{})

	showFullResponseBody(resp)

	if resp.StatusCode != 200 {
		log.Fatalf("RequestError: %v", resp.Status)
	}
	return resp
}

// UserGet - GET /{db}/user/{name} - Retrieves a specific user
func UserGet(username string) Response {
	log.Trace("zerotier UserGet")
	url := getEndpoint(fmt.Sprintf("user/%v", username))
	resp := Request("GET", url, []byte{})

	showFullResponseBody(resp)

	if resp.StatusCode != 200 {
		log.Fatalf("RequestError: %v", resp.Status)
	}
	return resp
}

// UserList - GET /{db}/user/ - Retrieves all users
func UserList() Response {
	log.Trace("zerotier UserList")
	url := getEndpoint("user")
	resp := Request("GET", url, []byte{})

	showFullResponseBody(resp)

	if resp.StatusCode != 200 {
		log.Fatalf("RequestError: %v", resp.Status)
	}
	return resp
}

// UserPut - PUT /{db}/user/{name} - Creates or updates a user
func UserPut(user User) Response {
	log.Trace("zerotier UserPut")
	marshalled, _ := json.Marshal(user)
	url := getEndpoint(fmt.Sprintf("user/%v", user.Name))
	resp := Request("PUT", url, []byte(marshalled))

	showFullResponseBody(resp)

	if resp.StatusCode > 201 {
		log.Fatalf("RequestError: %v", resp.Status)
	}
	return resp
}
