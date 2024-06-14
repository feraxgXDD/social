package main

import (
	"just/intrepritator"
)

func main() {
	manual := map[string]map[string]string{
		"key1": {
			"value1": "data1",
			"value2": "data2",
		},
	}
	intrepritator.WriteIntrFile("base.txt", manual)
}
