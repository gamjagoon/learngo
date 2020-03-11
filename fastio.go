/**
 * This snippet assumes large input and large output and that all the input contains numbers in the form of string
 */

package main

import (
	"bufio"
	"io/ioutil"
	"os"
	"strconv"
)

var bytes []byte
var l, max int

/**
 * Perform Quick Scan of 1 data(with no spaces in-between) on the buffered input
 */
func fastScan() int {
	b := bytes[l]

	// Check if the bytes is not a number
	for b < 48 || b > 57 {
		l++
		b = bytes[l]
	}

	result := 0
	for 47 < b && b < 58 {
		result = (result << 3) + (result << 1) + int(b-48)

		l++
		if l > max {
			return result
		}
		b = bytes[l]
	}
	return result
}

func main() {
	// Read all the input into Buffer
	bytes, _ = ioutil.ReadAll(bufio.NewReader(os.Stdin))
	max = len(bytes) - 1

	// Create a write buffer directed to STDOUT
	ws := bufio.NewWriter(os.Stdout)

	// Read the first param
	t := fastScan()

	for t > 0 {
		t--

		ws.WriteString(strconv.Itoa(t) + "\n")
	}
	ws.Flush()
}
