package memo

import (
	"context"
)

type Func func(context.Context, string) (interface{}, error)

type canceled interface {
	Cancelled() bool
}

type result struct {
	value interface{}
	err   error
}

type entry struct {
	res   result
	ready chan struct{} // closed when res is ready
}
type request struct {
	key      string
	response chan<- result // the client wants a single result
	ctx      context.Context
}

type Memo struct{ requests chan request }

func New(f Func) *Memo {
	memo := &Memo{
		requests: make(chan request),
	}
	go memo.server(f)
	return memo
}

// context
func CheckCanceled(err error) bool {
	c, ok := err.(canceled)
	return ok && c.Cancelled()
}

func (memo *Memo) Get(ctx context.Context, key string) (interface{}, error) {
	response := make(chan result)
	memo.requests <- request{key, response, ctx}
	res := <-response
	return res.value, res.err
}

func (memo *Memo) Close() { close(memo.requests) }

func (memo *Memo) server(f Func) {
	cache := make(map[string]*entry)
	for req := range memo.requests {
		e := cache[req.key]
		if e == nil {
			// This is the first request for this key.
			e = &entry{ready: make(chan struct{})}
			cache[req.key] = e
			go e.call(req.ctx, f, req.key) // call f(key)
		} else {
			select {
			case <-e.ready:
				if CheckCanceled(e.res.err) {
					delete(cache, req.key)
					e = nil
				}
			}
		}
		go e.deliver(req.response)
	}
}

func (e *entry) call(ctx context.Context, f Func, key string) {
	// Evaluate the function.
	e.res.value, e.res.err = f(ctx, key)
	// Broadcast the ready condition.
	close(e.ready)
}

func (e *entry) deliver(response chan<- result) {
	// Wait for the ready condition.
	<-e.ready
	// Send the result to the client.
	response <- e.res
}
