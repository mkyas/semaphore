# semaphore

An implementation of semaphores in Go.

This implementation is motivated by showing how to implemented semaphores in Go.

The implementation in "sync/semaphore" cannot be used to implement signalling semaphores. It checks if a semaphore is released more often than held, which is contrary to what we have to do for a signaling semaphore.  We need an alternative implementation. 
