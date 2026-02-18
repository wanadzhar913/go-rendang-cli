A small project for learning the basics of Golang.

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
go run main.go
```

# Resources
- Golang Tutorial for Beginners | Full Go Course: https://www.youtube.com/watch?v=yyUHQIec83I