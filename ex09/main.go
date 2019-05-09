package main

import (
	"bufio"
	"fmt"
	"strings"
)

func main() {

	s := "I felt so good like anything was possible\n I hit cruise control and rubbed my eyes\n The last three days\n the rain was un-stoppable\n It was always cold\n no sunshine\n Yeah runnin' down a dream\n That never would come to me\n Workin' on a mystery\n goin' wherever it leads\n Runnin' down a dream"

	scanner1 := bufio.NewScanner(strings.NewReader(s))
	scanner2 := bufio.NewScanner(strings.NewReader(s))
	scanner3 := bufio.NewScanner(strings.NewReader(s))

	for scanner1.Scan() {
		line := scanner1.Text()
		fmt.Println(line)
	}

	scanner2.Split(bufio.ScanWords) // by words

	for scanner2.Scan() {
		fmt.Printf("%s\n", scanner2.Text())
	}

	scanner3.Split(bufio.ScanRunes) // by char

	for scanner3.Scan() {
		fmt.Printf("%s\n", scanner3.Text())
	}

}
