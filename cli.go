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

var Version = "v2.5"

/*
 * @Description: Displays help and version information
 * @Param: [[]string]
 * @Return: [string]
 * @Author: RunThem
 * @Date: 2021/2/24
 */
func CommandLine(arguments []string) string {
	cli := flag.NewFlagSet("Zeller", flag.ExitOnError)
	help := cli.Bool("h", false, "show help")
	version := cli.Bool("v", false, "show version")

	cli.Parse(arguments)

	isTrue(*help, "Please enter a date in this format -- \033[31myyyy.mm.dd\033[0m\n$ Zeller 1949.10.1\n")
	isTrue(*version, fmt.Sprintf("Zeller %s - Calculate the day of the week\n"+
		"Copyright (c) RunThem 2020-xxxx, MIT Open Source Software.\n", Version))

	return cli.Arg(0)
}

func isTrue(expr bool, msg string) {
	if expr {
		fmt.Print(msg)
	}
}
