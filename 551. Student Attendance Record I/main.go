package main

// You are given a string s representing an attendance record for a student where each character signifies whether the student was absent, late, or present on that day. The record only contains the following three characters:

// 'A': Absent.
// 'L': Late.
// 'P': Present.
// The student is eligible for an attendance award if they meet both of the following criteria:

// The student was absent ('A') for strictly fewer than 2 days total.
// The student was never late ('L') for 3 or more consecutive days.
// Return true if the student is eligible for an attendance award, or false otherwise.

import (
	"fmt"
	"regexp"
	"strings"
)

func checkRecord(s string) bool {
	countA := 0
	str := strings.SplitAfter(s, "")

	match, _ := regexp.MatchString("L{3}", s)

	for i := 0; i < len(str); i++ {
		if str[i] == "A" {
			countA += 1
		}
	}

	if (match == true) || (countA >= 2) {
		return false
	} else {
		return true
	}
}

func main() {
	fmt.Println("Case 1")
	check := checkRecord("PPALLP")
	fmt.Println(check)
	fmt.Println("Case 2")
	check2 := checkRecord("PPALLL")
	fmt.Println(check2)
	fmt.Println("Case 3")
	check3 := checkRecord("LALL")
	fmt.Println(check3)
}
