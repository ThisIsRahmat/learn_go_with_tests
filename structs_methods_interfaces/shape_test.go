package main

import "testing"



func TestPerimeter(t *testing.T) {


	//t.run makes a subtest  and this is to test for perimeters of rectangles
	t.Run("rectangles", func(t *testing.T) {
		rectangle := Rectangle{12, 6}
		got := rectangle.Area()
		want := 72.0

		if got != want {
			t.Errorf("got %g want %g", got, want)
		}
	})

	//test for circle 

	t.Run("circles", func(t *testing.T) {

		//why do we need circle??? first
		circle := Circle{10}

		got := circle.Area()
		want := 314.1592653589793


		if got != want {
			t.Errorf("got %g want %g", got , want)
		}

	})

}


func TestArea(t *testing.T) {



	t.Run("rectangles", func (t *testing.T)  {

		rectangle := Rectangle{12.0, 6.0}
	got := rectangle.Area()
	want := 72.0

	
	if got != want {

		//f is for our float64 and .2 means print 2 decimal places 
		t.Errorf("got %2.f want %2.f", got, want)
	}
		
	})

	t.Run("circles", func(t *testing.T) {

		circle := Circle{10}

		got := circle.Area()
		want := 314.1592653589793

		if got != want {
			t.Errorf("got %g want %g", got, want)
		}
	})
	
}