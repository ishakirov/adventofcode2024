package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

type filesystem struct {
	blocks []*block
	files  []*file
}

type block struct {
	fileId int
}

type file struct {
	offset int
	length int
}

func newFilesystem(diskMap string) *filesystem {
	fs := filesystem{
		blocks: []*block{},
	}

	for i, val := range diskMap {
		length := int(val - '0')
		if i%2 == 0 {
			fileId := len(fs.files)
			fs.files = append(fs.files, &file{offset: len(fs.blocks), length: length})
			for j := 0; j < length; j++ {
				fs.blocks = append(fs.blocks, &block{fileId: fileId})
			}
		} else {
			for j := 0; j < length; j++ {
				fs.blocks = append(fs.blocks, nil)
			}
		}
	}

	return &fs
}

func (fs *filesystem) defragmentBlocks() {
	freePtr := 0
	for fs.blocks[freePtr] != nil {
		freePtr++
	}

	mvPtr := len(fs.blocks) - 1
	for mvPtr >= 0 && mvPtr > freePtr {
		if fs.blocks[mvPtr] != nil {
			fs.blocks[freePtr] = fs.blocks[mvPtr]
			fs.blocks[mvPtr] = nil
			for freePtr < len(fs.blocks) && fs.blocks[freePtr] != nil {
				freePtr++
			}
		}

		mvPtr--
	}
}

func (fs *filesystem) checksum() int {
	checksum := 0

	for i := 0; i < len(fs.blocks); i++ {
		if fs.blocks[i] != nil {
			checksum += fs.blocks[i].fileId * i
		}
	}

	return checksum
}

func (fs *filesystem) mvFile(fileId int, offset int) {
	for i := 0; i < fs.files[fileId].length; i++ {
		fs.blocks[i+offset] = fs.blocks[i+fs.files[fileId].offset]
		fs.blocks[i+fs.files[fileId].offset] = nil
	}
	fs.files[fileId].offset = offset
}

func (fs *filesystem) defragmentFiles() {
	for fileId := len(fs.files) - 1; fileId >= 0; fileId-- {
		ptr := 0
		for ptr < fs.files[fileId].offset {
			if fs.blocks[ptr] != nil {
				ptr += fs.files[fs.blocks[ptr].fileId].length
				continue
			}

			freePtr := ptr
			for ptr < len(fs.blocks) && fs.blocks[ptr] == nil {
				ptr++
			}

			if ptr-freePtr >= fs.files[fileId].length {
				log.Printf("moving file %d to offset %d", fileId, freePtr)
				fs.mvFile(fileId, freePtr)
				// fs.printBlocks()
				break
			}
		}
	}
}

func (fs *filesystem) printBlocks() {
	s := ""
	for i := 0; i < len(fs.blocks); i++ {
		if fs.blocks[i] != nil {
			s += fmt.Sprintf("%d", fs.blocks[i].fileId)
		} else {
			s += "."
		}
	}

	log.Println(s)
}

func part1(diskMap string) int {
	fs := newFilesystem(diskMap)
	fs.defragmentBlocks()
	return fs.checksum()
}

func part2(diskMap string) int {
	fs := newFilesystem(diskMap)
	// fs.printBlocks()
	fs.defragmentFiles()
	return fs.checksum()
}

func main() {
	input, err := os.Open("./input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer input.Close()

	diskMap := ""

	scanner := bufio.NewScanner(input)
	for scanner.Scan() {
		diskMap = scanner.Text()
	}

	log.Println(part1(diskMap))
	log.Println(part2(diskMap))
}
