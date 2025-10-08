package helpers

import "testing"

func TestContainsValue(t *testing.T) {
    testMap := map[string]string{"key":"value"}
    want := true
    received := ContainsValue(testMap,"value")

    if received != want{
        t.Errorf("Expected %v, Received %v", want, received)
    }
}