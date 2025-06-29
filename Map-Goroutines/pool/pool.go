package main

import (
	"context"
	"errors"
	"fmt"
	"io"
	"sync"
)

var ErrPoolClosed = errors.New("pool has been closed")

type Factory func(ctx context.Context) (io.Closer, error)

type Pool struct {
	m         sync.Mutex
	resources chan io.Closer
	factory   Factory
	closed    bool
}

func NewPool(fn Factory, size uint) (*Pool, error) {
	if size == 0 {
		return nil, errors.New("pool size must be greater than zero")
	}

	return &Pool{
		factory:   fn,
		resources: make(chan io.Closer, size),
	}, nil
}

func (p *Pool) Acquire(ctx context.Context) (io.Closer, error) {
	select {
	case <-ctx.Done():
		return nil, ctx.Err()
	case r, ok := <-p.resources:
		if !ok {
			return nil, ErrPoolClosed
		}
		return r, nil
	default:
		r, err := p.factory(ctx)
		if err != nil {
			return nil, fmt.Errorf("creating new resource: %w", err)
		}
		return r, nil
	}
}

func (p *Pool) Release(r io.Closer) error {
	p.m.Lock()
	defer p.m.Unlock()

	if p.closed {
		if err := r.Close(); err != nil {
			return fmt.Errorf("closing resource: %w", err)
		}
		return nil
	}

	select {
	case p.resources <- r:
		return nil
	default:
		if err := r.Close(); err != nil {
			return fmt.Errorf("closing resource: %w", err)
		}
		return nil
	}
}

func (p *Pool) Close() error {
	p.m.Lock()
	defer p.m.Unlock()

	if p.closed {
		return nil
	}
	close(p.resources)
	p.closed = true
	for r := range p.resources {
		if err := r.Close(); err != nil {
			return err
		}
	}
	return nil
}
