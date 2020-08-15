package site

type Site interface {
	FindOne(id int64) (*Model, error)
	//Demo(id int,name string) (string,error)
}
