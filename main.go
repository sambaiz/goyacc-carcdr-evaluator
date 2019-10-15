package main

import (
	"bufio"
	"fmt"
	"github.com/sambaiz/goyacc-carcdr-evaluator/parser"
	"io"
	"log"
	"os"
)

func main(){
	in := bufio.NewReader(os.Stdin)
	for {
		if _, err := os.Stdout.WriteString("> "); err != nil {
			log.Fatalf("WriteString: %s", err)
		}
		line, err := in.ReadBytes('\n')
		if err == io.EOF {
			return
		}
		if err != nil {
			log.Fatalf("ReadBytes: %s", err)
		}

		out, err := parser.Parse(string(line))
		if err != nil {
			fmt.Println(err.Error())
		} else {
			fmt.Println(out.String())
		}
	}
}
