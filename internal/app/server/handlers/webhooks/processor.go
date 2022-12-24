package webhooks

import (
	"github.com/FindHotel/emspy/internal/app/server/store"
)

type Processor struct {
	store store.Store
}

func NewProcessor(store store.Store) *Processor {
	return &Processor{store: store}
}

func (p *Processor) Capture(input []byte) error {
	return p.store.InsertWebhook(input)
}
