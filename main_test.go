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

var str = []struct {
	date    string
	num     []int
	weekday int
}{
	{"1582.10.4", []int{1582, 10, 4}, 1},
	{"1582.10.15", []int{1582, 10, 15}, 5},
	{"1949.10.1", []int{1949, 10, 1}, 6},
	{"2000.4.15", []int{2000, 4, 15}, 6},
	{"2000.09.16", []int{2000, 9, 16}, 6},
	{"2020.12.3", []int{2020, 12, 3}, 4},
	{"2112.9.3", []int{2112, 9, 3}, 6},
}

func Test_dateStrToNum(t *testing.T) {
	for _, v := range str {
		if numArray, _ := dateStrToNum(v.date); !reflect.DeepEqual(numArray, v.num) {
			t.Log(numArray, " != ", v.num)
		}
	}
}

func Test_zeller(t *testing.T) {
	for _, v := range str {
		if weekDay := zeller(v.num); weekDay != v.weekday {
			t.Log(weekDay, " != ", v.weekday)
		}
	}
}

func Benchmark_zeller(b *testing.B) {
	for i := 0; i < b.N; i++ {
		zeller(str[4].num)
	}
}
