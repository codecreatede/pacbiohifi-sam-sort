package main 

/*
Author Gaurav Sablok
Universitat Potsdam
Date 2024-8-27

structs and inheritances for making the pacbiohifi genome alignments easier
for the web browse and calling through the http.

*/

import (
   "bufio"
   "fmt"
   "log"
   "os"
   "strings"
)


func main () {

	// Task to complete today:
	// 1. Implement inteheritance
	// 2. route the CIGAR and the sequence alignment to the GMOD
	// 3. check for those and extract those reads using the form submitter



type readStore struct {
	idread string
	tag int
	seq string
}


readFile := os.Args[1:]
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
	          // calling the func with in the else loop to manage the faster rates and
			// not calling the iterator again and it is directly reading the line
			// from the previous Buffer.Scan(). Still checking this if i can implement this as noone has implemented this way.
		   func (r *readStore) add(line string) (seqStor string, err) {
			   r{
				   idread: strings.Split(line, "\t")[0]
				   tag : strings.Split(line, "\t")[5]
				   seq : strings.Split(line, "\t")[9]
			}

		}
		}
	}

}

}
