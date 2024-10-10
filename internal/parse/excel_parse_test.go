/*
 * @Author: gaoh
 * @Date: 2024-09-21 15:25:56
 * @LastEditTime: 2024-09-24 12:36:47
 */
package parse

import (
	"log"
	"testing"
)

func TestParseExcel(t *testing.T) {
	/* ProcessIndex(ProcessRowFunc(func(key int) ([]string, error) {
		return nil, nil
	}), 1) */
	var titles map[string][]string = make(map[string][]string)
	var oneRow map[string][]string = make(map[string][]string)
	var allData map[string][][]string = make(map[string][][]string)
	err := ParseExcel("../../涌金府水表号.xlsx", func(data [][]string, sheetName string) {
		allData[sheetName] = data
	}, ProcessRowFunc(func(s []string, sheetName string, i int) error {
		if i == 0 {
			titles[sheetName] = append(titles[sheetName], s...)
		}
		return nil
	}), ProcessRowFunc(func(s []string, sheetName string, i int) error {
		if i == 1 {
			oneRow[sheetName] = append(oneRow[sheetName], s...)
		}
		return nil
	}))
	log.Println(titles)
	log.Println(oneRow)
	log.Println(allData)
	t.Logf("%v", err)
}
