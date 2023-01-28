package main

func Sum(numbers []int) (sum int) {
    for _, number := range numbers {
        sum += number
    }
    return
}

func SumAll(arraysOfNumbers ...[]int) []int {
    var sum []int
    for _, arrayOfNumbers := range arraysOfNumbers {
        sum = append(sum, Sum(arrayOfNumbers))
    }
    return sum
}
