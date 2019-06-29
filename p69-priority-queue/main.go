package main

import "fmt"

type node struct {
	value    string
	priority int
}

type priorityQueue []*node

func (pq *priorityQueue) push(newNode *node) {
	*pq = append(*pq, newNode)

	chi, par := 0, len(*pq)-1
	for par > 0 {
		chi, par = par, (par-1)/2
		if (*pq)[chi].priority < (*pq)[par].priority {
			(*pq)[chi], (*pq)[par] = (*pq)[par], (*pq)[chi]
		}
	}
}

func (pq *priorityQueue) pop() *node {
	ret := (*pq)[0]

	(*pq)[0] = (*pq)[len(*pq)-1]
	*pq = (*pq)[:len(*pq)-1]

	par := 0
	for {
		chi1 := par*2 + 1
		chi2 := par*2 + 2
		if chi1 > len(*pq)-1 {
			break
		}

		child := (*pq)[chi1]
		altChi := chi1
		if chi2 < len(*pq) &&
			(*pq)[chi2].priority < child.priority {
			child = (*pq)[chi2]
			altChi = chi2
		}
		if (*pq)[par].priority < child.priority {
			break
		}

		(*pq)[par], (*pq)[altChi] = child, (*pq)[par]
		par = altChi
	}
	return ret
}

func main() {
	pq := &priorityQueue{}
	pq.push(&node{"hoge", 5})
	pq.push(&node{"piyo", 10})
	pq.push(&node{"foo", 7})
	pq.push(&node{"bar", 4})

	for i := range *pq {
		fmt.Println((*pq)[i].value, (*pq)[i].priority)
	}

	k := len(*pq)
	for i := 0; i < k; i++ {
		n := pq.pop()
		fmt.Println(n.value, n.priority)
	}
}
