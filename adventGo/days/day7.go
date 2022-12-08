package days

import (
	"fmt"
	"log"
	"sort"
	"strconv"
	"strings"
)

type File struct {
	Name   string
	Size   int64
	Parent *Directory
}

type Directory struct {
	Name     string
	File     []*File
	Children []*Directory
	Parent   *Directory
}

type DirectoryIndex struct {
	Name     string
	Size   int64
}

type DirectoryIndexes []DirectoryIndex

func (s DirectoryIndexes) Len() int {
	return len(s)
}

func (s DirectoryIndexes) Swap(i, j int) {
	s[i], s[j] = s[j], s[i]
}

func (s DirectoryIndexes) Less(i, j int) bool {
	return s[i].Size < s[j].Size
}

func getChildIndex(child string, pointer *Directory) int {
	// check if child has name
	for i, d := range pointer.Children {
		if d.Name == child {
			return i
		}
	}
	return -1
}

func makeNewDir(name string, parent *Directory) {
	newDir := new(Directory)
	newDir.Name = name
	newDir.Parent = parent
	parent.Children = append(parent.Children, newDir)
	log.Printf("Added Directory: %v\n", newDir.Name)
}

func makeNewDirAndNavToIt(name string, parent *Directory) *Directory {
	newDir := new(Directory)
	newDir.Name = name
	newDir.Parent = parent
	parent.Children = append(parent.Children, newDir)
	log.Printf("Added Directory and navigated to: %v\n", newDir.Name)
	return newDir
}

func addFile(input []string, director *Directory) {
	numb, err := strconv.ParseInt(input[0], 10, 0)
	if err != nil {
		log.Panic(err)
	}
	director.File = append(director.File, &File{Name: input[1], Size: numb, Parent: director})
	log.Printf("Added File: %v, Size: %v\n", director.File[len(director.File)-1].Name, director.File[len(director.File)-1].Size)
}

func nextOp(line string, pointer *Directory) *Directory {
	rvPointer := pointer
	input := strings.Split(line, " ")
	switch input[0] {
	case "$":
		switch input[1] {
		case "cd":
			if input[2] == ".." {
				return pointer.Parent
			}

			// go home
			if input[2] == "/" {
				for rvPointer.Name != "/" {
					rvPointer = rvPointer.Parent
				}
				return rvPointer
			}
			// nav to dir listed or make new
			if childIndex := getChildIndex(input[2], pointer); childIndex != -1 {
				return pointer.Children[childIndex]
			} else {
				// make the new directory, add it as a child
				return makeNewDirAndNavToIt(input[2], pointer)
			}
		}
	case "ls":
		return rvPointer
	case "dir":
		makeNewDir(input[1], pointer)
	default:
		addFile(input, pointer)
	}

	return rvPointer
}

func printDirectory(pointer *Directory, indent int) {
	log.Printf("%s- %v (dir)", strings.Repeat("  ", indent), pointer.Name)
	for _, d := range pointer.File {
		log.Printf("%s- %v (file, size=%v)", strings.Repeat("  ", indent+1), d.Name, d.Size)
	}
	indent++
	for _, d := range pointer.Children {
		printDirectory(d, indent)
	}
}

func printDirectoryFromBase(pointer *Directory) {
	newPointer := pointer

	for newPointer.Name != "/" {
		newPointer = newPointer.Parent
	}

	printDirectory(newPointer, 0)
}

func getDirectorySize(pointer *Directory) int64 {
	// always add the files up for this level of directory
	var rv int64
	rv = 0
	for _, d := range pointer.File {
		rv += d.Size
	}

	// terminal node, return all files added together
	if pointer.Children == nil {
		log.Printf("dir %s size %v", pointer.Name, rv)
		return rv
	}

	for _, d := range pointer.Children {
		addVal := getDirectorySize(d)
		rv += addVal
	}

	log.Printf("dir %s size %v", pointer.Name, rv)
	return rv
}

func getDirectorySizeStartingAtBase(pointer *Directory) int64{
	newPointer := pointer

	for newPointer.Name != "/" {
		newPointer = newPointer.Parent
	}

	return getDirectorySize(newPointer)
}

