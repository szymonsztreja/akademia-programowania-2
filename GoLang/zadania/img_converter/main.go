package main

import (
	"bytes"
	"encoding/base64"
	"fmt"
	"image"
	"image/gif"
	"image/jpeg"
	"image/png"
	"io"
	"net/http"
	"strings"
)

func main() {
	http.HandleFunc("/", handleIndex)
	http.HandleFunc("/convert", handleConvert)
	http.ListenAndServe(":8080", nil)
}

func handleIndex(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodGet {
		// Serve the HTML form for uploading an image
		html := `
		<!DOCTYPE html>
		<html>
		<head>
			<title>Image Format Converter</title>
		</head>
		<body>
			<h1>Image Format Converter</h1>
			<form action="/convert" method="post" enctype="multipart/form-data">
				<label for="choose_img">Choose an image:</label>
				<input type="file" name="image" id="choose_img" accept=".jpg, .jpeg, .png, .gif" required>
				<br>

				<label for="format_to">Choose a format to convert to:</label>
				<select name="format_to" id="format_to">
  				<option value="JPEG">JPEG</option>
  				<option value="PNG">PNG</option>
  				<option value="GIF">GIF</option>
				</select>
				<button type="submit">Convert</button>
			</form>
		</body>
		</html>
		`
		w.Header().Set("Content-Type", "text/html")
		w.Write([]byte(html))
	}
}

func handleConvert(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodPost {
		// Parse the multipart form containing the uploaded file
		err := r.ParseMultipartForm(10 << 20) // 10 MB max
		if err != nil {
			fmt.Println("Unable to parse form")
			http.Error(w, "Unable to parse form", http.StatusBadRequest)
			return
		}

		// Retrieve the uploaded image file
		file, _, err := r.FormFile("image")
		if err != nil {
			fmt.Println("Please upload an image file")
			http.Error(w, "Please upload an image file", http.StatusBadRequest)
			return
		}
		defer file.Close()

		// Retrieve fromat to convert
		formatTo := r.Form.Get("format_to")
		formatTo = strings.ToLower(formatTo)

		// Determine the file format (JPEG, PNG or GIF)
		img, format, err := image.Decode(file)
		if err != nil {
			fmt.Println("Failed to decode image")
			http.Error(w, "Failed to decode image", http.StatusInternalServerError)
			return
		}

		fmt.Println(format)

		// Encode the image in the "format_to" format
		var newEncoder func(io.Writer, image.Image) error
		switch strings.ToLower(formatTo) {
		case "jpeg", "jpg":
			newEncoder = func(w io.Writer, img image.Image) error {
				return jpeg.Encode(w, img, nil)
			}
		case "png":
			newEncoder = png.Encode
		case "gif":
			newEncoder = func(w io.Writer, img image.Image) error {
				return gif.Encode(w, img, nil)
			}
		}

		// Create bytes buffer and encode the image into the buffer
		var buf bytes.Buffer
		encoderError := newEncoder(&buf, img)
		if encoderError != nil {
			fmt.Println("Failed to convert image")
			http.Error(w, "Failed to convert image", http.StatusInternalServerError)
			return
		}

		// HTML response with image and download button
		html := fmt.Sprintf(`
			<!DOCTYPE html>
			<html>
			<head>
				<title>Converted Image</title>
			</head>
			<body>
				<h1>Converted Image</h1>
				<img src="data:image/%s;base64,%s" alt="Converted Image" style="max-width: 100%%; height: auto;">
				<br>
				<a href="/download" download="converted-image.%s"><button>Download Image</button></a>
			</body>
			</html>
		`, formatTo, base64.StdEncoding.EncodeToString(buf.Bytes()), formatTo)

		// Write HTML response to the client
		w.Header().Set("Content-Type", "text/html")
		w.Write([]byte(html))
	}
}
