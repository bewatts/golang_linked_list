package main

import "fmt"

type node struct {
	prev, next *node
	data       int
}

type list struct {
	head, tail *node
}

func (l *list) push(newNode node) {
	if l.head == nil {
		l.head = &newNode
		l.tail = &newNode
	} else {
		l.tail.next = &newNode
		newNode.prev = l.tail
		l.tail = &newNode
	}
}

func (l *list) pop() {
	fmt.Println(l.tail);
	if l.tail == nil {
		fmt.Println("Cannot pop an empty list")
	} else if l.tail.prev == nil {
		l.tail = nil;
	} else {
		l.tail = l.tail.prev
		l.tail.next = nil
	}
		fmt.Println("Popped!")
}

func (l *list) count() int {
	return recursiveCount(1, l.head)
}

func recursiveCount(count int, currentNode *node) int {
	// fmt.Println(currentNode);
	if currentNode.next == nil {
		return count
	} else {
		return recursiveCount(count+1, currentNode.next)
	}
}

func pushTest() {
	fmt.Println("*****pushTest*****")
	var nodeOne = node{data: 1}
	var nodeTwo = node{data: 2}
	var nodeThree = node{data: 3}
	var listOne = new(list)
	listOne.push(nodeOne)
	listOne.push(nodeTwo)
	listOne.push(nodeThree)

	fmt.Println("listOne.head.data is 1", listOne.head.data == 1)
	fmt.Println("listOne.head.next.data is 2", listOne.head.next.data == 2)
	fmt.Println("listOne.tail.prev.data is 2", listOne.tail.prev.data == 2)
	fmt.Println("listOne.tail.data == 3", listOne.tail.data == 3)
}

func popTest() {
	fmt.Println("*****popTest*****")
	var nodeOne = node{data: 1}
	var nodeTwo = node{data: 2}
	var nodeThree = node{data: 3}
	var nodeFour = node{data: 4}
	
	var listOne = new(list)
	listOne.push(nodeOne)
	listOne.push(nodeTwo)
	listOne.push(nodeThree)
	listOne.push(nodeFour)
	listOne.pop()
	listOne.pop()
	listOne.pop()
	listOne.pop()
}

func countTest() {
	fmt.Println("*****countTest*****")
	var nodeOne = node{data: 1}
	var nodeTwo = node{data: 2}
	var nodeThree = node{data: 3}
	var nodeFour = node{data: 4}
	var listOne = new(list)
	listOne.push(nodeOne)
	listOne.push(nodeTwo)
	listOne.push(nodeThree)
	fmt.Println("count is 3", listOne.count() == 3)
	listOne.push(nodeFour)
	fmt.Println("count is 4", listOne.count() == 4)
}

func main() {
	pushTest()
	countTest()
	popTest()
}