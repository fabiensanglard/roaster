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
	"image"
	"image/color"
	"image/draw"
	"image/png"
	"os"
)

type MenuImageChange struct {
	paletteOffset int
	imageOffset int
	bits_offset int
	filename string
	width int
	height int
	bpp int
	replace_with string
}

func (c *MenuImageChange) ParseParameters(pairs map[string]string) error {
	var err error

	c.paletteOffset,err = intFromString(pairs, "palette_offset")
	if err != nil {
		return err
	}

	c.imageOffset,err  = intFromString(pairs, "image_offset")
	if err != nil {
		return err
	}

	c.bits_offset,err  = intFromString(pairs, "bits_offset")
	if err != nil {
		return err
	}

	c.width,err  = intFromString(pairs, "width")
	if err != nil {
		return err
	}

	c.height,err  = intFromString(pairs, "height")
	if err != nil {
		return err
	}

	c.bpp,err  = intFromString(pairs, "bpp")
	if err != nil {
		return err
	}

	c.replace_with,err  = stringFromString(pairs, "replace_with")
	if err != nil {
		return err
	}

	return nil
}

func (c *MenuImageChange) Run() error {
	bytes, palette, err := loadPNG(c.replace_with, c.width, c.height, c.bpp)
	if err != nil {
		return err
	}
	writeBytes(gfxrom.mergedROM, c.imageOffset, c.bits_offset, bytes)
	writePlayerPalette(mainRom.mergedROM, c.paletteOffset, palette)
	return nil
}


// Debug function
func setPlayerPortraitPalette(bytes []byte, offset int, color byte) {
	for i := 0 ; i < 512 ; i++ { // xRGB555 palette entries are 2-byte long.
		bytes[offset + i] =  color
	}
}

// Debug function
func setImageColor(bytes []byte, offset int, w int, h int, c byte) {
	for i:=0 ; i < w * h ; i++ {
		bytes[offset + i] = c
	}
}

func writePlayerPalette(bytes []byte, offset int, palette color.Palette) {
	for i:= 0; i < len(palette) ; i++  {
		var r,g,b,_ = palette[i].RGBA()

		// RGBA to xRGB_555
		// x RRRRR GGGGG BBBBB
		var rgb555ex uint16 = 0

		var newR = (r >> 8) & 0xFF
		var newG = (g >> 8) & 0xFF
		var newB = (b >> 8) & 0xFF

		rgb555ex |= uint16((newR >> 3) <<10)
		rgb555ex |= uint16((newG >> 3) << 5)
		rgb555ex |= uint16((newB >> 3) << 0)

		bytes[offset + i * 2 + 0] = uint8(rgb555ex & 0x00FF >> 0)
		bytes[offset + i * 2 + 1] = uint8(rgb555ex & 0xFF00 >> 8)

	}
}

func loadPNG(filename string, width int, height int, bpp int) ([]uint8, color.Palette, error) {
	file, err := os.Open(filename)
	if err != nil {
		return nil,nil,err
	}

	defer file.Close()

	img, _ := png.Decode(file)

	if img.Bounds().Max.X != width {
		err := fmt.Sprintf("Error: Image %s must be %d pixels wide", filename, width)
		panic(err)
	}

	if img.Bounds().Max.Y != height {
		err := fmt.Sprintf("Error: Image %s must be %d pixels tall", filename, height)
		panic(err)
	}

	palettedImage := image.NewPaletted(img.Bounds(), img.ColorModel().(color.Palette))

	paletteSize := len(palettedImage.Palette)
	if paletteSize > (1 << uint(bpp)) {
		err := fmt.Sprintf("Error: Image %s palette size must be less or equals %d (found %d).", filename, 1 << uint(bpp), paletteSize)
		panic(err)
	}

	draw.Draw(palettedImage, palettedImage.Rect, img, img.Bounds().Min, draw.Over)
	return palettedImage.Pix, palettedImage.Palette, nil
}