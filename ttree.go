package ttreego

type TTree interface {
	// Add a string into trie tree
	Add([]byte) bool
	AddString(string) bool
	// Remove a string from trie tree
	Remove([]byte) error
	RemoveString(string) error
	// Check if a string has prefix
	Match([]byte) bool
	MatchString(string) bool
	// Get the add count
	GetCount() int
	// Reset the trie tree
	Reset()
}
