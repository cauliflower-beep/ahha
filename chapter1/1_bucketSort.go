package main

import "fmt"

type student struct {
	name  string
	grade int
	//rank  int
}

func bucketSort(s []student, n int, reverse bool) []student {
	buckets := make([][]student, n) // 分配桶资源

	for _, p := range s {
		buckets[p.grade] = append(buckets[p.grade], p)
	}
	var temp, res []student
	// 顺序输出
	for _, sp := range buckets {
		temp = append(temp, sp...)
	}

	if reverse {
		// 逆序输出
		for i := len(s) - 1; i >= 0; i-- {
			res = append(res, temp[i])
		}
		return res
	} else {
		// 顺序输出
		return temp
	}
}

func main() {
	var s1 = student{
		name:  "conan",
		grade: 5,
	}
	var s2 = student{
		name:  "goku",
		grade: 3,
	}
	var s3 = student{
		name:  "naruto",
		grade: 5,
	}
	var s4 = student{
		name:  "lufy",
		grade: 2,
	}
	var s5 = student{
		name:  "bike",
		grade: 8,
	}
	var s []student
	s = append(s, s1)
	s = append(s, s2)
	s = append(s, s3)
	s = append(s, s4)
	s = append(s, s5)

	res := bucketSort(s, 11, false)
	fmt.Println("顺序输出|", res)
	res = bucketSort(s, 11, true)
	fmt.Println("逆序输出|", res)
}
