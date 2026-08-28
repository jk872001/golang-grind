package main

import (
	"fmt"
	"time"
	// "math/rand"
	// "time"
)

// func processNum(numChan chan int){

// 	for num := range numChan{
// 	   fmt.Println("processing channel", num)
// 	   time.Sleep(time.Second)
// 	}

// }

// func sum(result chan int,num1 int,num2 int){
// 	add := num1 + num2
// 	result <- add
// }

func emailSender(emailChan chan string,done chan bool){
    defer func(){done <- true}()
	for email := range emailChan{
		fmt.Println("processing email",email)
		time.Sleep(time.Second)
	}
}


func main(){

   emailChan := make(chan string, 100)
   done := make(chan bool)

   go emailSender(emailChan,done)

   for i:=0;i<10;i++{
	 emailChan <- fmt.Sprintf("%d@gmail.com",i)
   }
   
   fmt.Println("done sending....")
   close(emailChan)
   <-done
   



    // result := make(chan int)
    // go sum(result,4,6)
	// ans := <- result
	// fmt.Println(ans)



    // numChan := make(chan int)
	// go processNum(numChan)

	// for {
	// 	numChan <- rand.Intn(100)
	// }

	// messageChannel := make(chan string)

	// messageChannel  <- "ping"

	// msg := <- messageChannel

	// fmt.Println(msg)
}
