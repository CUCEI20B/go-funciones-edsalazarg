package main

import "fmt"

func Burbuja(s []int64)  {
	for i := len(s); i > 0; i--{
		for j := 1; j < i;j++ {
			if s[j-1] > s[j] {
				intermediate := s[j]
				s[j] = s[j-1]
				s[j-1] = intermediate
			}
		}
	}
}

func main()  {
	s := []int64{9, 8, 7, 6, 5, 4, 3, 2, 1, 0, -1, -2}
	Burbuja(s)
	for i := 0; i < len(s); i++{
		fmt.Println(s[i])
	}
}