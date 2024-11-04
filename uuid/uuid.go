// Package UUID provides an easy to replace UUID package.
// Forks of Event Horizon can re-implement this package with a UUID library of choice.
package uuid

import (
	"github.com/google/uuid"
	"hash/fnv"
)

type AggretateID struct {
	aggregateId string
}

func NewAggregateID(s string) AggretateID {
	return AggretateID{aggregateId: s}
}

func (u AggretateID) ID() uint32 {
	h := fnv.New32a()
	h.Write([]byte(u.aggregateId))
	return h.Sum32()
}

// UUID is an alias type for github.com/google/uuid.UUID.
type UUID = AggretateID

func (u UUID) String() string {
	return string(u.aggregateId)
}

// Nil is an empty UUID.
var Nil = New(UUIDOption(func(id *AggretateID) {
	id.aggregateId = uuid.Nil.String()
}))

type UUIDOption func(id *AggretateID)

// New creates a new UUID.
func New(opts ...UUIDOption) UUID {
	if opts == nil {
		return NewAggregateID(uuid.New().String())
	}
	var m *AggretateID = &AggretateID{}
	for _, opt := range opts {
		opt(m)
	}
	return *m
}

// Parse parses a UUID from a string, or returns an error.
func Parse(s string) (UUID, error) {
	return NewAggregateID(s), nil
}

// MustParse parses a UUID from a string, or panics.
func MustParse(s string) UUID {
	return NewAggregateID(s)
}
