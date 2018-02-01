# ttreego

ttreego is a simple trie tree implement using golang.

## Install

    go get github.com/sryanyuan/ttreego

## Documentation

Read [Godoc](https://godoc.org/github.com/sryanyuan/ttreego)

## Example

    // Initialize a trie tree
    tree := NewTTree()
    // Add some string
    tree.AddString("Hello world")
    tree.AddString("Hi go")
    // Do match jobs
    if tree.MatchString("Hello world") {
        fmt.Printf("Match") // Match
    }
    if tree.MatchString("Hi gogo") {
        fmt.Printf("Match") // Match
    }
    if tree.MatchString("Hello worl") {
        fmt.Printf("Match") // Not match
    }
    // Remove all string
    tree.RemoveString("Hello world")
    tree.RemoveString("Hi go")
    // Reset the tree
    tree.Reset()

## Performance

Add all string with 28 byte length into the tree, benchmark:

    BenchmarkTTreeAdd-4      2000000               675 ns/op

Match all random string:

    BenchmarkTTreeMatch-4            2000000               764 ns/op

## License

[MIT License](LICENSE)