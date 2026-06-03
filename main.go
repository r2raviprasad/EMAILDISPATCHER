package main

import "sync"

type Recipient struct {
	Name  string
	Email string
}

func main() {
	recipientChannel := make(chan Recipient)
	workerCount := 5
	var wg sync.WaitGroup
	for i := 1; i <= workerCount; i++ {
		wg.Add(1)
		go emailWorker(i, recipientChannel, &wg)
	}
	go loadRecipients("emails.csv", recipientChannel)
	wg.Wait()

}
