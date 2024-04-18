package main

import (
	"fmt"
	"os"
	"strings"

	"github.com/Lutz-Pfannenschmidt/stunden-berechner/internal/csv"
	"github.com/Lutz-Pfannenschmidt/stunden-berechner/internal/date"
	"github.com/Lutz-Pfannenschmidt/stunden-berechner/internal/parser"
)

func main() {
	if len(os.Args) < 3 {
		fmt.Println("Usage:   " + os.Args[0] + " <pivot> <file>")
		fmt.Println("Example: " + os.Args[0] + " 01.01. file.csv")
		fmt.Println("Example: " + os.Args[0] + " 28.01. file.xlsx")
		fmt.Println("Supported file formats: XLAM / XLSM / XLSX / XLTM / XLTX / CSV")
		os.Exit(1)
	}

	pivot, err := date.ParseDate(os.Args[1])
	if err != nil {
		fmt.Println("Invalid date format")
		os.Exit(1)
	}

	filename := os.Args[2]
	file, err := os.OpenFile(filename, os.O_RDONLY, 0666)
	if err != nil || file.Close() != nil {
		fmt.Println("File not found or not readable")
		os.Exit(1)
	}

	pack, err := parser.ParseFile(filename, *pivot)
	if err != nil {
		panic(err)
	}

	split := strings.Split(filename, ".")
	outname := strings.Join(split[:len(split)-1], ".") + "_output.csv"

	err = csv.WriteToFile(outname, pack)
	if err != nil {
		panic(err)
	}

	fmt.Printf("Output written to %s with pivot %s\n", outname, pivot.String())
}
