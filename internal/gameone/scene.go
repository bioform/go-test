package gameone

import (
	"crypto/md5"
	"encoding/hex"
)

// Scene describes a game scene
type Scene struct {
	ID       string    `yaml:"id"`
	Desc     string    `yaml:"Описание"`
	Question *Question `yaml:"Вопрос"`
}

// UnmarshalYAML is custom unmarshaller used by go-yaml library
func (s *Scene) UnmarshalYAML(unmarshal func(interface{}) error) error {

	type plain Scene
	if err := unmarshal((*plain)(s)); err != nil {
		return err
	}

	hash := md5.Sum([]byte(s.Desc))
	s.ID = hex.EncodeToString(hash[:])

	return nil
}
