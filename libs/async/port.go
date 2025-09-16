package async

import "context"

type (
	Publisher interface {
		Publish(ctx context.Context, msg Envelope)
	}

	Enqueuer interface {
		Enqueue(ctx context.Context, msg Envelope)
	}

	Consumer interface {
		Handle(ctx context.Context, msg Envelope)
	}

	Deleter interface {
		Delete(ctx context.Context, r string)
	}

	DequeueBatch struct {
		Receipt string
		Msg     Envelope
	}

	Queue interface {
		Receive(ctx context.Context, max int, waitSeconds int) (DequeueBatch, error)
	}
)
