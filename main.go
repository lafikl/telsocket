package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	r := bufio.NewReader(os.Stdin)
	for {
		line, err := r.ReadString('\n')
		if err != nil && err.Error() != "unexpected newline" {
			fmt.Println(err)
			break
		}
		fmt.Println(line)
	}
}

