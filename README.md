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

Add all string(1024) with 28 byte length into the tree, benchmark:

    BenchmarkTTreeAdd-4      2000000               675 ns/op

Match all random string:

    BenchmarkTTreeMatch-4            2000000               764 ns/op

Add all string(10240) with 28 byte length into the tree, benchmark:

    BenchmarkTTreeAdd-4      2000000               2830 ns/op

Match all random string:

    BenchmarkTTreeMatch-4   	 1000000	      2197 ns/op	       0 B/op	       0 allocs/op

Match all random string (80960):

    BenchmarkTTreeMatch-4   	 1000000	      2267 ns/op	       0 B/op	       0 allocs/op

If > 10240 strings in the tree, the search cost time is abount 2267 ns (constant)

## License

[MIT License](LICENSE)
