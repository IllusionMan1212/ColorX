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

package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"

	"github.com/illusionman1212/colorx"
)

func main() {
	args := os.Args
	filePath := "example.jpg"

	if len(args) > 1 {
		filePath = args[1]
	}

	data, err := ioutil.ReadFile(filePath)
	if err != nil {
		log.Fatal(err)
	}

	color, err := colorx.GetProminentColor(data)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("The most prominent color is: %v\n", color)
}
