package main

import "testing"

func TestSolveFakeTest1(t *testing.T) {
	t.Run("fake test", func(t *testing.T) {
		HelloWorld()
		//t.Errorf("fail")
	})
}

func TestSolveFakeTest2(t *testing.T) {
	t.Run("fake test", func(t *testing.T) {
		main()
		t.Errorf("fail")
	})
}