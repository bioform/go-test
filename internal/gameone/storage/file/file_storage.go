package file

import (
	"bufio"
	"encoding/gob"
	"errors"
	"fmt"
	"os"
)

var magicNumber int = 0
var fileVersion int = 1

type FileStorage struct{}

func (storage FileStorage) WriteLastPosition(sceneId string) error {
	// Open config file
	file, err := os.Create(`last_position.txt`)
	if err != nil {
		return err
	}
	defer file.Close()

	writer := bufio.NewWriter(file)
	defer writer.Flush()

	encoder := gob.NewEncoder(writer)

	if err := encoder.Encode(magicNumber); err != nil {
		return err
	}

	if err := encoder.Encode(fileVersion); err != nil {
		return err
	}

	return encoder.Encode(sceneId)
}

func (storate FileStorage) ReadLastPosition() (*string, error) {

	// Open config file
	file, err := os.Open(`last_position.txt`)
	if err != nil {
		if os.IsNotExist(err) {
			return nil, nil
		} else {
			return nil, err
		}
	}
	defer file.Close()

	reader := bufio.NewReader(file)
	decoder := gob.NewDecoder(reader)

	var magic int
	if err := decoder.Decode(&magic); err != nil {
		return nil, err
	}

	if magic != magicNumber {
		return nil, errors.New(`cannot read non-last position gob file`)
	}

	var version int
	if err := decoder.Decode(&version); err != nil {
		return nil, err
	}
	if version > fileVersion {
		return nil, fmt.Errorf(`version %d is too new to read`, version)
	}

	var sceneId string
	err = decoder.Decode(&sceneId)

	return &sceneId, err
}