func getDirectorySizeStartingAtBaseWithCeil(pointer *Directory, size int64) {
	newPointer := pointer

	for newPointer.Name != "/" {
		newPointer = newPointer.Parent
	}

	_, allMatches := getDirectorySizeUnderVal(newPointer, size)
	fmt.Println("")
	addUp := int64(0)
	for _, d:= range allMatches {
		addUp +=d
	}
	log.Printf("size at root of all directories under %vk: %v",size/10000, addUp)
}

func getDirectorySizeUnderVal(pointer *Directory, size int64) (int64, []int64) {
	// always add the files up for this level of directory
	var totalSize int64
	totalSize = 0
	for _, d := range pointer.File {
		totalSize += d.Size
	}

	// terminal node, return all files added together
	if pointer.Children == nil {
		rvArr := make([]int64, 0)
		if totalSize <= size {
			return totalSize, append(rvArr, totalSize)
		}
		return totalSize, rvArr
	}

	dirSize := make([]int64, 0)
	for _, d := range pointer.Children {
		childSize, childDir := getDirectorySizeUnderVal(d, size)
		dirSize = append(dirSize, childDir...)
		totalSize += childSize
	}

	if totalSize > size {
		return totalSize, dirSize
	}

	log.Printf("dir %s size %v", pointer.Name, totalSize)
	return totalSize, append(dirSize, totalSize)
}

func getDirectorySizeStartingAtBaseWithFloor(pointer *Directory, size int64) {
	newPointer := pointer

	for newPointer.Name != "/" {
		newPointer = newPointer.Parent
	}

	_, allMatches := getDirectoriesAboveVal(newPointer, size)
	sort.Sort(allMatches)
	for _, d:= range allMatches {
		log.Printf("%v directory is above %vk: %v",d.Name, size/10000, d.Size)
	}
}

func getDirectoriesAboveVal(pointer *Directory, size int64) (int64, DirectoryIndexes) {
	// always add the files up for this level of directory
	var totalSize int64
	totalSize = 0
	for _, d := range pointer.File {
		totalSize += d.Size
	}

	// terminal node, return all files added together
	if pointer.Children == nil {
		rvArr := make([]DirectoryIndex, 0)
		if totalSize >= size {
			return totalSize, append(rvArr, DirectoryIndex{Name: pointer.Name, Size: totalSize})
		}
		return totalSize, rvArr
	}

	dirSize := make([]DirectoryIndex, 0)
	for _, d := range pointer.Children {
		childSize, childDir := getDirectoriesAboveVal(d, size)
		dirSize = append(dirSize, childDir...)
		totalSize += childSize
	}

	if totalSize < size {
		return totalSize, dirSize
	}

	return totalSize, append(dirSize, DirectoryIndex{Name: pointer.Name, Size: totalSize})
}

func Day7(sb string) {
	tsb := `$ cd /
$ ls
dir a
14848514 b.txt
8504156 c.dat
dir d
$ cd a
$ ls
dir e
29116 f
2557 g
62596 h.lst
$ cd e
$ ls
584 i
$ cd ..
$ cd ..
$ cd d
$ ls
4060174 j
8033020 d.log
5626152 d.ext
7214296 k`
	tsb = tsb
	output := strings.Split(sb, "\n")
	dir := new(Directory)
	dir.Name = "/"

	// go through each line
	pointer := dir
	for i := 0; i < len(output); i++ {
		if output[i] != "" {
			pointer = nextOp(output[i], pointer)
		}
	}
	// print file structure
	printDirectoryFromBase(pointer)

	// get sizes
	rootSize := getDirectorySizeStartingAtBase(pointer)

	getDirectorySizeStartingAtBaseWithCeil(pointer, 100000)

	fmt.Println("")
	neededSize := rootSize - (70000000 - 30000000)
	log.Printf("Total Space: %v; Space needed: %v",rootSize, neededSize  )
	getDirectorySizeStartingAtBaseWithFloor(pointer, neededSize)
}
