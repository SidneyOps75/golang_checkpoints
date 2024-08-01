package main

import "fmt"

func HashCode(dec string) string {
	var result string
	size := len(dec)
	for _, c := range dec {
		hashed := (int(c) + size) % 127
		if hashed < 33 {
			hashed += 33
		}
		result += string(rune(hashed))
	}
	return result
}

func main() {
	fmt.Println(HashCode("A"))
	fmt.Println(HashCode("AB"))
	fmt.Println(HashCode("BAC"))
	fmt.Println(HashCode("Hello World"))
}
