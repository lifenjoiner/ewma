package ewma

import (
	"math"
	"math/rand"
	"testing"
)

func TestEWMA(t *testing.T) {
	var ewma *EWMA
	// new
	slide := rand.Int()
	ewma = NewMovingAverage(slide)
	if ewma.slide != slide {
		t.Errorf("slide is %v != %v", ewma.slide, slide)
	}
	// count
	slide = 10
	ewma = NewMovingAverage(slide)
	for i := 0; i < 2 * slide; i++ {
		ewma.Add(1)
		if ewma.Value() != 1 {
			t.Errorf("Value() is NOT 1: %v", ewma.Value())
		}
		if i <= slide {
			if ewma.count != i + 1 {
				t.Errorf("%v: count = %v", i, ewma.count)
			}
		} else {
			if ewma.count <= slide {
				t.Errorf("%v: count = %v", i, ewma.count)
			}
		}
	}
	// Set
	slide = 10
	ewma = NewMovingAverage(slide)
	v := rand.Float64()
	ewma.Set(v)
	if ewma.value != v || ewma.count <= slide {
		t.Errorf("Set(%v): value = %v, count = %v", v, ewma.value, ewma.count)
	}
	// Preset samples, Value()
	var tests = [10]float64 {334, 373, 82, 141, 38, 443, 442, 397, 55, 60}
	var ewmas = [10]float64 {334, 341, 294, 266, 225, 264, 297, 315, 268, 230}
	ewma = NewMovingAverage(10)
	ewma.Set(tests[0])
	for i := 0; i < len(tests); i++ {
	    ewma.Add(tests[i])
	    if math.Round(ewma.Value()) != ewmas[i] {
	        t.Errorf("%v: value = %v, expected = %v", i, ewma.Value(), ewmas[i])
	    }
	}
}
