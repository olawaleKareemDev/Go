
package main

import (
    "fmt"
    "time"
)

// We can use channels to synchronize execution across goroutines.
//  Here’s an example of using a blocking receive to wait for a goroutine to finish.
// When waiting for multiple goroutines to finish, you may prefer to use a WaitGroup.

func worker(done chan bool) {
    fmt.Print("working...")
    time.Sleep(time.Second)
    fmt.Println("done")

    done <- true

}

func main() {

    done := make(chan bool, 1)
    go worker(done)

    <-done

}

// var/log/auth.log ... logs in all user interaction with the linux server
// timedatectl

