package role

import (
	"log"
	"testing"
)

func TestListByIDs(t *testing.T) {
	list, _ := Server2.List()
	ls := ListByIDs([]int64{1, 2}, list)
	for _, v := range ls {
		log.Printf("v=%+v\n", v)
	}
}

func TestListPowerIDsByIDs(t *testing.T) {
	list, _ := Server2.List()
	ls := ListPowerIDsByIDs([]int64{1, 2}, list)
	log.Printf("v=%+v\n", ls)
}
