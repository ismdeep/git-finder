package main

import (
	"fmt"
	"github.com/ismdeep/ismdeep-go-utils/args_util"
	"io/ioutil"
	"os"
)

func ShowHelpMsg() {
	fmt.Println("Usage: git-finder <path>")
}

func listAllFileByName(fileDir string) {
	files, _ := ioutil.ReadDir(fileDir)
	for _, item := range files {
		if item.IsDir() && item.Name() == ".git" {
			fmt.Println(fileDir)
			return
		}
	}

	for _, item := range files {
		if item.IsDir() {
			listAllFileByName(fmt.Sprintf("%v/%v", fileDir, item.Name()))
		}
	}
}

func main() {
	if args_util.Exists("--help") {
		ShowHelpMsg()
		return
	}

	if len(os.Args) < 2 {
		ShowHelpMsg()
		return
	}

	workDir := os.Args[1]

	listAllFileByName(workDir)
}
