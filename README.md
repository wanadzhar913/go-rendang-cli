A small project for learning the basics of Golang. Also can't wait for Hari Raya Aidilfitri! 🥳

# Getting Started
We first install Go from it's [website](https://go.dev/doc/install). I'm using WSL2.

```bash
wget https://go.dev/dl/go1.26.0.darwin-amd64.pkg
rm -rf /usr/local/go && tar -C /usr/local -xzf go1.26.0.linux-amd64.tar.gz
export PATH=$PATH:/usr/local/go/bin

go version # go1.25.6 linux/386
go env GOPATH # home/faiq0913/go
```

To start the CLI program:

```bash
go mod init go-ticketing-app
go run .

go fmt # to format your code!
```

# Notes
Regarding goroutines, the main goroutine **does NOT** wait for other goroutines so if the main goroutine exited, other goroutines e.g., sendConfirmationEmail will not be executed. To remedy this, we use a `WaitGroup`. Also, in comparison to other languages, creating thread is cheaper, with fast startup times, with minimal resources used.

Regarding `WaitGroup`s, these essentially wiat for the launched goroutines to finish.
- The package `sync` provides basic synchronization functionality.
- `Add` sets the number of goroutines to wait for (increases the counter by the provided number).
- `Wait` blocks until the `WaitGroup` counter is 0.
- `Done` **decrements** the `WaitGroup` counter by 1. It's called by the goroutine to indicate that it's finished.

Regarding `maps` in Go. Maps are key-value pairs. However, they support only 1 data type for each key & value.

```go
package main

import strconv

var bookings = make([]map[string]string, 0) // create empty slice of maps. The initial size is 0.

// create user map
var userData = make(map[string]string) // create empty map
userData["firstName"] = firstName
userData["lastName"] = lastName
userData["userEmail"] = userEmail
userData["userOrders"] = strconv.FormatUint(uint64(userOrders), 10) // convert string to integer
bookings = append(bookings, userData)

func printFirstNames() []string {
	// for-each loop
	// think of _ (blank identifier) as Python's enumerate() function
	firstNames := []string{} // we use {} when we want a non-nil slice

	for _, booking := range bookings {
		// var names = strings.Fields(booking) // when booking was a Slice
		firstNames = append(firstNames, booking["firstName"]) // append() returns a new slice with the new element added
	}
	return firstNames
}
```

# Resources
- Golang Tutorial for Beginners | Full Go Course: https://www.youtube.com/watch?v=yyUHQIec83I