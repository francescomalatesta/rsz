# Changelog

What's going on?

## Future ideas

* add a nice progress bar to give a better feedback to the user
* add a `--filter` flag to filter images that must be resized, by file name
* create executables for different platforms (Linux, Windows, OSX)
* write some tests

## 0.3

* added a `--verbose` flag that shows more info about the resize process
* added an error message if no images are found
* resized images are now saved in the "resized" folder as a default behavior

## 0.2

* added tiff support for input/output
* added `--in` flag to specify the output folder for resized images
* added `--only` flag to resize photos of a certain format only
* basic code refactoring
* use the mime type for image format recognition

## 0.1

* added `--width` flag to specify width. Default is 1024px;
* added `--height` flag to specify height;
* added `--to` flag to specify output format;
* added png encode and decode support;
* added jpg encode and decode support;
