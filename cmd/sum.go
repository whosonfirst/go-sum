package main

import (
	"bufio"
	"flag"
	"fmt"
	"github.com/dustin/go-humanize"
	"log"
	"os"
	"strconv"
)

func main() {

	var enhumanize = flag.Bool("humanize", false, "")

	flag.Parse()

	var total uint64
	total = 0
	
	scanner := bufio.NewScanner(os.Stdin)

	for scanner.Scan() {

		txt := scanner.Text()
		i, err := strconv.Atoi(txt)

		if err != nil {
			log.Fatal(err)
		}

		total += uint64(i)
		// log.Println(total)
	}

	if *enhumanize {
		fmt.Println(humanize.Bytes(total))
	} else {
		fmt.Println(total)
	}
}
