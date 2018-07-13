package main

import (
	"encoding/json"
	"fmt"
	"github.com/smartwalle/errors"
)

func main() {
	var err = errors.New("测试错误").Location()
	fmt.Println(err)

	mb, _ := json.Marshal(err)
	fmt.Println(string(mb))
}