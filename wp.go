package wp

import (
	"bufio"
	"fmt"
	"os"
)

func Search_and_replace(filename string, show_verbose_mode bool) {
	file, _ := os.Open(filename)
	defer file.Close()

	var lines []string
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}
	fmt.Printf("%v", lines)
}
