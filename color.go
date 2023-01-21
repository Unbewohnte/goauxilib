/*
The MIT License (MIT)

Copyright © 2022 Kasyanov Nikolay Alexeyevich (Unbewohnte)

Permission is hereby granted, free of charge, to any person obtaining a copy of this software and associated documentation files (the “Software”), to deal in the Software without restriction, including without limitation the rights to use, copy, modify, merge, publish, distribute, sublicense, and/or sell copies of the Software, and to permit persons to whom the Software is furnished to do so, subject to the following conditions:

The above copyright notice and this permission notice shall be included in all copies or substantial portions of the Software.

THE SOFTWARE IS PROVIDED “AS IS”, WITHOUT WARRANTY OF ANY KIND, EXPRESS OR IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY, FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM, OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN THE SOFTWARE.
*/

package goauxilib

import "math"

// Convert color channel to linear representation
func CToLinear(cc uint8) float64 {
	fcc := float64(cc) / 255.0

	if fcc <= 0.04045 {
		return fcc / 12.92
	} else {
		return math.Pow(((fcc + 0.055) / 1.055), 2.4)
	}
}

// Convert a complete RGB value into linear representation
func RGBToLinear(r uint8, g uint8, b uint8) (float64, float64, float64) {
	return CToLinear(r), CToLinear(g), CToLinear(b)
}

// Get approximate RGB color luminance (brightness) from 0 to 1
func LuminanceRGB(r uint8, g uint8, b uint8) float64 {
	rLin, gLin, bLin := RGBToLinear(r, g, b)
	return 0.2126*rLin + 0.7152*gLin + 0.0722*bLin
}
