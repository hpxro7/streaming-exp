// Reducer for the customary Word Count Map/Reduce pipeline.
package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	counts := make(map[string]int)
	for scanner.Scan() {
		var w string
		var cnt int
		_, err := fmt.Sscanf(scanner.Text(), "%s\t%d", &w, &cnt)
		if err != nil {
			fmt.Fprintln(os.Stderr, "Could not scan data: ", err)
		}
		counts[w] += cnt
	}
	if err := scanner.Err(); err != nil {
		fmt.Fprintln(os.Stderr, "Could not read data: ", err)
	}

	for w, cnt := range counts {
		fmt.Println(w, cnt)
	}
}
