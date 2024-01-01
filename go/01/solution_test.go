package main

import "testing"

func TestExample1(t *testing.T) {
	filename := "./test_input"
	expected := 142
	result := getCalibrationPart1(filename)

	if result != expected {
		t.Errorf("Expected %d, got %d", expected, result)
	}
}

func TestInput(t *testing.T) {
	filename := "./input"
	expected := 54708
	result := getCalibrationPart1(filename)

	if result != expected {
		t.Errorf("Expected %d, got %d", expected, result)
	}
}
