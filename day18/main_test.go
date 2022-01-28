package main

import (
	"testing"
)

func TestExplode1(t *testing.T) {
	s := "[[[[[9,8],1],2],3],4]"
	expected := "[[[[0,9],2],3],4]"
	r, _ := Explode(s)
	if r != expected {
		t.Errorf("Explode(%s) = %s, wanted %s", s, r, expected)
	}
}

func TestExplode2(t *testing.T) {
	s := "[7,[6,[5,[4,[3,2]]]]]"
	expected := "[7,[6,[5,[7,0]]]]"
	r, _ := Explode(s)
	if r != expected {
		t.Errorf("Explode(%s) = %s, wanted %s", s, r, expected)
	}
}

func TestExplode3(t *testing.T) {
	s := "[[6,[5,[4,[3,2]]]],1]"
	expected := "[[6,[5,[7,0]]],3]"
	r, _ := Explode(s)
	if r != expected {
		t.Errorf("Explode(%s) = %s, wanted %s", s, r, expected)
	}
}

func TestExplode4(t *testing.T) {
	s := "[[3,[2,[1,[7,3]]]],[6,[5,[4,[3,2]]]]]"
	expected := "[[3,[2,[8,0]]],[9,[5,[4,[3,2]]]]]"
	r, _ := Explode(s)
	if r != expected {
		t.Errorf("Explode(%s) = %s, wanted %s", s, r, expected)
	}
}

func TestExplode5(t *testing.T) {
	s := "[[3,[2,[8,0]]],[9,[5,[4,[3,2]]]]]"
	expected := "[[3,[2,[8,0]]],[9,[5,[7,0]]]]"
	r, _ := Explode(s)
	if r != expected {
		t.Errorf("Explode(%s) = %s, wanted %s", s, r, expected)
	}
}

func TestSplit1(t *testing.T) {
	s := "[[[[0,7],4],[15,[0,13]]],[1,1]]"
	expected := "[[[[0,7],4],[[7,8],[0,13]]],[1,1]]"
	r, _ := Split(s)
	if r != expected {
		t.Errorf("Split(%s) = %s, wanted %s", s, r, expected)
	}
}

func TestSplit2(t *testing.T) {
	s := "[[[[0,7],4],[[7,8],[0,13]]],[1,1]]"
	expected := "[[[[0,7],4],[[7,8],[0,[6,7]]]],[1,1]]"
	r, _ := Split(s)
	if r != expected {
		t.Errorf("Split(%s) = %s, wanted %s", s, r, expected)
	}
}

func TestAdd1(t *testing.T) {
	s1 := "[[[[4,3],4],4],[7,[[8,4],9]]]"
	s2 := "[1,1]"
	expected := "[[[[0,7],4],[[7,8],[6,0]]],[8,1]]"
	r := Add(s1, s2)
	if r != expected {
		t.Errorf("Add(%s + %s) = %s, wanted %s", s1, s2, r, expected)
	}
}

func TestMagnitude1(t *testing.T) {
	s := "[[1,2],[[3,4],5]]"
	expected := 143
	r := Magnitude(s)
	if r != expected {
		t.Errorf("Magnitude(%s) = %d, wanted %d", s, r, expected)
	}
}

func TestMagnitude2(t *testing.T) {
	s := "[[[[0,7],4],[[7,8],[6,0]]],[8,1]]"
	expected := 1384
	r := Magnitude(s)
	if r != expected {
		t.Errorf("Magnitude(%s) = %d, wanted %d", s, r, expected)
	}
}

func TestMagnitude3(t *testing.T) {
	s := "[[[[1,1],[2,2]],[3,3]],[4,4]]"
	expected := 445
	r := Magnitude(s)
	if r != expected {
		t.Errorf("Magnitude(%s) = %d, wanted %d", s, r, expected)
	}
}

func TestMagnitude4(t *testing.T) {
	s := "[[[[3,0],[5,3]],[4,4]],[5,5]]"
	expected := 791
	r := Magnitude(s)
	if r != expected {
		t.Errorf("Magnitude(%s) = %d, wanted %d", s, r, expected)
	}
}

func TestMagnitude5(t *testing.T) {
	s := "[[[[5,0],[7,4]],[5,5]],[6,6]]"
	expected := 1137
	r := Magnitude(s)
	if r != expected {
		t.Errorf("Magnitude(%s) = %d, wanted %d", s, r, expected)
	}
}

func TestMagnitude6(t *testing.T) {
	s := "[[[[8,7],[7,7]],[[8,6],[7,7]]],[[[0,7],[6,6]],[8,7]]]"
	expected := 3488
	r := Magnitude(s)
	if r != expected {
		t.Errorf("Magnitude(%s) = %d, wanted %d", s, r, expected)
	}
}
