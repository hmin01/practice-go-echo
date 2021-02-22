package utils

import (
	"encoding/json"
	"io/ioutil"
	"os"
	"path"
	"time"
)

func getAbsolutePath(filePath string) (string, error) {
	// Get workspace
	workspace, err := os.Getwd()
	if err != nil {
		PrintLog("error", err.Error())
		return "", err
	}
	// Return
	return path.Join(workspace, filePath), nil
}

func VerifyExistence(filePath string) bool {
	// Get absolute path
	absoluteFilePath, err := getAbsolutePath(filePath)
	if err != nil {
		PrintLog("error", err.Error())
		return false
	}
	
	// Check file existence
	if _, err := os.Stat(absoluteFilePath); os.IsNotExist(err) {
		PrintLog("error", err.Error())
		return false
	} else {
		return true
	}
}

func ReadFile(filePath string) ([]byte, bool) {
	isExist := VerifyExistence(filePath)
	if !isExist {
		return nil, false
	}
	// Open file
	absoluteFilePath, _ := getAbsolutePath(filePath)
	file, err := os.Open(absoluteFilePath)
	if err != nil {
		PrintLog("error", err.Error())
		return nil, false
	}
	// Get raw data
	rawData, err := ioutil.ReadAll(file)
	if err != nil {
		PrintLog("error", err.Error())
		return nil, false
	} else {
		return rawData, true
	}
}

/* Convert function */
func ConvertToJSON(rawData []byte) map[string]interface{} {
	var data map[string]interface{}
	json.Unmarshal(rawData, &data)
	return data
}

func ConvertToTime(strTime string) (time.Time, error) {
	time, err := time.Parse("2006-01-02 15:04:05", strTime)
	if err != nil {
		return time, err
	} else {
		return time, nil
	}
}