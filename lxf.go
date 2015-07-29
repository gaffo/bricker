package bricker

import (
	"archive/zip"
	"bufio"
	"encoding/xml"
	"errors"
	"fmt"
	"io/ioutil"
	"os"
	"strings"
)

type PartMap struct {
	ldd2bl map[string]string
	bl2ldd map[string]string
}

func (self *PartMap) LDD2BL(ldd string) string {
	return self.ldd2bl[ldd]
}

func ParsePartsMap(file string) (*PartMap, error) {
	m := PartMap{
		ldd2bl: make(map[string]string),
		bl2ldd: make(map[string]string),
	}
	fmt.Println("Parsing Parts File", file)
	f, err := os.Open(file)
	if err != nil {
		return nil, err
	}

	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		// LLDNAME LDDID BLID
		parts := strings.Split(scanner.Text(), "\t")
		// lddname := parts[0]
		lddid := parts[1]
		blid := parts[2]
		m.ldd2bl[lddid] = blid
		m.bl2ldd[blid] = lddid
		fmt.Printf("%s->%s\n", lddid, blid)
	}

	return &m, nil
}

type ColorMap struct {
	ldd2bl map[string]string
	bl2ldd map[string]string
}

func (self *ColorMap) LDD2BL(ldd string) string {
	return self.ldd2bl[ldd]
}

func ParseColorMap(file string) (*ColorMap, error) {
	m := ColorMap{
		ldd2bl: make(map[string]string),
		bl2ldd: make(map[string]string),
	}

	f, err := os.Open(file)
	if err != nil {
		return nil, err
	}

	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		// LDDID LDDNAME BLID BLNAME HEX
		parts := strings.Split(scanner.Text(), "\t")
		lddid := parts[0]
		// lddname := parts[1]
		blid := parts[2]
		// blname := parts[3]
		// hex := parts[4]
		m.ldd2bl[lddid] = blid
		m.bl2ldd[blid] = lddid
	}

	return &m, nil
}

type BLPart struct {
	ItemID   string
	Color    string
	Quantity int64
}

type LXFML struct {
	Bricks []LXFBrick `xml:"Bricks>Brick"`
}

func (self *LXFML) ConvertWithSources(colors *ColorMap, parts *PartMap) []BLPart {
	bl := make([]BLPart, 0, 1024)

	for _, brick := range self.Bricks {
		if len(brick.Parts) == 0 {
			continue
		}
		part := brick.Parts[0]
		design := part.DesignID
		materials := part.Materials
		materials = strings.Split(materials, ",")[0]
		bl = append(bl, BLPart{
			ItemID:   parts.LDD2BL(design),
			Color:    colors.LDD2BL(materials),
			Quantity: 1,
		})
	}

	return bl
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

		lxfml := LXFML{}
		err = xml.Unmarshal(b, &lxfml)
		if err != nil {
			return nil, err
		}
		return &lxfml, nil
	}
	return nil, errors.New("No LXFML File found in archive")
}
