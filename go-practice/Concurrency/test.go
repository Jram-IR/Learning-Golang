package main

// import (
// 	"fmt"
// 	"time"
// )

// type ImgStatus struct {
// 	data   []byte
// 	status int   // 200 is ok and 666 is ERROR !
// 	err    error // this is the error that occurred
// }

// func GetImageFormTheUser(data []byte) ImgStatus {
// 	var img ImgStatus

// 	// Process the img attachments sequentially
// 	start := time.Now()

// 	// Resize image
// 	img = ResizeImg(data)

// 	// Add watermark
// 	img = AddWaterMark(img.data)

// 	// Generate thumbnails
// 	img = GenThumbNails(img.data)

// 	fmt.Println("Logging Report")
// 	fmt.Println(img)
// 	if img.err != nil {
// 		fmt.Println("An error occurred while processing", img.err.Error())
// 	} else {
// 		fmt.Println("The image was successfully processed ")
// 	}

// 	fmt.Print("Done in", time.Since(start))

// 	return img
// }

// func ResizeImg(data []byte) ImgStatus {
// 	var img ImgStatus
// 	img.data = data

// 	start := time.Now()
// 	fmt.Println("Resizing Image", start)
// 	time.Sleep(time.Duration(time.Millisecond * 100)) // Simulate resizing
// 	img.err = nil
// 	img.status = 200
// 	fmt.Println("Resized Image", time.Since(start))

// 	return img
// }

// func AddWaterMark(data []byte) ImgStatus {
// 	var img ImgStatus
// 	img.data = data

// 	start := time.Now()
// 	fmt.Println("Applying WaterMark", start)
// 	time.Sleep(time.Duration(time.Millisecond * 120)) // Simulate watermarking
// 	img.err = nil
// 	img.status = 200
// 	fmt.Println("WaterMark Applied", time.Since(start))

// 	return img
// }

// func GenThumbNails(data []byte) ImgStatus {
// 	var img ImgStatus
// 	img.data = data

// 	start := time.Now()
// 	fmt.Println("Generating Thumbnail", start)
// 	time.Sleep(time.Duration(time.Millisecond * 150)) // Simulate thumbnail generation
// 	img.err = nil
// 	img.status = 200
// 	fmt.Println("Thumbnail generated", time.Since(start))

// 	return img
// }

// func main() {
// 	// The user image is converted into bytes first
// 	data := []byte("Some image data") // for example purpose

// 	// Process the image
// 	GetImageFormTheUser(data)
// }
