/*****************************************************************************
 * http_proxy.go                                                                 
 * Names: 
 * NetIds:
 *****************************************************************************/

 // TODO: implement an HTTP proxy
 


 //TODO: TAKE IN PORT# FROM CL ARG AND LISTEN FOR INCOMING CONNECTIONS ON THAT PORT
 //func Listen(network, address string) (Listener, error)
 package main
 
 import (
	 "fmt"
	 "os"
	 "net/http"
	 "net"
 )
 
 func handleError(err) {
	 
	 if err == nil {
		 return
	 }
 
	 fmt.Println("Error: ", err.Error())
	 os.Exit(1)
 }
 
 func handleRequest(conn) {
	// if request is not a GET request, return 500 status code
	// "Once a client has connected, the proxy should read data"
	// "from the client and then check for a properly-formatted HTTP"
	// "request. Use the http library to ensure that the proxy receives"
	// "a request that contains a 'valid request line'"
	 request, err := http.ReadRequest(bufio.NewReader(conn))
	 handleError(err)
	 
	 // if 'method' is empty or GET then we are good
	 fmt.Println("Request method: ", request.Method)
	 fmt.Println("Request protocol version: ", request.Proto)
 
 }
 
 
 func startServer(portNum) {
	 portStr := fmt.Sprintf(":%i", portNum)
	 // "Your proxy should listen on the port specified from the cmd line"
	 // "and wait for incoming client connections"
	 listener, err := net.Listen("tcp", portStr)
	 
	 handleError(err)
	
	 for {
		 
		 conn, err = listener.Accept()
 
		 handleError(err)
		 // "client requests should be handled concurrently, with a new"
		 // "go routine spawned to handle each request"
		 go handleRequest(conn)
	 }
 
 }
 
 
 func main() {
	 // Program Name is always the first (implicit) argument
	 cmd := os.Args[0]
	 port := os.Args[1]
	 
	 fmt.Println("this is to test the repo")
	 fmt.Println("Starting server")
	 startServer(port)
  
 }