/**
 * @Author:    iccy
 * @Mail:      iccy.fun@outlook.com
 * @Data:      2020/12/3 14:26
 */

package main

import (
	"errors"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
	"time"
)

func main() {
	var err error

	if len(os.Args) != 2 || os.Args[1] == "-h" || os.Args[1] == "--help" {
		fmt.Println("Please enter a date in this format -- \033[31myyyy.mm.dd\033[0m\n$ Zeller 1949.10.1")
		return
	}

	log.SetFlags(log.LstdFlags ^ log.Ldate ^ log.Ltime)

	if _, err = time.Parse("2006.1.2", os.Args[1]); err != nil {
		log.Fatalln(err)
	}

	var numArray []int
	if numArray, err = dateStrToNum(os.Args[1]); err != nil {
		log.Fatalln(err)
	}
	_ = numArray
}

/*
 * @Description: Cut the date string and turn it into a number
 * @Param: string
 * @Return: []int, error
 * @Author: iccy
 * @Date: 2020/12/3
 */
func dateStrToNum(date string) (num []int, err error) {
	strArray := strings.Split(date, ".")
	if len(strArray) != 3 {
		return []int{}, errors.New("incorrect parameter format")
	}

	num = make([]int, 3)
	for k, _ := range strArray {
		num[k], err = strconv.Atoi(strArray[k])
		check(&err)
	}
	if err != nil {
		return []int{}, errors.New("parameter error")
	}

	return num, nil
}

func check(err *error) {
	if *err != nil {
		log.Println(*err)
	}
}