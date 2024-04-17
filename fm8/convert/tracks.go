package main

import (
	"encoding/json"
	"fmt"
	"io"
	"os"
	"sort"
)

type Track struct {
	TrackOrdinal int
	Circuit      string
	Location     string
	IOC_Code     string
	Track        string
	Length_in_km float32
}

func readTracks() (map[int]Track, []int) {
	tracksJsonFile, err := os.Open("../tracks.json")
	if err != nil {
		fmt.Println(err)
	}
	defer tracksJsonFile.Close()
	fmt.Println("tracks.json opened")

	byteValue, err := io.ReadAll(tracksJsonFile)
	if err != nil {
		fmt.Println(err)
	}
	var tracks []Track
	json.Unmarshal(byteValue, &tracks)

	tracksWithKeys := make(map[int]Track)
	keys := make([]int, 0, len(tracks))
	for _, track := range tracks {
		tracksWithKeys[track.TrackOrdinal] = track
		keys = append(keys, track.TrackOrdinal)
	}
	sort.Ints(keys)

	return tracksWithKeys, keys
}

func writeTracksKeys(tracksWithKeys map[int]Track, keys []int) {
	fileJson, err := os.OpenFile("../tracks_keys.json", os.O_RDWR|os.O_CREATE|os.O_TRUNC, 0755)
	if err != nil {
		fmt.Println(err)
	}
	defer fileJson.Close()

	jsonString, err := json.MarshalIndent(tracksWithKeys, "", "    ")
	if err != nil {
		fmt.Println(err)
	}
	_, err = fileJson.Write(jsonString)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println("tracks_keys.json saved")
}

func writeTracksCsv(tracksWithKeys map[int]Track, keys []int) {
	fileCsv, err := os.OpenFile("../tracks.csv", os.O_RDWR|os.O_CREATE|os.O_TRUNC, 0755)
	if err != nil {
		fmt.Println(err)
	}
	defer fileCsv.Close()
	for _, key := range keys {
		fmt.Fprintf(
			fileCsv,
			fmt.Sprintf(
				"%d,%s,%s,%s,%s,%.2f\n",
				tracksWithKeys[key].TrackOrdinal,
				tracksWithKeys[key].Circuit,
				tracksWithKeys[key].Location,
				tracksWithKeys[key].IOC_Code,
				tracksWithKeys[key].Track,
				tracksWithKeys[key].Length_in_km,
			),
		)
	}
	fmt.Println("tracks.csv saved")
}
