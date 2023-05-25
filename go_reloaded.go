package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"strings"
)

/*
This function conversts given value from a numbering system into another.

	@parameters:
	  - s: a string that contains in the numbering system that want to change.
	  - b: the base of the new numbering sytem.
	@return:
	  - This function returns an int64 that represent s in the numbering system with base b
*/
func convertToInt(s string, b int) int64 {
	p := len(s) - 1
	intger := 0
	negative := false
	for i, char := range s {
		if char == '-' {
			negative = true
		}
		if negative {
			if char != '-' {
				intger = intger + charToInt(char)*power(p-i, b)
			}
		} else {
			intger = intger + charToInt(char)*power(p-i, b)
		}
	}
	if negative {
		return int64(intger * -1)
	}
	return int64(intger)
}

/*
This function conversts given value from a string form into integer form.

	@parameters:
	  - s: a string that contains in the numbering system that want to change.
	@return:
	  - This function returns an int64 that represent the value of s.
*/
func convertDicStrToInt(s string) int64 {
	p := len(s) - 1
	intger := 0
	negative := false
	for i, char := range s {
		if char == '-' {
			negative = true
		}
		if negative {
			if char != '-' {
				intger = intger + charToInt(char)*power(p-i, 10)
			}
		} else {
			intger = intger + charToInt(char)*power(p-i, 10)
		}
	}
	if negative {
		return int64(intger * -1)
	}
	return int64(intger)
}

/*
This function conversts given value from a char into int.

	@parameters:
	  - char: a string that contains in the numbering system that want to change.
	@return:
	  - This function returns an int64 that represent the values of s.
*/
func charToInt(char rune) int {
	if char >= '0' && char <= '9' {
		return int(char - 48)
	} else if char == 'A' {
		return 10
	} else if char == 'B' {
		return 11
	} else if char == 'C' {
		return 12
	} else if char == 'D' {
		return 13
	} else if char == 'E' {
		return 14
	} else if char == 'F' {
		return 15
	}
	return -1
}

/*
This function calculates the power of a number.

	@parameters:
	  - p: the power that base is raised for.
	  - b: the base of the power.
	@return:
	  - This function returns an int64 that represent result of caluculating b ^p.
*/
func power(p int, b int) int {
	r := 1
	for i := 1; i <= p; i++ {
		r = r * b
	}
	return r
}

/*
This function conversts given value from int into char.

	@parameters:
	  - i: an ineger value that the function will change.
	@return:
	  - This function returns a rune that represent the value of i.
*/
func intToChar(i int64) rune {
	x := '0'
	for j := int64(0); j < (i % 10); j++ {
		x++
	}
	return x
}

/*
This function conversts given value from int into string.

	@parameters:
	  - i: an ineger value that the function will change.
	@return:
	  - This function returns a string that represent the value of i.
*/
func intToString(i int64) string {
	str := ""
	counter := 0
	k := i
	if i < 0 {
		k = k * -1
	}
	for j := k; j >= 0; j = j / 10 {
		if j == 0 {
			if counter == 0 {
				char := intToChar(j)
				str = string(char) + str
			}
			break
		}
		char := intToChar(j)
		str = string(char) + str
		counter++
	}
	if i < 0 {
		str = "-" + str
	}
	return str
}

/*
This function extracts the number from the pattern used to format the text.

	@parameters:
	  - s: the patten with a num that represents the required number of repetetions.
	@return:
	  - This function returns an int64 that represent the required time of repetetions.
*/
func extractNum(s string) int64 {
	if s[4] == ',' {
		num := s[5:strings.IndexRune(s, ')')]
		return convertDicStrToInt(num)
	}
	num := s[4:strings.IndexRune(s, ')')]
	return convertDicStrToInt(num)

}

