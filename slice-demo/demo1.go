package main

import (
	"fmt"
	"strconv"
	"time"
)

type student struct {
	a1 string
	a2 string
	a3 string
	a4 string
	a5 string
	a6 string
	a7 string
	a8 string
	a9 string
	a10 string
	a11 string
	a12 string
	a13 string
	a14 string
	a15 string
	a16 string
	a17 string
	a18 string
	a19 string
	a20 string
}

func main() {
	var ss1 []*student
	var ss2 []student

	for i := 0; i < 100000; i++ {
		s := student{
			a1: "aaaaaaaaaaaaaaaa" + strconv.Itoa(i),
			a2: "aaaaaaaaaaaaaaaa" + strconv.Itoa(i),
			a3: "aaaaaaaaaaaaaaaa" + strconv.Itoa(i),
			a4: "aaaaaaaaaaaaaaaa" + strconv.Itoa(i),
			a5: "aaaaaaaaaaaaaaaa" + strconv.Itoa(i),
			a6: "aaaaaaaaaaaaaaaa" + strconv.Itoa(i),
			a7: "aaaaaaaaaaaaaaaa" + strconv.Itoa(i),
			a8: "aaaaaaaaaaaaaaaa" + strconv.Itoa(i),
			a9: "aaaaaaaaaaaaaaaa" + strconv.Itoa(i),
			a10: "aaaaaaaaaaaaaaaa" + strconv.Itoa(i),
			a11: "aaaaaaaaaaaaaaaa" + strconv.Itoa(i),
			a12: "aaaaaaaaaaaaaaaa" + strconv.Itoa(i),
			a13: "aaaaaaaaaaaaaaaa" + strconv.Itoa(i),
			a14: "aaaaaaaaaaaaaaaa" + strconv.Itoa(i),
			a15: "aaaaaaaaaaaaaaaa" + strconv.Itoa(i),
			a16: "aaaaaaaaaaaaaaaa" + strconv.Itoa(i),
			a17: "aaaaaaaaaaaaaaaa" + strconv.Itoa(i),
			a18: "aaaaaaaaaaaaaaaa" + strconv.Itoa(i),
			a19: "aaaaaaaaaaaaaaaa" + strconv.Itoa(i),
			a20: "aaaaaaaaaaaaaaaa" + strconv.Itoa(i),
		}
		ss1 = append(ss1, &s)
		ss2 = append(ss2, s)
	}

	res1 := 0
	res2 := 0
	for i := 0; i < 1000; i++ {
		start1 := time.Now()
		for _, s := range ss1 {
			s.a1 = "bbb"
		}
		t1 := time.Since(start1)

		start2 := time.Now()
		for _, s := range ss2 {
			s.a2 = "bbb"
		}
		t2 := time.Since(start2)

		if t1 == t2 {
			continue
		}

		if t1 > t2 {
			res2 ++
		} else {
			res1 ++
		}
	}

	fmt.Println(res1)
	fmt.Println(res2)
}

