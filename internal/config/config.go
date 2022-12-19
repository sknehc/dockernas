package config

import (
	"os"
	"path/filepath"
	"tinycloud/internal/utils"
)

func GetBasePath() string {
	basePath, err := os.Getwd()
	if err != nil {
		panic(err)
	}
	basePath = filepath.Join(basePath, "data")
	return basePath
}

func GetFullDfsPath(path string) string {
	basePath := GetBasePath()
	basePath = filepath.Join(basePath, "dfs", path)
	utils.CheckCreateDir(basePath)

	return basePath
}

func GetDBFilePath() string {
	basePath := GetBasePath()
	basePath = filepath.Join(basePath, "meta")
	utils.CheckCreateDir(basePath)
	return filepath.Join(basePath, "data.db3")
}

func GetAppLocalPath(instanceName string) string {
	basePath := GetBasePath()
	basePath = filepath.Join(basePath, "local", instanceName)
	return basePath
}

func GetLocalVolumePath(instanceName string, volumeName string) string {
	basePath := GetBasePath()
	basePath = filepath.Join(basePath, "local", instanceName, volumeName)
	utils.CheckCreateDir(basePath)
	return basePath
}
