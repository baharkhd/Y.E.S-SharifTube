package pending

import (
	"bytes"
	"encoding/json"
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

var toString = map[Status]string{
	PENDING:  "Pending",
	ACCEPTED: "Accepted",
	REJECTED: "Rejected",
}

var toID = map[string]Status{
	"Pending":  PENDING,
	"Accepted": ACCEPTED,
	"Rejected": REJECTED,
}

// MarshalJSON marshals the enum as a quoted json string
func (s Status) MarshalJSON() ([]byte, error) {
	buffer := bytes.NewBufferString(`"`)
	buffer.WriteString(toString[s])
	buffer.WriteString(`"`)
	return buffer.Bytes(), nil
}

// UnmarshalJSON unmarshalls a quoted json string to the enum value
func (s *Status) UnmarshalJSON(b []byte) error {
	var j string
	err := json.Unmarshal(b, &j)
	if err != nil {
		return err
	}
	// Note that if the string cannot be found then it will be set to the zero value, 'Created' in this case.
	*s = toID[j]
	return nil
}
