package ttreego

type TTree interface {
	// Add a string into trie tree
	Add([]byte) bool
	AddString(string) bool
	// Remove a string from trie tree
	Remove([]byte) error
	RemoveString(string) error
	// Match check if a string has prefix
	Match([]byte) bool
	MatchString(string) bool
	// MatchWhat return the matched prefix
	MatchWhat([]byte) string
	MatchWhatString(string) string
	// GetCount return the add count
	GetCount() int
	// Reset the trie tree
	Reset()
}
