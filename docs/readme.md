# Here I will add some notes about what Im learning reading the go lang documentation.

## I will try to keep this repository updated with the new things I learn about GoLang.

### What is GoLang?

Go is an open source programming language that makes it easy to build simple, reliable, and efficient software.
Focus on simplicity, safety, and performance.
Have resources like concurrency and garbage collection.
It is a compiled language, which means that the code is converted to machine code before being executed.
Statically typed language, which means that variables must be declared before they are used.
Garbage collection, which means that memory is managed automatically.
Limited structural typing, which means that types are defined by their structure, not by their name.
Reflection, which means that a program can inspect its own structure.
A rich standard library.
Built-in concurrency support.
Built-in testing support.
Maded by Robert Griesemer, Rob Pike, and Ken Thompson at Google in 2007.

Rob Pike: "Go is an attempt to combine the ease of programming of an interpreted, dynamically typed language with the efficiency and safety of a statically typed, compiled language. It also aims to be modern, with support for networked and multicore computing. Finally, working with Go is intended to be fast: it should take at most a few seconds to build a large executable on a single computer. To meet these goals required addressing a number of linguistic issues: an expressive but lightweight type system; concurrency and garbage collection; rigid dependency specification; and so on. These cannot be addressed well by libraries or tools; a new language was called for."

# hello world

create a folder called hello and inside it create a file called hello.go with the following content:

```go
// there are comments inside the code to explain what is happening.
package main

import "fmt"

func main() {
	fmt.Println("Hello, World!")
}

```

go needs modules to work, so we need to initialize the module with the command:

```bash
go mod init hello/world
```

then we can run the code with the command:

```bash
go run hello.go
```

or we can build the code with the command:

```bash
go build hello.go
```

The output will be a binary file called hello. We can run it with the command:

```bash
./hello
```

the output will be:

```bash
Hello, World!
```
