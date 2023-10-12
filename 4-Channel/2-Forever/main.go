package main

func main() {
	forever := make(chan bool)

	go func() {
		for i := 0; i < 10; i++ {
			println(i)
		}
		forever <- true
	}()

	<-forever

}
