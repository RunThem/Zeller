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
	if len(os.Args) != 2 || os.Args[1] == "-h" || os.Args[1] == "--help" {
		fmt.Println("Please enter a date in this format -- \033[31myyyy.mm.dd\033[0m\n$ Zeller 1949.10.1")
		return
	}

	log.SetFlags(log.LstdFlags ^ log.Ldate ^ log.Ltime)

	if _, err := time.Parse("2006.1.2", os.Args[1]); err != nil {
		log.Fatalln(err)
	}

	numArray, err := dateStrToNum(os.Args[1])
	if err != nil {
		log.Fatalln(err)
	}

	w := zeller(numArray)
	fmt.Println(w)
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

/*
 * @Description: Zeller formula
 * @Param: []int
 * @Return: int
 * @Author: iccy
 * @Date: 2020/12/4
 */
func zeller(nums []int) int {
	// January and February are the 13 and 14 months of the previous year, so a year begins in March
	if nums[1] > 0 && nums[1] < 3 {
		nums[1] += 12
		nums[0]--
	}

	var c, y, m, d int = int(nums[0] / 100), nums[0] % 100, nums[1], nums[2]

	// Dates before 1582-10-4 have different formulas
	if (nums[0] < 1582) || (nums[0] == 1582 && ((nums[1] < 10) || (nums[1] == 10 && nums[2] <= 4))) {
		// golang: Take the remainder of a negative number
		return ((y+int(y/4)-c+int((26*(m+1))/10)+d+4)%7 + 7) % 7
	}

	return ((y+int(y/4)+int(c/4)-2*c+int((26*(m+1))/10)+d-1)%7 + 7) % 7
}
