// Prüfdatei für das Morning-Theme: Kommentar (Comment, blau).
// TODO: Todo-Markierung, FIXME, XXX
package main

import (
	"errors"
	"fmt"
	"math"
	"os"
	"strings"
)

const (
	maxItems = 16      // Zahl (Constant, magenta)
	ratio    = 3.14159 // Float
	hexMask  = 0xFF
	label    = "Text\tmit \"Escapes\" und %d" // String (Constant, magenta)
	letter   = 'x'                            // Character
)

var ErrEmpty = errors.New("leer")

// Shape ist ein Interface (Type, grün, fett).
type Shape interface {
	Area() float64
	Name() string
}

type Point struct {
	X, Y float64 `json:"x"` // Struct-Tag
}

type Circle struct {
	Center Point
	R      float64
}

func (c Circle) Area() float64 { return math.Pi * c.R * c.R }
func (c Circle) Name() string  { return "circle" }

func sum[T int | float64](xs ...T) (total T) {
	for _, x := range xs {
		total += x
	}
	return
}

func describe(s Shape) (string, error) {
	if s == nil {
		return "", ErrEmpty
	}
	switch v := s.(type) {
	case Circle:
		return fmt.Sprintf("%s r=%.1f", v.Name(), v.R), nil
	default:
		return s.Name(), nil
	}
}

func main() {
	shapes := []Shape{Circle{Point{0, 0}, 2}, Circle{R: 1.5}}
	counts := make(map[string]int)
	ch := make(chan int, 2)

	for i, s := range shapes {
		if a := s.Area(); a > 1 && a != math.Inf(1) || i == 0 {
			counts[s.Name()]++
		}
	}

	go func() { ch <- 1 << 3; close(ch) }()
	defer fmt.Println("fertig", <-ch)

	if text, err := describe(shapes[0]); err != nil {
		fmt.Fprintln(os.Stderr, "Fehler:", err)
		os.Exit(1)
	} else {
		fmt.Println(strings.ToUpper(text), counts, sum(1, 2, 3), nil, true)
	}
}
