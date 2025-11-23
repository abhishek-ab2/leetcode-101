package main

import "fmt"

func longestPalindrome(s string) string {
 if len(s) < 2 {
        return s
    }

  expand := func(left, right int) (int, int) {
    for ;left >= 0 && left <= right && right < len(s); {
        if s[left] != s[right] {
            return left+1, right-1
        }
        left--
        right++
    }
    return left+1, right-1
  }


    finalLeft := 0
    finalRight := 0

   for i := 0 ; i < len(s); i ++ {
    l1, r1 := expand(i, i)

    l2 , r2 := expand(i, i+1)

    if r1 - l1 > finalRight - finalLeft {
        finalRight = r1
        finalLeft = l1
    }
    if r2 - l2 > finalRight - finalLeft {
        finalRight = r2
        finalLeft = l2
    }
   }

    return s[finalLeft:finalRight+1]
}

func main () {
	s := "abbcccaba"
	fmt.Println(longestPalindrome(s))
}
