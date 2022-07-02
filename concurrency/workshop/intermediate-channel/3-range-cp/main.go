package main

func squareWorker(input <-chan int, output chan<- int) {
	for i := range input {
		//mengambil nilai i dari channel
		output <- i * i
	}
}

func receiver(output chan<- int) {
	input := make(chan int) // mengirim 0-10 ke squareWorker
	go squareWorker(input, output)
	for i := 0; i < 10; i++ {
		//kirim nilai i ke channel
		// TODO: answer here
		input <- i
	}
}
