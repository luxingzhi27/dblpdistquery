package tagreader

import (
	"bufio"
	"io"
	"strings"
)

type TagReader struct {
	scanner *bufio.Scanner
	tags    []string
}

func NewTagReader(r io.Reader, tags []string) *TagReader {
	tr := &TagReader{
		scanner: bufio.NewScanner(r),
		tags:    tags,
	}

	tr.scanner.Split(tr.splitFunc)

	return tr
}

func (tr *TagReader) splitFunc(data []byte, atEOF bool) (advance int, token []byte, err error) {
	dataStr := string(data)
	if dataStr == "\n" {
		return 0, nil, nil
	}
	minStartIdx := len(data)
	minEndIdx := len(data)
	var minTag string
	for _, tag := range tr.tags {
		startTag := "<" + tag
		endTag := "</" + tag + ">"
		startIdx := strings.Index(dataStr, startTag)
		endIdx := strings.Index(dataStr, endTag)
		if startIdx != -1 && startIdx < minStartIdx && endIdx != -1 && endIdx < minEndIdx && startIdx < endIdx {
			minStartIdx = startIdx
			minEndIdx = endIdx
			minTag = tag
		}
	}

	if minTag != "" {
		// 返回找到的第一个标签及其内容
		return minEndIdx + len("</"+minTag+">"), data[minStartIdx : minEndIdx+len("</"+minTag+">")], nil
	}

	if atEOF {
		if len(data) > 0 {
			return len(data), data, nil
		}
		return 0, nil, nil
	}

	return 0, nil, nil
}

func (tr *TagReader) Scan() bool {
	return tr.scanner.Scan()
}

func (tr *TagReader) Text() string {
	return tr.scanner.Text()
}

func (tr *TagReader) Err() error {
	return tr.scanner.Err()
}
