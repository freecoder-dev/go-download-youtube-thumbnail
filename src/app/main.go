package main

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"strings"
)

func main() {
	// Prompt the user to enter a YouTube video URL
	fmt.Print("Enter the YouTube video URL: ")
	var videoURL string
	_, err := fmt.Scanln(&videoURL)
	if err != nil {
		log.Fatalf("Failed to read input: %v", err)
	}

	// Parse the video ID from the URL
	videoID := extractVideoID(videoURL)
	if videoID == "" {
		log.Fatal("Invalid YouTube video URL")
	}

	// Construct the thumbnail URL
	thumbnailURL := fmt.Sprintf("https://img.youtube.com/vi/%s/maxresdefault.jpg", videoID)

	// Download the thumbnail image
	response, err := http.Get(thumbnailURL)
	if err != nil {
		log.Fatalf("Failed to download thumbnail: %v", err)
	}
	defer response.Body.Close()

	// Extract the file name from the URL
	fileName := videoID + ".jpg"

	// Create a new file to save the image
	file, err := os.Create(fileName)
	if err != nil {
		log.Fatalf("Failed to create file: %v", err)
	}
	defer file.Close()

	// Copy the image data to the file
	_, err = io.Copy(file, response.Body)
	if err != nil {
		log.Fatalf("Failed to save image: %v", err)
	}

	fmt.Printf("Thumbnail saved: %s\n", fileName)
}

func extractVideoID(url string) string {
	splitURL := strings.Split(url, "?v=")
	if len(splitURL) != 2 {
		return ""
	}

	return splitURL[1]
}
