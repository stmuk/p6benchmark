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

	var num int64
	var target []string
	var lang string

	if len(os.Args) > 3 {
		num, _ = strconv.ParseInt(os.Args[1], 10, 64)
		lang = os.Args[2]
		target = os.Args[3:]
	} else {
		log.Fatal("number of iterations, lang, target program")
	}

	var i int64 // XXX

	start := time.Now()

	for i = 0; i < num; i++ {

		fmt.Printf("%d ", i)

		out, err := exec.Command(lang, target...).Output()

		if num == 1 {
			fmt.Fprintf(os.Stderr, "%s\n", out)
		}

		if err != nil {
			log.Fatal(err)
		}

	}

	elapsed := time.Since(start).Seconds()
	avg := elapsed / float64(num)
	fmt.Printf("\n\"%s %s\" %d run(s) wallclock=%0.2f avg=%0.2f sec(s)\n", lang, strings.Join(target, ""), num, float64(elapsed), avg)

}
