package model

type IDGenerator interface {
	Generate() string
}
