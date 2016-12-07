package service

import (
	"github.com/nairufan/yh-share/db/mongo"
	"github.com/nairufan/yh-share/model"
)

const (
	SequenceId = "document"
	collectionSequence = "sequence"
)

func Increase() int64 {
	session := mongo.Get()
	defer session.Close()
	sequence := &model.Sequence{}
	if err := session.FindId(collectionSequence, SequenceId, sequence); err != nil {
		sequence.Id = SequenceId
		sequence.Value = 1
		session.MustInsert(collectionSequence, sequence)
	} else {
		sequence.Value = sequence.Value + 1
		session.MustUpdateId(collectionSequence, SequenceId, sequence)
	}

	return sequence.Value
}

