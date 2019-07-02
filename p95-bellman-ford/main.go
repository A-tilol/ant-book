package main

import (
	"fmt"
	"math"
)

type edge struct {
	from int
	to   int
	cost int64
}

func bellmanFord(n, m, s int, edges []edge) {
	d := make([]int64, n)
	for i := range d {
		d[i] = math.MaxInt64
	}
	d[s] = 0

	for i := 0; i < n; i++ {
		for j := 0; j < m; j++ {
			e := edges[j]
			if d[e.from] == math.MaxInt64 {
				continue
			}
			if d[e.to] > d[e.from]+e.cost {
				d[e.to] = d[e.from] + e.cost
				if i == n-1 {
					fmt.Println("negative loop exists!")
					return
				}
			}
		}
	}
	fmt.Println(d)
}

func main() {

	bellmanFord(3, 3, 0, []edge{
		edge{0, 1, 4},
		edge{1, 2, 3},
		edge{0, 2, 5},
	})

	bellmanFord(2, 2, 0, []edge{
		edge{0, 1, -2},
		edge{1, 0, -1},
	})

}
