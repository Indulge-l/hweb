package util

import (
	"encoding/csv"
	"errors"
	"fmt"
	"io"
	"os"
)

type Csv interface {
	GetAllData() ([][]string, error)
	GetAllDataByFilePath(filePath string) ([][]string, error)
	GetRow(row int) ([]string, error)
	GetCol(col int) ([]string, error)
	RowLength() int
	GetPointData(row, col int) (string, error)
}
type CsvUtils struct {
	reader *csv.Reader
	writer *csv.Writer
	data   [][]string
}

func NewCsvUtils(r io.Reader) CsvUtils {
	var csvUtil CsvUtils
	csvUtil.reader = csv.NewReader(r)
	csvUtil.readAllData()
	return csvUtil
}

func (m *CsvUtils) GetAllData() ([][]string, error) {
	return m.reader.ReadAll()
}

func (m *CsvUtils) GetAllDataByFilePath(filePath string) ([][]string, error) {
	csvF, err := os.Open(filePath)
	if err != nil {
		return nil, err
	}
	m.reader = csv.NewReader(csvF)
	m.readAllData()
	return m.data, nil
}

func (m *CsvUtils) GetRow(row int) ([]string, error) {
	if err := m.checkRowRange(row); err != nil {
		return nil, err
	}
	return m.data[row], nil
}

func (m *CsvUtils) GetCol(col int) ([]string, error) {
	if err := m.checkColRange(col); err != nil {
		return nil, err
	}
	lens := m.RowLength()
	result := make([]string, lens)
	for i := 0; i < lens; i++ {
		result[i] = m.data[i][col]
	}
	return result, nil
}

func (m *CsvUtils) RowLength() int {
	return len(m.data)
}

func (m *CsvUtils) GetPointData(row, col int) (string, error) {
	if err := m.checkPointRange(row, col); err != nil {
		return "", err
	}
	return m.data[row][col], nil
}

// =================================================================
// internal methods
// =================================================================
func (m *CsvUtils) readAllData() {
	var err error
	m.data, err = m.reader.ReadAll()
	if err != nil {
		panic(errors.New(fmt.Sprintf("read all data failed,error(%+v)", err.Error())))
	}
}

func (m *CsvUtils) checkRowRange(row int) error {
	if row < 0 || row > m.RowLength() {
		return errors.New(fmt.Sprintf("row out of range"))
	}
	return nil
}

func (m *CsvUtils) checkColRange(col int) error {
	if col < 0 || col > m.RowLength() {
		return errors.New(fmt.Sprintf("col out of range"))
	}
	return nil
}

func (m *CsvUtils) checkPointRange(row, col int) error {
	if err := m.checkRowRange(row); err != nil {
		return err
	}
	if err := m.checkColRange(col); err != nil {
		return err
	}
	return nil
}
