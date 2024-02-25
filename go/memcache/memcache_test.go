package memcache

import (
	"fmt"
	"testing"
)

func TestInit(t *testing.T) {
	testCases := []struct {
		memoryLimit int
		algorithm   string
	}{
		{100, "FIFO"},
		{0, "FIFO"},
		{100, "LFU"},
	}

	for _, tc := range testCases {
		err := Initialize(tc.memoryLimit, tc.algorithm)
		if err != nil && err.Error() != "memory limit must be greater than 0" {
			t.Errorf("Error initializing memcache: %v", err)
		}
	}
}

func TestSet(t *testing.T) {
	Initialize(100, "FIFO")

	testCases := []struct {
		key   string
		value interface{}
	}{
		{"key1", "value1"},
		{"key2", 123},
		{"key3", 123.456},
	}

	for _, tc := range testCases {
		_, err := Set(tc.key, tc.value)

		if err != nil {
			t.Errorf("Error setting key: %s, value: %v", tc.key, tc.value)
		}
	}
}

func TestSetLFU(t *testing.T) {
	Initialize(100, "LFU")

	testCases := []struct {
		key   string
		value interface{}
	}{
		{"key1", "value1"},
		{"key2", 123},
		{"key3", 123.456},
	}

	for _, tc := range testCases {
		_, err := Set(tc.key, tc.value)

		if err != nil {
			t.Errorf("Error setting key: %s, value: %v", tc.key, tc.value)
		}
	}
}

func TestSetMemoryExceeded(t *testing.T) {
	Initialize(100, "FIFO")

	for i := 0; i < 10; i++ {
		_, err := Set("key"+fmt.Sprintf("%d", i), "value"+fmt.Sprintf("%d", i))

		if err != nil {
			t.Errorf("Error setting key: %s, value: %v", "key"+fmt.Sprintf("%d", i), "value"+fmt.Sprintf("%d", i))
		}
	}

	_, err := Set("key-random", "value-random")

	if err != nil {
		t.Errorf("Error setting key: %s, value: %v", "key-random", "key-random")
	}
}

func TestSetMemoryExceededLFU(t *testing.T) {
	Initialize(100, "LFU")

	for i := 0; i < 10; i++ {
		_, err := Set("key"+fmt.Sprintf("%d", i), "value"+fmt.Sprintf("%d", i))

		if err != nil {
			t.Errorf("Error setting key: %s, value: %v", "key"+fmt.Sprintf("%d", i), "value"+fmt.Sprintf("%d", i))
		}
	}

	_, err := Set("key-random", "value-random")

	if err != nil {
		t.Errorf("Error setting key: %s, value: %v", "key-random", "key-random")
	}
}

func TestGet(t *testing.T) {
	Initialize(100, "FIFO")

	testCases := []struct {
		key   string
		value interface{}
	}{
		{"key1", "value1"},
	}

	for _, tc := range testCases {
		_, err := Set(tc.key, tc.value)

		if err != nil {
			t.Errorf("Error setting key: %s, value: %v", tc.key, tc.value)
		}

		val := Get(tc.key)
		if val.key != tc.key || val.value != tc.value {
			t.Errorf("Error getting key: %s, value: %v", tc.key, val.value)
		}
	}
}

func TestGetLFU(t *testing.T) {
	Initialize(100, "LFU")

	testCases := []struct {
		key   string
		value interface{}
	}{
		{"key1", "value1"},
	}

	for _, tc := range testCases {
		_, err := Set(tc.key, tc.value)

		if err != nil {
			t.Errorf("Error setting key: %s, value: %v", tc.key, tc.value)
		}

		val := Get(tc.key)
		if val.key != tc.key || val.value != tc.value {
			t.Errorf("Error getting key: %s, value: %v", tc.key, val.value)
		}
	}
}

func TestDelete(t *testing.T) {
	Initialize(100, "FIFO")

	testCases := []struct {
		key   string
		value interface{}
	}{
		{"key1", "value1"},
		{"error", "error"},
	}

	for _, tc := range testCases {
		if tc.key != "error" {
			_, err := Set(tc.key, tc.value)

			if err != nil {
				t.Errorf("Error setting key: %s, value: %v", tc.key, tc.value)
			}
		}

		val := Delete(tc.key)
		if val != true && tc.key != "error" {
			t.Errorf("Error getting key: %s, value: %v", tc.key, val)
		}
	}
}

func TestDeleteLFU(t *testing.T) {
	Initialize(100, "LFU")

	testCases := []struct {
		key   string
		value interface{}
	}{
		{"key1", "value1"},
		{"error", "error"},
	}

	for _, tc := range testCases {
		if tc.key != "error" {
			_, err := Set(tc.key, tc.value)

			if err != nil {
				t.Errorf("Error setting key: %s, value: %v", tc.key, tc.value)
			}
		}

		val := Delete(tc.key)
		if val != true && tc.key != "error" {
			t.Errorf("Error getting key: %s, value: %v", tc.key, val)
		}
	}
}
