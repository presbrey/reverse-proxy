package main

import (
	"net/url"
	"testing"

	"github.com/drewolson/testflight"
	"github.com/stretchr/testify/assert"
)

var (
	u, _    = url.Parse(*next)
	handler = NewServer(u).Handler
)

func TestUserAgent(t *testing.T) {
	testflight.WithServer(handler, func(r *testflight.Requester) {
		response := r.Get("/user-agent")
		assert.Equal(t, 200, response.StatusCode)
		assert.Equal(t, "{\n  \"user-agent\": \"Go 1.1 package http\"\n}\n", response.Body)
	})
}
