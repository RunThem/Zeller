/**
 * Copyright 2021 RunThem. All Rights Reserved.
 *
 * Distributed under MIT license.
 *
 * See file LICENSE for detail or copy at https://opensource.org/licenses/MIT
 *
 * Author: RunThem
 * Email: iccy.fun@outlook.com
 * Created at 2021/2/24 17:12
 */

package main

import (
	"log"
	"os"
	"strconv"
	"strings"
	"time"
)

/*
 * @Description: Cut the date string and turn it into a number
 * @Param: string
 * @Return: []int
 * @Author: RunThem
 * @Date: 2020/12/3
 */
func strToNum(date string) []int {
	if date == "" {
		os.Exit(0)
	}

	_, err := time.Parse("2006.1.2", date)
	check(&err, log.Fatalln)

	strArray := strings.Split(date, ".")

	numArray := make([]int, 3)
	for i := range strArray {
		numArray[i], _ = strconv.Atoi(strArray[i])
	}

	return numArray
}

/*
 * @Description: Zeller formula
 * @Param: []int
 * @Return: int
 * @Author: RunThem
 * @Date: 2020/12/4
 */
func zeller(nums []int) int {
	// January and February are the 13 and 14 months of the previous year, so a year begins in March
	if nums[1] > 0 && nums[1] < 3 {
		nums[1] += 12
		nums[0]--
	}

	var c, y, m, d = nums[0] / 100, nums[0] % 100, nums[1], nums[2]

	// TODO: Debug -> Error calculating day of the week for dates 1582-10-4 and before
	// Dates before 1582-10-4 have different formulas
	if (nums[0] < 1582) || (nums[0] == 1582 && ((nums[1] < 10) || (nums[1] == 10 && nums[2] <= 4))) {
		// golang: Take the remainder of a negative number
		return ((y+y/4-c+(26*(m+1))/10+d+4)%7 + 7) % 7
	}

	return ((y+y/4+c/4-2*c+(26*(m+1))/10+d-1)%7 + 7) % 7
}