func main() {
	//This condetionnal statement is used to check if the file exists or not.
	if len(os.Args) == 3 {
		text, err := ioutil.ReadFile(os.Args[1])
		if err != nil {
			fmt.Print(err)
		}
		str := string(text)
		//This empty string is used to save the cleaned text (text without extra spaces)
		cStr := ""
		//This counter countes the nuber of spaces before each word.
		countSpace := 0
		//This flage used to define if the char is a quote or not.
		quote := false
		//This flag used to store the prev of each char, I intiated with M for  a secrete.
		prev := 'M'
		//This flag used to define if the single quote is the first or the second.
		firstS := true
		// //This flag used to define if the double quote is the first or the second.
		firstD := true
		for n, v := range str {
			//counting the nuber of spaces before each word.
			if v == ' ' {
				countSpace++
				// handling punctuation marks.
			} else if v == '.' || v == '?' || v == '!' || v == ':' || v == ';' || v == ',' || v == '"' || v == '\'' {
				if v == '\'' && (prev != ' ' || prev != '"') {
					if firstS {
						if str[n+1] == 't' && prev == 'n' {
							cStr += string(v)
							prev = v
							quote = false
						} else {
							cStr += " " + string(v)
							prev = v
							firstS = false
							quote = false
						}
					} else {
						cStr += string(v) + " "
						prev = v
						firstS = true
						quote = false
					}
				} else if v == '"' && (prev != ' ' || prev != '\'') {
					if firstD {
						cStr += " " + string(v)
						prev = v
						firstD = false
						quote = false
					} else {
						if str[n+1] == ' ' {
							cStr += string(v) + " "
							prev = v
							firstD = true
							quote = false
						} else {
							cStr += string(v) + " "
							prev = v
							firstD = true
							quote = false
						}
					}
				} else {
					cStr += string(v)
					prev = v
					quote = false
				}
				countSpace = 0
				if v == '"' || v == '\'' {
					quote = true
				}
				// handling normal letters
			} else {
				if !quote {
					if countSpace > 0 {
						for j := 1; j <= countSpace; j++ {
							cStr = cStr + " "
						}
						countSpace = 0
					}
				} else if v != ' ' {
					quote = false
					countSpace = 0
				}
				if (prev == '.' || prev == '?' || prev == '!' || prev == ':' || prev == ';' || prev == ',') && cStr[len(cStr)-1] != ' ' {
					cStr += " " + string(v)
					quote = false
					prev = v
					countSpace = 0
				} else {
					cStr += string(v)
					prev = v
				}

			}
		}
		println(cStr)
		// converting string into a list of strings.
		strArr := strings.Split(cStr, " ")
		// empty list of strings to store the updated list.
		UpdatedArr := []string{}
		// this counter used to store the difference in size between the original list and the updated one.
		counter := int64(0)
		//looping over the original list and caculate what is needed and creating the new list.
		for i := int64(0); i < int64(len(strArr)); i++ {
			if strings.Count(strArr[i], "(hex)") == 1 {
				counter++
				UpdatedArr[i-counter] = intToString(convertToInt(strArr[i-1], 16)) + strArr[i][5:]
				continue
			} else if strings.Count(strArr[i], "(bin)") == 1 {
				counter++
				UpdatedArr[i-counter] = intToString(convertToInt(strArr[i-1], 2)) + strArr[i][5:]
				continue
			} else if strings.Count(strArr[i], "(up)") == 1 {
				counter++
				UpdatedArr[i-counter] = strings.ToUpper(strArr[i-1]) + strArr[i][4:]
				continue
			} else if strings.Count(strArr[i], "(low)") > 0 {
				counter++
				UpdatedArr[i-counter] = strings.ToLower(strArr[i-1]) + strArr[i][5:]
				continue
			} else if strings.Count(strArr[i], "(cap)") > 0 {
				counter++
				UpdatedArr[i-counter] = strings.Title(strArr[i-1]) + strArr[i][5:]
				continue
			} else if strings.Count(strArr[i], "(cap,") > 0 {
				for j := extractNum(strArr[i] + strArr[i+1]); int64(j) >= 1; j-- {
					UpdatedArr[i-counter-j] = strings.Title(strArr[i-j])
				}
				UpdatedArr[len(UpdatedArr)-1] = UpdatedArr[len(UpdatedArr)-1] + strArr[i+1][strings.IndexRune(strArr[i+1], ')')+1:]
				i++
				counter += 2
				continue
			} else if strings.Count(strArr[i], "(low,") > 0 {
				for j := extractNum(strArr[i] + strArr[i+1]); int64(j) >= 1; j-- {
					UpdatedArr[i-counter-j] = strings.ToLower(strArr[i-j])
				}
				UpdatedArr[len(UpdatedArr)-1] = UpdatedArr[len(UpdatedArr)-1] + strArr[i+1][strings.IndexRune(strArr[i+1], ')')+1:]
				i++
				counter += 2
				continue
			} else if strings.Count(strArr[i], "(up,") > 0 {
				for j := extractNum(strArr[i] + strArr[i+1]); int64(j) >= 1; j-- {
					UpdatedArr[i-counter-j] = strings.ToUpper(strArr[i-j])
				}
				UpdatedArr[len(UpdatedArr)-1] = UpdatedArr[len(UpdatedArr)-1] + strArr[i+1][strings.IndexRune(strArr[i+1], ')')+1:]
				i++
				counter += 2
				continue
			} else if len(strArr[i]) > 0 && i > 0 {
				if strArr[i][0] == 'a' || strArr[i][0] == 'e' || strArr[i][0] == 'i' || strArr[i][0] == 'o' || strArr[i][0] == 'u' || strArr[i][0] == 'h' {
					if strArr[i-1] == "a" {
						UpdatedArr[i-counter-1] = "an"
					} else if strArr[i-1] == "A" {
						UpdatedArr[i-counter-1] = "An"
					}
				}

			}

			UpdatedArr = append(UpdatedArr, strArr[i])
		}
		// combine the list components to create the output string
		output := strings.Join(UpdatedArr, " ")
		// writing the output to the defined file.
		file, err := os.OpenFile(os.Args[2], os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
		if err != nil {
			fmt.Println("Could not open " + os.Args[2])
			return
		}
		defer file.Close()
		_, err2 := file.WriteString(output)
		if err2 != nil {
			fmt.Println("Could not write text to " + os.Args[2])
		} else {
			fmt.Println("Operation successful! Text has been appended to " + os.Args[2])
		}
		println(output)

	} else {
		println("Please enter the name of the file that want to format, please enter just one file name without spaces.")
	}
}
