package main

import (
	"encoding/csv"
	"fmt"
	"io"
	"math/big"
	"os"
	"sort"

	"strconv"
	"strings"
)

func main() {
	problem24()

}

func problem24() {
	/*
		A permutation is an ordered arrangement of objects. For example, 3124 is one possible permutation of the digits 1, 2, 3 and 4. If all of the permutations are listed numerically or alphabetically, we call it lexicographic order. The lexicographic permutations of 0, 1 and 2 are:
									012   021   102   120   201   210
		What is the millionth lexicographic permutation of the digits 0, 1, 2, 3, 4, 5, 6, 7, 8 and 9?
	*/
	// https://www.nayuki.io/page/next-lexicographical-permutation-algorithm
	array := []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}
	for i := 1; i <= 1000000; i++ {
		fmt.Println(array)
		nextPermutation(array)
	}
}

func nextPermutation(array []int) bool {
	//Find longest non-increasing suffix
	i := len(array) - 1
	for i > 0 && array[i-1] >= array[i] {
		//Now i is the head index of the suffix
		i--
	}
	//Are we at  the last permutation already?
	if i <= 0 {
		return false
	}
	//Let array[i - 1] be the pivot
	//Find rightmost element that exceeds the pivot
	j := len(array) - 1
	for array[j] <= array[i-1] {
		j--
	}
	//Now the value array[j] will become the new pivot
	//Assertion: j >= i

	//Swap the pivot with j
	temp := array[i-1]
	array[i-1] = array[j]
	array[j] = temp

	//Reverse the suffix
	j = len(array) - 1
	for i < j {
		temp = array[i]
		array[i] = array[j]
		array[j] = temp
		i++
		j--

	}
	//Successfully computed the next permutation
	return true
}

func problem23() {
	/*
		A perfect number is a number for which the sum of its proper divisors is exactly equal to the number. For example, the sum of the proper divisors of 28 would be 1 + 2 + 4 + 7 + 14 = 28, which means that 28 is a perfect number.
		A number n is called deficient if the sum of its proper divisors is less than n and it is called abundant if this sum exceeds n.
		As 12 is the smallest abundant number, 1 + 2 + 3 + 4 + 6 = 16, the smallest number that can be written as the sum of two abundant numbers is 24. By mathematical analysis, it can be shown that all integers greater than 28123 can be written as the sum of two abundant numbers. However, this upper limit cannot be reduced any further by analysis even though it is known that the greatest number that cannot be expressed as the sum of two abundant numbers is less than this limit.
		Find the sum of all the positive integers which cannot be written as the sum of two abundant numbers. */
	MAX := 28123
	abundandNumbers := []int{}
	result := [28124]bool{false}
	for i := 12; i < 28123; i++ {
		if IsAbundand(i) {
			abundandNumbers = append(abundandNumbers, i)
		}
	}

	for i := 0; i < len(abundandNumbers); i++ {
		for j := 0; j < len(abundandNumbers); j++ {
			sum := abundandNumbers[i] + abundandNumbers[j]
			if sum > MAX {
				break
			}
			result[sum] = true
		}
	}
	sumOfResult := 0
	for i, v := range result {
		if !v {
			sumOfResult += i
		}
	}
	fmt.Println(sumOfResult)
}

//IsAbundand returns true if the number is abundand
func IsAbundand(n int) bool {
	divs := getDivisors(n)
	sum := sumSlice(&divs)
	return sum > n
}

func problem22() {

	// Using names.txt (right click and 'Save Link/Target As...'), a 46K text file containing over five-thousand first names, begin by sorting it into alphabetical order. Then working out the alphabetical value for each name, multiply this value by its alphabetical position in the list to obtain a name score.
	// For example, when the list is sorted into alphabetical order, COLIN, which is worth 3 + 15 + 12 + 9 + 14 = 53, is the 938th name in the list. So, COLIN would obtain a score of 938 × 53 = 49714.
	// What is the total of all the name scores in the file?

	file, err := os.Open("p022_names.txt")
	if err != nil {
		fmt.Println("Error:", err)
		return
	}
	defer file.Close()
	reader := csv.NewReader(file)
	//names := []string{}
	var names []string
	for {
		name, err := reader.Read()
		if err == io.EOF {
			break
		} else if err != nil {
			fmt.Println("Error:", err)
			return
		}
		names = append(names, name...)
	}

	sort.Slice(names, func(i, j int) bool { return names[i] < names[j] })

	var numVal []int
	for i, name := range names {
		numVal = append(numVal, getNumericVal(name))
		numVal[i] *= (i + 1)
	}
	sumOfNumVal := sumSlice(&numVal)
	fmt.Println(sumOfNumVal)
}

func getNumericVal(s string) int {
	val := 0
	for i := range s {
		switch s[i] {
		case 'A':
			val++
			break
		case 'B':
			val += 2
			break
		case 'C':
			val += 3
			break
		case 'D':
			val += 4
			break
		case 'E':
			val += 5
			break
		case 'F':
			val += 6
			break
		case 'G':
			val += 7
			break
		case 'H':
			val += 8
			break
		case 'I':
			val += 9
			break
		case 'J':
			val += 10
			break
		case 'K':
			val += 11
			break
		case 'L':
			val += 12
			break
		case 'M':
			val += 13
			break
		case 'N':
			val += 14
			break
		case 'O':
			val += 15
			break
		case 'P':
			val += 16
			break
		case 'Q':
			val += 17
			break
		case 'R':
			val += 18
			break
		case 'S':
			val += 19
			break
		case 'T':
			val += 20
			break
		case 'U':
			val += 21
			break
		case 'V':
			val += 22
			break
		case 'W':
			val += 23
			break
		case 'X':
			val += 24
			break
		case 'Y':
			val += 25
			break
		case 'Z':
			val += 26
			break
		}
	}
	return val
}

func problem21() {
	sumOfAll := 0

	for index := 1; index < 10000; index++ {
		divsA := getDivisors(index)
		sumA := sumSlice(&divsA)
		divsB := getDivisors(sumA)
		sumB := sumSlice(&divsB)

		if index == sumB && index != sumA {
			fmt.Printf("Amicas Index: %d SumA: %d SumB: %d\n", index, sumA, sumB)
			sumOfAll += index
		}
	}
	fmt.Printf("Sum: %d", sumOfAll)
}

func sumSlice(slice *[]int) (number int) {
	sum := 0
	for _, s := range *slice {
		sum += s
	}
	return sum
}

func getDivisors(num int) (divisors []int) {
	for index := 1; index < (num/2)+1; index++ {
		if num%index == 0 {
			divisors = append(divisors, index)
		}
	}
	return divisors
}

func problem20() {
	result := big.NewInt(1)
	counter := big.NewInt(100)
	bigOne := big.NewInt(1)
	for index := 100; index > 1; index-- {
		result.Mul(result, counter)
		counter.Sub(counter, bigOne)
	}

	fmt.Printf("%d \n", result)

	stringResult := result.String()
	sum := 0
	numbers := strings.Split(stringResult, "")
	for _, number := range numbers {
		x, _ := strconv.Atoi(number)
		sum += x
	}

	fmt.Printf("%d \n", sum)
}
