package ds

import "fmt"

type IQueue interface {
	EnQueue(int)
	DeQueue()
	Display()
}

type Queue []int

func (q *Queue) EnQueue(i int) {
	*q = append(*q, i)
}

func (q *Queue) DeQueue() {
	*q = (*q)[1:len(*q)]
}

func (q *Queue) Display() {
	for _, v := range *q {
		fmt.Println(v)
	}
}
