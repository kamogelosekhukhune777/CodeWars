package kata

import "strings"

/*
Abbreviate a Two Word Name

Write a function to convert a name into initials. This kata strictly takes two words with one space in between them.

The output should be two capital letters with a dot separating them.

It should look like this:

 Sam Harris => S.H

 patrick feeney => P.F
*/
 

func AbbrevName(name string) string{
	nameParts := strings.Split(name, " ")

    if len(nameParts) != 2 {
        // Return an error or handle the input appropriately
        return "Invalid input"
    }

    // Take the first character of the first name and the first character of the last name
    initials := string(nameParts[0][0]) + "." + string(nameParts[1][0])

    // Convert to uppercase
    return strings.ToUpper(initials)
}
