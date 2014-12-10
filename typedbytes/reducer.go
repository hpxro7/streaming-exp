// Mapper for the customary typedbytes/SequenceFile experiments
package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		var key string
		var value string
		_, err := fmt.Sscanf(scanner.Text(), "%s\t%s", &key, &value)
		if err != nil {
			fmt.Fprintln(os.Stderr, "Could not scan data: ", err)
		}
	}
	if err := scanner.Err(); err != nil {
		fmt.Fprintln(os.Stderr, "Could not read data: ", err)
	}

}
