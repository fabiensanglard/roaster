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
	"fmt"
	"io/ioutil"
	"os"
	"strings"
)

func check(e error) {
	if e != nil {
		panic(e)
	}
}

var BASE_SRC = "nbajamte/"
var BASE_DST = ""

type EPROM struct {
	name string
	data []byte
}

var EPROM_SIZE = 524288

type GFXROM_UNIT struct {
	eproms    [4]EPROM
	mergedROM []byte
}

type GFXROM struct {
	roms      [4]GFXROM_UNIT
	mergedROM []byte
}

var gfxrom GFXROM

type MAINROM struct {
	rom       [2]EPROM
	mergedROM []byte
}

var mainRom MAINROM

func initialize() {
	gfxrom.roms[0].eproms[0].name = "l1_nba_jam_tournament_game_rom_ug14.ug14"
	gfxrom.roms[0].eproms[1].name = "l1_nba_jam_tournament_game_rom_uj14.uj14"
	gfxrom.roms[0].eproms[2].name = "l1_nba_jam_tournament_game_rom_ug19.ug19"
	gfxrom.roms[0].eproms[3].name = "l1_nba_jam_tournament_game_rom_uj19.uj19"

	gfxrom.roms[1].eproms[0].name = "l1_nba_jam_tournament_game_rom_ug16.ug16"
	gfxrom.roms[1].eproms[1].name = "l1_nba_jam_tournament_game_rom_uj16.uj16"
	gfxrom.roms[1].eproms[2].name = "l1_nba_jam_tournament_game_rom_ug20.ug20"
	gfxrom.roms[1].eproms[3].name = "l1_nba_jam_tournament_game_rom_uj20.uj20"

	gfxrom.roms[2].eproms[0].name = "l1_nba_jam_tournament_game_rom_ug17.ug17"
	gfxrom.roms[2].eproms[1].name = "l1_nba_jam_tournament_game_rom_uj17.uj17"
	gfxrom.roms[2].eproms[2].name = "l1_nba_jam_tournament_game_rom_ug22.ug22"
	gfxrom.roms[2].eproms[3].name = "l1_nba_jam_tournament_game_rom_uj22.uj22"

	gfxrom.roms[3].eproms[0].name = "l1_nba_jam_tournament_game_rom_ug18.ug18"
	gfxrom.roms[3].eproms[1].name = "l1_nba_jam_tournament_game_rom_uj18.uj18"
	gfxrom.roms[3].eproms[2].name = "l1_nba_jam_tournament_game_rom_ug23.ug23"
	gfxrom.roms[3].eproms[3].name = "l1_nba_jam_tournament_game_rom_uj23.uj23"

	mainRom.rom[0].name = "l4_nba_jam_tournament_game_rom_uj12.uj12"
	mainRom.rom[1].name = "l4_nba_jam_tournament_game_rom_ug12.ug12"

	if len(os.Args) < 1 {
		panic("Usage: go run src/main/main.go OUTPUT_DIRECTORY");
	}
	BASE_DST =os.Args[1]
	if !strings.HasSuffix(BASE_DST, "/") {
		BASE_DST += "/"
	}
}

func deinterlaceGFXROM() {
	fmt.Println("Deinterlacing gfxrom")

	// Load all EPROMs
	for i := 0; i < len(gfxrom.roms); i++ {
		for j := 0; j < len(gfxrom.roms[i].eproms); j++ {
			dat, err := ioutil.ReadFile(BASE_SRC + gfxrom.roms[i].eproms[j].name)
			check(err)
			gfxrom.roms[i].eproms[j].data = dat
		}
	}

	// Allocate 8MiB
	size := len(gfxrom.roms) * len(gfxrom.roms[0].eproms) * EPROM_SIZE
	gfxrom.mergedROM = make([]byte, size)

	for i := 0; i < len(gfxrom.roms); i++ {
		deinterlaceROM_UNIT(&gfxrom.roms[i])
		offset := EPROM_SIZE * 4 * i
		copy(gfxrom.mergedROM[offset:], gfxrom.roms[i].mergedROM)
	}

	f, err := os.Create("gfxrom.bin")
	check(err)
	defer f.Close()
	f.Write(gfxrom.mergedROM)
}

