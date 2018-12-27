package main

import (
	"flag"
	"fmt"
	"log"
	"os"
)

func main() {
	usage := fmt.Sprintf(`%s action --mode MODE`, os.Args[0])
	actionCMD := flag.NewFlagSet("action", flag.ExitOnError)
	mode := actionCMD.String("mode", "GO", "mode parameter,default:GO")
	if len(os.Args) < 2 {
		fmt.Printf("usage:%s\n", usage)
		os.Exit(0)
	}
	switch os.Args[1] {
	case "action":
		err := actionCMD.Parse(os.Args[2:])
		if err != nil {
			log.Fatal(err)
		}
		if actionCMD.Parsed() {
			fmt.Printf("action mode:%s\n", *mode)
		}
	default:
		fmt.Printf("usage:%s\n", usage)
	}
}
