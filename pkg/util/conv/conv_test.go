package conv

import (
	"log"
	"testing"
)

import "fmt"

func TestFunc_StringArrayToString(t *testing.T) {
	v := StringArrayToString("[a,b,c]")
	println(v)
}

func TestFunc_StringToStringArray(t *testing.T) {
	v := StringToStringArray("[a,b,c]")
	fmt.Printf("%+v\n", v)
}

func Test_ArrayToString(t *testing.T) {
	list := []int64{1, 2, 3, 4}
	//list := []int64{1}
	rs := ArrayToString(list, ",")
	fmt.Println(rs)
	// 1,2,3,4
}

func Test_ArrayIntToString(t *testing.T) {
	list := []int{1, 2, 3, 4}
	//list := []int64{1}
	rs := ArrayIntToString(list, ",")
	fmt.Println(rs)
	// 1,2,3,4
}

func TestStringToInt(t *testing.T) {
	v := StringToInt("1110",-1)
	println("v=",v)
}

type ModelUser struct {
	UserID int64
	UserName string
}
/*
=== RUN   TestName_StructToJsonString
{"i":110,"p":"","n":"name","e":"","t":false,"c":0}
*/
func TestName_StructToJsonString(t *testing.T) {
	us := new(ModelUser)
	us.UserID= 110
	us.UserName ="name"
	v1:= StructToJsonString(us)
	println(v1)
}

func TestName_JsonStringToStruct(t *testing.T) {
	st := `{"i":110,"p":"pp","n":"name","e":"eee","t":false,"c":0}`
	us := new(ModelUser)
	JsonStringToStruct(st,us)
	log.Printf("v= %+v \n", us)
}

func TestNilArrayStringTChange(t *testing.T) {
	NilArrayStringTChange(nil)
}

func TestNilArrayIntTChange(t *testing.T) {
	NilArrayIntTChange(nil)
}

func TestArrayInt64Contains(t *testing.T) {
	ArrayInt64Contains(nil,1)
}

func TestArrayIntContains(t *testing.T) {
	println( ArrayIntContains(nil,2))
}

func TestArrayInt64Join(t *testing.T) {
	a1 := []int64{1,2,3}
	a2 := []int64{1,2,3}
	join := ArrayInt64Join(a1,a2)
	log.Printf("v=%+v\n ",join)
}


func TestArrayIntJoin(t *testing.T) {
	a1 := []int{1,2,3}
	a2 := []int{1,2,3,4}
	join := ArrayIntJoin(a1,a2)
	log.Printf("v=%+v\n ",join)
}

func TestArrayStringJoin(t *testing.T) {
	a1 := []string{"1","2","3" }
	a2 := []string{"1","2","3"}
	join := ArrayStringJoin(a1,a2)
	log.Printf("v=%+v\n ",join)
}

func TestArrayStringContains(t *testing.T) {
	a1 := []string{"1","2","3" }
	a2 := []string{"1","2","3"}
	join := ArrayStringJoin(a1,a2)
	tag := "300"
	f := ArrayStringContains(join,tag)
	log.Printf("v=%+v\n ",f)
}

func TestStringToIntArray(t *testing.T) {
	ls := StringToIntArray("[1,2,3,4]")
	log.Printf("ls=%+v\n",ls)
}