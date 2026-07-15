package csm

import "testing"

func TestNewContextVariable(t *testing.T) {
	_, err := NewContextVariable("a", 1)
	if err != nil {
		t.Errorf("Expected no error but received: %v", err)
	}
	_, err = NewContextVariable("", 1)
	if err == nil {
		t.Errorf("Expected error but received: %v", err)
	}
}
