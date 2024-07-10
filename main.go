package main

import (
	"flag"
	"fmt"
	"io"
	"os"
	"sort"
	"strings"

	"github.com/fatih/color"
)

func main() {
	var maxwidth int
	var filetype string
	flag.StringVar(&filetype, "filetype", "*", "specify the filetype to char count on")
	flag.IntVar(&maxwidth, "width", 80, "maximum width of bar chart")

	flag.Parse()

	if flag.NArg() < 1 {
		fmt.Println("Error: missing dirpath argument")
		printUsage()
		os.Exit(1)
	}

	dirPath := flag.Arg(0)

	cf := make(CharFreq)

	cf.ProcessDir(dirPath, filetype)

	cf.orderPrint(maxwidth)
}

func (cf CharFreq) ProcessDir(path string, filetype string) error {

	dir, err := os.Open(path)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	defer dir.Close()

	fileInfos, err := dir.ReadDir(-1)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	for _, fileInfo := range fileInfos {
		path := fmt.Sprintf("%s/%s", path, fileInfo.Name())
		if fileInfo.IsDir() {
			err := cf.ProcessDir(path, filetype)
			if err != nil {
				return err
			}
			continue
		}

		ft := fileType(fileInfo.Name())

		if filetype == "*" || ft == filetype {
			err := cf.readContent(path)
			if err != nil {
				return err
			}
		}
	}
	return nil
}

func fileType(name string) string {
	split := strings.Split(name, ".")
	if len(split) == 1 {
		return ""
	}
	return split[len(split)-1]

}

type CharFreq map[rune]int

func (cf CharFreq) orderPrint(maxwidth int) {
	type kv struct {
		key rune
		val int
	}
	list := []kv{}
	for k, v := range cf {
		if v > 0 && k != ' ' && k != '\n' && k != '\t' {
			list = append(list, kv{k, v})
		}
	}

	sort.Slice(list, func(i, j int) bool {
		return list[i].val > list[j].val
	})

	red := color.New(color.BgRed)
	blue := color.New(color.BgBlue)
	green := color.New(color.BgGreen)
	cyan := color.New(color.BgWhite)

	var scaleDown int

	for i, freq := range list {

		if i == 0 {
			scaleDown = freq.val / maxwidth
		}

		fmt.Printf("%c %7d ", freq.key, freq.val)
		bar := fmt.Sprintf("%s", strings.Repeat(" ", freq.val/scaleDown))
		switch {
		case freq.key >= '0' && freq.key <= '9':
			// numerical
			red.Print(bar)
		case (freq.key >= 'A' && freq.key <= 'Z') || (freq.key >= 'a' && freq.key <= 'z'):
			// alphabet
			blue.Print(bar)
		case freq.key >= 33 && freq.key <= 47:
			// symbol
			green.Print(bar)
		case freq.key >= 58 && freq.key <= 64:
			// punctuation
			cyan.Print(bar)
		case freq.key >= 91 && freq.key <= 96:
			// symbols
			green.Print(bar)
		case freq.key >= 123 && freq.key <= 126:
			// puntuation
			cyan.Print(bar)
		default:
		}
		fmt.Println("")

	}
}

func (cr CharFreq) readContent(filepath string) error {

	file, err := os.Open(filepath)
	if err != nil {
		return err
	}
	defer file.Close()

	var buf [1]byte

	for {
		n, err := file.Read(buf[:])
		if err != nil {
			if err == io.EOF {
				break
			}
			return err
		}

		if n > 0 {
			cr[rune(buf[0])]++
		}
	}

	return nil
}

func printUsage() {
	fmt.Println("Usage: charcnt dirpath")
}
