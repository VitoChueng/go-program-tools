package main

import (
	"github.com/VitoChueng/go-program-tools/ch1/tour/cmd"
	"log"
)

func main(){
	err := cmd.Execute()
	if err != nil{
		log.Fatalf("cmd Execute err: %v",err)
	}
}