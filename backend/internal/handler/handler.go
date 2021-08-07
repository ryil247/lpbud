package handler

import (
	"errors"
	"fmt"

	"github.com/ryil247/lpbud/backend/internal/event"
)

type Handler struct {
	Subjects []event.AddSubjectRequest
}

// func (h *Handler) AddSubjectHandler(subject event.AddSubjectRequest) {
// 	h.Subjects = append(h.Subjects, subject)
// 	fmt.Printf("new subject: %q\n", subject.Name)
// }

func (h *Handler) AddSubjectHandler(request event.AddSubjectRequest) (result event.SubjectAddedEvent, err error) {
	h.Subjects = append(h.Subjects, request)
	fmt.Printf("new subject: %q\n", request.Name)
	if request.Name == "" {
		return result, errors.New("no name provided")
	}

	return event.SubjectAddedEvent{
		Name: request.Name,
	}, nil
}

/*
class Handler {
	public Array<AddSubjectRequest> subjects;

	public void AddSubjectHandler(event.AddSubjectRequest subject)
}
*/
