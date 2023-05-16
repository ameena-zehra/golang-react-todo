package main

import (
	"fmt" 
	"log" 
	"net/http" 
	"github.com/ameena-zehra/golang-react-todo/router")
func main(){
	r:= router.Router()
	fmt.Println("starting the server on port 9000")
	log.Fatal(http.ListenandServe(":9000", r))

}