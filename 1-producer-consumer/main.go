//////////////////////////////////////////////////////////////////////
//
// Given is a producer-consumer scenario, where a producer reads in
// tweets from a mockstream and a consumer is processing the
// data. Your task is to change the code so that the producer as well
// as the consumer can run concurrently
//

package main

import (
	"fmt"
	"time"
)

func producer(stream Stream, pipe chan<- *Tweet) (tweets []*Tweet) {
	for {
		tweet, err := stream.Next()
		if err == ErrEOF {
			close(pipe)
			return tweets
		}

		pipe <- tweet
		// tweets = append(tweets, tweet)
	}
}

func consumer(pipe <-chan *Tweet) {
	// for _, t := range tweets {
	for t := range pipe {

		if t.IsTalkingAboutGo() {
			fmt.Println(t.Username, "\ttweets about golang")
		} else {
			fmt.Println(t.Username, "\tdoes not tweet about golang")
		}
	}
}

func main() {
	start := time.Now()

	pipe := make(chan *Tweet)

	stream := GetMockStream()

	// Producer
	go producer(stream, pipe)

	// Consumer
	consumer(pipe)

	fmt.Printf("Process took %s\n", time.Since(start))
}
