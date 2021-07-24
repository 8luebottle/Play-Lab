package main

import "fmt"

func main() {
	splitLine()
	// [NOTE] Different output results
	// Interface 空接口 | Slice 切片
	var colors = []interface{}{"BLACK", "GOLD", "PINK"}

	// Parameter 参数
	fmt.Println(colors)          // [BLACK GOLD PINK]
	fmt.Printf("%#v \n", colors) // []interface {}{"BLACK", "GOLD", "PINK"}

	// ...notation 切片展开
	fmt.Println(colors...) // BLACK GOLD PINK
	for _, color := range colors {
		fmt.Printf("%#v \n", color)
		/*
		   "BLACK"
		   "GOLD"
		   "PINK"
		*/
	}

	splitLine()
	// [NOTE] Can not adjustable 无法修改
	years := [3]int{2019, 2020, 2021}

	// Function 函数
	func(ls [3]int) {
		ls[0] = 2023
		fmt.Println(ls) // [2023 2020 2021]
	}(years)
	// Outside the scope
	fmt.Println(years) // [2019 2020 2021]

	years[0] = 2023
	fmt.Println(years) // [2023 2020 2021]

	splitLine()
	// [NOTE] Map iteration order is random
	type Year int
	const (
		Year2020 Year = iota + 2020
		Year2021
		Year2022
		Year2023
	)

	ages := map[Year]int{ // map randomness:  map中的key无序
		Year2020: 19,
		Year2021: 20,
		Year2022: 21,
		Year2023: 22,
	}

	for year, age := range ages {
		fmt.Println(year, age)
	}
}

func splitLine() {
	fmt.Printf("\n************************************ [CASE STUDY] ************************************\n")
}
