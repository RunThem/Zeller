/**
 * @Author:    iccy
 * @Mail:      iccy.fun@outlook.com
 * @Data:      2020/12/3 14:26
 */

package main

import (
	"fmt"
	"log"
	"os"
	"time"
)

func main() {
	if len(os.Args) != 2 || os.Args[1] == "-h" || os.Args[1] == "--help" {
		fmt.Println("Please enter a date in this format -- \033[31myyyy.mm.dd\033[0m\n$ Zeller 1949.10.1")
		return
	}

	log.SetFlags(log.LstdFlags ^ log.Ldate ^ log.Ltime)

	if _, err := time.Parse("2006.1.2", os.Args[1]); err != nil {
		log.Fatalln(err)
	}

}