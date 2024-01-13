package main

func loop(i int) int {
    idx := i
    for idx < 10 {
        if idx < 5 {
            idx++
        } else {
            idx += 2
        }
    }
    return 0
}

func main() {}
