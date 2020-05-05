package orelang_test

import (
	"encoding/json"
	"fmt"
	"log"

	"github.com/d-tsuji/orelang"
)

func Example() {
	engine := orelang.NewEngine()

	input := `
["step",
  ["set", "i", 10],
  ["set", "sum", 0],
  ["until", ["=", ["get", "i"], 0], [
    "step",
    ["set", "sum", ["+", ["get", "sum"], ["get", "i"]]],
    ["set", "i", ["+", ["get", "i"], -1]]
  ]],
  ["get", "sum"]
]`

	var i interface{}
	if err := json.Unmarshal([]byte(input), &i); err != nil {
		log.Fatalf("json unmershal: %v", err)
	}

	fmt.Println(engine.Eval(i))
	// Output:
	// 55
}
