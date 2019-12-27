package client

import (
	"log"
	"net/rpc"
	"spider/engine"
)

func ItemSaver(chanClient chan *rpc.Client, index string, serviceMethod string) (chan engine.Item) {

	out := make(chan engine.Item)
	go func() {
		count := 1
		for {
			item := <-out
			log.Printf("God item saver: %d# save: %+v", count, item)
			result := ""
			client := <- chanClient
			err := client.Call(serviceMethod, item, &result)
			if err != nil || result != "ok"{
				log.Printf("God item saver: %d# err: %v", count, err)
				continue
			}
			count++
		}
	}()
	return out
}