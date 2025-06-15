package converter

import (
	"fmt"
	"log"
	"sync"
)

type WorkerPool struct {
	Images      []string
	Concurrency int
	sendChan    chan string
	wg          sync.WaitGroup
}

func (wp *WorkerPool) Worker(outputPath, convertionType string) {
	for img := range wp.sendChan {

		if convertionType == "JPG to PNG" {
			err := ConvertToPNG(img, outputPath)
			if err != nil {
				log.Printf("error converting %s to png %v", img, err)
			}
		} else if convertionType == "PNG to JPG" {
			err := ConvertToJPG(img, outputPath)
			if err != nil {
				log.Printf("error converting %s to jpg %v", img, err)
			}
		}
		fmt.Printf("convertion of '%s completed\n", img)
		wp.wg.Done()

	}
}

func (wp *WorkerPool) Run(convertionType, outputPath string) {

	wp.sendChan = make(chan string, len(wp.Images))
	for i := 0; i < wp.Concurrency; i++ {
		go wp.Worker(convertionType, outputPath)

	}

	wp.wg.Add(len(wp.Images))
	for _, image := range wp.Images {
		wp.sendChan <- image
	}

	close(wp.sendChan)
	wp.wg.Wait()

}
