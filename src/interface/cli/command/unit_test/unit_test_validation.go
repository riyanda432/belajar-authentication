package command_unit_test

import (
	"bufio"
	"errors"
	"math"
	"os"
	"sort"
	"strconv"
	"strings"

	"github.com/sirupsen/logrus"
)

var minCoverInPercent = 100.0

var float64EqualityThreshold = 0.1

var maxShowLowestCoverage = 10

var osOpen = os.Open

func Validate() error {
	var minCover string
	if len(os.Args) > 2 {
		minCover = os.Args[2]
		if x, err := strconv.ParseFloat(minCover, 64); err == nil {
			minCoverInPercent = x
		}

	}

	var lowestCov []suggestion
	lowestCov = make([]suggestion, 0)
	// jsonFile, err := os.Open("../../../test-report")
	jsonFile, err := osOpen("test-report")
	// if we os.Open returns an error then handle it
	if err != nil {
		return err
	}
	scanner := bufio.NewScanner(jsonFile)
	for scanner.Scan() {
		x := strings.ReplaceAll(scanner.Text(), "\x00", "")
		x = strings.ReplaceAll(x, "\t", " ")
		x = strings.ReplaceAll(x, "\r", "")

		if strings.Contains(x, "total") {
			idx := strings.Index(x, "%")
			i, err := getCoverage(x, idx, 5)
			if err != nil {
				return err
			}
			total := float64(*i)
			logrus.Println("====================")
			logrus.Println("The Lowest Coverage:")
			if len(lowestCov) == 0 {
				logrus.Println("all code already cover")
			} else {
				for _, cov := range lowestCov {
					logrus.Println(cov.functionName, *cov.cover)
				}
			}

			logrus.Println("====================")

			x1 := strconv.FormatFloat(float64(math.Floor(total*10)/10), 'f', 1, 64)
			x := strconv.FormatFloat(float64(minCoverInPercent), 'f', 1, 64)
			if math.Abs(total-float64(minCoverInPercent)) > float64EqualityThreshold &&
				total < float64(minCoverInPercent) {
				errMsg := "failed minimum cover: "
				errMsg += x
				errMsg += " cover: "
				errMsg += x1

				return errors.New(errMsg)
			}
			logrus.Println("Success Unit Test with Minimum cover:", x, ". cover:", x1)

		} else {
			s := strings.Fields(x)
			if len(s) == 3 {
				name := strings.ReplaceAll(
					s[0],
					"github.com/riyanda432/belajar-authentication/",
					"",
				)
				name = strings.ReplaceAll(
					name,
					"\xff",
					"",
				)
				name = strings.ReplaceAll(
					name,
					"\xfe",
					"",
				)
				covStr := strings.ReplaceAll(s[2], "%", "")
				cov, err := strconv.ParseFloat(covStr, 64)
				if err != nil {
					return err
				}
				if cov >= 100 {
					continue
				}
				fileName := strings.Split(name, ":")
				exist := false
				for i, low := range lowestCov {
					if strings.Contains(low.functionName, fileName[0]) {
						exist = true
						x1 := float64(cov)
						x2 := *lowestCov[i].cover
						if x2 > x1 {
							lowestCov[i].cover = &x1
							lowestCov[i].functionName = name + " " + s[1]
						}
						break
					}
				}
				if exist {
					continue
				}
				if len(lowestCov) < maxShowLowestCoverage {
					x1 := float64(cov)
					lowestCov = append(lowestCov, suggestion{
						functionName: name + " " + s[1],
						cover:        &x1,
					})
				} else {
					x1 := *lowestCov[len(lowestCov)-1].cover
					x2 := float64(cov)
					if float64(x1) > cov {
						lowestCov[len(lowestCov)-1].cover = &x2
						lowestCov[len(lowestCov)-1].functionName = name + " " + s[1]
					}
				}

				sort.Slice(lowestCov, func(p, q int) bool {
					x1 := *lowestCov[p].cover
					x2 := *lowestCov[q].cover
					if x1 == x2 {
						switch strings.Compare(lowestCov[p].functionName, lowestCov[q].functionName) {
						case -1:
							return true
						case 1:
							return false
						}
					}
					return x1 < x2
				})

			}

		}
	}
	defer jsonFile.Close()
	// defer os.Remove("test-report")
	return nil
}

func getCoverage(result string, idx int, count int) (*float64, error) {
	asRunes := []rune(result)
	s := string(asRunes[idx-count : idx])
	i, err := strconv.ParseFloat(s, 32)
	if err != nil {
		if count <= 0 {
			return nil, err
		}
		return getCoverage(result, idx, count-1)
	}
	x := float64(i)
	return &x, nil
}

type suggestion struct {
	functionName string
	cover        *float64
}
