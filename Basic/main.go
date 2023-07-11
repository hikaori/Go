package go-basic

import "fmt"

func main(){
	fmt.Println("Hello World")
	fmt.Println(Hello("Kaori"))
}

func Hello(name string) string {
    // Return a greeting that embeds the name in a message.
    message := fmt.Sprintf("Hi, %v. Welcome!", name)
    return message
}