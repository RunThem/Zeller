/**
 * Copyright 2021 RunThem. All Rights Reserved.
 *
 * Distributed under MIT license.
 *
 * See file LICENSE for detail or copy at https://opensource.org/licenses/MIT
 *
 * Author: RunThem
 * Email: iccy.fun@outlook.com
 * Created at 2021/2/24 19:38
 */

package main

import (
	"testing"
)

var example = []struct {
	date    string
	weekday int
}{
	{"1582.10.4", 1},
	{"1582.10.15", 5},
	{"1949.10.1", 6},
	{"2000.4.15", 6},
	{"2000.09.16", 6},
	{"2020.12.3", 4},
	{"2112.9.3", 6},
}

func Test_Zeller(t *testing.T) {
	for _, v := range example {
		if weekDay, _ := Zeller(v.date); weekDay != v.weekday {
			t.Log(v.date, weekDay, " != ", v.weekday)
		}
	}
}

func Benchmark_Zeller(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Zeller(example[4].date)
	}
}
