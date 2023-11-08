package kata

import (
	"fmt"
	"strings"
)

//help the bookseller!
/*
A bookseller has lots of books classified in 26 categories labeled A, B, ... Z.
Each book has a code c of 3, 4, 5 or more characters. The 1st character of a code
is a capital letter which defines the book category.

In the bookseller's stocklist each code c is followed by a space and by a positive
integer n (int n >= 0) which indicates the quantity of books of this code in stock.

For example an extract of a stocklist could be:

L = {"ABART 20", "CDXEF 50", "BKWRK 25", "BTSQZ 89", "DRTYM 60"}.
or
L = ["ABART 20", "CDXEF 50", "BKWRK 25", "BTSQZ 89", "DRTYM 60"] or ....
You will be given a stocklist (e.g. : L) and a list of categories in capital letters e.g :

M = {"A", "B", "C", "W"}
or
M = ["A", "B", "C", "W"] or ...
and your task is to find all the books of L with codes belonging to each category of M
and to sum their quantity according to each category.

For the lists L and M of example you have to return the string
(in Haskell/Clojure/Racket/Prolog a list of pairs):

(A : 20) - (B : 114) - (C : 50) - (W : 0)
where A, B, C, W are the categories, 20 is the sum of the unique book of category A,
114 the sum corresponding to "BKWRK" and "BTSQZ", 50 corresponding to "CDXEF" and 0 to category
'W' since there are no code beginning with W.

If L or M are empty return string is "" (Clojure/Racket/Prolog should return an empty array/list instead).

Notes:
In the result codes and their values are in the same order as in M.
See "Samples Tests" for the return.
*/

func StockList(listArt []string, listCat []string) string {
	if len(listArt) == 0 || len(listCat) == 0 {
		return ""
	}

	categorySum := make(map[string]int)

	// Initialize the categorySum map with 0 for all categories
	for _, cat := range listCat {
		categorySum[cat] = 0
	}

	// Iterate through the stocklist
	for _, book := range listArt {
		parts := strings.Fields(book)
		if len(parts) < 2 {
			continue
		}

		code := parts[0]
		quantity := 0

		// Convert the quantity to an integer
		fmt.Sscanf(parts[1], "%d", &quantity)

		// Get the category from the first character of the code
		category := string(code[0])

		// Check if the category is in the list of categories
		if _, exists := categorySum[category]; exists {
			categorySum[category] += quantity
		}
	}

	// Create the result string
	result := ""
	for i, category := range listCat {
		if i > 0 {
			result += " - "
		}
		result += fmt.Sprintf("(%s : %d)", category, categorySum[category])
	}

	return result
}
