package main

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
	depth      int
	showHidden bool
	prefix     string
}

func init() {
	flag.StringVar(&rootPath, "p", "", "The path of target directory.")
}

func fileSize(size int64) string {
	str := strconv.FormatInt(size, 10)

	return "[" + str + "]"
}

func listDirectory(basePath string, option *Option) error {
	// depth = depth ? depth : 1
	if option.depth != 0 {

	} else {
		option.depth = 1
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
	// var prefix string
	if len(fis) < 1 {
		// red := color.New(color.FgRed).SprintFunc()

		// fmt.Printf("this is a %s and this is %s.\n", yellow("warning"), red("error"))
		// fmt.Println(option.prefix, "└─", "(null)"+" "+strconv.Itoa(option.depth))
		fmt.Println(option.prefix, FgHiBlack("└─ (null) "))
	} else {
		for index, v := range fis {
			filePath := v.Name()

			tmpOpt := new(Option)
			*tmpOpt = *option

			if strings.HasPrefix(filePath, ".") && !option.showHidden {
				continue
			}
			if v.IsDir() {
				absFp := path.Join(basePath, filePath)

				if index == len(fis)-1 {
					fmt.Println(tmpOpt.prefix, "└─┬", hiCyan(filePath)+" "+fileSize(v.Size()))
				} else {
					fmt.Println(tmpOpt.prefix, "├─┬", hiCyan(filePath)+" "+fileSize(v.Size()))
				}
				if index == len(fis)-1 {
					tmpOpt.prefix += "  "
				} else {
					tmpOpt.prefix += " │"
				}
				err = listDirectory(absFp, tmpOpt)
				if err != nil {
					return err
				}
			} else {
				if index == len(fis)-1 {
					fmt.Println(tmpOpt.prefix, "└──", filePath+" "+fileSize(v.Size()))
				} else {
					fmt.Println(tmpOpt.prefix, "├──", filePath+" "+fileSize(v.Size()))
				}
				// tmpOpt.prefix += "  "
			}
		}
	}

	return nil
}

func main() {
	flag.Parse()

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

	option := Option{depth: 0, showHidden: false, prefix: ""}

	err := listDirectory(rootPath, &option)
	if err != nil {
		fmt.Println("err:", err)
		return
	}

	// for b := 40; b <= 47; b++ { // 背景色彩 = 40-47
	// 	for f := 30; f <= 37; f++ { // 前景色彩 = 30-37
	// 		for d := range []int{0, 1, 4, 5, 7, 8} { // 显示方式 = 0,1,4,5,7,8
	// 			fmt.Printf(" %c[%d;%d;%dm%s(f=%d,b=%d,d=%d)%c[0m ", 0x1B, d, b, f, "", f, b, d, 0x1B)
	// 			fmt.Println("")
	// 		}
	// 	}
	// 	fmt.Println("")
	// }

}
