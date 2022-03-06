package main

import (
	filesystem "alphabet-sort/FileSystem"
	"alphabet-sort/types"
	"log"
)

var (
	input  = "./list.txt"
	output = "./sorted.txt"
)

func main() {
	fs := filesystem.NewFS(input, output)
	list, err := fs.ReadFile()
	if err != nil {
		log.Fatalln(err)
	}

	sortedList := types.StrArray(list).SortAlphabeticalAsc()
	if err := fs.WriteFile(sortedList); err != nil {
		log.Fatalln(err)
	}
}
