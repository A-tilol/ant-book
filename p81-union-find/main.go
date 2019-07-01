package main

import "fmt"

type unionFind struct {
	par  []int
	rank []int
}

func (u *unionFind) init(n int) {
	u.par = make([]int, n+1)
	u.rank = make([]int, n+1)
	for i := range u.par {
		u.par[i] = i
		u.rank[i] = 0
	}
}

func (u *unionFind) root(x int) int {
	if u.par[x] == x {
		return x
	}
	u.par[x] = u.root(u.par[x])
	return u.par[x]
}

func (u *unionFind) unite(x, y int) {
	xroot := u.root(x)
	yroot := u.root(y)
	if xroot == yroot {
		return
	}
	if u.rank[xroot] > u.rank[yroot] {
		u.par[yroot] = xroot
	} else {
		u.par[xroot] = yroot
		if u.rank[xroot] == u.rank[yroot] {
			u.rank[yroot]++
		}
	}
}

func (u *unionFind) same(x, y int) bool {
	if u.root(x) == u.root(y) {
		return true
	}
	return false
}

func main() {
	n := 7
	u := &unionFind{}
	u.init(n)
	u.unite(1, 2)
	u.unite(5, 1)
	u.unite(6, 4)
	u.unite(4, 7)
	fmt.Println(u.par)
	u.unite(1, 6)
	fmt.Println(u.par)
}
