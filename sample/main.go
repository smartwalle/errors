package main

import (
	"encoding/json"
	"fmt"
	"github.com/smartwalle/errors"
)

func main() {
	var e1 = errors.New("1", "error1")
	var e2 = errors.New("2", "error2").WithError(e1)
	var e3 = errors.New("3", "error3").WithError(e2).Location()
	var e4 = errors.WithData(e3)
	var e5 = e4.Stack()

	fmt.Println(e3)

	mb, _ := json.Marshal(e3)
	fmt.Println(string(mb))

	mb, _ = json.Marshal(e4)
	fmt.Println(string(mb))

	mb, _ = json.Marshal(e5)
	fmt.Println(string(mb))
}
