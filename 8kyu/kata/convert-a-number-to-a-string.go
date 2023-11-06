package kata 

import "fmt"

/*
Convert a number to a string

We need a function that can transform a number (integer) into a string.

What ways of achieving this do you know?

Examples (input --> output):
123  --> "123"
999  --> "999"
-100 --> "-100"

func NumberToString(n int) string {
  return ""
}
*/
func NumberToString(n int) string {
    return fmt.Sprintf("%d", n)
}