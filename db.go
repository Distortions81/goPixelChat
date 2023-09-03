package main

import (
	"bytes"
	"encoding/json"
	"image/color"
	"log"
	"os"
)

func WriteDB() {
	tempPath := dbFile + ".tmp"
	finalPath := dbFile

	outbuf := new(bytes.Buffer)
	enc := json.NewEncoder(outbuf)
	enc.SetIndent("", "\t")

	itemList := []color.RGBA{}
	for _, item := range theGrid {
		r, g, b, _ := item.RGBA()
		itemList = append(itemList, color.RGBA{R: uint8(r), G: uint8(g), B: uint8(b), A: 0xff})
	}

	if err := enc.Encode(itemList); err != nil {
		log.Fatal("WriteGCfg: enc.Encode failure")
		return
	}

	_, err := os.Create(tempPath)

	if err != nil {
		log.Fatal("WriteGCfg: os.Create failure")
		return
	}

	err = os.WriteFile(tempPath, outbuf.Bytes(), 0644)

	if err != nil {
		log.Fatal("WriteGCfg: WriteFile failure")
		return
	}

	err = os.Rename(tempPath, finalPath)

	if err != nil {
		log.Fatal("Couldn't rename Gcfg file.")
		return
	}
}
