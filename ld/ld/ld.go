package ld

import (
	"flag"
	"fmt"
	"os"
	"path"
	"strconv"
	"strings"

	"github.com/fatih/color"
)

// Const Variable
const (
	INDENT = "  "
)

var (
	rootPath string
)

// Option is
type Option struct {
	Depth      int
	ShowHidden bool
	Prefix     string
}

func init() {
	flag.StringVar(&rootPath, "p", "", "The path of target directory.")

}

func fileSize(size int64) string {
	var str string
	if size > 1024*1024*1024 {
		size /= 1024 * 1024 * 1024
		str = strconv.FormatInt(size, 10) + "GB"
	} else if size > 1024*1024 {
		size /= 1024 * 1024
		str = strconv.FormatInt(size, 10) + "MB"
	} else if size > 1024 {
		size /= 1024
		str = strconv.FormatInt(size, 10) + "KB"
	} else {
		str = strconv.FormatInt(size, 10) + "B"
	}

	return "[" + str + "]"
}

// ListDirectory aha
func ListDirectory(basePath string, option *Option) error {
	// Depth = Depth ? Depth : 1
	if option.Depth != 0 {

	} else {
		option.Depth = 1
	}
	pFile, err := os.Open(basePath)
	if err != nil {
		return err
	}
	fis, err := pFile.Readdir(-1)
	if err != nil {
		return err
	}

	hiCyan := color.New(color.FgHiCyan).SprintFunc()
	FgHiBlack := color.New(color.FgHiBlack).SprintFunc()

	// var Prefix string
	if len(fis) < 1 {
		fmt.Println(option.Prefix, FgHiBlack("└─ (null) "))
	} else {
		for index, v := range fis {
			filePath := v.Name()

			tmpOpt := new(Option)
			*tmpOpt = *option

			if strings.HasPrefix(filePath, ".") && !option.ShowHidden {
				continue
			}
			if v.IsDir() {
				absFp := path.Join(basePath, filePath)

				if index == len(fis)-1 {
					fmt.Println(tmpOpt.Prefix, "└─┬", hiCyan(filePath)+" "+fileSize(v.Size()))
				} else {
					fmt.Println(tmpOpt.Prefix, "├─┬", hiCyan(filePath)+" "+fileSize(v.Size()))
				}
				if index == len(fis)-1 {
					tmpOpt.Prefix += "  "
				} else {
					tmpOpt.Prefix += " │"
				}
				err = ListDirectory(absFp, tmpOpt)
				if err != nil {
					return err
				}
			} else {
				if index == len(fis)-1 {
					fmt.Println(tmpOpt.Prefix, "└──", filePath+" "+fileSize(v.Size()))
				} else {
					fmt.Println(tmpOpt.Prefix, "├──", filePath+" "+fileSize(v.Size()))
				}
			}
		}
	}
	return nil
}

// Run begin
func Run(option Option) {
	if len(rootPath) == 0 {
		defaultPath, err := os.Getwd()
		if err != nil {
			fmt.Println("GetwdError:", err)
			return
		}
		rootPath = defaultPath
	}

	fmt.Printf("Now in directory of: %s\n", rootPath)
	fmt.Println(" ┬")

	// option := ld. Option{Depth: 0, ShowHidden: false, Prefix: ""}

	err := ListDirectory(rootPath, &option)
	if err != nil {
		fmt.Println("err:", err)
		return
	}
}
