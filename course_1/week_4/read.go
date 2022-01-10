package main

import (
	"fmt"
	"bufio"
	"os"
	"strings"
)

type Name struct {
    fname string
    lname string
}

func main() {
	var fileName string
	sli := make([]Name, 0)

	fmt.Println("Enter a file name:")
	fmt.Scan(&fileName)

	file, err := os.Open(fileName)
    if err != nil {
        fmt.Println(err)
    }

	scanner := bufio.NewScanner(file)
    for scanner.Scan() {
		namearr := strings.Split(scanner.Text(), " ")
		name := Name{fname: namearr[0], lname: namearr[1]}
		sli = append(sli, name)
    }

	file.Close()

	for _, n := range sli {
		fmt.Println(n.fname, n.lname)
	}

}

