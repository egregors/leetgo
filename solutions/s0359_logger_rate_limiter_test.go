package solutions

import (
	"testing"
)

func TestNewLogger(t *testing.T) {
	//	Example 1:
	//
	//	Input
	// ["Logger", "shouldPrintMessage", "shouldPrintMessage", "shouldPrintMessage", "shouldPrintMessage", "shouldPrintMessage", "shouldPrintMessage"]
	//	[[], [1, "foo"], [2, "bar"], [3, "foo"], [8, "bar"], [10, "foo"], [11, "foo"]]
	//	Output
	//	[null, true, true, false, false, false, true]
	//
	//	Explanation
	//	Logger logger = new Logger();
	//	logger.shouldPrintMessage(1, "foo");  // return true, next allowed timestamp for "foo" is 1 + 10 = 11
	//	logger.shouldPrintMessage(2, "bar");  // return true, next allowed timestamp for "bar" is 2 + 10 = 12
	//	logger.shouldPrintMessage(3, "foo");  // 3 < 11, return false
	//	logger.shouldPrintMessage(8, "bar");  // 8 < 12, return false
	//	logger.shouldPrintMessage(10, "foo"); // 10 < 11, return false
	//	logger.shouldPrintMessage(11, "foo"); // 11 >= 11, return true, next allowed timestamp for "foo" is 11 + 10 = 21
	logger := NewLogger()
	if !logger.ShouldPrintMessage(1, "foo") {
		t.Errorf("Expected true, got false")
	}
	if !logger.ShouldPrintMessage(2, "bar") {
		t.Errorf("Expected true, got false")
	}
	if logger.ShouldPrintMessage(3, "foo") {
		t.Errorf("Expected false, got true")
	}
	if logger.ShouldPrintMessage(8, "bar") {
		t.Errorf("Expected false, got true")
	}
	if logger.ShouldPrintMessage(10, "foo") {
		t.Errorf("Expected false, got true")
	}
	if !logger.ShouldPrintMessage(11, "foo") {
		t.Errorf("Expected true, got false")
	}
}
