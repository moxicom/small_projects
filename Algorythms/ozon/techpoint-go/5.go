package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"os"
	"strings"
)

type Folder struct {
	Dir    string   `json:"dir"`
	Files  []string `json:"files"`
	Folers []Folder `json:"folders"`
}

var Reset = "\033[0m"
var Red = "\033[31m"

func virusesAmount(f Folder) int {
	filesAmount := len(f.Files)
	flag := false
	for i := 0; i < filesAmount; i++ {
		if strings.HasSuffix(f.Files[i], ".hack") {
			flag = true
			break
		}
	}
	if flag {
		return filesAmount
	}
	return 0
}

func DFS(f Folder, hasCursedFiles bool) int {
	cursedFiles := 0

	if hasCursedFiles {
		cursedFiles = len(f.Files)
	} else {
		cursedFiles = virusesAmount(f)
		if cursedFiles > 0 {
			hasCursedFiles = true
		}
	}

	for _, subFolder := range f.Folers {
		cursedFiles += DFS(subFolder, hasCursedFiles)
	}

	return cursedFiles
}

func main() {
	var in *bufio.Reader
	var out *bufio.Writer
	in = bufio.NewReader(os.Stdin)
	out = bufio.NewWriter(os.Stdout)
	defer out.Flush()

	var setsAmount int

	fmt.Fscan(in, &setsAmount)

	for i := 0; i < setsAmount; i++ {
		var lines int
		var sb strings.Builder

		fmt.Fscan(in, &lines)

		for j := 0; j <= lines; j++ {
			s, _ := in.ReadString('\n')
			s = strings.TrimSpace(s)
			sb.WriteString(s)
		}

		var rootFolder Folder

		err := json.Unmarshal([]byte(sb.String()), &rootFolder)
		if err != nil {
			panic(err)
		}

		fmt.Println(DFS(rootFolder, false))

	}
}

