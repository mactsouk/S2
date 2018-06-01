package main

import ( 
	"fmt"
	"os"
)

func main() {
	fmt.Println("Hello World!")
	fmt.Println("SEMAPHORE_PIPELINE_ID:", os.Getenv("SEMAPHORE_PIPELINE_ID"))
}

