package main

import "fmt"

func isSubString(ref string, str string) bool {
	if ref == "" {
		return false
	}
	var pos =0;
	for i := 0; i < len(ref); i++ {
		found := false
		for j := pos; j < len(str); j++ {
			if ref[i] == str[j] {
				pos = j+1;
				found = true
				break
			}
		}
		if !found {
			return false
		}
	}
	return true
}

func main() {
	var ref string
	var num int
	var arr [100]string
	var str string

	fmt.Scanln(&ref)
	fmt.Scanln(&num)

	for i := 0; i < num && i < len(arr); i++ {
		fmt.Scanln(&str)
		arr[i] = str
	}

	count := 0
	for i := 0; i < num && i < len(arr); i++ {
		if isSubString(ref, arr[i]) {
			count++
		}
	}

	fmt.Println(count)
}