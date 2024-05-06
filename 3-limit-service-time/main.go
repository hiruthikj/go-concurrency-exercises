//////////////////////////////////////////////////////////////////////
//
// Your video processing service has a freemium model. Everyone has 10
// sec of free processing time on your service. After that, the
// service will kill your process, unless you are a paid premium user.
//
// Beginner Level: 10s max per request
// Advanced Level: 10s max per user (accumulated)
//

package main

import (
	"time"
)

// User defines the UserModel. Use this to check whether a User is a
// Premium user or not
type User struct {
	ID        int
	IsPremium bool
	TimeUsed  int64 // in seconds
}

var timeLimit = 10 * time.Second

// HandleRequest runs the processes requested by users. Returns false
// if process had to be killed
func HandleRequest(process func(), u *User) bool {
	timer := time.NewTimer(timeLimit)

	processDone := make(chan interface{})

	go func() {
		process()
		close(processDone)
	}()

	for {
		select {
		case <-timer.C:
			return u.IsPremium
		case <-processDone:
			return true
		}
	}
}

func main() {
	RunMockServer()
}
