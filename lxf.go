package bricker

import (
	"archive/zip"
	"encoding/xml"
	"errors"
	"fmt"
	"io/ioutil"
	"strings"
)

type LXFML struct {
	Bricks []LXFBrick `xml:"Bricks>Brick"`
}

type LXFBrick struct {
	XMLName xml.Name  `xml:"Brick"`
	Parts   []LXFPart `xml:"Part"`
}

type LXFPart struct {
	DesignID  string `xml:"designID,attr"`
	Materials string `xml:"materials,attr"`
}

type LFXParser struct {
}

func (self *LFXParser) Parse(file string) (*LXFML, error) {
	r, err := zip.OpenReader(file)
	if err != nil {
		return nil, err
	}
	defer r.Close()

	for _, zf := range r.File {
		if !strings.HasSuffix(zf.Name, ".LXFML") {
			continue
		}

		f, err := zf.Open()
		if err != nil {
			return nil, err
		}

		b, err := ioutil.ReadAll(f)
		if err != nil {
			return nil, err
		}
		fmt.Printf(string(b))

		lxfml := LXFML{}
		err = xml.Unmarshal(b, &lxfml)
		if err != nil {
			return nil, err
		}
		return &lxfml, nil
	}
	return nil, errors.New("No LXFML File found in archive")
}
