package concepts

import (
	"fmt"
	"log"
	"math/rand"
	"sync"
	"time"
)

var db = []int{}
var mutex = &sync.Mutex{}

func Insert(val int) {
	mutex.Lock()
	db = append(db, val)
	mutex.Unlock()
	log.Println("Value appended: ", val)
}

func Select(index int) {
	val := db[index]
	log.Printf("Value is %d at index %d \n", val, index)
}

func Transaction(transactionType string, value int) {
	switch transactionType {
	case "Read":
		Select(value)
	case "Write":
		Insert(value)
	}
}

func seedData() {
	rand.Seed(time.Hour.Nanoseconds())
	i := 0
	for i < 100 {
		num := rand.Intn(100)
		Transaction("Write", num)
		i++
	}
	wg.Done()
}

func readData() {
	i := 0
	for i < 100 {
		rand.Seed(time.Hour.Nanoseconds())
		max := len(db)
		Transaction("Read", rand.Intn(max))
		i++
	}
	wg.Done()
}

func GenerateTransactions() {
	wg.Add(10000)
	for i := 0; i < 10000; i++ {
		go seedData()
	}
	wg.Wait()
	fmt.Println(len(db) == 10000*100)
	fmt.Println(len(db))
}
