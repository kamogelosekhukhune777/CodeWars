package kata

/*
altERnaTIng cAsE <=> ALTerNAtiNG CaSe
Define String.prototype.toAlternatingCase 
(or a similar function/method such as 
to_alternating_case/toAlternatingCase/ToAlternatingCase
in your selected language; see the initial solution for details)
such that each lowercase letter becomes uppercase and each uppercase
letter becomes lowercase. For example:

ToAlternatingCase("hello world"); // => "HELLO WORLD"
ToAlternatingCase("HELLO WORLD"); // => "hello world"
ToAlternatingCase("hello WORLD"); // => "HELLO world"
ToAlternatingCase("HeLLo WoRLD"); // => "hEllO wOrld"
ToAlternativeCase("12345"); // => "12345" (Non-alphabetical characters are unaffected)
ToAlternativeCase("1a2b3c4d5e"); // => "1A2B3C4D5E"
ToAlternativeCase("String.prototype.toAlternatingCase"); // => "sTRING.PROTOTYPE.TOaLTERNATINGcASE"
As usual, your function/method should be pure, i.e. it should not mutate the original string.
*/

func ToAlternatingCase(str string) string {
	result := make([]rune, len(str))
	  for i, char := range str {
		  if 'a' <= char && char <= 'z' {
			  // Convert lowercase to uppercase
			  result[i] = char - 'a' + 'A'
		  } else if 'A' <= char && char <= 'Z' {
			  // Convert uppercase to lowercase
			  result[i] = char - 'A' + 'a'
		  } else {
			  // Keep non-letter characters as they are
			  result[i] = char
		  }
	  }
	  return string(result)
  }