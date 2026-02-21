# Simple Rendang Ordering CLI in Go
A small project for learning the basics of Golang. Also can't wait for Hari Raya Aidilfitri! 🥳

## 1.0 Getting Started

We first install Go from it's [website](https://go.dev/doc/install). I'm using WSL2. In case you get the wrong Go interpreter path, on VSCode/Cursor, press `CTRL + SHIFT + P` > **Go: Choose Go Environment** > /usr/local/go/bin/go


```bash
wget https://go.dev/dl/go1.26.0.darwin-amd64.pkg
rm -rf /usr/local/go && tar -C /usr/local -xzf go1.26.0.linux-amd64.tar.gz
export PATH=$PATH:/usr/local/go/bin

go version # go1.25.6 linux/386
go env GOPATH # home/faiq0913/go
```

To start the CLI program:

```bash
# go mod init go-rendang-cli
go run .

go fmt # to format your code if you make changes!
```

## 2.0 How The Program Works & Outputs

On a high level, when using the CLI, we can essentially fake order Rendang as the CLI tool will ask us for first/last name, email & how many orders we want to make.

```
Welcome to our Rendang Factory!
Get your Rendang stock here!
We have 100.00 packets of rendang available.
Enter your first name:
Faiq
Enter your last name:
Adzlan
Enter your email address:
faiq@gmail.com
Enter your orders:
54
The first names of bookings are: [Faiq]
Thank you for your order!
You're in Malaysia.
Enter your first name:
Zahirah
Enter your last name:
Yayah
Enter your email address:
yayah@gmail.###########
Sending confirmation email: Faiq Adzlan has ordered 54 packets of rendang.
To: faiq@gmail.com...
###########
com
Enter your orders:
46
The first names of bookings are: [Faiq Zahirah]
Thank you for your order!
Our rendang is sold out! See you next Ramadan & Selamat Hari Raya!
###########
Sending confirmation email: Zahirah Yayah has ordered 46 packets of rendang.
To: yayah@gmail.com...
###########
All confirmation emails have been sent. List of all bookings:
[{Faiq Adzlan faiq@gmail.com 54} {Zahirah Yayah yayah@gmail.com 46}]
```

Here's the general flow more technically:

```
Main Goroutine
    ↓
Order Channel
    ↓
Worker Pool (multiple goroutines)
    ↓
Mutex-protected stock update
    ↓
Async Email Sender (Mocked)
```

## 3.0 Notes
Regarding goroutines, the main goroutine **does NOT** wait for other goroutines so if the main goroutine exited, other goroutines e.g., sendConfirmationEmail will not be executed. To remedy this, we use a `WaitGroup`. Also, in comparison to other languages, creating thread is cheaper, with fast startup times, with minimal resources used.

**Channels in Go are typed conduits used to send and receive values**, allowing goroutines to synchronize execution and communicate without explicit locks. They provide safe data transfer, blocking by default until both sender and receiver are ready.

Regarding `WaitGroup`, these essentially wait for the launched goroutines to finish.
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

## 4.0 Resources
- [Golang Tutorial for Beginners | Full Go Course](https://www.youtube.com/watch?v=yyUHQIec83I)
- [Go Channels](https://go.dev/tour/concurrency/2)
- [Go Goroutines](https://go.dev/tour/concurrency/1)
- [A Tour of Go](https://go.dev/tour/list)
- [Advanced Golang: Channels, Context and Interfaces Explained](youtube.com/watch?si=IHmrmKsboSbbMQP2&v=VkGQFFl66X4&feature=youtu.be)
