package main

/*
Author Gaurav Sablok
Universitat Potsdam
Date 2024-8-27

structs and inheritances for making the pacbiohifi genome alignments easier for the web browse and calling through the http. It takes the sam alignment files, parses the alignments and then makes a struct based JSON for easy feed into the database and also writes a CSV for direct import into the SQLITE, MongoDB and PostGresSQL.

*/

import (
	"bufio"
	"log"
	"os"
	"strings"
)

func main() {
	// struct for the sam file define

	type readStore struct {
		idread string
		tag    string
		seq    string
	}

	// reading the fileIO
	openFile, err := os.Open("readFile")
	if err != nil {
		log.Fatal(err)
	}

	// bufferscan and feeding into the struct
	readBuffer := bufio.NewScanner(openFile)
	for readBuffer.Scan() {
		line := readBuffer.Text()
		if strings.HasPrefix(line, "@") {
			continue
		} else {
			data := []readStore{}
			data = append(data, readStore{
				idread: strings.Split(line, "\t")[0],
				tag:    strings.Split(line, "\t")[5],
				seq:    strings.Split(line, "\t")[9]
			})
		}
	}
	// writing the csv for direct ingestion into the SQL, SQLite3, POSTGresSQL
	create, err := os.Create("writefile")
	if err != nil {
		log.Fatal(err)
	}
	readSam := bufio.NewScanner(openFile)
	for readSam.Scan() {
		line := readSam.Text()
		if strings.HasPrefix(line, "@") {
			continue
		} else {
			data := []readStore{}
			data = append(data, readStore{
				idread: strings.Split(line, "\t")[0],
				tag:    strings.Split(line, "\t")[5],
				seq:    strings.Split(line, "\t")[9]
			})
		}
	}
	defer create.Close()
	for i := range data {
		create.WriteString("The mapped alignment reads along with the CIGAR values are:")
		create.WriteString("%s, %s, %s", data.idread, data.tag, data.seq)
		create.WriteString("The file for the sql injection and transaction has been written")
	}
}
