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
	"strconv"

	"image/color"
	_ "image/jpeg"
	_ "image/png"

	_ "golang.org/x/image/webp"
)

// downsampleImage downsamples images larger than 20000px for faster image processing.
func downsampleImage(img image.Image) image.Image {
	// TODO: implement downsampling
	return img
}

// calculateColorMap loops through the image pixels and creates a map-
// of all unique pixel colors and how many times they occur.
func calculateColorMap(img image.Image, width int, height int) map[color.Color]int {
	colorMap := map[color.Color]int{}

	for x := 0; x < width; x++ {
		for y := 0; y < height; y++ {
			color := img.At(x, y)
			// if the alpha is bigger than 100 then add the color to the map
			if _, _, _, a := color.RGBA(); a > 100 {
				colorMap[color] = colorMap[color] + 1
			}
		}
	}

	return colorMap
}

// GetProminentColor gets the most prominent color in the image bytes reader that gets passed in.
// returns a hex color as a string.
func GetProminentColor(data []byte) (string, color.Color, error) {
	buf := bytes.NewBuffer(data)
	img, _, err := image.Decode(buf)
	if err != nil {
		return "", nil, err
	}

	width := img.Bounds().Max.X
	height := img.Bounds().Max.Y

	// amount of pixels in image
	pixels := height * width

	// image is too big, downsample to a reasonable size
	if pixels > 20000 {
		img = downsampleImage(img)
	}

	colorMap := calculateColorMap(img, width, height)

	var maxColor color.Color
	maxAmmount := 0

	for color, amount := range colorMap {
		if amount > maxAmmount {
			maxColor = color
			maxAmmount = amount
		}
	}

	RGBAColor := color.RGBAModel.Convert(maxColor).(color.RGBA)

	red := strconv.FormatUint(uint64(RGBAColor.R), 16)
	green := strconv.FormatUint(uint64(RGBAColor.G), 16)
	blue := strconv.FormatUint(uint64(RGBAColor.B), 16)

	// prepend 0 if the hex number has less than 2 digits
	if len(red) < 2 {
		red = "0" + red
	}

	if len(green) < 2 {
		green = "0" + green
	}

	if len(blue) < 2 {
		blue = "0" + blue
	}

	hexColor := fmt.Sprintf("#%2s%2s%2s", red, green, blue)

	return hexColor, maxColor, nil
}
