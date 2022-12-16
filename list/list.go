package list

import "fmt"

type List struct {
	Len       int64
	FirstNode *Node
}

type Node struct {
	Data     int64
	NextNode *Node
}

func NewList() (l *List) {
	l = &List{
		Len:       0,
		FirstNode: nil,
	}
	return
}

func (l *List) Add(data int64) (index int64) {
	if l.FirstNode == nil {
		n := &Node{}
		n.Data = data
		l.FirstNode = n
		l.Len++
		return l.Len - 1
	}
	cn := l.FirstNode
	for {
		if cn.NextNode == nil {
			break
		}
		cn = cn.NextNode
	}
	n := &Node{}
	n.Data = data
	cn.NextNode = n
	l.Len++
	return l.Len - 1
}

func (l *List) Print() {
	if l.FirstNode == nil {
		fmt.Println("Empty list")
		return
	}
	cn := l.FirstNode
	for {
		if cn.NextNode == nil {
			fmt.Println(cn.Data)
			break
		}
		fmt.Println(cn.Data)
		cn = cn.NextNode
	}
}

func (l *List) Delete(index int64) (ok bool) {
	if index > l.Len-1 {
		return false
	}

	if index == 0 {
		l.FirstNode = l.FirstNode.NextNode
		l.Len--
		return true
	}

	var lastNode *Node
	currentNode := l.FirstNode
	for n := int64(0); n < index; n++ {
		lastNode = currentNode
		currentNode = lastNode.NextNode
	}
	nextNode := currentNode.NextNode
	lastNode.NextNode = nextNode
	l.Len--
	return true
}

func (l *List) SortInc() {
	if l.FirstNode == nil {
		fmt.Println("Empty list")
		return
	}

	for i := int64(0); i < l.Len-1; i++ {
		currentNode := l.FirstNode
		for {
			if currentNode.NextNode == nil {
				break
			}
			if currentNode.Data > currentNode.NextNode.Data {
				currentNode.NextNode.Data, currentNode.Data = currentNode.Data, currentNode.NextNode.Data
			}
			currentNode = currentNode.NextNode
		}
	}
}

func (l *List) SortIncLink() {
	if l.FirstNode == nil {
		fmt.Println("Empty list")
		return
	}

	for i := int64(0); i < l.Len-1; i++ {
		if l.FirstNode.Data > l.FirstNode.NextNode.Data {
			n := l.FirstNode
			l.FirstNode = l.FirstNode.NextNode
			n.NextNode = l.FirstNode.NextNode
			l.FirstNode.NextNode = n
		}
		lastNode := l.FirstNode
		currentNode := l.FirstNode.NextNode
		for {
			if currentNode.NextNode == nil {
				break
			}
			if currentNode.Data > currentNode.NextNode.Data {
				n := currentNode
				lastNode.NextNode = currentNode.NextNode
				n.NextNode = currentNode.NextNode.NextNode
				lastNode.NextNode.NextNode = n
				currentNode = lastNode.NextNode
			}
			lastNode = lastNode.NextNode
			currentNode = currentNode.NextNode
		}
	}
}

func (l *List) SortDec() {
	if l.FirstNode == nil {
		fmt.Println("Empty list")
		return
	}

	for i := int64(0); i < l.Len-1; i++ {
		currentNode := l.FirstNode
		for {
			if currentNode.NextNode == nil {
				break
			}
			if currentNode.Data < currentNode.NextNode.Data {
				currentNode.NextNode.Data, currentNode.Data = currentNode.Data, currentNode.NextNode.Data
			}
			currentNode = currentNode.NextNode
		}
	}
}

func (l *List) SortDecLink() {
	if l.FirstNode == nil {
		fmt.Println("Empty list")
		return
	}

	for i := int64(0); i < l.Len-1; i++ {
		if l.FirstNode.Data < l.FirstNode.NextNode.Data {
			n := l.FirstNode
			l.FirstNode = l.FirstNode.NextNode
			n.NextNode = l.FirstNode.NextNode
			l.FirstNode.NextNode = n
		}
		lastNode := l.FirstNode
		currentNode := l.FirstNode.NextNode
		for {
			if currentNode.NextNode == nil {
				break
			}
			if currentNode.Data < currentNode.NextNode.Data {
				n := currentNode
				lastNode.NextNode = currentNode.NextNode
				n.NextNode = currentNode.NextNode.NextNode
				lastNode.NextNode.NextNode = n
				currentNode = lastNode.NextNode
			}
			lastNode = lastNode.NextNode
			currentNode = currentNode.NextNode
		}
	}
}
