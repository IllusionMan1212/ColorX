// ColorX library project
// Copyright (C) 2021 IllusionMan1212 and contributors
//
// This program is free software; you can redistribute it and/or modify
// it under the terms of the GNU General Public License as published by
// the Free Software Foundation; either version 3 of the License, or
// (at your option) any later version.
//
// This program is distributed in the hope that it will be useful,
// but WITHOUT ANY WARRANTY; without even the implied warranty of
// MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE. See the
// GNU General Public License for more details.
//
// You should have received a copy of the GNU General Public License along
// with this program; if not, see https://www.gnu.org/licenses.

package colorx

import (
	"bytes"
	"fmt"
	"image"

	_ "image/jpeg"
	_ "image/png"
)

// GetProminentColor gets the most prominent color in the image bytes reader that gets passed in.
// returns a hex color as a string.
func GetProminentColor(data []byte) (string, error) {
	buf := bytes.NewBuffer(data)
	img, _, err := image.Decode(buf)
	if err != nil {
		return "", err
	}

	width := img.Bounds().Max.X
	height := img.Bounds().Max.Y

	pixels := height * width

	fmt.Printf("Pixels in image: %v\n", pixels)
	fmt.Printf("Width: %v\n", width)
	fmt.Printf("Height: %v\n", height)

	return "", nil
}
