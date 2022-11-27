package main

import (
	"bufio"
	"flag"
	"fmt"
	"log"
	"os"
	"sort"
	"strings"
)

func main() {

	var HeaderField = flag.Bool("h", false, "The first line is a header that must be ignored during sorting but included in the output.")
	var InputFileName = flag.String("i", "", "Use a file with the name file-name as an input.")
	var OutputFileName = flag.String("o", "", "Use a file with the name file-name as an output.")
	var ReverseSorting = flag.Bool("r", false, "Sort input lines in reverse order.")
	var SortingField = flag.Int("f", 0, "Sort input lines by value number N.")
	
	flag.Parse()

	var write *bufio.Scanner
	if *InputFileName == "" {
		fmt.Println("Array:")
		write = bufio.NewScanner(os.Stdin)
	} 

	else {
		f, err := os.Open(*InputFileName)
		if err != nil {
			panic(err)
		}

		defer f.Close()
		write = bufio.NewScanner(f)

		fmt.Println("\nFile data:")
	}

	array := [][]string{}
	title := [][]string{}
	row := 0
	n := 0
	

	for write.Scan() {

		if *InputFileName != "" {
			fmt.Println(write.Text())
		}

		line := write.Text()
		splitLine := strings.Split(strings.ReplaceAll(line, " ", ""), ",")

		if n == 0 {
			n = len(splitLine)
		}

		if n != len(splitLine) && line != "" {
			log.Fatalf("Error: row has %d columns, but must has %d\n", len(splitLine), n)
		}

		if *HeaderField == true {
			if line != "" && row > 0 {
				array = append(array, splitLine)
			} 
			
			else {
				title = append(array, splitLine)
			}

		} 
		else {
			if line != "" {
				array = append(array, splitLine)
			}
		}

		if line == "" {
			break
		}
		row++
	}

	index := *SortingField

	if *ReverseSorting == false {
		sort.Slice(array, func(i, j int) bool { return array[i][index] < array[j][index] })
	} 

	else {
		sort.Slice(array, func(i, j int) bool { return array[i][index] > array[j][index] })
	}

	if *OutputFileName == "" {
		fmt.Println("\nSorted:")
		n = len(title)

		for i := 0; i < n; i++ {
			fmt.Printf("%v\n", title[i])
		}

		n = len(array)
		for i := 0; i < n; i++ {
			fmt.Printf("%v\n", array[i])
		}

	} 
	
	else {

		f, err := os.OpenFile(*OutputFileName, os.O_WRONLY|os.O_CREATE, 0666)
		if err != nil {
			fmt.Println("File doesn`t exist")
			os.Exit(1)
		}

		defer f.Close()

		o := bufio.NewWriter(f)

		fmt.Printf("\nSorted result has been writen.")

		n = len(title)
		for i := 0; i < n; i++ {
			fmt.Fprintf(o, "%v\n", title[i])
		}

		n = len(array)
		for i := 0; i < n; i++ {
			fmt.Fprintf(o, "%v\n", array[i])
		}
		
		o.Flush()
	}
}
