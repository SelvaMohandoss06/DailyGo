package  main

import "fmt"

func main() {
	messages := make(chan string)
	go func() {
		fmt.Println("waiting to write to messages")
		messages <- "hello!"
		fmt.Println("wrote to messages")
	}()

	fmt.Println("waiting to read from messages")
	msg := <- messages
	fmt.Println("read from messages")
	fmt.Println(msg)

	bufferedMessages := make(chan string,2)
	bufferedMessages <- "hellohghggh!"
	bufferedMessages <- "hello Again !"
	 //bufferedMessages <- "Dfdsfds"

	fmt.Println(<-bufferedMessages)
	fmt.Println(<-bufferedMessages)

	bufferedMsgs := make(chan string,10)

	go func() {
      for i := 0 ; i <10 ; i++ {
      	msg := fmt.Sprintf("message %d",i)
		  bufferedMsgs <- msg
      	  fmt.Printf("sent : %s \n",msg)
      }
	}()

	m :=  <- bufferedMsgs
	fmt.Printf("received: %s \n", m)
	c := make(chan string)
	quit := make(chan struct{})

	go func(messages  [] string) {
		for _, s := range messages {
			c  <- s
		}
		close(quit)
	}([]string {"hi","bye"})

	for {
		select {
		 case  messsage := <- c:
		 	fmt.Println(messsage)
		 case  <- quit:
		 	fmt.Println("shutting down ")
		}
	}


}
