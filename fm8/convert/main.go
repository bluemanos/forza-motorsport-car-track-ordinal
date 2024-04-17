package main

func main() {
	carsWithKeys, keys := readCars()
	writeCarsKeys(carsWithKeys, keys)
	writeCarsCsv(carsWithKeys, keys)

	tracksWithKeys, keys := readTracks()
	writeTracksKeys(tracksWithKeys, keys)
	writeTracksCsv(tracksWithKeys, keys)
}
