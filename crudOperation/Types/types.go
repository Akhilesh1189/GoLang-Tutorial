package main

import("fmt")

func Add(x float64,y float64) float64 {
	return x+y
}
func Multiply(a string,b string)(string,string) {
	return a,b
}

func main(){
	var num1 float64 = 5.7
	var num2 float64 =8.9
	fmt.Println("Addition:-", Add(num1,num2))
	 str1 := "akhilesh"
	str2 := "pandey"
	fmt.Println(Multiply(str1,str2))
}