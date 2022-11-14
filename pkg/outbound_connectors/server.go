package outbound_connectors

import (
	"context"
	"fmt"

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
		kgo.ConsumerGroup("my-group-identifier"),
		kgo.ConsumeTopics("foo"),
	)
	if err != nil {
		panic(err)
	}
	defer cl.Close()

	ctx := context.Background()

	// 2.) Consuming messages from a topic
	for {
		fetches := cl.PollFetches(ctx)
		if errs := fetches.Errors(); len(errs) > 0 {
			// All errors are retried internally when fetching, but non-retriable errors are
			// returned from polls so that users can notice and take action.
			panic(fmt.Sprint(errs))
		}

		// We can iterate through a record iterator...
		iter := fetches.RecordIter()
		for !iter.Done() {
			record := iter.Next()
			fmt.Println(string(record.Value), "from an iterator!")
		}

		// or a callback function.
		fetches.EachPartition(func(p kgo.FetchTopicPartition) {
			for _, record := range p.Records {
				fmt.Println(string(record.Value), "from range inside a callback!")
			}

			// We can even use a second callback!
			p.EachRecord(func(record *kgo.Record) {
				fmt.Println(string(record.Value), "from a second callback!")
			})
		})
	}

	return nil
}
