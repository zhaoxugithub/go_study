package main

import "testing"

/**
 *ClassName triangle_test
 *Description TODO
 *Author 11931
 *Date 2023-07-07:1:00
 *Version 1.0
 **/

func testTriangle(t *testing.T) {
	test := []struct{ a, b, c int }{
		{3, 4, 5},
		{5, 12, 13},
		{8, 15, 17},
		{12, 35, 37},
		{30000, 40000, 50000},
	}

	for _, tt := range test {
		if actual := calcTriangle(tt.a, tt.b); actual != tt.c {
			t.Errorf("calcTringle(%d,%d); "+
				"got %d; expected %d", tt.a, tt.b, actual, tt.c,
			)
		}
	}
}

func main() {
	//testTriangle()
}
