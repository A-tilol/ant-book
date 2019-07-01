package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strconv"
)

var (
	readString func() string
	readBytes  func() []byte
)

func init() {
	readString, readBytes = newReadString(os.Stdin)
}

func newReadString(ior io.Reader) (func() string, func() []byte) {
	r := bufio.NewScanner(ior)
	r.Buffer(make([]byte, 1024), int(1e+11))
	r.Split(bufio.ScanWords)

	f1 := func() string {
		if !r.Scan() {
			panic("Scan failed")
		}
		return r.Text()
	}
	f2 := func() []byte {
		if !r.Scan() {
			panic("Scan failed")
		}
		return r.Bytes()
	}
	return f1, f2
}

func readInt() int {
	return int(readInt64())
}

func readInt64() int64 {
	i, err := strconv.ParseInt(readString(), 10, 64)
	if err != nil {
		panic(err.Error())
	}
	return i
}

func readFloat64() float64 {
	f, err := strconv.ParseFloat(readString(), 64)
	if err != nil {
		panic(err.Error())
	}
	return f
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func minInt64(a, b int64) int64 {
	if a < b {
		return a
	}
	return b
}

func maxInt64(a, b int64) int64 {
	if a > b {
		return a
	}
	return b
}

func abs(a int) int {
	if a < 0 {
		return -a
	}
	return a
}

func absInt64(a int64) int64 {
	if a < 0 {
		return -a
	}
	return a
}

func sum(a []int) int {
	var ret int
	for i := range a {
		ret += a[i]
	}
	return ret
}

func sumInt64(a []int64) int64 {
	var ret int64
	for i := range a {
		ret += a[i]
	}
	return ret
}

func sumFloat64(a []float64) float64 {
	var ret float64
	for i := range a {
		ret += a[i]
	}
	return ret
}

func gcdInt64(m, n int64) int64 {
	for m%n != 0 {
		m, n = n, m%n
	}
	return n
}

func lcmInt64(m, n int64) int64 {
	return m / gcdInt64(m, n) * n
}

// sort ------------------------------------------------------------

type int64Array []int64

func (s int64Array) Len() int           { return len(s) }
func (s int64Array) Swap(i, j int)      { s[i], s[j] = s[j], s[i] }
func (s int64Array) Less(i, j int) bool { return s[i] < s[j] }

type xxx struct {
	x int
}

type sortArray []xxx

func (s sortArray) Len() int           { return len(s) }
func (s sortArray) Swap(i, j int)      { s[i], s[j] = s[j], s[i] }
func (s sortArray) Less(i, j int) bool { return s[i].x < s[j].x }

// -----------------------------------------------------------------
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
	n := readInt()
	k := readInt()

	diff := make([][]int, n+1)
	for i := range diff {
		diff[i] = make([]int, n+1)
		for j := range diff[i] {
			diff[i][j] = -1
		}
	}
	u := &unionFind{}
	u.init(n)

	var ans int
	for i := 0; i < k; i++ {
		t := readInt()
		x := readInt()
		y := readInt()
		if x < 1 || x > n || y < 1 || y > n {
			ans++
			continue
		}
		if t == 1 {
			if diff[x][y] == 1 {
				ans++
			} else {
				u.unite(x, y)
			}
		}
		if t == 2 {
			if u.root(x) == u.root(y) {
				ans++
			}
		} else {
			diff[x][y] = 1
			diff[y][x] = 1
		}
	}

	fmt.Println(ans)
}
