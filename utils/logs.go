package utils

import (
	"fmt"
	"go-starterkit-project/config"
	"io"
	"io/ioutil"
	"os"
	"time"
)

func CreateSqlLog() io.Writer {
	filePath := config.Config("LOG_LOCATION") + config.Config("LOG_SQL_ERROR_FILENAME")
	getLastDateLog := getLastLineWithSeek(filePath, 1)[0:10]
	currentDate := time.Now().Format("2006-01-02")
	fmt.Println(changeDateLogToDate(getLastDateLog) + " " + currentDate)

	rescueStdout := os.Stdout
	r, w, _ := os.Pipe()
	os.Stdout = w

	w.Close()
	out, _ := ioutil.ReadAll(r)
	os.Stdout = rescueStdout

	fmt.Printf("Captured: %s", out)

	logFile := CreateFile(config.Config("LOG_LOCATION"), config.Config("LOG_SQL_ERROR_FILENAME"))
	multiOutput := MultiWrite(os.Stdout, logFile)

	return multiOutput
}

func getLastLineWithSeek(filepath string, lineFromBottom int) string {
	fileHandle, err := os.Open(filepath)

	if err != nil {
		panic("Cannot open file")
	}
	defer fileHandle.Close()

	line := ""
	var cursor int64 = 0
	stat, _ := fileHandle.Stat()
	filesize := stat.Size()
	for {
		cursor -= 1
		fileHandle.Seek(cursor, io.SeekEnd-lineFromBottom)

		char := make([]byte, 1)
		fileHandle.Read(char)

		if cursor != -1 && (char[0] == 10 || char[0] == 13) { // stop if we find a line
			break
		}

		line = fmt.Sprintf("%s%s", string(char), line) // there is more efficient way

		if cursor == -filesize { // stop if we are at the begining
			break
		}
	}

	if lineFromBottom > 0 {
		return reverseString(line)
	}

	return line
}

func reverseString(str string) string {
	byte_str := []rune(str)
	for i, j := 0, len(byte_str)-1; i < j; i, j = i+1, j-1 {
		byte_str[i], byte_str[j] = byte_str[j], byte_str[i]
	}
	return string(byte_str)
}

func changeDateLogToDate(date string) string {
	layout := "2006/01/02"
	t, _ := time.Parse(layout, date)

	return t.Format("2006-01-02")
}
