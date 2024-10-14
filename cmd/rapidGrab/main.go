package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
)

func main() {
	downloadFile("/Users/ashour/Dev/GoProjects/rapidGrab/test1.mp4", "https://videos.pexels.com/video-files/7710516/7710516-hd_1920_1080_25fps.mp4")
}

func downloadFile(file string, url string) (err error) {
	out, err := os.Create(file)
	if err != nil {
		return err
	}
	defer out.Close()
	client := &http.Client{}
	req, _ := http.NewRequest("GET", url, nil)
	req.Header.Set("Range", "bytes=0-1000")

	req2, _ := http.NewRequest("HEAD", url, nil)
	req.Header.Set("Accept-Ranges", "bytes")
	res1, _ := client.Do(req2)
	fmt.Println(res1.StatusCode)

	// head, err := http.Head(url)
	res, _ := client.Do(req)
	// fmt.Println(head.ContentLength)
	if err != nil {
		return err
	}

	defer res.Body.Close()

	io.Copy(out, res.Body)
	return nil
}
