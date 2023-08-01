package main

import (
	"fmt"
	"geektime/first/delete_slice"
)

func main() {
	slice := delete_slice.Slice[int]{}

	slice.SliceArray = append(slice.SliceArray, 1, 2, 3, 4, 5, 6, 7, 8)
	fmt.Printf("删除前 slice数组的值为: %v, 长度为：%d, 容量为：%d \n",
		slice.SliceArray, len(slice.SliceArray), cap(slice.SliceArray))
	slice.DeleteByIndex(7)
	fmt.Printf("删除后 slice数组的值为: %v, 长度为：%d, 容量为：%d \n",
		slice.SliceArray, len(slice.SliceArray), cap(slice.SliceArray))
	slice.DeleteByIndex(1)
	fmt.Printf("删除后 slice数组的值为: %v, 长度为：%d, 容量为：%d \n",
		slice.SliceArray, len(slice.SliceArray), cap(slice.SliceArray))
	slice.DeleteByIndex(1)
	fmt.Printf("删除后 slice数组的值为: %v, 长度为：%d, 容量为：%d \n",
		slice.SliceArray, len(slice.SliceArray), cap(slice.SliceArray))
	slice.DeleteByIndex(1)
	fmt.Printf("删除后 slice数组的值为: %v, 长度为：%d, 容量为：%d \n",
		slice.SliceArray, len(slice.SliceArray), cap(slice.SliceArray))

}
