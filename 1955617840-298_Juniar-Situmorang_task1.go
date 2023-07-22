package main 

import ( 
	"fmt"
	"strconv"
)

func main() { 
	var i int = 21 
	var j bool = true
	russia := 'Я'
	var k float64 = 123.456

	// menampilkan nilai dari variabel i
	fmt.Printf("%v\n", i) 

	// menampilkan tipe data dari variabel i 
	fmt.Printf("%T\n", i) 

	// menampilkan tanda %
	fmt.Print("%\n") 

	// menampilkan nilai j 
	fmt.Printf("%v\n", j) 

	// menampilkan unicode russia 
	fmt.Printf("%+U\n", russia)

	// menampilkan nilai base 2 dari 21
	fmt.Println(strconv.FormatInt(21, 2))

	// menampilkan nilai base 10 dari 21
	fmt.Println(strconv.FormatInt(21, 10))

	// menampilkan nilai base 8 dari 21
	fmt.Println(strconv.FormatInt(21, 8))

	// menampilkan nilai base 16 dari f, F, 13
	fmt.Printf("%x\n", "f")
	fmt.Printf("%x\n", "F") 
	fmt.Printf("%x\n", 13)

	// menampilkan karakter Я dengan unicode-nya 
	fmt.Printf("%c\n", '\u042F')

	// menampilkan float k
	fmt.Printf("%.6f\n", k)

	// menampilkan float scientific k 
	fmt.Printf("%E\n", k)

}