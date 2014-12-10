// typebiter writes a series of strings to a file in a typed byte format
package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Split(bufio.ScanWords)
	for scanner.Scan() {
		toTypedPair(scanner.Text())
	}

	if scanner.Err() != nil {
		fmt.Errorf("Could not scan from std in.\n")
	}
}

func toTypedPair(s string) {

}
