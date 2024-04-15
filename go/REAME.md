```go
func main() {
	l, err := net.Listen("tcp", "0.0.0.0:4221")
	if err != nil {
		fmt.Println("Failed to bind to port 4221")
		os.Exit(1)
	}

	fmt.Println("GoServe Listening on 4221")

	for {
		conn, err := l.Accept()
		if err != nil {
			fmt.Println("Error accepting connection: ", err.Error())
			os.Exit(1)
		}

		rep := []byte("HTTP/1.1 200 OK\r\n")
		nBytesSent, err := conn.Write(rep)
		if err != nil {
			fmt.Println("Error accepting connection: ", err.Error())
		}
	}
}
```
