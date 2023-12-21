package main

import (
	tagreader "dblpdistquery/parserecord"
	"fmt"
	"os"
	"strings"
	"sync"
)

var tags = []string{"article", "inproceedings", "proceedings", "book", "incollection", "phdthesis", "mastersthesis", "www", "data"}

const blockNum = 207
const recordSize = 50000

func getBlockName(i int) string {
	return fmt.Sprintf("dblp_blocks/block_%d.xml", i)
}

func queryBlock(blockName string, author string, wg *sync.WaitGroup) {
	defer wg.Done()
	dblp, err := os.Open(blockName)
	if err != nil {
		panic(err)
	}
	defer dblp.Close()

	cnt := 0
	tr := tagreader.NewTagReader(dblp, tags)

	for tr.Scan() {
		if strings.Contains(tr.Text(), author) {
			cnt++
		}
	}
	fmt.Println(blockName, cnt)
}

func main() {
	var author string
	fmt.Scanln(&author)
	fmt.Println("你输入的作者名字是：", author)
	var wg sync.WaitGroup
	for i := 1; i < 20; i++ {
		wg.Add(1)
		go queryBlock(getBlockName(i), author, &wg)
	}
	wg.Wait()
}
