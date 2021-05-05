package main

import (
	"github.com/astravexton/logger"
	photoprism "github.com/astravexton/photoprism-client-go"
)

func main() {
	logger.Level = 4
	client := photoprism.New("http://localhost:8080")
	err := client.Auth(photoprism.NewClientAuthLogin("admin", "missy"))
	if err != nil {
		halt(4, "Error logging into API: %v", err)
	}
	logger.Always("Login Success!")
}
