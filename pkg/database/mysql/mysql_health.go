package mysql

import "log"

type ModelTable struct {
	ID int `gorm:"column:ID"`
}

func Health() int {
	println("Mysql:Server.URL=",Server.URL)
	item := new(ModelTable)
	sql := "select 1 as ID  "
	if Server.IsConnect{
		Server.Engine().Raw(sql).Scan(item)
	}else {
		Server.NewEngine()
	}
	log.Printf("[INFO] =%+v\n",item)
	return item.ID
}
