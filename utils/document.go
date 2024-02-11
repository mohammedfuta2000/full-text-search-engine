package utils

import (
	"compress/gzip"
	"encoding/xml"
	"os"
)

type document struct {
	Title string `xml:"title"`
	Url   string `xml:"url"`
	Text  string `xml:"abstract"`
	ID    int
}
// convert from .xml.gz to document struct
func LoadDocuments(path string)([]document, error){
	f, err:= os.Open(path); if err!=nil {return nil, err}
	defer f.Close()
	gz, err:=gzip.NewReader(f); if err!=nil {return nil, err}
	defer gz.Close()
	dec:= xml.NewDecoder(gz)
	dump:= struct{
		Document []document `xml:"doc"`
	}{}
	if err:= dec.Decode(&dump); err!=nil {return nil, err}
	docs := dump.Document
	for i:= range docs{ 
		docs[i].ID = i
	}
	return docs, err
}
