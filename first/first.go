package main

import "fmt"

type st struct {
	stInt int
	stString string
}

func t(Val *st,a string)() {
	fmt.Println(Val.stInt,"\n"+Val.stString)
	Val.stString = a
	fmt.Println(Val.stInt,"\n"+Val.stString)
}

func tt(test, test2 int) (int) {
	return test+test2
}


func main(){
	var TestST = new(st)
	a := "newString"
	TestST.stInt = 27
	TestST.stString = "testtest"
	t(TestST,a)
	fmt.Println(TestST.stInt,"\n"+TestST.stString)
	
	nums := []int{1,2,3}
	for index,value := range nums {
		fmt.Println(index,value)
	}

}