func deinterlaceROM_UNIT(rom *GFXROM_UNIT) {
	rom.mergedROM = make([]byte, EPROM_SIZE*4)

	var cursor = 0
	for i := 0; i < EPROM_SIZE; i++ {
		rom.mergedROM[cursor] = rom.eproms[0].data[i];
		cursor += 1
		rom.mergedROM[cursor] = rom.eproms[1].data[i];
		cursor += 1
		rom.mergedROM[cursor] = rom.eproms[2].data[i];
		cursor += 1
		rom.mergedROM[cursor] = rom.eproms[3].data[i];
		cursor += 1
	}
}

func deinterlaceMAINROM() {
	fmt.Println("Deinterlacing mainrom")
	for i := 0; i < len(mainRom.rom); i++ {
		dat, err := ioutil.ReadFile(BASE_SRC + mainRom.rom[i].name)
		check(err)
		mainRom.rom[i].data = dat
		if EPROM_SIZE != len(dat) {
			return
		}
	}

	f, err := os.Create("mainrom.bin")
	check(err)
	defer f.Close()

	mainRom.mergedROM = make([]byte, EPROM_SIZE*2)
	var cursor = 0
	for i := 0; i < EPROM_SIZE; i++ {
		mainRom.mergedROM[cursor] = mainRom.rom[0].data[i];
		cursor += 1
		mainRom.mergedROM[cursor] = mainRom.rom[1].data[i];
		cursor += 1
	}
	f.Write(mainRom.mergedROM)
}

func interlaceMainROM() {
	var cursor = 0
	for ; cursor < EPROM_SIZE; cursor++ {
		mainRom.rom[0].data[cursor] = mainRom.mergedROM[cursor*2    ]
		mainRom.rom[1].data[cursor] = mainRom.mergedROM[cursor*2+1]
	}

	for i := 0; i < len(mainRom.rom); i++ {
		f, err := os.Create(BASE_DST + mainRom.rom[i].name)
		check(err)
		defer f.Close()
		_, err = f.Write(mainRom.rom[i].data)
	}
}

func interlaceGFXROM_UNIT(rom *GFXROM_UNIT) {
	for cursor := 0; cursor < EPROM_SIZE; cursor++ {
		rom.eproms[0].data[cursor] = rom.mergedROM[cursor*4+0]
		rom.eproms[1].data[cursor] = rom.mergedROM[cursor*4+1]
		rom.eproms[2].data[cursor] = rom.mergedROM[cursor*4+2]
		rom.eproms[3].data[cursor] = rom.mergedROM[cursor*4+3]
	}

	for i := 0; i < len(rom.eproms); i++ {
		f, err := os.Create(BASE_DST + rom.eproms[i].name)
		check(err)
		defer f.Close()
		f.Write(rom.eproms[i].data)
	}
}

func interlaceGFXROMs() {
	for i := 0; i < len(gfxrom.roms); i++ {
		var start = EPROM_SIZE * 4 * i
		gfxrom.roms[i].mergedROM = gfxrom.mergedROM[start : start+EPROM_SIZE*4]
		interlaceGFXROM_UNIT(&gfxrom.roms[i])
	}
}

func writeUint16(bytes []byte, offset int, value uint16) {
	bytes[offset+0] = uint8((value & 0xFF00) >> 8)
	bytes[offset+1] = uint8((value & 0x00FF) >> 0)
}

func writeBytes(dst []byte, offset int, bitOffset int, src []byte) {

	//if bitOffset != 0 {
	//	for i := 0; i < len(src); i++ {
	//		dst[offset+i] = 16 //src[i]
	//	}
	//	return
	//}
	//// Working for non bitOffset = 0
	//for i := 0; i < len(src); i++ {
	//	dst[offset + i] = src[i]
	//}

	var shift = uint(8 - bitOffset)
	var mask = uint16(0xFF) << shift

	for i := 0; i < len(src); i++ {
		var p_src uint16
		p_src = uint16(src[i])
		p_src = p_src << shift

		p_dst := uint16(uint16(dst[offset + i] + 0) << 8 ) | (uint16(dst[offset + i + 1]) )
		p_dst = (p_dst & ^mask) | (p_src & mask)

		dst[offset + i + 0] = uint8(p_dst >> 8)
		dst[offset + i + 1] = uint8(p_dst)
	}
}

func Roast() {
	initialize()

	deinterlaceGFXROM()
	deinterlaceMAINROM()

	Dispatch()

	interlaceMainROM()
	interlaceGFXROMs()
}
