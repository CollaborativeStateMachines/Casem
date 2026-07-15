package csm

import "testing"

func TestNewContextVariable(t *testing.T) {
	_, err := New("a", 1)
	if err != nil {
		t.Errorf("Expected no error but received: %v", err)
	}
	_, err = New("", 1)
	if err == nil {
		t.Errorf("Expected error but received: %v", err)
	}
}
