package ttreego

import (
	"crypto/rand"
	"testing"
)

var (
	randomStrings [80960][]byte
)

func init() {
	// Init random strings
	const stringLength = 28

	for i := 0; i < len(randomStrings); i++ {
		stringBuffer := make([]byte, stringLength)
		if _, err := rand.Read(stringBuffer[:]); nil != err {
			panic(err)
		}
		randomStrings[i] = stringBuffer
	}
}

func BenchmarkTTreeAdd(b *testing.B) {
	tree := NewTTree()
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		tree.Add(randomStrings[i%len(randomStrings)])
	}
	b.Logf("Add count %v", tree.GetCount())
}

func BenchmarkTTreeMatch(b *testing.B) {
	tree := NewTTree()
	for i := 0; i < b.N; i++ {
		tree.Add(randomStrings[i%len(randomStrings)])
	}

	b.ResetTimer()
	nMatch := 0
	for i := 0; i < b.N; i++ {
		if tree.Match(randomStrings[i%len(randomStrings)]) {
			nMatch++
		}
	}
	// Test one not in the tree
	if tree.MatchString("..!!@@##$$%%^^&&**(())__") {
		b.Errorf("Match a not match string")
	}
	b.Logf("%v matched, %v total, %v in tree", nMatch, b.N, tree.GetCount())
}

func BenchmarkTTreeRemove(b *testing.B) {
	tree := NewTTree()
	for i := 0; i < b.N; i++ {
		tree.Add(randomStrings[i%len(randomStrings)])
	}

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		if err := tree.Remove(randomStrings[i%len(randomStrings)]); nil != err {
			if err != ErrTTreeNodeNotFound {
				b.Errorf("Remove %v failed, err = %v", i, err)
			}
		}
	}
	// Test one not in the tree
	if tree.MatchString("..!!@@##$$%%^^&&**(())__") {
		b.Errorf("Match a not match string")
	}
	if tree.GetCount() != 0 {
		b.Errorf("Removed all, but still have %v in tree", tree.GetCount())
	}
}
