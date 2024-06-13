package writerFunc

import (
	"bufio"
	"fmt"
	"os"
)

func Write(file string) []string {

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
