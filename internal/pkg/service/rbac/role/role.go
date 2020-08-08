package role

import "github.com/go-liam/util/conv"

//import "grape/pkg/util/conv"

// ListByIDs :
func ListByIDs(ls []int64, list []*Model) []*Model {
	lsB := make([]*Model, 0)
	for _, v := range list {
		for _, i := range ls {
			if v.ID == i {
				lsB = append(lsB, v)
			}
		}
	}
	return lsB
}

// ListPowerIDsByIDs :
func ListPowerIDsByIDs(ls []int64, list []*Model) []int {
	var idS []int
	for _, v := range list {
		for _, i := range ls {
			if v.ID == i {
				idS = conv.ArrayIntJoin(idS, conv.StringToIntArray(v.PowerIDS))
			}
		}
	}
	return idS
}
