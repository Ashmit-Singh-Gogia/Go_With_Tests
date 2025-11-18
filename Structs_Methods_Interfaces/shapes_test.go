package main

import (
	"testing"
)

func TestArea(t *testing.T) {

	areaTests := []struct {
		shape Shape
		want  float64
	}{
		{shape: Rectangle{width: 12, height: 6}, want: 72.0},
		{shape: Circle{radius: 10}, want: 314.1592653589793},
		{shape: Triangle{Base: 10, Height: 10}, want: 50.0},
	}
	for _, tt := range areaTests {
		got := tt.shape.Area()
		if got != tt.want {
			t.Errorf("Test for struct %#v got %g want %g", tt.shape, got, tt.want)
		}
	}

}
