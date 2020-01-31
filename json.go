package main

import (
	"bytes"
	"io"
	"io/ioutil"
	"path"
)

const JSONFILE = "data.dat"

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

func (j *JSON) AddPath(path string) {
	j.PATH = append(j.PATH, path)
}

func (j *JSON) DelPath(pos uint) {
	j.PATH = append(j.PATH[:pos], j.PATH[pos+1:]...)
}

func (j *JSON) SetEnv(name, value string) {
	if len(value) == 0 {
		delete(j.Envi, name)
	} else {
		j.Envi[name] = value
	}
}

func (j *JSON) SetAlias(name, value string) {
	if len(value) == 0 {
		delete(j.Alias, name)
	} else {
		j.Alias[name] = value
	}
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
	return newJSON(), nil
}

func WriteJSON(json *JSON) error {
	buffer := bytes.NewBuffer([]byte{})
	if err := json.Encode(buffer); err != nil {
		return err
	}
	return ioutil.WriteFile(path.Join(GetRootDir(), JSONFILE), buffer.Bytes(), 0x0755)
}
