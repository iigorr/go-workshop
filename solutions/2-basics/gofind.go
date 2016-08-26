package main

import (
	"fmt"
	"io/ioutil"
	"strings"
)

func main() {

	if cntGo, cntIdiom, err := gofind(); err != nil {
		fmt.Printf("Error! %v\n.", err)
	} else {
		fmt.Printf("Found %d occurrences of'go' and %d of 'idiom'.\n", cntGo, cntIdiom)
	}

}

func gofind() (cntGo int, cntIdiom int, err error) {
	data, err := ioutil.ReadFile("effective-go.txt")
	if err != nil {
		return 0, 0, err
	}

	lines := strings.Split(string(data), "\n")

	for _, line := range lines {
		lowLine := strings.ToLower(line)
		if strings.Contains(lowLine, "go") {
			cntGo++
		}
		if strings.Contains(lowLine, "idiom") {
			cntIdiom++
		}
	}

	return
}
