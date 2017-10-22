package main

import (
	"flag"
	// "go-utility/ld"
	"./ld"
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

// func init() {
// 	flag.StringVar(&rootPath, "p", "", "The path of target directory.")
// }

func main() {
	flag.Parse()

	ld.Run(ld.Option{Depth: 0, ShowHidden: false, Prefix: ""})
	// ld.Run(Option{Depth: 0, ShowHidden: false, Prefix: ""})

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
