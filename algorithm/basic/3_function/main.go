package main

import (
	"fmt"
)

type Node struct {
	Name string
	Age int
}

func (n *Node) String() string {
	return fmt.Sprintf("%s is age = %d", n.Name, n.Age)
}

func main() {
	node := &Node{
		Name: "zhangsan",
		Age: 10,
	}

	fmt.Println("node content = ", node)

	// f := fibo()
	// for i:=3;i<=9;i++{
	// 	fmt.Printf("i: %d, fibo = %d \n", i, f())
	// }

	// jpgSuffix := suffixFactory(".jpg")
	// txtSuffix := suffixFactory(".txt")

	// fmt.Printf("jpgSuffix(\"a\"): %v\n", jpgSuffix("a"))
	// fmt.Printf("txtSuffix(\"a\"): %v\n", txtSuffix("a"))
}


// func fibo() func() int{
// 	p1,p2 := 1,1 //每次在拿取函数时初始化一个仅对该返回func可见的全局变量
// 	start := time.Now()
// 	return func() int {
// 		p1,p2 = p2, p1+p2
// 		fmt.Printf("time.Now().Sub(start): %v\n", time.Since(start)) //计算调用耗时
// 		return p2
// 	}
// }

// //工厂函数:函数逻辑相似，仅初始化入参不一样，可以将差异入参作为闭包入参
// func suffixFactory(suffix string) func(string) string{
// 	return func(name string) string{
// 		return name+suffix
// 	}
// }
