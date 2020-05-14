package conv

import (
	"encoding/json"
	"fmt"
	"log"
	"strconv"
	"strings"
)

// StringArrayToString :
func StringArrayToString(rs interface{}) string {
	v := fmt.Sprintf("%v", rs)
	return strings.Trim(strings.Replace(v, " ", ",", -1), "")
}

// StringToStringArray : [a,b,c]
func StringToStringArray(rs string) []string {
	if rs == "" {
		return []string{}
	}
	s0 := strings.Trim(rs, "[]")
	s := strings.Split(s0, ",")
	return s
}
// StringToIntArray : "[1,2,3,4]"
func StringToIntArray(rs string) []int {
	if rs == "" {
		return []int{}
	}
	s0 := strings.Trim(rs, "[]")
	s := strings.Split(s0, ",")
	var ls []int
	for _,v:= range s{
		ls = append(ls,StringToInt(v,0))
	}
	return ls
}

// ArrayToString :
func ArrayToString(list []int64, delim string) string {
	return strings.Trim(strings.Replace(fmt.Sprint(list), " ", delim, -1), "[]")
}

// ArrayIntToString :
func ArrayIntToString(list []int, delim string) string {
	return strings.Trim(strings.Replace(fmt.Sprint(list), " ", delim, -1), "[]")
}

// StringToInt :
func StringToInt(st string, defaultInt int) int {
	if st == "" {
		return defaultInt
	}
	vInt, err := strconv.Atoi(st)
	if err != nil {
		println("[ERROR]string到int:", err)
		return defaultInt
	}
	return vInt
}

// StringToInt64 :
func StringToInt64(st string, defaultInt int64) int64 {
	if st == "" {
		return defaultInt
	}
	vInt, err := strconv.ParseInt(st, 10, 64)
	if err != nil {
		println("[ERROR]string到int64:", err)
		return defaultInt
	}
	return vInt
}

// StructToJsonString :
func StructToJsonString(info interface{}) string {
	if info == nil {
		return ""
	}
	jsons, errs := json.Marshal(info) //转换成JSON返回的是byte[]
	if errs != nil {
		log.Println("[ERROR]", errs.Error())
		//return ""
	}
	return string(jsons) //byte[]转换成string
}

// JsonStringToStruct :
func JsonStringToStruct(jsonSt string, obj interface{}) {
	if jsonSt == "" || obj == nil {
		return
	}
	errs := json.Unmarshal([]byte(jsonSt), obj)
	if errs != nil {
		log.Println("[ERROR]", errs.Error())
	}
}

func NilArrayStringTChange(info []string) []string {
	if info == nil {
		return make([]string, 0)
	}
	return info
}

// NilArrayIntTChange :
func NilArrayIntTChange(info []int) []int {
	if info == nil {
		return make([]int, 0)
	}
	return info
}

// ArrayInt64Contains :array 是否存在某值
func ArrayInt64Contains(list []int64, item int64) bool {
	if len(list) <= 0 {
		return false
	}
	for _, v := range list {
		if v == item {
			return true
		}
	}
	return false
}
// ArrayIntContains :
func ArrayIntContains(list []int, item int) bool {
	if len(list) <= 0 {
		return false
	}
	for _, v := range list {
		if v == item {
			return true
		}
	}
	return false
}


func ArrayStringContains(list []string, item string) bool {
	if len(list) <= 0 {
		return false
	}
	for _, v := range list {
		if strings.Compare(v, item) == 0{
			return true
		}
	}
	return false
}

// ArrayInt64Join :合并多个 arry
func ArrayInt64Join(arr1 []int64, arr2 []int64) []int64 {
	return append(arr1, arr2...)
}

// ArrayIntJoin :
func ArrayIntJoin(arr1 []int, arr2 []int) []int {
	return append(arr1, arr2...)
}

// ArrayStringJoin :
func ArrayStringJoin(arr1 []string, arr2 []string) []string {
	return append(arr1, arr2...)
}