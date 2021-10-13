package main

import (
	"errors"
	"flag"
	"fmt"
	"log"
)

func main(){
}

type Vito string

func (v *Vito) String()string{
	return fmt.Sprint(*v)
}

func (v *Vito) Set(value string)error{
	if len(*v) > 0{
		return errors.New("hello,new flag")
	}
	*v = Vito("hello friend: "+value)
	return nil
}



func flagVar(){
	var name string

	flag.StringVar(&name,"name","a name","help info")
	flag.Parse()

	log.Printf("name %s",name)
}

// 准确识别不同的子命令
func flagCli(){
	var name string
	flag.Parse()
	goCmd := flag.NewFlagSet("go",flag.ExitOnError)
	phpCmd := flag.NewFlagSet("php",flag.ExitOnError)

	goCmd.StringVar(&name,"name","golang","help info")
	phpCmd.StringVar(&name,"n","phplang","help info")

	args := flag.Args()
	switch args[0]{
	case "go":
		goCmd.Parse(args[1:])
	case "php":
		phpCmd.Parse(args[1:])
	}
	log.Printf("name %s",name)
}


