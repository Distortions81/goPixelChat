package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"image/color"
	"log"
	"os"
)

func WriteDB() {
	tempPath := dbFile + ".tmp"
	finalPath := dbFile

	outbuf := new(bytes.Buffer)
	enc := json.NewEncoder(outbuf)

	var itemList []itemData
	for pos, item := range theGrid {
		if pos.X < 1 || pos.Y < 1 {
			continue
		}
		if item.R == 0 && item.G == 0 && item.B == 0 {
			continue
		}
		r, g, b, _ := item.RGBA()
		itemList = append(itemList, itemData{X: pos.X, Y: pos.Y, R: uint8(r), G: uint8(g), B: uint8(b)})
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

/* Read in cached list of Discord players with specific roles */
func readDB() {

	_, err := os.Stat(dbFile)
	notfound := os.IsNotExist(err)

	if !notfound { /* Otherwise just read in the config */
		file, err := os.ReadFile(dbFile)

		if file != nil && err == nil {
			var itemList []itemData

			fmt.Println("Reading db.")

			err := json.Unmarshal([]byte(file), &itemList)
			if err != nil {
				log.Fatal("Readcfg.RoleList: Unmarshal failure")
			}

			gridLock.Lock()
			defer gridLock.Unlock()

			for _, item := range itemList {
				if item.X < 1 && item.Y < 1 {
					continue
				}
				if item.X > int(gridSize) || item.Y > int(gridSize) {
					continue
				}
				if item.R == 0 || item.G == 0 || item.B == 0 {
					continue
				}
				theGrid[XY{X: item.X, Y: item.Y}] = color.NRGBA{R: item.R, G: item.G, B: item.B, A: 0xFF}
				fmt.Println(item)
			}
		} else {
			fmt.Println("No database file.")
		}
	}
}
