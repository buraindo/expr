package main

func loop(i, j int) int {
  sum := 0
  for k := 0; k < 100; k++ {
    for l := 0; l < 100; l++ {
      if k+l < i+j {
        sum += k + l
      }
    }
  }
  return sum
}

func main() {}

