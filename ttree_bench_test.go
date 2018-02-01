package ttreego

import (
	"crypto/rand"
	"testing"
)

var (
	randomStrings [1024][]byte
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
}

func BenchmarkTTreeMatch(b *testing.B) {
	tree := NewTTree()
	for i := 0; i < b.N; i++ {
		tree.Add(randomStrings[i%len(randomStrings)])
	}

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		tree.Match(randomStrings[i%len(randomStrings)])
	}
}
