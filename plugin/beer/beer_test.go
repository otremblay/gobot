package beer

import (
	"testing"
	"time"
)

func TestBeerPluginName(t *testing.T) {
	p := Beer{}
	expected := "Beer v1.0"
	name := p.Name()
	if name != expected {
		t.Errorf("Test: '%s' does not equal '%s'", name, expected)
	}

}

func TestItsTimeForBeer(t *testing.T) {
	// reply should be in ayes slice
	time := time.Date(2015, time.October, 30, 23, 0, 0, 0, time.UTC)
	reply := beer(time)
	for _, v := range ayes {
		if v == reply {
			return
		}
	}
	t.Errorf("Test: '%s' not found in ayes slice: %v", reply, ayes)
}

func TestNotTimeForBeer(t *testing.T) {
	// reply should be in nays slice
	time := time.Date(2009, time.November, 10, 23, 0, 0, 0, time.UTC)
	reply := beer(time)
	for _, v := range nays {
		if v == reply {
			return
		}
	}
	t.Errorf("Test: '%s' not found in nays slice: %v", reply, nays)
}

func TestChristmasEve(t *testing.T) {
	time := time.Date(2009, time.December, 24, 23, 0, 0, 0, time.UTC)
	reply := beer(time)
	expected := "Merry Beermas!"
	if reply == expected {
		return
	}
	t.Errorf("Test: '%s' does not equal %s", reply, expected)
}
