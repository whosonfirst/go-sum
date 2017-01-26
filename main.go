package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
)

func main() {

	total := 0
	scanner := bufio.NewScanner(os.Stdin)

	for scanner.Scan() {

		txt := scanner.Text()
		i, err := strconv.Atoi(txt)

		if err != nil {
			log.Fatal(err)
		}

		total += i
		// log.Println(total)
	}

	fmt.Println(total)
}
