In Golang, empty struct doesn't occupy any memory
struct{}

You can use "unsafe" library to check a variable's memory size
unsafe.Sizeof(struct{})

https://zhuanlan.zhihu.com/p/67580188
https://allenwu.itscoder.com/set-in-go

1. array_based
2. map_based (hash)

variants 
1. ordered
2. unordered

Golang's Generic for declaration of Set interface:
https://go.dev/doc/tutorial/generics

Comparing 2 maps, use "reflect" package's DeepEqual func