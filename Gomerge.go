package main

// app name is genix
import (
	"errors"
	"fmt"
	"os"
	"path/filepath"
	"strings"
)

var log = fmt.Println
var stamp string = "[genix]"

const DUMB_FILE = "dump_data"

const MERGE_FILE = "test_data/yt_video.mp4"

// const MERGE_FILE = "test_data/cmd.jpg"

func main() {
	filename_f := GetFileName(MERGE_FILE)
	str := []string{DUMB_FILE, "/", filename_f, stamp}
	FolderName := strings.Join(str, "")
	ext, err := GetFileExtension(MERGE_FILE)
	if err != nil {
		log("problem getting file extension : ", err)
	}
	Merge(FolderName, DUMB_FILE, ext)

	// Split(MERGE_FILE, DUMB_FILE, 5)
}
func GetFileExtension(fileName string) (string, error) {
	cheap := '.'
	if strings.ContainsRune(fileName, cheap) {
		extIndex := strings.LastIndex(fileName, string(cheap))
		return fileName[extIndex+1:], nil
	} else {
		return "", errors.New("you should only use this on file names with extensions")
	}
}
func Split(file string, Dump string, amount int) {
	FileName := GetFileName(file)
	// log(FileName)
	file_house := Create_Template(FileName, Dump)
	log(file_house)
	var kb float64 = 1024
	data, err := os.ReadFile(file)
	if err != nil {
		log("** problem ", err)
	} else {
		// data Display just for data can be removed afterwards
		var chunk_size_byte int = len(data) / amount
		log(FileName, ":", float64(len(data))/kb, "kb")
		log("split amount:", amount, ", each:", float64(chunk_size_byte), "byte")

		var chunks [][]byte
		for i := 0; i < len(data); i += chunk_size_byte {
			end := i + chunk_size_byte
			if end > len(data) {
				end = len(data)
			}
			chunks = append(chunks, data[i:end])
		}
		for i, chunk := range chunks {
			// ext, err := GetFileExtension(FileName)
			// if err != nil {
			// log("problmen in getting the file extension", err)
			// }
			name_elements := []string{fmt.Sprintf("%d", i), "[genix]"}
			filename_rander := strings.Join(name_elements, "")
			filename := filepath.Join(file_house, filename_rander)

			f_err := os.WriteFile(filename, chunk, os.ModePerm)
			if f_err != nil {
				fmt.Println("Error writing chunk:", err)
				return
			}
			fmt.Println("Written:", filename)
			// paths = append(paths, filename)
		}
		// log(paths)
	}
}
func GetFileName(filePath string) string {
	cheap := '/'
	if strings.ContainsRune(filePath, cheap) {
		splitName := strings.Split(filePath, "/")
		fileName := splitName[len(splitName)-1]
		return fileName
	}
	return filePath // if there's no '/', return the entire path
}
func Create_Template(file_path string, home string) string {
	create_temp := func(name string) string {
		el := []string{name, "[genix]"}
		folder_name := filepath.Join(home, strings.Join(el, ""))
		err := os.Mkdir(folder_name, os.ModePerm)
		if err != nil {
			log("There was a problem in Creating a folder for buffer", err)
		}
		return folder_name
	}
	// " " = for string ,, ' ' = for character
	cheap := '/'
	if strings.ContainsRune(file_path, cheap) {
		split_name := strings.Split(file_path, "/")
		file_name := split_name[len(split_name)-1]
		create_temp(file_name)
		el := []string{home, "/", file_name, stamp}
		path := strings.Join(el, "")
		return path
	} else {
		file_path := create_temp(file_path)
		return file_path
	}
}

func Merge(folder string, dump string, ext string) {
	var data_buffer []byte
	files, err := os.ReadDir(folder)
	if err != nil {
		log("problem reading data folder :", err)
	}
	for i := range len(files) {
		el := []string{folder, "/", files[i].Name()}
		file_name := strings.Join(el, "")
		f_data, err := os.ReadFile(file_name)
		if err != nil {
			log("error reading file : ", err)
		} else {
			data_buffer = append(data_buffer, f_data...)
		}
	}
	// log(data_buffer)
	M_f := []string{folder, "[Merged]", ".", ext}
	MergedFileName := strings.Join(M_f, "")
	err = os.WriteFile(MergedFileName, data_buffer, os.ModePerm)
	if err != nil {
		log("problem writing merged file : ", err)
	} else {
		log("worked i guess ? ")
	}
}

// [
// 	Ranked_filed:1file,2fiel,3file,4file
// 	Orge_file_name:"something",
// 	Orge_file_size:102129
// ]
