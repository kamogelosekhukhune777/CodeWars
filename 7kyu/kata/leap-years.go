package kata

/*
Leap Years

In this kata you should simply determine, whether a given year is a leap year or not. In case you don't know the rules, here they are:

	Years divisible by 4 are leap years,
	but years divisible by 100 are not leap years,
	but years divisible by 400 are leap years.
Tested years are in range 1600 ≤ year ≤ 4000.
*/

func IsLeapYear(year int) bool {
	if year%4 == 0 {
		// If the year is divisible by 4, it might be a leap year.
		if year%100 == 0 {
			// If the year is divisible by 100, it's not a leap year unless it's also divisible by 400.
			return year%400 == 0
		}
		return true
	}
	return false
}
