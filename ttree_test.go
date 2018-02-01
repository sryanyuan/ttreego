package ttreego

import (
	"testing"
)

func TestTTreeMatch0(t *testing.T) {
	tree := NewTTree()
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
