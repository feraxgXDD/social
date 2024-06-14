package intrepritator

import (
	"fmt"

	basedata "just/BaseData"

)

func initIntr(fileRead string) map[string]map[string]string {
	fmt.Println("intrepritator package initialized")

	var returnedMap = map[string]map[string]string{}

	file := basedata.Read(fileRead)

	for i := 0; i < len(file); i++ {
		fileLexing := basedata.Lexer2(file[i])

		if fileLexing[i] == "->" {
			if fileLexing[i+1] == "{" {
				for j := len(file) - i; j > i; j-- {
					if fileLexing[j] == "}" {
						fmt.Println(file[i+1 : j])
						break
					}
					if fileLexing[j+1] == "=" {
						returnedMap[fileLexing[i-1]] = map[string]string{fileLexing[j+2]: fileLexing[j+3]}
						break
					}
				}
			}

		}
	}
	return returnedMap
}
