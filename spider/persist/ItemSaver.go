package persist

import (
	"context"
	"github.com/olivere/elastic"
	"github.com/pkg/errors"
	"log"
	"spider/engine"
)

func ItemSaver(index string) (chan engine.Item, error) {
	client, err := elastic.NewClient(
		elastic.SetURL("http://www.vowcloud.cn:9200"),
		elastic.SetSniff(false),
	)
	if err != nil {
		return nil, err
	}
	out := make(chan engine.Item)
	go func() {
		count := 1
		for {
			item := <-out
			log.Printf("God item saver: %d# save: %+v", count, item)
			err := Save(client, index, item)
			if err != nil {
				log.Printf("God item saver: %d# err: %v", count, err)
				continue
			}
			count++
		}
	}()
	return out, nil
}

func Save(client *elastic.Client, index string, item engine.Item) (err error) {
	if item.Type == "" {
		return errors.New("Must supply Type")
	}

	indexService := client.Index().
		Index(index).
		Type(item.Type).
		BodyJson(item)

	if item.Id != "" {
		indexService.Id(item.Id)
	}

	_, err = indexService.
		Do(context.Background())
	if err != nil {
		return err
	}

	return nil
}