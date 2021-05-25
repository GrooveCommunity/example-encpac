package pack

import (
	"archive/zip"
	"bytes"
	"io"
	"io/ioutil"
	"log"
	"os"
	"strconv"
)

func checkErr(err error) {
	if err != nil {
		log.Fatal(err)
	}
}

func CreateStructure(targetName string) string {
	dir, err := ioutil.TempDir("", ".ks-"+targetName)
	checkErr(err)

	checkErr(os.Chdir(dir))

	return dir
}

func zipFiles(fileName string, files []string) {
	newFile, err := os.Create(fileName)
	checkErr(err)

	defer newFile.Close()

	zipWriter := zip.NewWriter(newFile)
	defer zipWriter.Close()

	for _, file := range files {
		zipFile, err := os.Open(file)
		checkErr(err)

		defer zipFile.Close()

		info, err := zipFile.Stat()
		checkErr(err)

		header, err := zip.FileInfoHeader(info)
		checkErr(err)

		header.Method = zip.Deflate
		writer, err := zipWriter.CreateHeader(header)
		checkErr(err)

		_, err = io.Copy(writer, zipFile)
		checkErr(err)
	}
}

func Compress(keysPath string, currentPath string, amountKeys int) {
	files := []string{}

	for i := 0; i < amountKeys; i++ {
		files = append(files, "k"+strconv.Itoa(i+1))
		files = append(files, "v"+strconv.Itoa(i+1))
	}

	output := currentPath + "/..."

	zipFiles(output, files)

	checkErr(os.Chdir(currentPath))

	os.RemoveAll(keysPath)
}

func DoPackage(content bytes.Buffer, name string) {
	file, err := os.Create(name)
	checkErr(err)

	file.Write(content.Bytes())

	defer file.Close()
}
