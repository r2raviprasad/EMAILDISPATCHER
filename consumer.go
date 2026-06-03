package main

import (
	"fmt"
	"sync"
)

func emailWorker(id int, ch chan Recipient, wg *sync.WaitGroup) {
	defer wg.Done()
	for recipient := range ch {
		fmt.Println("Processed by worker ", id, ": ", recipient)
		
		// SMTP email sending logic would go here
	}

}





