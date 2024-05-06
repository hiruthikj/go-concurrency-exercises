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

var timeLimitInSeconds int64 = 10

// HandleRequest runs the processes requested by users. Returns false
// if process had to be killed
func HandleRequest(process func(), u *User) bool {
	processDone := make(chan interface{})
	ticker := time.NewTicker(1 * time.Second)

	go func() {
		process()
		close(processDone)
	}()

	for {
		select {
		case <-ticker.C:
			// fmt.Printf("tick %v\n", time.Now().Unix() )
			u.TimeUsed += 1
			if u.TimeUsed >= timeLimitInSeconds {
				return u.IsPremium
			}
		case <-processDone:
			return true
		}
	}
}

func main() {
	RunMockServer()
}
