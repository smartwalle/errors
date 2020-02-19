package main

import (
	"encoding/json"
	"fmt"
	"github.com/smartwalle/errors"
)

func main() {
	var e1 = errors.New(1, "error1")
	fmt.Println(e1)
	var e2 = errors.Parse(e1.Error())
	fmt.Println(e2)

	var e3 = errors.New(3, "error3").Location().WithData("sss")

	fmt.Println(e3)

	mb, _ := json.Marshal(e3)
	fmt.Println(string(mb))

	var e5 = errors.Parse(`{"code":3,"message":"error3","file":"/Users/yang/go/src/github.com/smartwalle/errors/sample/main.go","line":12,"func":"main.main"}`)
	fmt.Println(e5)
}
