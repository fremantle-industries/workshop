package main

import (
	"log"
	"os"
	"sync"

	"github.com/fremantle-industries/workshop/pkg/web"
)

var defaultUiPort = "8080"
var defaultAPIPort = "8081"

func main() {
	apiPort := os.Getenv("API_HTTP_PORT")
	if apiPort == "" {
		apiPort = defaultAPIPort
	}
	apiServer := web.NewAPIServer(apiPort)
	uiPort := os.Getenv("UI_HTTP_PORT")
	if uiPort == "" {
		uiPort = defaultUiPort
	}
	uiServer := web.NewUIServer(uiPort)

	wg := sync.WaitGroup{}
	wg.Add(2)
	go func() {
		defer wg.Done()
		log.Fatal(apiServer.ListenAndServe())
	}()
	go func() {
		defer wg.Done()
		log.Fatal(uiServer.ListenAndServe())
	}()
	wg.Wait()
}
