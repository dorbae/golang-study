//
// main.go
// package main: Complile Entry Point
// package ~= Java package + class
//package main

package main

import (
	"fmt"
	"strings"
)

//
// Function 만들기
//
// 파라미터 변수 타입 정의 필요
// 리턴 타입도 지정
// func <func_name> (<p1 name> <p1 type>, <p2 name> <p2 type> ...) <return_type> {}
// 파라미터 변수 타입이 같으면 하나로 정의 가능
// func <func_name> (<p1, <p2> <type>) <return_type> {}
//func multiply(a int, b int) int {
func multiply(a, b int) int {
	return a * b
}

//
// Multiple Return values Function
//
// GO는 여러 타입의 다중 값을 리턴 가능 very sexy
func lenAndToUpper(name string) (int, string) {
	return len(name), strings.ToUpper(name)
}

//
// Naked Return values Function
//
func lenAndToUpperNaked(name string) (length int, uppercae string) {
	length = len(name)
	uppercae = strings.ToUpper(name)
	return
}

//
// defer
//
// Function return 후 실행
// Resource Close 등으로 사용
func lenAndToUpperDefer(name string) (int, string) {
	defer fmt.Println("I'm done!!")
	return len(name), strings.ToUpper(name)
}

//
// Loop: for
//
// GO는 for문만 존재. foreach, for in 이런거 없음
func superAdd(numbers ...int) int {
	fmt.Print("[DEBUG] numbers=")
	fmt.Println(numbers)

	total := 0
	// Index값도 확인 가능
	for index, number := range numbers {
		fmt.Printf("[DEBUG] index=%d, num=%d\n", index, number)
	}

	for i := 0; i < len(numbers); i++ {
		//
	}

	for number := range numbers {
		//
		total += number
	}

	return total
}

//
// if/else
//
// GO는 if문에서 조건 비교 전에 변수 설정 가능
// if문에서만 쓸 변수임을 명시적으로 표현 가능
func canDrink(age int) bool {
	// if age < 18 {
	// 	return false
	// } else if age > 120 {
	// 	return true
	// } else {
	// 	return true
	// }

	if koreanAge := age + 2; koreanAge < 18 {
		return false
	} else if koreanAge > 120 {
		return true
	} else {
		return true
	}
}

//
// Switch
//
// GO는 case절에서 조건문을 사용 가능
func canSmoke(age int) bool {
	// Basic
	switch age {
	case 10:
		//
	case 20:
		//
	default:
		//
	}

	// Switch If/Else문
	switch {
	case age < 18:
		//
	case age == 18:
		//
	case age > 120:
		//
	}

	switch koreanAge := age + 2; koreanAge {
	case 10:
		return false
	case 20:
		return true
	default:
		return false
	}
}

//
// Struct
//
// Method 포함 가능
// Constructor 개념 없음
type person struct {
	name    string
	age     int
	favFood []string
}

//
// Entry function
//
func main() {
	fmt.Println("Hello World")

	//
	// Package 개념 & Import 모듈 사용
	//
	// Import된 func 사용 시, func명 대문자로 시작
	// something.SayHello()
	// something.SayBye()

	//
	// 변수
	//
	// 타입 없는 상수
	const nameConstNoType = "nameConstNoType"
	fmt.Println(nameConstNoType)
	// 타입 지정 상수
	const nameConstTypeDefined string = "nameConstTypeDefined"
	fmt.Println(nameConstTypeDefined)
	// 타입 없는 변수
	var nameVarNoType = "nameVarNoType"
	fmt.Println(nameVarNoType)
	// 타입 지정 변수
	var nameVarTypeDefined string = "nameVarTypeDefined"
	fmt.Println(nameVarTypeDefined)
	nameVarTypeDefined = "nameVarTypeDefinedChanged"
	fmt.Println(nameVarTypeDefined)
	// 변수 축약형 (GO가 자동으로 타입 매핑)
	name := "name" // var name string = "name"
	fmt.Println(name)

	fmt.Println()

	// Call Function
	fmt.Println(multiply(2, 3))

	// Call multiple return values function
	totalLength, upperName := lenAndToUpper("myName")
	fmt.Println(totalLength, upperName)
	fmt.Println(lenAndToUpper("myName"))

	// 다중 리턴 값 무시 Binding
	// 리턴 값 개수에 맞게 변수 정의 필요. But, _ 로 무시 가능
	// totalLength := lenAndToUpper("myName") -> Compile Error
	totalLength2, _ := lenAndToUpper("myName2")
	fmt.Println(totalLength2)

	// Naked return
	fmt.Println(lenAndToUpperNaked("abcde"))

	// Defer
	fmt.Println(lenAndToUpperDefer("Defer"))

	// For loop
	total := superAdd(1, 2, 3, 4, 5, 6)
	fmt.Printf("[INFO ] total=%d\n", total)

	// if/else
	age := 14
	fmt.Printf("[INFO ] Can someone who is %d years old drink? -> %t\n", age, canDrink(age))

	// Switch
	fmt.Printf("[INFO ] Can someone who is %d years old smoke? -> %t\n", age, canSmoke(age))

	//
	// Pointer
	//
	a := 2
	b := a
	a = 10
	fmt.Println(a, b)
	fmt.Println(&a, &b)

	// 주소 참조
	c := 2
	cAddress := &c
	fmt.Println(c, cAddress, *cAddress)

	// 주소 참조 -> 값 변경
	d := 2
	dAddress := &d
	*dAddress = 20
	fmt.Println(d)

	//
	// Arrays
	//
	names := [5]string{"a", "b", "c", "d", "e"}
	names[3] = "ddd"
	// names[5] = "ffff" => OutOfBoundary
	fmt.Println(names)

	//
	// Slice
	//
	// 길이 제한 없는 Array
	namesForSlice := []string{"a", "b", "c"}
	fmt.Println(namesForSlice)

	// append는 arg array값을 변형시키지 않음. 리턴 받아야함
	namesForSlice = append(namesForSlice, "def")
	fmt.Println(namesForSlice)

	//
	// Maps
	//
	// map[<key_type>]<value_type>
	// Object 이용하고 싶으면 Struct(Class) 활용
	dorbae := map[string]string{"name": "dorbae", "age": "32"}
	fmt.Println(dorbae)

	for key, value := range dorbae {
		fmt.Printf("key=%s, value=%s\n", key, value)
	}

	for _, value := range dorbae {
		fmt.Printf("value=%s\n", value)
	}

	//
	// Struct
	//
	favFood := []string{"김치", "ramen"}
	// person1 := person{"dorbae", 32, favFood}
	// Labeling
	person1 := person{name: "dorbae", age: 32, favFood: favFood}
	fmt.Println(person1.name)
}
