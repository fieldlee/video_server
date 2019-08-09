package main

import "log"

type ConnLimiter struct {
	concurrentConn int
	bucket chan int
}

func NewConnLimiter(c int) *ConnLimiter{
	return &ConnLimiter{
		concurrentConn:c,
		bucket:make(chan int,c),
	}
}

func (c *ConnLimiter)GetConn()bool{
	if len(c.bucket) >= c.concurrentConn{
		log.Printf("limiter %d",len(c.bucket))
		return false
	}
	c.bucket <- 1
	return true
}

func (c *ConnLimiter)ReleaseConn(){
	cb := <- c.bucket
	log.Printf("new connection coming :%d",cb)
}