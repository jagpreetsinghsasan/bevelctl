package main

import (
	"bevelctl/utils"
	"os"
	"fmt"
)

func main(){
	args := os.Args[1:]
	var output string = utils.CreateNetworkConfig(args[0],"fabric")
	fmt.Println(output)
}