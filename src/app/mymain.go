package main

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"strconv"
	"strings"
	"sync"

	"github.com/pioardi/goproof/src/user"
)

var (
	trace *log.Logger // Just about anything
	debug *log.Logger // Just about anything
	info  *log.Logger // Important information
	warn  *log.Logger // Be concerned
	err   *log.Logger // Critical problem
)

func port() string {
	port := os.Getenv("PORT")
	if len(port) == 0 {
		port = "8080"
	}
	return ":" + port
}

func supergo(wg *sync.WaitGroup, i int, c chan string) {

	log.Println("Goroutine " + strconv.Itoa(i))
	c <- strconv.Itoa(i)
	wg.Done()
}

// init is called prior to main.
func init() {
	log.Println(user.Export)
	user.Run("Ok")
	var user = user.MyCustomType{Username: "ok", Password: "ok"}
	log.Println(user)
	log.SetPrefix("TRACE : ")
	log.SetFlags(log.Ldate | log.Lmicroseconds | log.Lshortfile)
	file, err := os.OpenFile("server.log",
		os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
	if err != nil {
		log.Fatalln("Not able to open log files")
	}
	// multi writer to write on more outptut.
	writer := io.MultiWriter(os.Stdout, file)
	// how to have a custom logger in go.
	info = log.New(writer, "INFO: ", log.Ldate|log.Ltime|log.Lmicroseconds|log.Lshortfile)
	debug = log.New(writer, "DEBUG: ", log.Ldate|log.Ltime|log.Lmicroseconds|log.Lshortfile)
	warn = log.New(writer, "WARN: ", log.Ldate|log.Ltime|log.Lmicroseconds|log.Lshortfile)
	trace = log.New(writer, "TRACE: ", log.Ldate|log.Ltime|log.Lmicroseconds|log.Lshortfile)
}

// main is the entry point for the program.
func main() {
	// endof htt server
	var wg sync.WaitGroup
	// 10 + web server goroutine
	wg.Add(11)
	go startServer()
	var results = make(chan string)
	for index := 0; index < 10; index++ {
		go supergo(&wg, index, results)
		i := <-results
		log.Println("Received a result into the channel ", i)
	}

	// wait after start goroutines , not before
	wg.Wait()
	go func() {
		log.Println("Closing channel")
		close(results)
	}()
}

func multireturn(i, j string) (x, y string) {
	return "primo", "secondo"
}

func startServer() {
	// http server
	var port = port()
	log.Println("Starting web server , stop to play start to work really on go")
	log.Println("Web server starting on port ", port, "and listening for requests")
	http.HandleFunc("/", sayHello)
	log.Fatal(http.ListenAndServe(port, nil))
}

func sayHello(w http.ResponseWriter, r *http.Request) {
	message := r.URL.Path
	message = strings.TrimPrefix(message, "/")
	message = "Hello " + message
	info.Println(message)
	// write response back
	fmt.Fprintf(w, message)
}
