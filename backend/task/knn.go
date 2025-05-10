package task

import (
	"time"
)

type KNN25P struct{}

func (r *KNN25P) Run(input *Input) time.Duration {
	_, t := input.Index.KNNTimed(input.RandomPoint, input.Amount/4)

	return t
}

type KNN90P struct{}

func (r *KNN90P) Run(input *Input) time.Duration {
	_, t := input.Index.KNNTimed(input.RandomPoint, input.Amount*90/100)

	return t
}

type KNN1P struct{}

func (r *KNN1P) Run(input *Input) time.Duration {
	_, t := input.Index.KNNTimed(input.RandomPoint, input.Amount/100)

	return t
}

type KNN10 struct{}

func (r *KNN10) Run(input *Input) time.Duration {
	_, t := input.Index.KNNTimed(input.RandomPoint, 10)

	return t
}

type KNN100 struct{}

func (r *KNN100) Run(input *Input) time.Duration {
	_, t := input.Index.KNNTimed(input.RandomPoint, 100)

	return t
}
