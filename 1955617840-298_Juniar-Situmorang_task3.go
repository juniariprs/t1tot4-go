package main 

import ( 
	"fmt"
	"strings"
)

func main() { 
	sentence := "selamat malam" 
	var key [7]string 
	key = [7]string{" ", "a", "e", "l", "m", "s", "t"} 
	alpha := make(map[string]int) 

	// loop letter per letter of sentence
	for i := range sentence { 
		fmt.Printf("%c\n", sentence[i])
	}
	
	// strings.Count
	for i := range key {
		value := strings.Count(sentence, key[i])
		
		alpha[key[i]] = value
	}

	fmt.Println(alpha)

}