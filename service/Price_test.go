package service

import "testing"

func TestPrice(t *testing.T ){
	var expected int = 14
	var actual = Price(1,2,3,4)
	if actual != expected{
		t.Errorf("expected=%d,actual=%d",expected,actual)
	}
}