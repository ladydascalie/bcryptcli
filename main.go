package main

import (
	"bufio"
	"flag"
	"fmt"
	"log"
	"os"

	"golang.org/x/crypto/bcrypt"
)

var cost int

func main() {
	log.SetFlags(log.Lshortfile | log.LstdFlags)

	flag.IntVar(&cost, "cost", 0, "bcryptcli -cost <cost: value between 4 and 31>")
	flag.Parse()

	if flag.NFlag() == 0 {
		cost = bcrypt.DefaultCost
	}

	if cost < bcrypt.MinCost || cost > bcrypt.MaxCost {
		log.Println("invalid cost value provided, using default:", bcrypt.DefaultCost)
	}

	var (
		password []byte
		err      error
	)

	r := bufio.NewReader(os.Stdin)
	password, err = r.ReadBytes('\n')
	if err != nil {
		log.Fatal(err)
	}

	hash, err := bcrypt.GenerateFromPassword(password, cost)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(string(hash))
}
