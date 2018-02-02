package ttreego

import (
	"testing"
)

func TestTTreeMatch0(t *testing.T) {
	tree := &ttreeImpl{}
	tree.AddString("Hello world")
	tree.AddString("Hello my honey")
	tree.AddString("Go golang")

	if !tree.MatchString("Hello world") {
		t.Errorf("Error")
	}
	if !tree.MatchString("Hello world hello") {
		t.Errorf("Error")
	}
	if tree.MatchString("Hello worl") {
		t.Errorf("Error")
	}

	if !tree.MatchString("Hello my honey") {
		t.Errorf("Error")
	}
	if !tree.MatchString("Hello my honey honey") {
		t.Errorf("Error")
	}
	if tree.MatchString("Hello my hoNey") {
		t.Errorf("Error")
	}

	if !tree.MatchString("Go golang") {
		t.Errorf("Error")
	}
	if !tree.MatchString("Go golang ") {
		t.Errorf("Error")
	}
	if tree.MatchString("o golang") {
		t.Errorf("Error")
	}

	if err := tree.RemoveString("Hello world"); nil != err {
		t.Error(err)
		t.FailNow()
	}
	if tree.MatchString("Hello world") {
		t.Errorf("Error")
	}
	if !tree.MatchString("Hello my honey") {
		t.Errorf("Error")
	}
	if !tree.MatchString("Go golang") {
		t.Errorf("Error")
	}

	if err := tree.RemoveString("Hello my honey"); nil != err {
		t.Error(err)
		t.FailNow()
	}
	if tree.MatchString("Hello my honey") {
		t.Errorf("Error")
	}
	if !tree.MatchString("Go golang") {
		t.Errorf("Error")
	}

	if err := tree.RemoveString("Go golang"); nil != err {
		t.Error(err)
		t.FailNow()
	}
	if tree.MatchString("Go golang") {
		t.Errorf("Error")
	}

	if nil != tree.nodes {
		t.Errorf("node not nil")
	}
}

func TestTTreeRemove(t *testing.T) {
	tree := NewTTree()
	for i := 0; i < 102400; i++ {
		tree.Add(randomStrings[i%len(randomStrings)])
	}

	for i := 0; i < 10240; i++ {
		if err := tree.Remove(randomStrings[i%len(randomStrings)]); nil != err {
			if err != ErrTTreeNodeNotFound {
				t.Errorf("Remove %v failed, err = %v", i, err)
			}
		}
	}
	// Test one not in the tree
	if tree.MatchString("..!!@@##$$%%^^&&**(())__") {
		t.Errorf("Match a not match string")
	}
	if tree.GetCount() != 0 {
		t.Errorf("Removed all, but still have %v in tree", tree.GetCount())
	}
}
