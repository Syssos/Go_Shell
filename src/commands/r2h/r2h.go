package main

import (
	"fmt"
	"os"
	"strings"
	"io/ioutil"
)

func main() {
	if len(os.Args) > 2 {
		readme := os.Args[1]
		htmlfile := os.Args[2]

		dat, err := ioutil.ReadFile(readme)
		for {
			for _, char := range dat {
				
				if char == 35 {
					dat = addHeadHtmlToFile(htmlfile, dat)
					// fmt.Println(dat)
					break
				} else if char == 10 {
					dat = addParaHTMLToFile(htmlfile, dat)
					// fmt.Println(dat)
					break
				}
			}
			if len(dat) < 2 {
				break
			}
		}

		check(err)

		// fmt.Println(dat, htmlfile)
	} else {
		fmt.Println("Please enter readme filename to use and html filename to generate")
	}
}

func check(err error) {
	if err != nil {
		fmt.Println(err)
	}
}

func addHeadHtmlToFile(file string, data []byte) []byte{
	word := []string{}
	count := 0
	headcount := 0

	for x, char := range data {
		if data[x] == 10 {
			f, err := os.OpenFile(file, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
			check(err)
			defer f.Close()
			_, werr := f.WriteString(fmt.Sprintf("<h%d>%v</h%d>\n", headcount, strings.Join(word, ""), headcount))
			check(werr)
			fmt.Println(headcount)
			return data[count+1:]

		} else if char == 35 {
			headcount = headcount + 1
			count = count + 1

		} else if char == 32 && data[x-1] == 35 {count=count+1} else {
			word = append(word, string(char))
			count = count + 1
		}
	}
	return data[:len(data)-2]
}

func addParaHTMLToFile(file string, data []byte) []byte {
	word := []string{}
	count := 0
	for x, char := range data[1:] {
		if char == 10 && x == 0 || data[x+1] == 35{
			return data[1:]
		} else if char == 10{
			f, err := os.OpenFile(file, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
			check(err)
			defer f.Close()
			_, werr := f.WriteString(fmt.Sprintf("\t<p>%v</p>\n", strings.Join(word, "")))
			check(werr)
			return data[count+1:]

		} else {
			word = append(word, string(char))
			count = count + 1
		}
	}
	return data[:len(data)-2]

}