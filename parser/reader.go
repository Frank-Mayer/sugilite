package parser

import (
	"bufio"
	"log"
	"os"

	"sugilite.frank-mayer.io/utils"
)

type reader struct {
	File        *os.File
	Line        uint64
	LineContent string
	Col         uint64
	Current     rune
	eof         bool
	Scanner     *bufio.Scanner
}

func Reader(f string) reader {
	r := reader{
		File:        nil,
		Line:        1,
		LineContent: utils.EMPTY_STRING,
		Col:         1,
		Current:     '\000',
		Scanner:     nil,
	}
	var err error
	r.File, err = os.OpenFile(f, os.O_RDONLY, os.ModePerm)
	if err != nil {
		log.Fatal(err)
	}
	defer r.File.Close()

	r.Scanner = bufio.NewScanner(r.File)
	r.Scanner.Scan()
	r.LineContent = r.Scanner.Text()
	return r
}

func (self *reader) Read() bool {
	if self.eof {
		return false
	}
	if self.Col > uint64(len(self.LineContent)) {
		if self.nextLine() {
			self.Current = '\n'
		} else {
			return false
		}
	} else {
		self.Current = rune(self.LineContent[self.Col-1])
		self.Col++
	}
	return true
}

func (self *reader) nextLine() bool {
	self.Line++
	self.Col = 1
	if self.Scanner.Scan() {
		self.LineContent = self.Scanner.Text()
		return true
	} else {
		self.eof = true
		return false
	}
}

func (self *reader) Peak(i uint64) rune {
	if self.Col+i >= uint64(len(self.LineContent)) {
		return '\000'
	} else {
		return rune(self.LineContent[self.Col+i])
	}
}
