package static

import (
	"strconv"
	"time"

	"github.com/ryil247/lpbud/backend/internal/event"
	"github.com/ryil247/lpbud/backend/internal/subscription"
)

type Static struct {
	spawnAfter          time.Ticker
	subjectAddedHandler subscription.AddSubjectHandler
}

func New(opts ...option) *Static {
	c := &config{
		spawnAfter: 1 * time.Second,
	}
	for _, opt := range opts {
		opt(c)
	}
	sub := &Static{
		spawnAfter: *time.NewTicker(c.spawnAfter),
	}

	go func() {
		for i := 0; ; i++ {
			<-sub.spawnAfter.C
			if sub.subjectAddedHandler != nil {
				subject := event.AddSubjectRequest{
					Name: "subject " + strconv.Itoa(i),
				}
				sub.subjectAddedHandler(subject)
			}
		}
	}()

	return sub
}

func (s *Static) RegisterSubjectAddedHandler(handler subscription.AddSubjectHandler) {
	s.subjectAddedHandler = handler
}
