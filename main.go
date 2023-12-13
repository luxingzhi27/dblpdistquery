package main

import (
	tagreader "dblpdistquery/parserecord"
	"fmt"
	"os"
	"strconv"
	"strings"
)

var tags = []string{"article", "inproceedings", "proceedings", "book", "incollection", "phdthesis", "mastersthesis", "www", "data"}

const xmlFile = "../dblp.xml"
const recordSize = 50000

func main() {
	os.MkdirAll("dblp_blocks", 0777)
	dblp, err := os.Open(xmlFile)
	if err != nil {
		panic(err)
	}
	defer dblp.Close()

	cnt := 0

	tr := tagreader.NewTagReader(dblp, tags)

	var buf strings.Builder
	for tr.Scan() {
		cnt++
		buf.WriteString(tr.Text() + "\n")
		if cnt%recordSize == 0 {
			block, err := os.OpenFile("dblp_blocks/block_"+strconv.Itoa(cnt/recordSize)+".xml", os.O_CREATE|os.O_WRONLY, 0777)
			if err != nil {
				panic(err)
			}
			block.WriteString(buf.String())
			buf.Reset()
			block.Close()
			fmt.Println("Block", cnt/recordSize, "done")
		}
	}
	fmt.Println("Total records:", cnt)
}
