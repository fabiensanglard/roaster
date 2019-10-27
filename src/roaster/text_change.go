/*
Copyright 2019 Google LLC

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

https://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

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
