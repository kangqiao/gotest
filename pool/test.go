package pool

import (
	"io"
	"log"
	"math/rand"
	"sync"
	"sync/atomic"
	"time"
)

const (
	maxCoroutines   = 25
	pooledResources = 5
)

type dbConnection struct {
	ID int32
}

func (dbConn *dbConnection) Close() error {
	log.Println("Close: Connection", dbConn.ID)
	return nil
}

var idCounter int32

func createConnection() (io.Closer, error) {
	id := atomic.AddInt32(&idCounter, 1)
	log.Println("Create: New Connection", id)

	return &dbConnection{ID: id}, nil
}

func Test() {
	var wg sync.WaitGroup
	wg.Add(maxCoroutines)

	p, err := New(createConnection, pooledResources)
	if err != nil {
		log.Println(err)
	}

	for query := 0; query < maxCoroutines; query++ {
		go func(q int) {
			preformQueries(q, p)
			wg.Done()
		}(query)
	}

	wg.Wait()
	log.Println("Shutdown Program.")
	p.Close()
}

func preformQueries(query int, p *Pool) {
	conn, err := p.Acquire()
	if err != nil {
		log.Println(err)
		return
	}

	defer p.Release(conn)

	time.Sleep(time.Duration(rand.Intn(1000)) * time.Millisecond)
	log.Printf("QID[%d] CID[%d]\n", query, conn.(*dbConnection).ID)
}
