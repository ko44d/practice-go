package main

import (
	"bufio"
	"io/ioutil"
	"net"
	"net/http"
	"net/http/httputil"
	"strings"

	"github.com/sirupsen/logrus"
)

func main() {
	listener, err := net.Listen("tcp", "localhost:8888")
	if err != nil {
		panic(err)
	}
	logrus.Info("Server is running at localhost:8888")
	for {
		conn, err := listener.Accept()
		if err != nil {
			panic(err)
		}
		go func() {
			logrus.Infof("Accept %v\n", conn.RemoteAddr())
			request, err := http.ReadRequest(
				bufio.NewReader(conn))
			if err != nil {
				panic(err)
			}
			dump, err := httputil.DumpRequest(request, true)
			if err != nil {
				panic(err)
			}
			logrus.Info(string(dump))
			response := http.Response{
				StatusCode: 200,
				ProtoMajor: 1,
				ProtoMinor: 0,
				Body: ioutil.NopCloser(
					strings.NewReader("Hello world\n")),
			}
			response.Write(conn)
			conn.Close()
		}()
	}
}
