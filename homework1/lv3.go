package main

import "fmt"

func romanToInt(roman string) int {
	sum := 0
	nums := make([]int, len(roman))
	maps := map[rune]int{
		'I': 1,
		'V': 5,
		'X': 10,
		'L': 50,
		'C': 100,
		'D': 500,
		'M': 1000,
	}
	for i, romanNum := range roman {
		nums[i] = maps[romanNum]
	}
	fmt.Println(nums)
	for i, value := range nums {
		if i < len(nums)-1 {
			if value*5 == nums[i+1] || value*10 == nums[i+1] {
				sum -= value
			} else {
				sum += value
			}
		} else {
			sum += value
		}
	}
	return sum
}
func main() {
	fmt.Println(romanToInt("III"))
	fmt.Println(romanToInt("MCMXCIV"))
	fmt.Println(romanToInt("LVIII"))
	fmt.Println(romanToInt("IX"))

}
