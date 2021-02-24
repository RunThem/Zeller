/**
 * Copyright 2021 RunThem. All Rights Reserved.
 *
 * Distributed under MIT license.
 *
 * See file LICENSE for detail or copy at https://opensource.org/licenses/MIT
 *
 * Author: RunThem
 * Email: iccy.fun@outlook.com
 * Created at 2020/12/3 14:26
 */

package main

import (
	"fmt"
	"log"
	"os"
)

var DayNames = []string{
	"Sunday",
	"Monday",
	"Tuesday",
	"Wednesday",
	"Thursday",
	"Friday",
	"Saturday",
}

func init() {
	log.SetFlags(log.LstdFlags ^ log.Ldate ^ log.Ltime)
}

func main() {
	date := CommandLine(os.Args[1:])
	if date == "" {
		return
	}

	day, err := Zeller(date)
	check(&err, log.Fatalln)

	fmt.Printf("%s it's %s\n", date, DayNames[day])
}

func check(err *error, f func(...interface{})) {
	if *err != nil {
		f(*err)
	}
}
