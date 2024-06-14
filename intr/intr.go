package intr

import (
	"fmt"

	basedata "just/BaseData"
)

func InitIntr(fileRead string) map[string]map[string]string {
	fmt.Println("intrepritator package initialized")

	var returnedMap = map[string]map[string]string{}

	file := basedata.Read(fileRead)

	var currentKey string
	var valueMap = map[string]string{}

	for _, line := range file {
		tokens := basedata.Lexer2(line)
		fmt.Println(tokens)

		if len(tokens) > 0 {
			if tokens[0] == "->" && tokens[1] == "{" {
				currentKey = tokens[2]
			} else if currentKey != "" {
				for j := 0; j < len(tokens); j++ {
					if tokens[j] == "}" {
						returnedMap[currentKey] = valueMap
						valueMap = map[string]string{} // Очищаем valueMap для следующего ключа
						currentKey = ""
						break
					}
					if tokens[j] == "=" {
						if j+2 < len(tokens) {
							valueMap[tokens[j+1]] = tokens[j+2]
						}
					}
				}
			}
		}
	}
	return returnedMap
}
