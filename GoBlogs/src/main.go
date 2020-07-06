package main

import (
	"GoBlogs/routers"
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
)

func init() {
	log.SetPrefix("TRACE: ")
	log.SetFlags(log.Ldate | log.Lmicroseconds | log.Llongfile)
	log.Println("init started")
}

//Execution starts from main function
func main() {

	e := godotenv.Load()
	if e != nil {
		fmt.Print(e)
	}

	r := routers.SetupRouter()

	port := os.Getenv("port")

	// For run on requested port
	if len(os.Args) > 1 {
		reqPort := os.Args[1]
		if reqPort != "" {
			port = reqPort
		}
	}

	if port == "" {
		port = "9090" //localhost
	}
	type Job interface {
		Run()
	}

	r.Run(":" + port)

}
