package main

import (
	"testing"
)

func TestAdd(t *testing.T) {
	got := Add(2,3)

	if got != 5 {
		t.Error("expected 5")
	}
	t.Log("testing complete!")
}

// 1. t.Log 
// - Dear Test Runner,একটা note লিখে রাখো।

// 2. t.Fail  
// - Mark test as FAILED
// - BUT keep executing

// 3. t.Error 
// - Mark test as FAILED
// - still executing
// - t.Log + t.Fail


// 4. t.FailNow 
// - Mark test as FAILED
// - Stop execution now

// 5. t.Fatal 
// - Mark test as FAILED
// - Stop execution now
// - t.Log + t.FailNow


// 6. t.Skip 
// - Don't run this test.
// - Not pass.
// - Not fail.
// - Ignore.

// 7. t.SkipNow
// - Don't run this test.
// - Not pass.
// - Not fail.
// - Ignore.
// - Stop execution now