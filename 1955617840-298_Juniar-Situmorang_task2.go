package main 

import "fmt" 

func main() { 
	// loop for i starts from zero to four
	for i := 0; i < 5; i++ { 
		fmt.Printf("Nilai i = %d\n", i) 
		// loop for j starts from zero to ten
		for j := 0; j <= 10; j++ {
			if j == 5 { 
			// loop for unicode
				char := "САШАРВО"
				for pos, val := range char {
					fmt.Printf("character %U '%c' starts at byte position %d\n", val, val, pos)
				}
			} else { 
				fmt.Printf("Nilai j = %d\n", j)
			}
		}
	}
	
}