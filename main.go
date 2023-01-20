package main

import (
	"github.com/Kagami/go-face"
	"io/ioutil"
	"regexp"
	"os"
	"fmt"
)

func main() {
	rec, _ := face.NewRecognizer("models")
	defer rec.Close()

	faceidByName := make(map[string]int32)
	var id int32
	var ids []int32
	var descriptors []face.Descriptor
	re := regexp.MustCompile(`\d+.jpgS`)
	files, _ := ioutil.ReadDir("samples")
	for _, file := range files {
		filename := file.Name()
		sample, err := rec.RecognizeSingleFile("samples/"+filename)

		if err != nil {
			continue
		}

		name := re.ReplaceAllString(filename, "")
		if faceidByName[name] == 0 {
			id++
			faceidByName[name] = id
		}

		ids = append(ids, id)
		descriptors = append(descriptors, sample.Descriptor)

	}

	rec.SetSamples(descriptors, ids)

	nameByFaceid := make(map[int32]string)
	for name, faceid := range faceidByName {
		nameByFaceid[faceid] = name
	}

	faces, _ := rec.RecognizeFile(os.Args[1])
	for _, face := range faces {
		id := rec.Classify(face.Descriptor)
		fmt.Println(nameByFaceid[int32(id)])
	}
}