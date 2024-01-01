package main

import "testing"

func TestPart1Example(t *testing.T) {
	filename := "./test_input"
	expected := 142
	result := getCalibrationPart1(filename)

	if result != expected {
		t.Errorf("Expected %d, got %d", expected, result)
	}
}

func TestPart1Input(t *testing.T) {
	filename := "./input"
	expected := 54708
	result := getCalibrationPart1(filename)

	if result != expected {
		t.Errorf("Expected %d, got %d", expected, result)
	}
}

func TestPart2Example(t *testing.T) {
	filename := "./test_input_2"
	expected := 281
	result := getCalibrationPart2(filename)

	if result != expected {
		t.Errorf("Expected %d, got %d", expected, result)
	}
}
func TestPart2Input(t *testing.T) {
	filename := "./input"
	expected := 281
	result := getCalibrationPart2(filename)

	if result != expected {
		t.Errorf("Expected %d, got %d", expected, result)
	}
}
