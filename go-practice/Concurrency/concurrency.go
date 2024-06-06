// // Imagine you're building a social media platform where users can upload posts
// // containing images. After a user uploads a post, the system needs to process the
// // image attachments in the background to perform tasks such as resizing images for
// // different display sizes, adding watermarks, and generating thumbnails.
// // Users should be able to upload posts and continue using the platform without
// // waiting for image processing to complete.

// // channels goroutines waitgroups
package main

import (
	"fmt"
	"sync"
	"time"
)

type ImgStatus struct {
	data   []byte
	status int   // 200 is ok and 666 is ERROR !
	err    error // this is the error that occured

}

var wg sync.WaitGroup

//wg.Add, wg.Done, wg.Wait

func GetImageFormTheUser(ch chan ImgStatus, data []byte) {
	//process the img attachments

	ProcessImageInBg(ch, data)

}

// log report
func log(ch chan ImgStatus) {
	defer wg.Done()
	img := <-ch // blocking statement
	fmt.Println("Logging Report")
	fmt.Println(img)
	if img.err != nil {
		fmt.Println("An error occured while processing", img.err.Error())
	} else {
		fmt.Println("The image was succesfully processed ")
	}
	// based on the final status of the img struct you can do more operations if neccsary
}

func ProcessImageInBg(ch chan ImgStatus, data []byte) {
	start := time.Now()
	//add the data to the struct
	img := ImgStatus{
		data:   data,
		status: 200,
		err:    nil,
	}
	// add the updated image Struct to the channel

	wg.Add(3)
	//resize
	go ResizeImg(ch)
	//add watermarks
	go AddWaterMark(ch)
	//generate thumbnails
	go GenThumbNails(ch)
	ch <- img

	//add the log
	wg.Add(1)
	go log(ch)
	wg.Wait()
	fmt.Print("Done in", time.Since(start))

}

func ResizeImg(ch chan ImgStatus) {
	defer wg.Done()
	img := <-ch
	data := img.data //get the data like this then
	//resize the image
	start := time.Now()
	fmt.Println("Resizing Image", start)
	time.Sleep(time.Duration(time.Millisecond * 100)) //sim
	var err error = nil                               // this actually would be the error encountered while resizing

	//update the err in img when encountered and change status
	img.err = err
	img.status = 666

	//else
	img.data = data // update the data
	ch <- img       //update the img in the channel
	fmt.Println("Resized Image", time.Since(start))

}

func AddWaterMark(ch chan ImgStatus) {
	defer wg.Done()

	img := <-ch
	data := img.data //get the data like this then
	//add watermark to the image
	start := time.Now()
	fmt.Println("Applying WaterMark", start)
	time.Sleep(time.Duration(time.Millisecond * 120)) //sim
	var err error = nil                               // this actually would be the error encountered while watermarking

	//update the err in img when encountered and change status
	img.err = err
	img.status = 666

	//else
	img.data = data // update the data
	ch <- img       //update the img in the channel
	fmt.Println("WaterMark Applied", time.Since(start))

}

func GenThumbNails(ch chan ImgStatus) {
	defer wg.Done()
	img := <-ch
	data := img.data //get the data like this then
	//generate the thumbnails
	start := time.Now()
	fmt.Println("Generating Thumbnail", start)
	time.Sleep(time.Duration(time.Millisecond * 150)) //sim
	var err error = nil                               // replace actual error encountered while generating thumbnail

	//update the err in img when encountered and change status
	img.err = err
	img.status = 666

	//else
	img.data = data // update the data
	ch <- img       //update the img in the channel
	fmt.Println("Thumbnail generated", time.Since(start))

}

func main() {

	//the user image is converted into bytes first
	data := []byte("Some image data") // for exmple purpose

	//make the channel
	mChannel := make(chan ImgStatus)
	GetImageFormTheUser(mChannel, data) // user data is passed

}
