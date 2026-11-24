package main

import "fmt"

// GenerateParens returns every well-formed string of n bracket pairs.
// Backtracking with two counters: an open bracket is legal while opens
// remain, a close while it wouldn't overtake the opens.
func GenerateParens(n int) []string {
	out := []string{}
	buf := make([]byte, 0, 2*n)
	var build func(open, close int)
	build = func(open, close int) {
		if len(buf) == 2*n {
			out = append(out, string(buf))
			return
		}
		if open < n {
			buf = append(buf, '(')
			build(open+1, close)
			buf = buf[:len(buf)-1]
		}
		if close < open {
			buf = append(buf, ')')
			build(open, close+1)
			buf = buf[:len(buf)-1]
		}
	}
	build(0, 0)
	return out
}

func main() {
	fmt.Println(GenerateParens(3))
}
