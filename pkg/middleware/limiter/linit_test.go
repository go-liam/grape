package limiter

import (
	"fmt"
	"strings"
	"testing"
)

func TestLimit_1(t *testing.T) {
	fmt.Println(strings.Compare("192.168.31.235", ":::1"))
	fmt.Println(strings.Compare("192.168.31.235", "192.168.31.235"))

	fmt.Println(strings.Compare("golang", "GoLang"))
	fmt.Println(strings.Compare("GoLang", "golang"))
	fmt.Println(strings.Compare("golang", "golang"))

	fmt.Println(strings.EqualFold("GoLang", "golang"))
	fmt.Println(strings.EqualFold("golang", "GoLang"))
}
