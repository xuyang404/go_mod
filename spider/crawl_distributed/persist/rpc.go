package persist

import (
	"github.com/olivere/elastic"
	"log"
	"spider/engine"
	"spider/persist"
)

type ItemSaverService struct {
	Client *elastic.Client
	Index string
}

func (s *ItemSaverService) Save(item engine.Item, result *string) error {
	err := persist.Save(s.Client, s.Index, item)
	log.Printf("Save Item is: %v, ", item)
	if err == nil{
		*result = "ok"
	}else {
		log.Printf("Save Item error: %v, ", err)
	}
	return err
}