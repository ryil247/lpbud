package event

type AddSubjectRequest struct {
	Name string
}

type SubjectAddedEvent struct {
	Name string
}

type Example struct {
	//private variable because lower case "n"
	name string
	//public variable because uper case "h"
	Hodor int
}
