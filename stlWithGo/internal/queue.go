package internal

import "fmt"

type Queue struct {
	Elements []int
}

func NewQueue() *Queue {
	return &Queue{
		Elements: make([]int, 0),
	}
}

func (q *Queue) Enqueue(element int) {
	q.Elements = append(q.Elements, element)
}

func (q *Queue) Dequeue() int {
	if len(q.Elements) == 0 {
		return 0
	}
	element := q.Elements[0]
	q.Elements = q.Elements[1:]
	return element
}

func (q *Queue) Front() int {
	if len(q.Elements) == 0 {
		return 0
	}
	return q.Elements[0]
}

func (q *Queue) IsEmpty() bool {
	return len(q.Elements) == 0
}

func (q *Queue) Size() int {

	return len(q.Elements)
}

func (q *Queue) Print() {
	for _, element := range q.Elements {
		fmt.Println(element)
	}
}
