package main

import (
	"flag"
	"github.com/olivere/elastic"
	"log"
	"spider/crawl_distributed/config"
	"spider/crawl_distributed/persist"
	"spider/crawl_distributed/rpcsupport"
)

var host = flag.String("host", ":1234", "ItemSaver host")

func main() {
	flag.Parse()
	log.Fatal(ServeRpc(*host, config.ElasticIndex))
}

func ServeRpc(host string, index string) error {
	client, err := elastic.NewClient(
		elastic.SetSniff(false),
		elastic.SetURL(config.ElasticHost),
	)
	if err != nil{
		return err
	}

	return rpcsupport.ServeRpc(host, &persist.ItemSaverService{
		Client: client,
		Index:  index,
	})
}