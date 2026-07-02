package main

import "sync"

func fanIn(channels ...chan int)chan int{
	 var wg sync.WaitGroup
		muxedStream := make(chan int)
		output := func(c chan int){//sub-worker to get values from each channel
			for n := range c { //range over the channel
				muxedStream <- n
			}
			wg.Done()
		}
		wg.Add(len(channels))
		for _, channel := range channels{//ranges the pool of channels
			go output(channel)//creates sub-workers
		}
		go func(){
			wg.Wait()
		}()
		return muxedStream
	} 
