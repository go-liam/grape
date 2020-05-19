package role

import "github.com/go-liam/util/conv"

//import "grape/pkg/util/conv"

// ListByIDs :
func ListByIDs(ls []int,list []*ModelRole) []*ModelRole {
	lsB := make( []*ModelRole,0)
	for _,v := range list{
		for _, i:= range ls {
			if v.RoleID == i {
				lsB = append(lsB,v)
			}
		}
	}
	return lsB
}
// ListPowerIDsByIDs :
func ListPowerIDsByIDs(ls []int,list []*ModelRole) []int {
	var idS []int
	for _,v := range list{
		for _, i:= range ls {
			if v.RoleID == i {
				idS = conv.ArrayIntJoin(idS, conv.StringToIntArray(v.PowerIDS))
			}
		}
	}
	return idS
}