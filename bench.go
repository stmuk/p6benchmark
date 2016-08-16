package main

import (
	"fmt"
	"log"
	"os"
	"os/exec"
	"strconv"
	"strings"
	"time"
)

func main() {

	if len(os.Args) < 3 {
		Usage()
	}

	repeatMax, err := strconv.Atoi(os.Args[1])
	if err != nil {
		Usage()
	}

	program := os.Args[2]
	opts := os.Args[3:]

	start := time.Now()

	for i := 0; i < repeatMax; i++ {

		fmt.Printf("%d ", i)

		out, err := exec.Command(program, opts...).Output()
		if err != nil {
			log.Fatal(err)
		}

		if repeatMax == 1 {
			fmt.Fprintf(os.Stderr, "%s\n", out)
		}

	}

	elapsed := time.Since(start).Seconds()
	avg := elapsed / float64(repeatMax)
	fmt.Printf("\n\"%s %s\" %d run(s) total=%0.2f avg=%0.2f sec(s)\n", program, strings.Join(opts, " "), repeatMax, float64(elapsed), avg)

}

func Usage() {
	log.Fatalf("Usage: %s <# iterations> <program path>\n", os.Args[0])
}
