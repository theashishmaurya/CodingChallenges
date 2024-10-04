package main

import (
	"bufio"
	"fmt"
	"io"
	"net"
	"os"
	"strings"
)

type HTTPRequest struct {
	Method  string
	Path    string
	Version string
	Headers map[string]string
	Body    string
}

func parseRequest(reader *bufio.Reader) (HTTPRequest, error) {
	request := HTTPRequest{
		Headers: make(map[string]string),
	}

	// Read the request line
	requestLine, err := reader.ReadString('\n')
	if err != nil {
		return request, fmt.Errorf("error reading request line: %v", err)
	}
	requestLine = strings.TrimSpace(requestLine)

	// Parse the request line
	parts := strings.Split(requestLine, " ")
	if len(parts) != 3 {
		return request, fmt.Errorf("invalid request line: %s", requestLine)
	}
	request.Method = parts[0]
	request.Path = parts[1]
	request.Version = parts[2]

	// Read headers
	for {
		line, err := reader.ReadString('\n')
		if err != nil {
			return request, fmt.Errorf("error reading header line: %v", err)
		}
		line = strings.TrimSpace(line)
		if line == "" {
			break // End of headers
		}
		parts := strings.SplitN(line, ":", 2)
		if len(parts) != 2 {
			return request, fmt.Errorf("invalid header line: %s", line)
		}
		key := strings.TrimSpace(parts[0])
		value := strings.TrimSpace(parts[1])
		request.Headers[key] = value
	}

	// Read body if Content-Length is set
	if contentLength, ok := request.Headers["Content-Length"]; ok {
		length := 0
		fmt.Sscanf(contentLength, "%d", &length)
		if length > 0 {
			body := make([]byte, length)
			_, err := io.ReadFull(reader, body)
			if err != nil {
				return request, fmt.Errorf("error reading body: %v", err)
			}
			request.Body = string(body)
		}
	}

	return request, nil
}

func handleConnections(conn net.Conn) {
	fmt.Println("Connection accepted from:", conn.RemoteAddr())

	// Close the connection
	defer conn.Close()

	// Listen to the connection
	reader := bufio.NewReader(conn)
	httpRequest, err := parseRequest(reader)
	fmt.Println(httpRequest, "httpRequest")

	// rawRequest := ""
	// for {
	// 	fmt.Println("Attempting to read a line...")
	// 	line, err := reader.ReadString('\n')

	// 	if err != nil {
	// 		if err == io.EOF {
	// 			fmt.Println("EOF reached. Raw request so far:", rawRequest)
	// 		} else {
	// 			fmt.Println("Error reading from connection:", err)
	// 		}
	// 		return
	// 	}

	// 	rawRequest += line

	// 	if line == "\r\n" {
	// 		fmt.Println("End of headers reached")
	// 		break
	// 	}
	// }

	// fmt.Println("Full raw request:")
	// fmt.Println(rawRequest)

	response := "HTTP/1.1 404 Not Found\r\n\r\n"

	_, err = conn.Write([]byte(response))
	if err != nil {
		fmt.Println("Something went wrong while sending response", err)
	}

}

func main() {

	l, err := net.Listen("tcp", "0.0.0.0:4221")
	if err != nil {
		fmt.Println("Failed to bind to port 4221")
		os.Exit(1)
	}

	fmt.Println("Server Started at http://localhost:4221")

	conn, err := l.Accept()
	if err != nil {
		fmt.Println("Error accepting connection: ", err.Error())
		os.Exit(1)
	}

	handleConnections(conn)

}
