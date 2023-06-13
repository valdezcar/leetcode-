package main

// https://leetcode-cn.com/problems/kids-with-the-greatest-number-of-candies/

func kidsWithCandies(candies []int, extraCandies int) []bool {
	max := 0
	for _, candy := range candies {
		max = maxFunc(max, candy)
	}
	result := make([]bool, len(candies))

	for i, candy := range candies {
		result[i] = candy+extraCandies >= max
	}

	return result
}

func maxFunc(i, j int) int {
	if i > j {
		return i
	}
	return j
}

func main() {

}
