package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"strings"
)

func convertHexStrToInt(s string) int64 {
	p := len(s) - 1
	intger := 0
	negative := false
	for i, char := range s {
		if char == '-' {
			negative = true
		}
		if negative {
			if char != '-' {
				intger = intger + charToInt(char)*power16(p-i)
			}
		} else {
			intger = intger + charToInt(char)*power16(p-i)
		}
	}
	if negative {
		return int64(intger * -1)
	}
	return int64(intger)
}

func convertBinStrToInt(s string) int64 {
	p := len(s) - 1
	intger := 0
	negative := false
	for i, char := range s {
		if char == '-' {
			negative = true
		}
		if negative {
			if char != '-' {
				intger = intger + charToInt(char)*power2(p-i)
			}
		} else {
			intger = intger + charToInt(char)*power2(p-i)
		}
	}
	if negative {
		return int64(intger * -1)
	}
	return int64(intger)
}

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
				intger = intger + charToInt(char)*power10(p-i)
			}
		} else {
			intger = intger + charToInt(char)*power10(p-i)
		}
	}
	if negative {
		return int64(intger * -1)
	}
	return int64(intger)
}

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

func power10(p int) int {
	r := 1
	for i := 1; i <= p; i++ {
		r = r * 10
	}
	return r
}

func power2(p int) int {
	r := 1
	for i := 1; i <= p; i++ {
		r = r * 2
	}
	return r
}

func power16(p int) int {
	r := 1
	for i := 1; i <= p; i++ {
		r = r * 16
	}
	return r
}

func intToChar(i int64) rune {
	x := '0'
	for j := int64(0); j < (i % 10); j++ {
		x++
	}
	return x
}

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

func extractNum(s string) int64 {
	if s[4] == ',' {
		num := s[5:strings.IndexRune(s, ')')]
		return convertDicStrToInt(num)
	}
	num := s[4:strings.IndexRune(s, ')')]
	return convertDicStrToInt(num)

}

func main() {
	if len(os.Args) == 2 {
		text, err := ioutil.ReadFile(os.Args[1])
		if err != nil {
			fmt.Print(err)
		}
		str := string(text)
		cStr := ""
		countSpace := 0
		quote := false
		prev := 'M'
		for _, v := range str {
			if v == ' ' {
				countSpace++
			} else if v == '.' || v == '?' || v == '!' || v == ':' || v == ';' || v == ',' || v == '"' || v == '\'' {
				if (v == '\'' || v == '"') && (prev == '.' || prev == '?' || prev == '!' || prev == ':' || prev == ';' || prev == ',') {
					cStr += " " + string(v)
					prev = v
					quote = false
				} else {
					cStr += string(v)
					prev = v
					quote = false
				}
				countSpace = 0
				if v == '"' || v == '\'' {
					quote = true
				}
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
				cStr += string(v)
				prev = v
			}
		}
		println(cStr)

		strArr := strings.Split(cStr, " ")
		UpdatedArr := []string{}
		counter := int64(0)

		for i := int64(0); i < int64(len(strArr)); i++ {
			if strings.Count(strArr[i], "(hex)") > 0 {
				counter++
				UpdatedArr[i-counter] = intToString(convertHexStrToInt(strArr[i-1])) + strArr[i][5:]
				continue
			} else if strings.Count(strArr[i], "(bin)") > 0 {
				counter++
				UpdatedArr[i-counter] = intToString(convertBinStrToInt(strArr[i-1])) + strArr[i][5:]
				continue
			} else if strings.Count(strArr[i], "(up)") > 0 {
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
				i++
				counter += 2
				continue
			} else if strings.Count(strArr[i], "(low,") > 0 {
				for j := extractNum(strArr[i] + strArr[i+1]); int64(j) >= 1; j-- {
					UpdatedArr[i-counter-j] = strings.ToLower(strArr[i-j])
				}
				i++
				counter += 2
				continue
			} else if strings.Count(strArr[i], "(up,") > 0 {
				for j := extractNum(strArr[i] + strArr[i+1]); int64(j) >= 1; j-- {
					UpdatedArr[i-counter-j] = strings.ToUpper(strArr[i-j])
				}
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

		output := strings.Join(UpdatedArr, " ")
		println(output)

	} else {
		println("Please enter the name of the file that want to format, please enter just one file name without spaces.")
	}
}
