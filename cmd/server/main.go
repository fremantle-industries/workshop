package main

import (
	"log"
	"os"
	"sync"

	"github.com/fremantle-industries/workshop/pkg/config"
	"github.com/fremantle-industries/workshop/pkg/inbound_connectors"
	"github.com/fremantle-industries/workshop/pkg/outbound_connectors"
	"github.com/fremantle-industries/workshop/pkg/web"
)

var defaultUIPort = "8080"
var defaultUIEnabled = false
var defaultAPIPort = "8081"
var defaultAPIEnabled = false
// var configFile = "workshop.yaml"

func main() {
	// c := config.NewFromYAMLFile(configFile)
	c := config.New([]byte(``))
	wg := sync.WaitGroup{}

	if os.Getenv("UI_ENABLED") == "true" {
		// if c.UI.Enabled {
		uiPort := os.Getenv("UI_HTTP_PORT")
		if uiPort == "" {
			uiPort = defaultUIPort
		}
		s := web.NewUIServer(uiPort)
		wg.Add(1)
		go func() {
			defer wg.Done()
			log.Fatal(s.ListenAndServe())
		}()
	}

	if os.Getenv("API_ENABLED") == "true" {
		apiPort := os.Getenv("API_HTTP_PORT")
		if apiPort == "" {
			apiPort = defaultAPIPort
		}
		s := web.NewAPIServer(apiPort)
		wg.Add(1)
		go func() {
			defer wg.Done()
			log.Fatal(s.ListenAndServe())
		}()
	}

	if os.Getenv("INBOUND_CONNECTORS_VENUES_ENABLED") == "true" {
		s := inbound_connectors.NewServer(c.InboundConnectors)
		wg.Add(1)
		go func() {
			defer wg.Done()
			log.Fatal(s.Start())
		}()
	}

	if os.Getenv("OUTBOUND_CONNECTORS_VENUES_ENABLED") == "true" {
		s := outbound_connectors.NewServer(c.OutboundConnectors)
		wg.Add(1)
		go func() {
			defer wg.Done()
			log.Fatal(s.Start())
		}()
	}

	wg.Wait()
}
