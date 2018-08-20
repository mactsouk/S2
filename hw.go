package main

import ( 
	"fmt"
	"os"
)

func main() {
	fmt.Printf("Hello World!\n")
	fmt.Println("SEMAPHORE_PIPELINE_ID:", os.Getenv("SEMAPHORE_PIPELINE_ID"))
}

