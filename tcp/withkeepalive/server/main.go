package main

import (
	"bufio"
	"io"
	"io/ioutil"
	"net"
	"net/http"
	"net/http/httputil"
	"strings"
	"time"

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
			defer conn.Close()
			logrus.Infof("Accept %v\n", conn.RemoteAddr())
			for {
				conn.SetReadDeadline(time.Now().Add(5 * time.Second))
				request, err := http.ReadRequest(bufio.NewReader(conn))
				if err != nil {
					neterr, ok := err.(net.Error)
					if ok && neterr.Timeout() {
						logrus.Info("Timeout")
						break
					} else if err == io.EOF {
						break
					}
					panic(err)
				}
				dump, err := httputil.DumpRequest(request, true)
				if err != nil {
					panic(err)
				}
				logrus.Info(string(dump))
				contents := "Hello World\n"
				response := http.Response{
					StatusCode:    200,
					ProtoMajor:    1,
					ProtoMinor:    1,
					ContentLength: int64(len(contents)),
					Body: ioutil.NopCloser(
						strings.NewReader(contents)),
				}
				response.Write(conn)
			}
		}()
	}
}
