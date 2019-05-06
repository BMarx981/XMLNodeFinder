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
	Program Program  `xml:"program"`
	Payload Payload  `xml:"payload"`
}

//Program : has the top level data
type Program struct {
	XMLName xml.Name `xml:"program"`
	Values  []Value  `xml:",any"`
}

//Payload : Contains all of the data
type Payload struct {
	XMLName xml.Name `xml:"payload"`
	Values  []Value  `xml:"any"`
}

//Value : contains the nodes of the payload
type Value struct {
	Node    string
	Content string
}

func main() {
	fmt.Print("Enter the file name: ")
	reader := bufio.NewReader(os.Stdin)

	fileName, _, err := reader.ReadLine()
	if err != nil {
		fmt.Println("Error reading file")
	}

	xmlFile, err := ioutil.ReadFile(string(fileName))
	if err != nil {
		fmt.Println("Error opening the xml file", err)
	}
	processXML(&xmlFile)
	fmt.Println("End of program")
} // C:\Clients\YRC\BASELINE2\B.CCS985\BASELINE_JELLYSTONE.xml

//ProcessXML : the XML nodes.
func processXML(xmlFile *[]byte) (i int) {
	var root Root
	xml.Unmarshal(*xmlFile, &root)
	inputs := root.DataInput
	for i, inp := range inputs {
		j := 0
		fmt.Println(i)
		for n := 0; n < len(inp.Payload.Values); n++ {
			fmt.Println(inp.Payload.Values[n].Content)
		}

		j = j + 1
	}
	return 5
}
