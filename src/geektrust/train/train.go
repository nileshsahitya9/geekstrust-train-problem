package train

import (
	"fmt"
	"geektrust/validation"
	"sort"
	"strings"
)

var stationAfterHyderabad = map[string]int{
	"NGP": 400,
	"ITJ": 700,
	"BPL": 800,
	"AGA": 1300,
	"NDL": 1700,
	"PTA": 1800,
	"NJP": 2200,
	"GHY": 2700,
}

func Train(train [][]string) {

	if len(train) == 0 {
		fmt.Println("JOURNEY_ENDED")
		return
	}

	trainA, trainAList := coachFiltering(train[0], stationAfterHyderabad)
	trainB, trainBList := coachFiltering(train[1], stationAfterHyderabad)

	if validation.IsListEmpty(trainAList) && validation.IsListEmpty(trainBList) {
		fmt.Println("JOURNEY_ENDED")
		return
	}

	trainAB := mergeTrainAB(trainA, trainB, stationAfterHyderabad)

	formattedOutput(trainAList, "ARRIVAL  TRAIN_A ENGINE")
	formattedOutput(trainBList, "ARRIVAL  TRAIN_B ENGINE")
	formattedOutput(trainAB, "DEPARTURE  TRAIN_AB ENGINE ENGINE")

}

func coachFiltering(train []string, stationAfterHyderabad map[string]int) (map[string]int, []string) {
	if !validation.IsListEmpty(train) {
		train = train[2:]
	}

	trainAtHYBMap := make(map[string]int)
	trainAtHYBList := []string{}
	for _, station := range train {
		if _, ok := stationAfterHyderabad[station]; ok {
			trainAtHYBMap[station] += 1
			trainAtHYBList = append(trainAtHYBList, station)
		}
	}
	return trainAtHYBMap, trainAtHYBList
}

func mergeTrainAB(trainA, trainB, stationAfterHyderabad map[string]int) []string {
	finalOutput := []string{}
	stations := sortStation(stationAfterHyderabad)
	for index, _ := range stations {
		totalStations := 0
		stationName := stations[len(stations)-1-index]
		if _, ok := trainA[stationName]; ok {
			totalStations += trainA[stationName]
		}
		if _, ok := trainB[stationName]; ok {
			totalStations += trainB[stationName]
		}
		initialVal := 0
		for initialVal < totalStations {
			finalOutput = append(finalOutput, stationName)
			initialVal += 1
		}
	}
	return finalOutput
}

func sortStation(stationAfterHyderabad map[string]int) []string {
	stations := make([]string, 0, len(stationAfterHyderabad))
	for station := range stationAfterHyderabad {
		stations = append(stations, station)
	}

	sort.Slice(stations, func(i, j int) bool {
		return stationAfterHyderabad[stations[i]] < stationAfterHyderabad[stations[j]]
	})
	return stations
}

func formattedOutput(train []string, scheduleType string) {
	fmt.Println(scheduleType, strings.Join(train[:], " "))
}
