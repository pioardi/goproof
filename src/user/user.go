package user

import (
	"log"
)

// Export is a proof
var Export = "JAM a vre"

func init() {
	log.SetPrefix("user.go")
}

// notifier is an interface that defined notification
// type behavior.
type notifier interface {
	notify()
}

// MyCustomType my first custom type
type MyCustomType struct {
	username string
	password string
	isRoot   bool
}

func (user *MyCustomType) notify() {
	log.Println(user)
}

func (user *MyCustomType) method() {
	log.Println("Invoked a custom method of a custom type " + user.username)
}

// func main() {
// 	var user = MyCustomType{username: "username", password: "password"}
// 	user.method()
// 	provaInterfaccia(&user)
// }

func provaInterfaccia(n notifier) {
	n.notify()
}
