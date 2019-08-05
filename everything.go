package main

import(
	"fmt"
	"os"
)

func main(){
	fmt.Println("Hello World")
	nextHello()
}

func nextHello(){
	var x int //variable declaration
	x=3
	y:=4 //short declration to infer the type, declare and assign
	sum,prod:=learnMultiple(x,y)
	fmt.Println("sum: ",sum,"prod: ",prod)
	learnTypes()
}

// functions with parameters must also contain the return types 
// func <name>(input params)(return params)
func learnMultiple(x,y int)(sum,prod int){
	return x+y,x*y
}

// some extra variable types
func learnTypes(){
	str:="Learn Go!" //string 
	s2:=`Raw String
	with breaks and spaces` //raw string

	g := 'Σ' //UTF-8 source code
	f:=3.14195 //float
	c:=3+4i

	var u uint=7 //unsigned int
	var pi float32=22./7

	n:=byte('\n') //byte is also uint8

	var a4 [4] int // array declaration
	a5:= [...]int{3,1,5,10,100} // array initialization 

	s3:= []int{4,5,9}
	s4:=make([]int,4)
	var d2 [][]float64
	bs:=[]byte("a slice")

	s := []int{1, 2, 3}     // Result is a slice of length 3.
    s = append(s, 4, 5, 6)  // Added 3 elements. Slice now has length of 6.
	fmt.Println(s)
	
	s = append(s, []int{7, 8, 9}...) // Second argument is a slice literal.
	fmt.Println(s)
	
	m := map[string]int{"three": 3, "four": 4}
	m["one"] = 1
	
	//file writing underscore denotes the error variable which is being discarded
	file, _ := os.Create("output.txt")
    fmt.Fprint(file, "This is how you write to a file, by the way")
    file.Close()

    // Unused variables are an error in Go.
    // The underscore lets you "use" a variable but discard its value.
	_, _, _, _, _, _, _, _, _, _ = str, s2, g, f, u, pi, n, a5, s4, bs
	
	fmt.Println(s, c, a4, s3, d2, m)

}