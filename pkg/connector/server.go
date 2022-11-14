package connector

import (
	"context"
	"fmt"
	"sync"

	"github.com/twmb/franz-go/pkg/kgo"
)

func NewServer(config Config) *Server {
	return &Server{
		seeds: config.Seeds,
	}
}

type Config struct {
	Enabled bool     `yaml:"enabled"`
	Seeds   []string `yaml:"seeds"`
}

type Server struct {
	seeds []string
}

func (s *Server) Start() error {
	// One client can both produce and consume!
	// Consuming can either be direct (no consumer group), or through a group. Below, we use a group.
	cl, err := kgo.NewClient(
		kgo.SeedBrokers(s.seeds...),
	)
	if err != nil {
		panic(err)
	}
	defer cl.Close()

	ctx := context.Background()

	// 1.) Producing a message
	// All record production goes through Produce, and the callback can be used
	// to allow for synchronous or asynchronous production.
	var wg sync.WaitGroup
	wg.Add(1)
	record := &kgo.Record{Topic: "foo", Value: []byte("bar")}
	cl.Produce(ctx, record, func(_ *kgo.Record, err error) {
		defer wg.Done()
		if err != nil {
			fmt.Printf("record had a produce error: %v\n", err)
		}
	})
	wg.Wait()

	return nil
}
