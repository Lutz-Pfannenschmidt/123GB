package parser

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/Lutz-Pfannenschmidt/stunden-berechner/internal/csv"
	"github.com/Lutz-Pfannenschmidt/stunden-berechner/internal/date"
)

func ParseFile(path string, pivotDate date.Date) (map[string][2]float64, error) {
	var firstSem2Day = date.FromInt(pivotDate.GetInt() + 1)

	lines, err := csv.ReadAnyFileToCSV(path)
	if err != nil {
		return nil, err
	}

	var result = map[string][2]float64{}

	var currName string
	var currYear [2][]uint
	var isSem2 bool
	var foundTable bool

	for _, line := range *lines {
		line = strings.TrimSpace(line)
		if isNameLine(line) {
			currName = strings.ReplaceAll(line, ",", "")
			currYear = [2][]uint{}
			isSem2 = false
			foundTable = false
		} else if isHeaderLine(line) {
			foundTable = true
		} else if foundTable && currName != "" {
			values := strings.Split(line, ",")

			if values[0] == "" {
				avg1 := avg(currYear[0])
				avg2 := avg(currYear[1])
				result[currName] = [2]float64{avg1, avg2}
				foundTable = false
				continue
			}

			dates := strings.Split(values[1], "-")
			start, err1 := date.ParseDate(dates[0])
			end, err2 := date.ParseDate(dates[1])
			if err1 != nil || err2 != nil {
				return nil, err1
			}

			if start.Compare(firstSem2Day) == 0 {
				isSem2 = true
			}

			if strings.Contains(line, "Ferien") {
				if pivotDate.Compare(start) >= 0 && pivotDate.Compare(end) <= 0 {
					isSem2 = true
				}
				continue
			}

			c, err := strconv.Atoi(values[4])
			classCount := uint(c)

			if err != nil {
				return nil, err
			}

			if start.DaysUntil(end) > 6 {
				var from, to int
				fmt.Sscanf(values[0], "%d-%d", &from, &to)

				if pivotDate.Compare(start) >= 0 && pivotDate.Compare(end) < 0 {
					fmt.Printf("Values for %s are probably wrong because the pivot is in a range of weeks\n", currName)
				}

				for i := from; i <= to; i++ {
					if isSem2 {
						currYear[1] = append(currYear[1], classCount)
					} else {
						currYear[0] = append(currYear[0], classCount)
					}
				}

			} else {
				if isSem2 {
					currYear[1] = append(currYear[1], classCount)
				} else {
					currYear[0] = append(currYear[0], classCount)
				}
			}

			if pivotDate.Compare(start) >= 0 && pivotDate.Compare(end) <= 0 {
				isSem2 = true
			}

		}

	}

	return result, nil
}

func isHeaderLine(line string) bool {
	return strings.Contains(line, "Woche") && strings.Contains(line, "Periode") && strings.Contains(line, "Soll")
}

// // expandLines expands lines if they contain a range of weeks.
// func expandLines(lines []string) []string {
// 	var expanded []string
// 	for _, line := range lines {
// 		values := strings.Split(line, ",")
// 		if strings.Contains(values[0], "-") {
// 			var from, to int
// 			fmt.Sscanf(values[0], "%d-%d", &from, &to)
// 			var fromDate, toDate string
// 			fmt.Sscanf(values[1], "%s-%s", &fromDate, &toDate)
// 			fromD := date.MustParseDate(fromDate)
// 			toD := date.MustParseDate(toDate)
// 			for i := from; i <= to; i++ {
// 				expanded = append(expanded, fmt.Sprintf("%d,%d.%d.-%d.%d.,%s,", i, strings.Join(values[2:], ",")))
// 			}
// 		} else {
// 			expanded = append(expanded, line)
// 		}
// 	}
// 	return expanded
// }

// isNameLine checks if a line is a name line by checking if it contains only letters and commas and has at least 3 letters.
func isNameLine(line string) bool {
	line = strings.ReplaceAll(line, ",", "")
	if len(line) < 3 {
		return false
	}

	for i := 0; i < len(line); i++ {
		if !isLetter(line[i]) {
			return false
		}
	}

	return true
}

func isLetter(char byte) bool {
	return char >= 'a' && char <= 'z' || char >= 'A' && char <= 'Z' || char == 'ä' || char == 'ö' || char == 'ü' || char == 'ß'
}

// avg calculates the average of an array of uints.
// If the array is empty, -1 is returned.
func avg(arr []uint) float64 {
	var sum uint = 0
	for _, val := range arr {
		sum += val
	}
	if len(arr) == 0 {
		return -1
	}
	return float64(sum) / float64(len(arr))
}
