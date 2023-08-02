package main

import "fmt"

func main() {
	arr := []int{1, 2, 3, 4, 5}
	newArr := removeElement(arr, 3)
	fmt.Println("new Arr:", newArr)
}

/*
使用append方法，将要删除的元素前面那段切片和后面那段切片拼接起来
*/
func removeElement(s []int, index int) []int {
	if index > len(s)-1 {
		return nil
	}
	return append(s[:index], s[index+1:]...)
}
