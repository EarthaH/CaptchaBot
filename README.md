# CaptchBot
Text-based Image Captcha Bot RESTful API that makes use of the [gosseract](github.com/otiai10/gosseract/v2) library to read and return basic text in an image.

## Requirements
- go 1.19+
- tesseract 5.2.0

Get started with Go [Here](https://go.dev/learn/)

## Build and Run
Build the main.go file into an executable called `main`
```
go build cmd/main/main.go
```
Now you can run `main` which will make the API available at `http://localhost:8080`
```
./main
```

## Usage
### Read Specified Image `/image/:filename`
This returns a json response containing the names of all files saved in the set directory.
```
curl http://localhost:8080/image/sample0.png
"Hello World"
```
If the image does not exist, HTTP response 500 is returned with an empty string and the error is logged.
