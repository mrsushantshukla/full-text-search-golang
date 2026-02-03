package utils

import (
	"compress/gzip"
	"encoding/xml"
	"os"
)

type document struct {
	Title string `xml:"logtitle"`
	Text  string `xml:"comment"`
	URL   string `xml:"url"`
	ID    int
}

func LoadDocuments(dumpPath string) ([]document, error) {
	file, err := os.Open(dumpPath)
	if err != nil {
		return nil, err
	}
	defer file.Close()
	gz, err := gzip.NewReader(file)
	if err != nil {
		return nil, err
	}
	defer gz.Close()
	xdec := xml.NewDecoder(gz)
	dump := struct {
		Documents []document `xml:"logitem"`
	}{}
	if err := xdec.Decode(&dump); err != nil {
		return nil, err
	}
	docs := dump.Documents
	for i := range docs {
		docs[i].ID = i
	}
	return docs, nil
}
