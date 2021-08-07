package subscription

import "github.com/ryil247/lpbud/backend/internal/event"

type AddSubjectHandler func(event.AddSubjectRequest) (event.SubjectAddedEvent, error)

type Subscription interface {
	RegisterAddSubjectHandler(AddSubjectHandler)
}
