package kata

import "strings"

/*
Build Tower
Build a pyramid-shaped tower, as an array/list of strings, given a positive integer number of floors. A tower block is represented with "*" character.

For example, a tower with 3 floors looks like this:

[
  "  *  ",
  " *** ",
  "*****"
]
And a tower with 6 floors looks like this:

[
  "     *     ",
  "    ***    ",
  "   *****   ",
  "  *******  ",
  " ********* ",
  "***********"
]
func TowerBuilder(nFloors int) []string {
  return []string{}
}
*/

func TowerBuilder(nFloors int) []string {
	tower := make([]string, nFloors)
	width := nFloors*2 - 1

	for i := 0; i < nFloors; i++ {
		stars := (i*2 + 1) // Number of '*' in each row
		spaces := (width - stars) / 2
		row := strings.Repeat(" ", spaces) + strings.Repeat("*", stars) + strings.Repeat(" ", spaces)
		tower[i] = row
	}
	return tower
}
