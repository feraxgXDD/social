package intrepritator

import (
	basedata "just/BaseData"
)

func WriteIntrFile(writeFile string, manual map[string]map[string]string) {
	for k, _ := range manual {
		basedata.Write(writeFile, k+" -> {\n")
		for k2, v := range manual[k] {
			basedata.Write(writeFile, "\t"+k2+" = "+v+"\n")

		}
	}
}
