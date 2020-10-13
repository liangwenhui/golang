package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	scanner :=bufio.NewScanner(os.Stdin)
	//阻塞
	for scanner.Scan() {
		text := scanner.Text()
		fmt.Println(strings.ToUpper(text))
	}


}
