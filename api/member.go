package api

import (
	"encoding/json"
	"fmt"

	log "github.com/sirupsen/logrus"
)

// MemberCreate - POST  /network/{id}/member/ - Creates a new member
func MemberCreate(member Member) Response {
	log.Trace("MemberCreate")
	marshalled, _ := json.Marshal(member)
	url := getEndpoint("member")
	resp := Request("POST", url, []byte(marshalled))

	showFullResponseBody(resp)

	if resp.StatusCode != 201 {
		log.Fatalf("RequestError: %v", resp.Status)
	}
	return resp
}

// MemberDelete - DELETE  /network/{id}/member/{name} - Deletes a member
func MemberDelete(network string, membername string) Response {
	log.Trace("MemberDelete")
	url := getEndpoint(fmt.Sprintf("/network/%v/member/%v", network, membername))
	resp := Request("DELETE", url, []byte{})

	showFullResponseBody(resp)

	if resp.StatusCode != 200 {
		log.Fatalf("RequestError: %v", resp.Status)
	}
	return resp
}

// MemberGet - GET  /network/{id}/member/{name} - Retrieves a specific member
func MemberGet(network string, membername string) Response {
	log.Trace("MemberGet")
	url := getEndpoint(fmt.Sprintf("/network/%v/member/%v", network, membername))
	resp := Request("GET", url, []byte{})

	showFullResponseBody(resp)

	if resp.StatusCode != 200 {
		log.Fatalf("RequestError: %v", resp.Status)
	}
	return resp
}

// MemberList - GET /network/{id}/member/ - Retrieves all members
func MemberList(network string) Response {
	log.Trace("MemberList")
	url := getEndpoint(fmt.Sprintf("/network/%v/member", network))
	resp := Request("GET", url, []byte{})

	showFullResponseBody(resp)

	if resp.StatusCode != 200 {
		log.Fatalf("RequestError: %v", resp.Status)
	}
	return resp
}

// MemberPut - PUT  /network/{id}/member/{name} - Creates or updates a member
func MemberPut(network string, member Member) Response {
	log.Trace("MemberPut")
	marshalled, _ := json.Marshal(member)
	url := getEndpoint(fmt.Sprintf("/network/%v/member/%v", network, member.ID))
	resp := Request("PUT", url, []byte(marshalled))

	showFullResponseBody(resp)

	if resp.StatusCode > 201 {
		log.Fatalf("RequestError: %v", resp.Status)
	}
	return resp
}
