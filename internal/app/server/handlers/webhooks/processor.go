package webhooks

import (
	"context"

	"github.com/FindHotel/emspy/internal/app/store"
	"golang.org/x/sync/errgroup"
)

type Processor struct {
	stores []store.Store
	source string
}

func NewProcessor(stores []store.Store, source string) *Processor {
	return &Processor{source: source, stores: stores}
}

func (p *Processor) Capture(ctx context.Context, input []byte) error {
	g, _ := errgroup.WithContext(ctx)
	for _, store := range p.stores {
		s := store
		g.Go(func() error {
			return s.InsertWebhook(p.source, input)
		})
	}
	return g.Wait()
}
