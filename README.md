# ColorX
ColorX is a library to determine the most prominent color in an image.
ColorX doesn't use any sort of complex algorithms to calculate the prominent color, it simply loops over the image pixels and returns the color that occurs the most.

The `getProminentColor` function returns the color in 2 formats, a hex string such as: `#abfecd` and a variable of type `color.Color`.

# Example
An example on how to use this library is written in [example/example.go](https://github.com/illusionman1212/colorx/blob/master/example/example.go)

# Supported Image Formats
Currently ColorX supports these image formats:
- png
- jpeg
- webp
