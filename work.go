package main

type WorkRequest struct {
	Execute func(config interface{}) error
}
