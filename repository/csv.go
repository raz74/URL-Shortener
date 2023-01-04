package repository

import (
	"encoding/csv"
	"log"
	"os"
	"shortened_link/model"
	"time"
)

//type CSVRepository interface {
//	ReadCSVFile(filePath string) (map[string]string, error)
//	WriteCSVFile(MyMap map[string]string, outputPath string) error
//}
//
//type CSVRepositoryImpl struct {
//}

//func CreateCSVFile() {
//	//creating csv file
//	CSVFile, err := os.Create("shorted.CSV")
//	if err != nil {
//		panic(err)
//	}
//	defer func(CSVFile *os.File) {
//		err := CSVFile.Close()
//		if err != nil {
//
//		}
//	}(CSVFile)
//
//}

func ReadCSVFile(filePath string) (map[string]model.ShortedUrl, error) {
	CSVFile, err := os.Open(filePath)
	if err != nil {
		log.Fatal(err)
	}
	//defer CSVFile.Close()
	defer func(CSVFile *os.File) {
		err := CSVFile.Close()
		if err != nil {

		}
	}(CSVFile)

	csvReader := csv.NewReader(CSVFile)
	records, err := csvReader.ReadAll()
	if err != nil {
		log.Fatal("Unable to parse file as CSV for " + filePath)
	}
	//fmt.Print(records)
	m := make(map[string]model.ShortedUrl)
	for _, row := range records {
		expiredAt, _ := time.Parse("2023-01-11 16:17:57.189364484 +0330 +0330 m=+604805.251998686", row[2])
		m[row[0]] = model.ShortedUrl{
			Id:         0,
			LongUrl:    row[1],
			ShortedUrl: row[0],
			ExpiredAt:  expiredAt,
		}
		//log.Fatal(row)
	}

	return m, err
}

func WriteCSVFile(MyMap map[string]model.ShortedUrl, outputPath string) error {
	CSVFile, err := os.Create(outputPath)
	if err != nil {
		log.Fatal(err)
	}
	Writer := csv.NewWriter(CSVFile)
	defer Writer.Flush()
	var data [][]string
	i := 0
	for key, record := range MyMap {
		data = append(data, []string{key, record.LongUrl, record.ExpiredAt.String()})
		i++
	}
	err = Writer.WriteAll(data)

	return err
}
