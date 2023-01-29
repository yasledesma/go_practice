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

func SumAllTails(arraysOfNumbers ...[]int) []int {
    var sum []int
    for _, arrayOfNumbers := range arraysOfNumbers {
       if len(arrayOfNumbers) > 1 {
           sum = append(sum, Sum(arrayOfNumbers[1:]))
       } else {
           sum = append(sum, 0)
       } 
    }
    return sum
}
