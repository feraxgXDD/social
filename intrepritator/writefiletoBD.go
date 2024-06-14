package intrepritator

import (
	"strings"

	basedata "just/BaseData"
)

func WriteIntrFile(writeFile string, manual map[string]map[string]string) {
	var content strings.Builder

	for key, subMap := range manual {
		content.WriteString(key + " -> {\n")
		for subKey, value := range subMap {
			content.WriteString("\t" + subKey + " = " + value + "\n")
		}
		content.WriteString("}\n")
	}

	basedata.Write(writeFile, content.String())
}
