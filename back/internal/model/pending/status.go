package pending

import (
	"yes-sharifTube/graph/model"
)

type Status int

const (
	PENDING Status = iota
	ACCEPTED
	REJECTED
)

func (s Status) String() string {
	return toString[s]
}

func (s Status) Reshape() model.Status {
	switch s {
	case ACCEPTED:
		return model.StatusAccepted
	case REJECTED:
		return model.StatusRejected
	default:
		return model.StatusPending
	}
}

func NewStatus(s model.Status) Status {
	switch s {
	case model.StatusAccepted:
		return ACCEPTED
	case model.StatusRejected:
		return REJECTED
	default:
		return PENDING
	}
}

var toString = map[Status]string{
	PENDING:  "Pending",
	ACCEPTED: "Accepted",
	REJECTED: "Rejected",
}