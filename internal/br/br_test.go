package br

import (
	"testing"
)

func TestRun(t *testing.T) {
	teardown := setup(t)
	defer teardown()

	if err := Run(); err != nil {
		t.Errorf("run() failed: %v", err)
	}
}

func setup(t *testing.T) func() {
	t.Log("setup")

	return func() {
		t.Log("teardown")
	}
}
