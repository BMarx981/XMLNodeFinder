package main

import (
	"bufio"
	"encoding/xml"
	"fmt"
	"io/ioutil"
	"os"
)

//Root : The main doc
type Root struct {
	XMLName   xml.Name    `xml:"root"`
	DataInput []DataInput `xml:"dataInput"`
}

//DataInput : contain the meat of the data
type DataInput struct {
	XMLName xml.Name `xml:"dataInput"`
	Hdr     Hdr      `xml:"hdr"`
	Payload Payload  `xml:"payload"`
}

//Hdr : has the top level data
type Hdr struct {
	XMLName xml.Name `xml:"hdr"`
	// Values    []Value   `xml:",any"`
	ProgramID string `xml:"programId"`
	Action    string `xml:"action"`
	EsbFormat string `xml:"esbFormat"`
	DateTime  string `xml:"dateTime"`
	CalledBy  string `xml:"calledBy"`
	User      string `xml:"user"`
}

//Payload : Contains all of the data
type Payload struct {
	XMLName xml.Name `xml:"payload"`
	Record  Record   `xml:"record"`
}

//Control : Has data related to the control of the program
type Control struct {
	XMLName xml.Name `xml:"control"`
	Values  []Value  `xml:",any"`
}

//Record : Contains all of the very important data
type Record struct {
	XMLName xml.Name `xml:"record"`
	Values  []string `xml:",any"`
	RecKey  string   `xml:"RECORD.KEY"`
}

//Value : contains the nodes of the payload
type Value struct {
	Node    string
	Content string
}

func main() {
	processXML(findFile())
} // /Users/brianmarx/Desktop/baselineFake.txt

func findFile() []byte {
	fmt.Print("Enter the file name: ")
	reader := bufio.NewReader(os.Stdin)

	fileName, _, error := reader.ReadLine()
	if error != nil {
		fmt.Println("Error reading file")
	}

	xmlFile, err := ioutil.ReadFile(string(fileName))
	if err != nil {
		fmt.Println("Error opening the xml file", err)
	}
	return xmlFile
}

//ProcessXML : the XML nodes.
func processXML(xmlFile []byte) {
	var root Root

	error := xml.Unmarshal(xmlFile, &root)
	if error != nil {
		fmt.Println(error)
	}
	_, er := processDataInput(root.DataInput)
	if er != nil {
		fmt.Println(er)
	}
	// fmt.Println(m)
}

func processDataInput(input []DataInput) (map[string]DataInput, error) {
	m := make(map[string]DataInput)
	n := make(map[string]Hdr)
	for index, element := range input {
		fmt.Println("New DataInput")
		m[element.Payload.Record.RecKey] = input[index]
		n[element.Payload.Record.RecKey] = input[index].Hdr
	}
	fmt.Println(n)
	fmt.Printf(" There are %d elements in the map\n", len(m))
	return m, nil
}
