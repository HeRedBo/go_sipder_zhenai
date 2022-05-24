package persist

import (
	"GoSpider/ConcurrenceTask/persist"
	"GoSpider/ConcurrenceTask/zhenai/model"
	"github.com/olivere/elastic"
	"log"
)

type ItemSaveService struct {
	 Client *elastic.Client
	 Index string
}

func (s *ItemSaveService) Save(item model.Member, result *string) error {
	id, err := persist.Save(s.Client, s.Index, item)
	if err ==  nil {
		*result = "ok"
		log.Printf("id : %s, Item: %v;  Saved", id,item)
	} else {
		log.Printf("Save item %v Error: %v", item, err)
	}
	return err
}