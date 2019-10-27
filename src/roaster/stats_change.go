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

type StatsChange struct {
	offset int
	speed uint16
	tpts uint16
	dunks uint16
	pass uint16
	power uint16
	steal uint16
	block uint16
	cltch uint16
}


func (c *StatsChange) ParseParameters(pairs map[string]string) error {
	var err error

	c.offset, err = intFromString(pairs,"offset")
	if err != nil {
		return err
	}

	c.speed, err  = uint16FromString(pairs, "speed")
	if err != nil {
		return err
	}

	c.tpts, err   = uint16FromString(pairs, "3pts")
	if err != nil {
		return err
	}

	c.dunks, err  = uint16FromString(pairs, "dunks")
	if err != nil {
		return err
	}

	c.pass, err   = uint16FromString(pairs, "pass")
	if err != nil {
		return err
	}

	c.power, err  = uint16FromString(pairs, "power")
	if err != nil {
		return err
	}

	c.steal, err  = uint16FromString(pairs, "steal")
	if err != nil {
		return err
	}

	c.block, err  = uint16FromString(pairs, "block")
	if err != nil {
		return err
	}

	c.cltch, err  = uint16FromString(pairs, "cltch")
	if err != nil {
		return err
	}

	return err
}

func (c *StatsChange) Run() error {
	bytes := mainRom.mergedROM
	writeUint16(bytes, c.offset + 0x0, c.speed) // Speed
	writeUint16(bytes, c.offset + 0x2, c.tpts)  // 3 pts
	writeUint16(bytes, c.offset + 0x4, c.dunks) // dunks
	writeUint16(bytes, c.offset + 0x6, c.pass)  // pass
	writeUint16(bytes, c.offset + 0x8, c.power) // power
	writeUint16(bytes, c.offset + 0xA, c.steal) // steal
	writeUint16(bytes, c.offset + 0xC, c.block) // block
	writeUint16(bytes, c.offset + 0xE, c.cltch) // cltch

	return nil
}