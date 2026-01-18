package main

import (
	"fmt"
	"strconv"
	"strings"
)

func main() {
    str := "你好hello"
    fmt.Println(len(str)) // num of byte
    strRune := []rune(str)
    fmt.Println(len(strRune)) // num of char

    fmt.Println(strings.HasPrefix(str, "你好"))
    fmt.Println(strings.HasSuffix(str, "llo"))
    fmt.Printf("strings.Contains(str, \"hello\"): %v\n", strings.Contains(str, "hello"))

    str += ", demo1, demo2"
    fmt.Printf("strings.Count(str, \"o\"): %v\n", strings.Count(str, "o"))
    fmt.Printf("strings.ToUpper(str): %v\n", strings.ToUpper(str))
    fmt.Printf("strings.ToLower(str): %v\n", strings.ToLower(str))

    strSlice := strings.Split(str, ",")
    fmt.Printf("strSlice: %v\n", strSlice)
    fmt.Printf("strings.Join(strSlice, \"|\"): %v\n", strings.Join(strSlice, "|"))

    num := 123
    numStr := strconv.Itoa(num)
    num2, _ := strconv.Atoi(numStr)
    fmt.Printf("numStr: %v\n", numStr)
    fmt.Printf("num2: %v\n", num2)

    fmt.Printf("strings.Map(callback, str): %v\n", strings.Map(callback, str))
}

func callback(c rune) rune {
    if c > 127 {
        return '?'
    }
    return c
}