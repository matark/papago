package lolcat

import "fmt"
import "math"

type Color struct {
  Red   int
  Green int
  Blue  int
}

func Rainbow(data []rune) {
  for i, value := range data {
    color := makeColor(0.25, (200 + i) / 8.0)
    fmt.Printf("\033[38;2;%d;%d;%dm%c\033[0m", color.Red, color.Green, color.Blue, value)
  }
}

func makeColor(freq float64, i int) Color {
  var color Color

  color.Red = int(math.Sin(freq * float64(i) + 0) * 127 + 128)
  color.Blue = int(math.Sin(freq * float64(i) + 4 * math.Pi / 3) * 127 + 128)
  color.Green = int(math.Sin(freq * float64(i) + 2 * math.Pi / 3) * 127 + 128)

  return color
}
