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

## Authors
photosorcery is a project by <a href="http://dawnlabs.io/"><img width=8% src="https://cloud.githubusercontent.com/assets/10369094/25406306/dacebd4c-29cb-11e7-8e1c-468687cde495.png"></a>
- Brian Dennis ([@briandennis](https://github.com/briandennis))
