# Retina

A command-line utility that inspects and identifies JPG images. The same utility can be used as an S3 Lambda trigger to examine every photo uploaded to a bucket. The resulting analysis is a structured JSON object, which can be forwarded to a configurable endpoint for saving into a database. It's ideal for web developers who want to automate creating and cataloguing images uploaded to a website.

## Features

* Extract all [EXIF](https://en.wikipedia.org/wiki/Exif) data in a structured format.
* Return the 5 most prominent colours within the image, ideal for using as a background colour to a HTML `picture` element.
* Calculate the [BlurHash](https://blurha.sh/) value
* Generate a blurred thumbnail, ideal for using as a background image to a HTML `picture` element.
* Optionally use [AWS Rekognition](https://aws.amazon.com/rekognition/) to identify image contents, ideal for auto tagging images. * Additional costs apply.
* Auto-generate a [webp](https://developers.google.com/speed/webp) version of the file.

## Usage

When run from the command-line, you specify the input file like:

``` bash
./retina -i sunset.jpg
```

## Image Identification

![Bird](/test-images/bird.jpg?raw=true)

For the example above image, the following identifications would returned:

* Bird
* Animal
* Jay, Finch
* Blue Jay
