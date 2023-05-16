package main

import "fmt"

func main() {
	srt := "IgiuguibvIUKFGIUKVkujVBugFVUKIJbJgfvOIUVLJvbIUYvLJbKJvIUbvLJKB"
	srtNew := ""
	for i := 0; i < len(srt); i++ {
		if len(srt) != i+1 {
			srtNew += (string(srt[i]) + "*")
		} else {
			srtNew += (string(srt[i]))
		}
	}
	fmt.Println(srtNew)
}
