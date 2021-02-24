/**
 * Copyright 2021 RunThem. All Rights Reserved.
 *
 * Distributed under MIT license.
 *
 * See file LICENSE for detail or copy at https://opensource.org/licenses/MIT
 *
 * Author: RunThem
 * Email: iccy.fun@outlook.com
 * Created at 2021/2/23 22:29
 */

package main

import (
	"flag"
	"fmt"
)

var Version = "v1.5"

/*
 * @Description: Displays help and version information
 * @Param: void
 * @Return: string
 * @Author: RunThem
 * @Date: 2021/2/24
 */
func commandLine() string {
	help := flag.Bool("h", false, "show help")
	version := flag.Bool("v", false, "show version")

	flag.Parse()

	if *help {
		fmt.Println("Please enter a date in this format -- \033[31myyyy.mm.dd\033[0m\n$ Zeller 1949.10.1")
	}

	if *version {
		fmt.Printf("Zeller %s - Calculate the day of the week\n"+
			"Copyright (c) RunThem 2020-xxxx, MIT Open Source Software.\n", Version)
	}

	return flag.Arg(0)
}
