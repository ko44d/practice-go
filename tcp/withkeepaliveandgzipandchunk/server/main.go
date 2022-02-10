package main

import (
	"bufio"
	"fmt"
	"github.com/sirupsen/logrus"
	"io"
	"net"
	"net/http"
	"net/http/httputil"
	"strings"
)

var contents = []string{
	"あああああああああああああああああああああああああああああああああああああああああ",
	"いいいいいいいいいいいいいいいいいいいいいいいいいいいいいいいいいいいいいいいいい",
	"ううううううううううううううううううううううううううううううううううううううううう",
	"えええええええええええええええええええええええええええええええええええええええええ",
	"おおおおおおおおおおおおおおおおおおおおおおおおおおおおおおおおおおおおおおおおお",
}

func main() {
	listen, err := net.Listen("tcp", "localhost:8888")
	if err != nil {
		panic(err)
	}
	logrus.Println("Server is running at localhost:8888")
	for {
		conn, err := listen.Accept()
		if err != nil {
			panic(err)
		}
		go processSession(conn)
	}
}

func processSession(conn net.Conn) {
	logrus.Printf("%v\n", conn.RemoteAddr())
	defer conn.Close()
	for {
		request, err := http.ReadRequest(bufio.NewReader(conn))
		if err != nil {
			if err == io.EOF {
				break
			}
			panic(err)
		}
		dump, err := httputil.DumpRequest(request, true)
		if err != nil {
			panic(err)
		}
		logrus.Println(string(dump))

		fmt.Fprintf(conn, strings.Join([]string{
			"HTTP/1.1 200 OK",
			"Contents-Type: text/plain",
			"Transfer-Encoding: chunked",
			"", ""}, "\r\n"))
		for _, content := range contents {
			bytes := []byte(content)
			fmt.Fprintf(conn, "%x\r\n%s\r\n", len(bytes), content)
		}
		fmt.Fprintf(conn, "0\r\n\r\n")
	}
}

func isGZipAcceptable(request *http.Request) bool {
	return strings.Index(strings.Join(request.Header["Accpt-Encoding"], ","), "gzip") != -1
}
