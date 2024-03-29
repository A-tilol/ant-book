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

func dfs(i, kk int) bool {
	if kk == k {
		return true
	}
	if kk > k || i == n {
		return false
	}
	if dfs(i+1, kk+a[i]) || dfs(i+1, kk) {
		return true
	}
	return false
}

var n, k int
var a []int

func main() {
	n = readInt()
	k = readInt()
	a = make([]int, n)
	for i := range a {
		a[i] = readInt()
	}
	if dfs(0, 0) {
		fmt.Println("Yes")
	} else {
		fmt.Println("No")
	}
	fmt.Println()
}
