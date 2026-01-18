package main

import (
	"fmt"
	"reflect"
)

type Animal interface {
	Speak(string)
}

type Dog struct {
	Name string
}

func (d *Dog) Speak(s string) {
	fmt.Println(d.Name + " is speaking " + s)
}

type Cat struct {
	Name string
}

func (c *Cat) Speak(s string) {
	fmt.Println(c.Name + " is speaking " + s)
}

func main() {
	var animal Animal
	dog := &Dog{"Dog"}
	cat := &Cat{"Cat"}
	
	animal = dog
	animal.Speak("main test")
	animal = cat
	animal.Speak("main test")

	AnimalSpeak(animal)
}

func AnimalSpeak(ani Animal) {
	//1.类型推断
	if v, ok := ani.(*Dog); ok {
		v.Speak("test1")
	}
	if v, ok := ani.(*Cat); ok {
		v.Speak("test1")
	}

	//2.swith-type
	switch t := ani.(type) {
	case *Dog:
		t.Speak("test2")
	case *Cat:
		t.Speak("test2")
	default:
		fmt.Println("unknow is speaking")
	}

	//3.reflect 类型判断，方法直接调用，直接访问字段
	t := reflect.TypeOf(ani)
	v := reflect.ValueOf(ani)
	var baseType reflect.Type
	var baseValue reflect.Value
	if t.Kind() == reflect.Ptr {
		//指针类型需要借助Elem访问到底层的存储对象
		baseType = t.Elem()   // 剥离指针，拿到 Dog/Cat 类型
		baseValue = v.Elem()  // 剥离指针，拿到 Dog/Cat 值
	} else if t.Kind() == reflect.Struct {
		baseType = t
		baseValue = v
	}

	if baseType.Name()  == "Dog" {
		fmt.Println("it is a dog")
	} else if baseType.Name() == "Cat" {
		fmt.Println("it is a cat")
	}


	//直接调用方法:指针可以直接调用
	v.MethodByName("Speak").Call([]reflect.Value{reflect.ValueOf("test3")})
	//直接访问字段:指针不可直接访问，必须剥离
	if nameField := baseValue.FieldByName("Name"); nameField.IsValid() {
		fmt.Println(nameField.String() + " is speaking test4")
	}
}
