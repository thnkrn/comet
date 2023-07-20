package util

import (
	"strconv"
	"strings"
)

const SPLITTER_CHAR = "_"

func AddTrailingSlash(str string) string {
	isSlashed := strings.HasSuffix(str, "/")
	if isSlashed {
		return str
	}
	return str + "/"
}

func GetBaseWithSplitter(str string) (string, error) {
	base, err := GetBase(str)
	if err != nil {
		return "", err
	}
	return strconv.Itoa(base) + SPLITTER_CHAR, nil
}

func GetBase(str string) (int, error) {
	base := strings.Split(GetLastPath(str), SPLITTER_CHAR)[0]
	i, err := strconv.Atoi(base)
	if err != nil {
		return -1, err
	}
	return i, nil
}

func GetIncremental(str string) (int, error) {
	increment := strings.Split(GetLastPath(str), SPLITTER_CHAR)[1]
	i, err := strconv.Atoi(increment)
	if err != nil {
		return -1, err
	}
	return i, nil
}

func GetLastPath(str string) string {
	array := strings.Split(str, "/")
	if len(array) > 0 {
		return array[len(array)-1]
	}
	return str
}
