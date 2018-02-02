package ttreego

type TTree interface {
	Add([]byte) bool
	AddString(string) bool
	Remove([]byte) error
	RemoveString(string) error
	Match([]byte) bool
	MatchString(string) bool
	GetCount() int
	Reset()
}
