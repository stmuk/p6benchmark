package main

import (
	"fmt"
	"log"
	"os"
	"os/exec"
	"strconv"
	"time"
)

func main() {

	const DEBUG = 0
	var num int64
	var target string
	var lang string

	if len(os.Args) > 3 {
		num, _ = strconv.ParseInt(os.Args[1], 10, 64)
		lang = os.Args[2]
		target = os.Args[3]
	} else {
		log.Fatal("number of iterations, lang, target program")
	}

	var i int64 // XXX

	start := time.Now()

	for i = 0; i < num; i++ {

		if DEBUG == 1 {
			fmt.Printf("%d\n", i)
		}

		err := exec.Command(lang, target).Run()

		if err != nil {
			log.Fatal(err)
		}

	}

	elapsed := time.Since(start).Seconds()
	avg := elapsed / float64(num)
	fmt.Printf("\"%s %s\" %d run(s)  avg=%0.2f sec(s)", lang, target, num, avg)

}
