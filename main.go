package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
	"reflect"
	"regexp"
	"strconv"
	"strings"
	"time"
)

func main() {
	//Open file
	file, err := os.Open("input.json")
	if err != nil {
		fmt.Println(err)
	}
	defer file.Close()

	bytes, _ := ioutil.ReadAll(file)

	pkg := make(map[string]any)

	json.Unmarshal(bytes, &pkg)

	output := convert(pkg)

	outMarsh, err := json.Marshal(output)
	if err != nil {
		fmt.Println(err)
	}

	os.WriteFile("output.json", outMarsh, 0644)
  fmt.Println(output)
}

func convert(input any) any {

	if reflect.TypeOf(input).Kind() == reflect.Map {
		output := make(map[string]any)
		input, err := input.(map[string]any)
		if !err {
			return nil
		}

		for k, v := range input {
			k = strings.TrimSpace(k)
			if len(k) == 0 || v == nil {
				continue
			}

			switch k {
			case "S":
				if reflect.TypeOf(v).Kind() == reflect.String {
					return convertString(v.(string))
				}
				return nil
			case "N":
				if reflect.TypeOf(v).Kind() == reflect.String {
					return convertNum(v.(string))
				}
				return nil

			case "BOOL":
				if reflect.TypeOf(v).Kind() == reflect.String {
					return convertBool(v.(string))
				}
				return nil

			case "NULL":
				if reflect.TypeOf(v).Kind() == reflect.String {
					return convertNull(v.(string))
				}
				return nil
			case "M":
				if reflect.TypeOf(v).Kind() == reflect.Map {
					res := convert(v)

					if res != nil && reflect.TypeOf(res).Kind() == reflect.Map {
						return res
					}
				}
				return nil
			case "L":
				if reflect.TypeOf(v).Kind() == reflect.Slice {
					res := convert(v)

					if res != nil && reflect.TypeOf(res).Kind() == reflect.Slice {
						return res
					}
				}
				return nil
			default:
				if reflect.TypeOf(v).Kind() == reflect.Map {
					res := convert(v)
					if res != nil {
						if res == "null" {
							res = nil
						}
						output[k] = res
					}
				}

			}
		}

		if len(output) > 0 {
			return output
		}
		return nil
	} else if reflect.TypeOf(input).Kind() == reflect.Slice {
		var output []any
		input, err := input.([]any)
		if !err {
			return nil
		}

		for _, v := range input {
			if reflect.TypeOf(v).Kind() == reflect.Map {
				res := convert(v.(map[string]any))
				if res != nil {
					output = append(output, res)
				}
			}
		}

		if len(output) > 0 {
			return output
		}
		return nil
	}

	return nil
}

func convertString(in string) any {
	//Sanitize string
	in = strings.TrimSpace(in)
	//Attempt to convert to epoch
	parse, err := time.Parse(time.RFC3339, in)
	if err == nil {
		return parse.Unix()
	}

	//Return input if not blank string
	if len(in) > 0 {
		return in
	}

	return nil
}

func convertNum(in string) any {
	//Sanitize
	in = strings.TrimSpace(in)
	in = strings.TrimLeft(in, "0")

	//Check if float or num
	matchFloat, _ := regexp.MatchString("^[0-9]+[.]{1}[0-9]+$", in)
	matchInt, _ := regexp.MatchString("^-?[0-9]+$", in)
	if matchFloat {
		//If matched float regex, parse, if no error, return output
		out, err := strconv.ParseFloat(in, 64)
		if err == nil {
			return out
		}
		return nil
	} else if matchInt {
		out, err := strconv.Atoi(in)
		if err == nil {
			return out
		}
		return nil
	}

	return nil
}

func convertBool(in string) any {
	//Sanitize
	in = strings.TrimSpace(in)
	out, err := strconv.ParseBool(in)

	//If no error, return corresponding boolean
	if err == nil {
		return out
	}
	return nil
}

func convertNull(in string) any {
	//Sanitize
	in = strings.TrimSpace(in)
	out, err := strconv.ParseBool(in)
	if err == nil {
		if out {
			return "null"
		}
	}

	return nil
}
