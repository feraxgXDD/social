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
	f, err := os.OpenFile(file, os.O_WRONLY|os.O_CREATE|os.O_APPEND, 0777)
	if err != nil {
		fmt.Println("Error opening file:", err)
		return
	}
	defer f.Close()

	_, err = f.WriteString(text)
	if err != nil {
		fmt.Println("Error writing to file:", err)
	}
}
