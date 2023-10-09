package database

import (
	"bufio"
	"log"
	"os"
	"strconv"
	"strings"
)

const FILENAME = "database/products.txt"

func Initialize() error {
	if _, err := os.Stat(FILENAME); os.IsNotExist(err) {
		file, err := os.Create(FILENAME)
		if err != nil {
			return errCannotCreateFile
		}
		defer file.Close()
	}
	return nil
}

func AppendProduct(info string) error {
	file, err := os.OpenFile(FILENAME, os.O_APPEND|os.O_WRONLY, 0644)
	if err != nil {
		return errCannotOpenFile
	}
	defer file.Close()
	file.WriteString(info + "\n")
	return nil
}

func DeleteProductByID(idFind int) error {
	file, err := os.OpenFile(FILENAME, os.O_RDONLY, 0644)
	if err != nil {
		return errCannotOpenFile
	}
	defer file.Close()

	tempFile, err := os.CreateTemp("", "tempfile.txt")
	if err != nil {
		return errCannotCreateTempFile
	}
	defer tempFile.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		fields := strings.Split(line, ",")
		if len(fields) == 4 {
			idStr := strings.TrimSpace(fields[0])
			id, err := strconv.Atoi(idStr)
			if err != nil {
				log.Println(errCannotConvertStrToInt.Error())
				continue
			}
			if id == idFind {
				continue
			}
		}

		_, err := tempFile.WriteString(line + "\n")
		if err != nil {
			return errCannotWriteFile
		}
	}

	if err := scanner.Err(); err != nil {
		return err
	}

	err = os.Rename(tempFile.Name(), file.Name())
	if err != nil {
		return err
	}

	log.Println("Deleted successfully")
	return nil
}

func OpenFile() error {
	return errCannotOpenFile
}

func GetProducts() (string, error) {
	open, err := os.ReadFile(FILENAME)
	if err != nil {
		return "", errCannotOpenFile
	}

	return string(open), nil
}
