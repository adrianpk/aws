package async

import "context"

type ConsumerFunc func(ctx context.Context, msg Envelope) error

func (f ConsumerFunc) Handle(ctx context.Context, msg Envelope) error

type Middleware func(Consumer) Consumer

func Chain(c Consumer, m ...Middleware) Consumer {
	w := c

	for i := len(m) - 1; i >= 0; i-- {
		w = m[i](w)
	}

	return w
}
