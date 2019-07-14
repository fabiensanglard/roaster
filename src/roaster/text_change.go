package roaster

import (
	"errors"
	"fmt"
)

type TextChange struct {
	offset int
	maxLength int
	text string
}

func (c *TextChange) ParseParameters(pairs map[string]string) error {
	var err error

	c.offset,err = intFromString(pairs, "offset")
	if err != nil {
		return err
	}

	c.maxLength,err = intFromString(pairs, "max_length")
	if err != nil {
		return err
	}

	c.text, err = stringFromString(pairs, "new_text")
	if err != nil {
		return err
	}

	if (len(c.text) > c.maxLength) {
		error := fmt.Sprintf("Error, cannot change text with '%s' (not enough space)", c.text)
		return errors.New(error)
	}

	return nil
}

func (c *TextChange) Run() error {
	bytes := mainRom.mergedROM

	var byteArray = []byte(c.text)
	for i := 0 ; i < len(byteArray) ; i++ {
		bytes[c.offset + i] = byteArray[i]
	}
	bytes[c.offset + len(byteArray)] = 0

	return nil
}
