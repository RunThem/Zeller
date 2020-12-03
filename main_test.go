/**
 * @Author:    iccy
 * @Mail:      iccy.fun@outlook.com
 * @Data:      2020/12/3 21:36
 */

package main

import (
	"reflect"
	"testing"
)

func Test_dateStrToNum(t *testing.T) {
	str := []struct {
		date string
		num  []int
	}{
		{"1949.10.1", []int{1949, 10, 1}},
		{"2000.4.15", []int{2000, 4, 15}},
		{"2000.09.16", []int{2000, 9, 16}},
		{"2020.12.3", []int{2020, 12, 3}},
	}

	for _, v := range str {
		if numArray, _ := dateStrToNum(v.date); !reflect.DeepEqual(numArray, v.num) {
			t.Log(numArray, " != ", v.num)
		}
	}
}