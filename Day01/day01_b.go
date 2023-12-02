package Day01

import (
	_ "embed"
	"fmt"
	"strconv"
	"strings"
)

//go:embed inputB.txt
var inputB string

func Day01B() {
	lines := strings.Split(inputB, "\n")
	// fmt.Printf("lines: %#v", lines)
	// Totals := []int{}
	sum := 0

	for _, line := range lines {
		tmp := ""
		for _, char := range line {
			if char
			if char >= '0' && char <= '9' {
				// fmt.Printf("char is %c", char)
				tmp = string(char)
				break
			}
		}
		for i := len(line) - 1; i >= 0; i-- {
			if line[i] >= '0' && line[i] <= '9' {
				// fmt.Printf("end char is %c", line[i])
				tmp += string(line[i])
				break
			}
		}
		num, err := strconv.Atoi(tmp)
		if err != nil {
			fmt.Println("something done went wrong")
		}
		// Totals = append(Totals, num)
		sum += num
	}

	// fmt.Printf("%#v", Totals)
	fmt.Println(sum)

}
