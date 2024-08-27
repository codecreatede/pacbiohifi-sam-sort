package main 

/*
Author Gaurav Sablok
Universitat Potsdam
Date 2024-8-27

structs and inheritances for making the pacbiohifi genome alignments easier
for the web browse and calling through the http. It takes the alignment files,
parses the alignments and then makes a struct based JSON for easy feed into the
database and also writes a CSV for direct import into the SQLITE, MongoDB and PostGresSQL.

*/

import (
   "bufio"
   "fmt"
   "log"
   "os"
   "strings"
)


func main () {

type readStore struct {
	idread string
	tag int
	seq string
}


readFile := os.Args[1:]
writefile := os.Args[2:]
if len(readFile) == nil {
	panic(err)
	log.Fatal(err)
} else {
	openFile, err := os.Open("readFile")
	if err != nil {
		panic(err)
		log.Fatal(err)
	}
	readBuffer := bufio.NewScanner(openFile)
	for readBuffer.Scan() {
		line := readBuffer.Text()
		if strings.HasPrefix(line, "@") {
			continue
		} else {
			storeRead := []readStore{}
			storeRead = append(storeRead, readStore{idread: strings.Split(line, "\t")[0],
				    tag: strings.Split(line, "\t")[5],
				    seq: strings.Split(line,"\t")[9],
			}
		}
	}
        create, err := os.createFile("writefile")
	if err != nil {
		panic(err)
	}
	defer create.Close()
	for i := range storeRead {
		create.WriteString("The mapped alignment reads along with the CIGAR values are:")
		create.WriteString("%s, %s, %s", storeRead.idread, storeRead.tag, storeRead.seq)
		create.WriteString("The file for the sql injection and transaction has been written")
	}

}
}
