# Rsz!

## A really simple CLI image resizer.

Hi there, folks!

This is **rsz**, a simple image resize tool, with some cool features and a fluent syntax, that I built for two main reasons:

* I am learning Go, so I need some simple training projects;
* my girlfriend needed an image resize tool for her university projects;

Take a look to the "Usage" section to see rsz in action!

If you want to learn more about features, go to the CHANGELOG.md file. I am going to update it at every release, also with future ideas. However, don't take it too much seriously: it's just a learning project.

### Compile

Nothing really special, just go ahead with `go install`, or `go install -ldflags '-s'` to strip out debugging data and reduce final executable size.

### Usage

#### Simple Resize

This will resize every image in the folder to a width of 1024. The height will be calculated accordingly.

    rsz

#### Choose Resized Image Size

When not specified, a dimension will be calculated by maintaining proportions. If neither of them were specified, the width will be 1024px automatically and the height will be calculated accordingly.

    rsz --width "800"
    rsz --height "400"
    rsz --width "800" --height "400"

#### Choose Output format

You can also convert the resized images to a certain format, if you want. Just use the "to" flag.

    rsz --to "jpg"
    rsz --to "png"
    rsz --to "tiff"

#### Choose Output Sub-Directory

Maybe you have tons of image to resize and you just don't want to put everything in the same folder.

    rsz --in "resized"

The resized images will be saved in `current_dir/resized` folder. This also works recursively:

    rsz --in "resized/images/here"

#### Filter Input by Format

Let's say you have some JPEGs and PNGs, but you just want to resize the PNG files. No problem, just use

    rsz --only "png"

and you're good to go!
