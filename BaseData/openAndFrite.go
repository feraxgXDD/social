package basedata

import (
	"bufio"
	"fmt"
	"os"
)

func Read(file string) []string {

	f, err := os.Open(file)
	if err != nil {
		fmt.Println(err)
		return nil
	}

	var lines []string

	scanner := bufio.NewScanner(f)

	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}

	return lines

}

func Write(file string, text string) {
	files, err := os.Open(file)
	if err != nil {
		fmt.Println(err)
		return
	}
	defer files.Close()

	files.WriteString(text)
}
