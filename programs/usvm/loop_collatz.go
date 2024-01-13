package main

func loop(i int) int {
  if i <= 0 || i >= 100 {
    return 0
  }

  j := i
  loopCount := 0

  for j != 1 {
    oldJ := j
    j = j * 3 + 1
    if oldJ % 2 == 0 {
      j = oldJ / 2
    }
    loopCount++
  }

  if loopCount == 17 {
    return 1
  }
  return 2
}

func main() {}

