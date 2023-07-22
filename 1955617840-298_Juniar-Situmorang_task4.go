package main 

import ( 
	"fmt" 
	"os" 
	"strconv"
)

type Biodata struct { 
	name string 
	address string 
	job string 
	go_reason string
}

func main() {
	var friendsBiodata = []Biodata { 
		{name: " Indah", address: " Jln. Melati No. 36", job: " Frontend Engineer", go_reason: " To obtain a job"}, 
		{name: " Wahyu", address: " Jln. Mawar No. 17", job: " Software Engineer", go_reason: " To get new insight"}, 
		{name: " Muhammad", address: " Jln. Kenanga No. 8", job: " Backend Engineer", go_reason: " To learn about Go deeply"}, 
		{name: " John", address: " Jln. Anggrek No. 12", job: " Web Developer", go_reason: " Go for backend web development"},
	}

	var err error
	var no int
	
	args := os.Args[1:]

	ints := make([]int, len(args))

    // Iterate through each element of the original slice and convert to int
    for i, arg := range args {
		if ints[i], err = strconv.Atoi(arg); err != nil {
			panic(err)
		} 
		no = ints[i] - 1
    }

	fmt.Printf("%+v\n", friendsBiodata[no])
}