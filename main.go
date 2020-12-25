package main

import (
	"github.com/gliderlabs/ssh"
	"io"
	"log"
)

func main(){
	ssh.Handle(func(s ssh.Session){
		io.WriteString(s, "Hello World\n")
	})
	log.Fatal(ssh.ListenAndServe(":1049",nil))
}
