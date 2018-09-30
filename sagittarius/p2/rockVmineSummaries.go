package p2

import (
	"bufio"
	"fmt"
	"net/http"
	"strings"
)

var llist [][]string

// RockVmineSummaries vv
func RockVmineSummaries() {
	resp, err := http.Get("https://archive.ics.uci.edu/ml/machine-learning-databases/undocumented/connectionist-bench/sonar/sonar.all-data")

	if err != nil {
		fmt.Println(err)
	}
	defer resp.Body.Close()
	// fmt.Println(reflect.TypeOf(resp.Body))
	// body, err := ioutil.ReadAll(resp.Body)

	reader := bufio.NewReader(resp.Body)

	for {
		line, isPrefix, err := reader.ReadLine()

		if err != nil {
			break
		}
		if !isPrefix {
			llist = append(llist, strings.Split(strings.TrimSpace(string(line)), ","))
		}
	}
	fmt.Println("Number of Rows of Data =", len(llist))
	fmt.Println("Number of Columns of Data =", len(llist[1]))
}
