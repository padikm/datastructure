package ds

import "fmt"

type Istack interface {
	Push(i int)
	Pop()
	Display()
}

type Stack []int

func (s *Stack) Push(i int) {
	*s = append(*s, i)
}

func (s *Stack) Pop() {
	*s = (*s)[0 : len(*s)-1]
}

func (s *Stack) Display() {
	for _, v := range *s {
		fmt.Println(v)
	}
}
