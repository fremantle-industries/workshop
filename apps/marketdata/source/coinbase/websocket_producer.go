package coinbase

import (
	// "log"

	// "github.com/ergo-services/ergo/etf"
	"github.com/ergo-services/ergo/gen"
	"github.com/fremantle-industries/tabletop/pkg/producer"
	// "github.com/twmb/franz-go/pkg/kgo"
	// "nhooyr.io/websocket"
	// "nhooyr.io/websocket/wsjson"
)

// var (
// 	coinbaseWebsocketReadLimit int64 = 1048576 // 1024 Kb
// )

func NewWebsocketProducer(url string, seedBrokers []string) gen.ServerBehavior {
	return producer.NewWebsocketProducer("coinbaseWebsocketProducer", url, seedBrokers)
}

type CoinbaseWebsocketMsgSubscribe struct {
	Type       string        `json:"type"`
	ProductIDs []string      `json:"product_ids"`
	Channels   []interface{} `json:"channels"`
}

type CoinbaseWebsocketChannel struct {
	Name       string   `json:"name"`
	ProductIDs []string `json:"product_ids"`
}

// type coinbaseWebsocketProducer struct {
// 	producer.WebsocketProducer
// }
//
// func (p *coinbaseWebsocketProducer) Init(process *gen.ServerProcess, args ...etf.Term) error {
// 	wsConn, _, err := websocket.Dial(process.Context(), p.opts.url, nil)
// 	if err != nil {
// 		return err
// 	}
// 	wsConn.SetReadLimit(coinbaseWebsocketReadLimit)
// 	p.wsConn = wsConn
//
// 	kClient, err := kgo.NewClient(
// 		kgo.SeedBrokers(p.opts.seeds...),
// 	)
// 	if err != nil {
// 		return err
// 	}
// 	p.kClient = kClient
//
// 	go p.readPump(process)
// 	go p.writePump(process)
//
// 	subMsg := coinbaseWebsocketSubscribe{
// 		Type:       "subscribe",
// 		ProductIDs: []string{"ETH-USD", "ETH-EUR"},
// 		Channels: []interface{}{
// 			"level2",
// 			"heartbeat",
// 			coinbaseWebsocketChannel{
// 				Name:       "ticker",
// 				ProductIDs: []string{"ETH-BTC", "ETH-USD"},
// 			},
// 		},
// 	}
// 	p.writeWS <- subMsg
//
// 	return nil
// }
//
// func (p *coinbaseWebsocketProducer) HandleCast(process *gen.ServerProcess, message etf.Term) gen.ServerStatus {
// 	typeAndMsg := message.([]interface{})
// 	msgType := typeAndMsg[0].(websocket.MessageType)
// 	msg := typeAndMsg[1].([]byte)
//
// 	log.Printf("HandleCast type=%v, msg=%s, producer=%s\n", msgType, msg, p.name)
// 	return gen.ServerStatusOK
// }
//
// func (p *coinbaseWebsocketProducer) HandleCall(process *gen.ServerProcess, from gen.ServerFrom, message etf.Term) (etf.Term, gen.ServerStatus) {
// 	log.Printf("HandleCall producer=%s\n", p.name)
// 	return process.State, gen.ServerStatusOK
// }
//
// func (p *coinbaseWebsocketProducer) Terminate(process *gen.ServerProcess, reason string) {
// 	defer func() {
// 		p.kClient.Close()
// 		err := p.wsConn.Close(websocket.StatusAbnormalClosure, reason)
// 		if err != nil {
// 			log.Printf("Terminate wsConn.Close error=%s, producer=%s", err, p.name)
// 		}
// 	}()
// 	log.Printf("Terminate reason=%s, producer=%s\n", reason, p.name)
// }
//
// func (p *coinbaseWebsocketProducer) readPump(process *gen.ServerProcess) {
// 	for {
// 		msgType, msg, err := p.wsConn.Read(process.Context())
// 		if err != nil {
// 			log.Printf("readPump error=%v, producer=%s\n", err, p.name)
// 			break
// 		}
// 		process.Cast(process.Self(), []interface{}{msgType, msg})
// 	}
// }
//
// func (p *coinbaseWebsocketProducer) writePump(process *gen.ServerProcess) {
// 	for {
// 		select {
// 		case <-process.Context().Done():
// 			break
// 		case msg := <-p.writeWS:
// 			err := wsjson.Write(process.Context(), p.wsConn, msg)
// 			if err != nil {
// 				log.Printf("writePump error=%s, producer=%s\n", err, p.name)
// 				break
// 			}
// 		}
// 	}
// }
