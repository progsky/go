package main

import (
	"fmt"
	"os"
	"regexp"
	"strings"
)

func main() {
	if len(os.Args) != 2 {
		fmt.Fprintf(os.Stderr, "you should specify one argument - knight place\n")
		os.Exit(-1)
	}
	knightPlace := os.Args[1]
	rightPlaceFormat, _ := regexp.MatchString("^[a-h][1-8]$", knightPlace)
	if !rightPlaceFormat {
		fmt.Fprintf(os.Stderr, "wrong place\n")
		os.Exit(-2)
	}
	knightMoves := knightMoves(knightPlace)
	fmt.Println(strings.Join(knightMoves, "\n"))
}

func knightMoves(place string) []string {
	col, row := decodePlace(place)
	places := make([]string, 0)
	if col-2 > 0 && row+1 < 9 {
		places = append(places, encodePlace(col-2, row+1))
	}
	if col-2 > 0 && row-1 > 0 {
		places = append(places, encodePlace(col-2, row-1))
	}
	if col-1 > 0 && row-2 > 0 {
		places = append(places, encodePlace(col-1, row-2))
	}
	if col+1 < 9 && row-2 > 0 {
		places = append(places, encodePlace(col+1, row-2))
	}
	if col+2 < 9 && row-1 > 0 {
		places = append(places, encodePlace(col+2, row-1))
	}
	if col+2 < 9 && row+1 < 9 {
		places = append(places, encodePlace(col+2, row+1))
	}
	if col+1 < 9 && row+2 < 9 {
		places = append(places, encodePlace(col+1, row+2))
	}
	if col-1 > 0 && row+2 < 9 {
		places = append(places, encodePlace(col-1, row+2))
	}
	return places
}

func decodePlace(place string) (col int, row int) {
	return int(place[0] - 96), 9 - int(place[1]-48)
}

func encodePlace(col int, row int) string {
	return string([]byte{byte(col + 96), byte(9 - row + 48)}[:])
}
