package config

import (
	dtos "github.com/dranikpg/dto-mapper"
)

func NewMapper () *dtos.Mapper {
	mapper := &dtos.Mapper{}
	return mapper
}
