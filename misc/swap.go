package main

func Swap(x, y *int) { 
    *x, *y = *y, *x
}
