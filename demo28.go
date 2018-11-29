package main

import (
	"fmt"
)

func main () {
	complexArray1 := [3][]string{
		[]string{"d", "e", "f"},
		[]string{"g", "h", "i"},
		[]string{"j", "k", "l"},
	}	
	fmt.Printf("array : %v\n",complexArray1)
	fmt.Printf("array address: %p\n",&complexArray1)

	array1 := modifyArray(complexArray1)
	fmt.Printf("origin array : %v\n",complexArray1)
	fmt.Printf("origin address: %p\n",&complexArray1)
	fmt.Printf("modify array : %v\n",array1)
	fmt.Printf("modify address: %p\n",&array1)

}
/// 拷贝的是对象引用指针 浅拷贝
func modifyArray (a [3][]string) [3][]string {
	fmt.Printf("method params address: %p\n",&a)
	a[0]=[]string{"a","b","c","d"}
	a[1][0] = "1"
	return a;
}

/**
array : [[d e f] [g h i] [j k l]]
array address: 0xc4200880f0
method params address: 0xc4200881e0
origin array : [[d e f] [1 h i] [j k l]]
origin address: 0xc4200880f0
modify array : [[a b c d] [1 h i] [j k l]]
modify address: 0xc420088190
 */