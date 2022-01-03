package main

import "fmt"

type Location struct {
	Lat       float64 `json:"lat"`
	Lng       float64 `json:"lng"`
	Timestamp int64   `json:"timestamp"`
}

func main() {
	var locs []Location
	sectionLength := 4
	locs = append(locs, Location{Lat: 1.0, Lng: 2.0, Timestamp: 3}, Location{Lat: 4.0, Lng: 5.0, Timestamp: 6}, Location{Lat: 7.0, Lng: 8.0, Timestamp: 9}, Location{Lat: 1.0, Lng: 2.0, Timestamp: 3}, Location{Lat: 4.0, Lng: 5.0, Timestamp: 6}, Location{Lat: 7.0, Lng: 8.0, Timestamp: 9})
	slices := SplitSlice(locs, sectionLength)
	fmt.Println(slices)
}

//  Split a slice into multiple slices of a given length. last slice may be smaller than the given length.
func SplitSlice(slice []Location, length int) [][]Location {
	var result [][]Location
	for i := 0; i < len(slice); i += length {
		end := i + length
		if end > len(slice) {
			end = len(slice)
		}
		result = append(result, slice[i:end])
	}
	return result
}
