# Changelog

What's going on?

## Future ideas

* add a nice progress bar to give a better feedback to the user
* add a `--stats` flag to show some conversion data after the entire process
* add a `--filter` flag to filter images that must be resized, by file name
* create executables for different platforms (Linux, Windows, OSX)
* separate code and responsibilities in a better way
* write some tests

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
