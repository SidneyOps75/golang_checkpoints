package main

import "fmt"

func HashCode(dec string) string {
	result := ""
	size := len(dec)

	for _,c := range dec {
		hashedchar := (int(c) + size)%127
		if hashedchar < 33 {
			hashedchar += 33
		}
		result += string(rune(hashedchar))
	}
	return result
}

func main() {
	fmt.Println(HashCode("A"))
	fmt.Println(HashCode("AB"))
	fmt.Println(HashCode("BAC"))
	fmt.Println(HashCode("Hello World"))
}
