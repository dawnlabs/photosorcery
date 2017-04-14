![gandalf](http://38.media.tumblr.com/d2d97e94e4b5bd5d9334cf2f85f88ddf/tumblr_n8tqcqPhPB1thowlgo1_250.gif)

# photosorcery

Small CLI for converting between image types and merging images into PDFs.

Converts images concurrently for added speed.

Developed for use with [Alchemy](https://github.com/dawnlabs/alchemy)

## API

### Merge

`photosorcery merge -out [outputPath] [...imagePaths]`

### Convert

`photosorcery convert -type [imageType] -out [outputDirectory] [...imagePaths]`

## File Types Supported

### Convert

  To and from: `jpg`, `png`, `gif`, `bmp`, `tiff`
  
### Merge 

  To: `pdf`
  
  From: `jpg`, `png`, `gif`
