package wp

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"strconv"
	"strings"
)

func Search_and_replace(filename string, show_verbose_mode bool) {
	file, _ := os.Open(filename)
	defer file.Close()

	scanner := bufio.NewScanner(file)
	var cdata bool

	for scanner.Scan() {
		line := scanner.Text()
		if strings.Contains(line, "<content:encoded>") {
			cdata = true
		} else if strings.Contains(line, "</content:encoded>") {
			cdata = false
		}
		if cdata {
			//headings shortcode to html
			re := regexp.MustCompile("size=\"h?")
			loc := re.FindStringIndex(line)
			if loc != nil {
				//headings
				hval := string(line[loc[1]]) //get the h val e.g. <h3> = 3

				line = strings.Replace(line, "[/heading]", "</h"+hval+">", 1)
				line = strings.Replace(line, "[heading size=\"h"+hval+"\"]", "<h"+hval+">", 1)
				line = strings.Replace(line, "h3>", "h2>", -1) //keep it semantic

				//fmt.Println(string(line[loc[1]]))

			}
			//lose the whitespace
			for i := 0; i < 120; i++ {
				line = strings.Replace(line, "[whitespace height=\""+strconv.Itoa(i)+"\"]", "", -1)
			}

		}
		line = strings.Replace(line, "<wp:post_type>page</wp:post_type>", "<wp:post_type>case-history</wp:post_type>", 1)

		fmt.Println(line)
	}

}
