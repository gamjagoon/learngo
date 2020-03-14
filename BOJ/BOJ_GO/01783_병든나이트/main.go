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

type global struct {
	N,M int64
	result int64
}

var input global

func Min(a,b int64) int64{
	if a > b {
		return b
	} 
	return a
}

func main() {
	defer writer.Flush()
	scanf("%d %d\n",&input.N,&input.M)
	if input.N == 1{
		input.result = 1
	} else if input.N == 2 {
		input.result = Min(4,(input.M+1)/2)
	} else if input.N >= 3 && input.M <= 6 {
		input.result = Min(4,input.M)
	} else {
		input.result = input.M-2
	}
	printf("%d",input.result)
}