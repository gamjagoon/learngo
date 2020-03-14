package main

import (
	"bufio"
	"fmt"
	"os"
)

var reader *bufio.Reader = bufio.NewReader(os.Stdin)
var writer *bufio.Writer = bufio.NewWriter(os.Stdout)
func printf(f string, a ...interface{}) { fmt.Fprintf(writer, f, a...) }
func scanf(f string, a ...interface{}) { fmt.Fscanf(reader, f, a...) }

var (
	str string
	yarr [50001]int
	res []int
)

func main() {
	defer writer.Flush()
	scanf("%s",&str)
	max := len(str)
	for i := 0; i < max-1; i++ {
		if str[i+1] == str[i] {
			switch str[i] {
			case '(':
				res = append(res,i)
			case ')':
				yarr[i]++
			}
		}
		yarr[i+1] = yarr[i]
	}
	result := len(res)*yarr[max-1]
	for i := 0; i < len(res); i++ {
		result -= yarr[res[i]]
	}
	printf("%d",result)
}