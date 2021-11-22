package opensea

import (
	"context"
	"fmt"
	"log"
	"net/url"
	"os"
	"os/signal"
	"sync"
	"time"

	"github.com/google/uuid"
	"github.com/gorilla/websocket"
	"github.com/twmb/franz-go/pkg/kgo"
)

var (
	producerName = "connectors/opensea"
)

func NewProducer(seeds []string, apiKey string, opts ...Opt) (*Producer, error) {
	c := defaultConfig()
	for _, opt := range opts {
		opt.apply(&c)
	}
	id, err := uuid.NewRandom()
	if err != nil {
		log.Printf("could not generate uuid error='%v', producer=%s\n", err, producerName)
	}

	return &Producer{
		id:     id,
		seeds:  seeds,
		apiKey: apiKey,
		config: c,
	}, nil
}

type Producer struct {
	id     uuid.UUID
	seeds  []string
	apiKey string
	config config
}

func (s *Producer) Start() error {
	log.Printf("starting producer=%s\n", producerName)
	cl, err := kgo.NewClient(
		kgo.SeedBrokers(s.seeds...),
	)
	if err != nil {
		log.Printf("could not create client error='%v', producer=%s\n", err, producerName)
		return err
	}
	defer cl.Close()

	wg := sync.WaitGroup{}
	if s.config.wsEventsEnabled {
		interrupt := make(chan os.Signal, 1)
		signal.Notify(interrupt, os.Interrupt)
		URL := url.URL{
			Scheme:   s.config.wsScheme,
			Host:     s.config.wsHost,
			Path:     s.config.wsPath,
			RawQuery: fmt.Sprintf("token=%s", s.apiKey),
		}
		c, _, err := websocket.DefaultDialer.Dial(URL.String(), nil)
		if err != nil {
			log.Printf("dial error=%v, producer=%s\n", err, producerName)
			return err
		}
		defer c.Close()
		done := make(chan struct{})
		wg.Add(1)
		go func() {
			defer func() {
				close(done)
				wg.Done()
			}()

			ctx := context.Background()
			for {
				_, message, err := c.ReadMessage()
				if err != nil {
					log.Printf("read message error=%v, producer=%s\n", err, producerName)
					return
				}
				log.Printf("received message=%s\n", message)
				record := &kgo.Record{Topic: ".workshop.connectors.opensea.events.websocket", Value: message}
				cl.Produce(ctx, record, func(_ *kgo.Record, err error) {
					defer wg.Done()
					if err != nil {
						log.Printf("record had a produce error='%v', producer=%s\n", err, producerName)
					}
				})
			}
		}()

		for {
			select {
			case <-done:
				return nil
			case <-interrupt:
				log.Printf("caught interrupt signal - quitting!\n")
				err := c.WriteMessage(websocket.CloseMessage, websocket.FormatCloseMessage(websocket.CloseNormalClosure, ""))

				if err != nil {
					log.Printf("write close error=%v, producer=%s\n", err, producerName)
					return err
				}
				select {
				case <-done:
				case <-time.After(2 * time.Second):
				}
				return nil
			}
		}
	}
	// if s.config.pollEvents {
	// 		wg.Add(1)
	// 		record := &kgo.Record{Topic: ".workshop.connectors.opensea.events.poll", Value: []byte("bar")}
	// 		cl.Produce(ctx, record, func(_ *kgo.Record, err error) {
	// 			defer wg.Done()
	// 			if err != nil {
	// 				log.Printf("record had a produce error='%v', producer=%s\n", err, producerName)
	// 			}
	// 		})
	// }
	// if s.config.pollListings {
	// 		wg.Add(1)
	// 		record := &kgo.Record{Topic: ".workshop.connectors.opensea.listings.poll", Value: []byte("bar")}
	// 		cl.Produce(ctx, record, func(_ *kgo.Record, err error) {
	// 			defer wg.Done()
	// 			if err != nil {
	// 				log.Printf("record had a produce error='%v', producer=%s\n", err, producerName)
	// 			}
	// 		})
	// }
	// if s.config.pollOffers {
	// 		wg.Add(1)
	// 		record := &kgo.Record{Topic: ".workshop.connectors.opensea.offers.poll", Value: []byte("bar")}
	// 		cl.Produce(ctx, record, func(_ *kgo.Record, err error) {
	// 			defer wg.Done()
	// 			if err != nil {
	// 				log.Printf("record had a produce error='%v', producer=%s\n", err, producerName)
	// 			}
	// 		})
	// }
	// if s.config.pollCollections {
	// 		wg.Add(1)
	// 		record := &kgo.Record{Topic: ".workshop.connectors.opensea.collections.poll", Value: []byte("bar")}
	// 		cl.Produce(ctx, record, func(_ *kgo.Record, err error) {
	// 			defer wg.Done()
	// 			if err != nil {
	// 				log.Printf("record had a produce error='%v', producer=%s\n", err, producerName)
	// 			}
	// 		})
	// }
	// if s.config.pollCollectionStats {
	// 		wg.Add(1)
	// 		record := &kgo.Record{Topic: ".workshop.connectors.opensea.collectionstats.poll", Value: []byte("bar")}
	// 		cl.Produce(ctx, record, func(_ *kgo.Record, err error) {
	// 			defer wg.Done()
	// 			if err != nil {
	// 				log.Printf("record had a produce error='%v', producer=%s\n", err, producerName)
	// 			}
	// 		})
	// }
	wg.Wait()

	return nil
}

// config is the configuration for the OpenSea producer.
//
// It streams the following resources from the OpenSea API:
// - events (websocket)
// - events (poll)
// - listings (poll)
// - offers (poll)
// - collections (poll)
// - collection stats (poll)
type config struct {
	wsScheme                   string
	wsHost                     string
	wsPath                     string
	wsEventsEnabled            bool
	wsEventsCollections        []string
	pollEventsEnabled          bool
	pollListingsEnabled        bool
	pollOffersEnabled          bool
	pollCollectionsEnabled     bool
	pollCollectionStatsEnabled bool
}

// Opt is an option to configure a producer.
type Opt interface {
	apply(*config)
}

type WebsocketEventsOpt struct {
	Enabled     bool
	Collections []string
}

func (o *WebsocketEventsOpt) apply(c *config) {
	c.wsEventsEnabled = true
	c.wsEventsCollections = o.Collections
}

type PollEventsOpt struct {
	Enabled bool
}

func (o *PollEventsOpt) apply(c *config) {
	c.pollEventsEnabled = true
}

type PollListingsOpt struct {
	Enabled bool
}

func (o *PollListingsOpt) apply(c *config) {
	c.pollListingsEnabled = true
}

type PollOffersOpt struct {
	Enabled bool
}

func (o *PollOffersOpt) apply(c *config) {
	c.pollOffersEnabled = true
}

type PollCollectionsOpt struct {
	Enabled bool
}

func (o *PollCollectionsOpt) apply(c *config) {
	c.pollCollectionsEnabled = true
}

type PollCollectionStatsOpt struct {
	Enabled bool
}

func (o *PollCollectionStatsOpt) apply(c *config) {
	c.pollCollectionStatsEnabled = true
}

func defaultConfig() config {
	return config{
		wsScheme:                   "wss",
		wsHost:                     "stream.openseabeta.com",
		wsPath:                     "/socket/websocket",
		wsEventsEnabled:            true,
		wsEventsCollections:        []string{"*"},
		pollEventsEnabled:          true,
		pollListingsEnabled:        true,
		pollOffersEnabled:          true,
		pollCollectionsEnabled:     true,
		pollCollectionStatsEnabled: true,
	}
}
