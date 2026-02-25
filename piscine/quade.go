package piscine

import (
	"github.com/01-edu/z01"
)

func QuadE(x, y int) {
	if x <= 0 || y <= 0 {
		return
	}

	for row := 1; row <= y; row++ {
		for col := 1; col <= x; col++ {
			if row == 1 {
				// Top row
				if col == 1 {
					z01.PrintRune('A')
				} else if col == x {
					z01.PrintRune('C')
				} else {
					z01.PrintRune('B')
				}
			} else if row == y {
				// Bottom row
				if col == 1 {
					z01.PrintRune('C')
				} else if col == x {
					z01.PrintRune('A')
				} else {
					z01.PrintRune('B')
				}
			} else {
				// Middle rows
				if col == 1 || col == x {
					z01.PrintRune('B')
				} else {
					z01.PrintRune(' ')
				}
			}
		}
		z01.PrintRune('\n')
	}
}
