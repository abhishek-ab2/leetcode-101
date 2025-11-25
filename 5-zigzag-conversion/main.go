package main

import "fmt"

func combine(grid [][]rune) string {
    out := make([]rune, 0)

    for _, row := range grid {
        for _, r := range row {
            if r != 0 {
                out = append(out, r)
            }
        }
    }

    return string(out)
}


func convert(s string, numRows int) string {
	if numRows == 1{
		return s
	}
	var result [][]rune
	for i:=0; i < numRows; i++{
		result = append(result, []rune{})
	}
	var sameRow bool = true
	var row int = 0
	for _, char := range s {

		result[row] = append(result[row], char)

		// if !sameRow{
		// 	for i:=0; i < numRows; i++{
		// 		if (i != row) {
		// 			result[i] = append(result[i], rune(0))
		// 		}
		// 	}			
		// }
		if sameRow{
			row++
			sameRow = (row <= numRows - 1) || (numRows <= 2)
			if row == numRows {
				row--
				row--
			}
		} else {
			row = max(row-1, 0)
			sameRow = row == 0
		}
	}
	return combine(result)
}

func main() {
	s := "abcd"
	fmt.Println(convert(s, 2))
}
