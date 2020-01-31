package main

import (
	"bytes"
	"fmt"
	"io"
	"io/ioutil"
	"path"
)

const JSONFILE = "data.dat"

var ErrFileNotFound = fmt.Errorf("file not found")

type JSON struct {
	PATH  []string          `json:"path"`
	Envi  map[string]string `json:"envi"`
	Alias map[string]string `json:"alias"`
}

func newJSON() *JSON {
	return &JSON{
		Envi:  make(map[string]string),
		Alias: make(map[string]string),
	}
}

func (j *JSON) AddPath(path string) error {
	return nil
}

func (j *JSON) SetEnvi(name, value string) error {
	return nil
}

func (j *JSON) SetAlias(name, value string) error {
	return nil
}

func (j *JSON) Decode(r io.Reader) error {
	return nil
}

func (j *JSON) Encode(w io.Writer) error {
	return nil
}

func LoadJSON() (*JSON, error) {
	if file := path.Join(GetRootDir(), JSONFILE); FileExist(file) {
		byts, err := ioutil.ReadFile(file)
		if err != nil {
			return nil, err
		}
		json := newJSON()
		if err := json.Decode(bytes.NewReader(byts)); err != nil {
			return nil, err
		}
		return json, nil
	}
	return nil, ErrFileNotFound
}

func WriteJSON(json *JSON) error {
	buffer := bytes.NewBuffer([]byte{})
	if err := json.Encode(buffer); err != nil {
		return err
	}
	return ioutil.WriteFile(path.Join(GetRootDir(), JSONFILE), buffer.Bytes(), 0x755)
}
