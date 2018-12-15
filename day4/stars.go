package day4

import (
	"fmt"
	"sort"
	"strings"
	"time"
)

type actionType int

const (
	beginsShift = actionType(iota)
	endsShift
	fallsAsleep
	wakesUp
)

type event struct {
	guardId int
	action  actionType
}

type record struct {
	year, month, day int
	hour, minute     int
	e                event
}

type ByTimeSorter []record

func (a ByTimeSorter) Len() int      { return len(a) }
func (a ByTimeSorter) Swap(i, j int) { a[i], a[j] = a[j], a[i] }
func (a ByTimeSorter) Less(i, j int) bool {
	ti := time.Date(a[i].year, time.Month(a[i].month), a[i].day, a[i].hour, a[i].minute, 0, 0, time.UTC)
	tj := time.Date(a[j].year, time.Month(a[j].month), a[j].day, a[j].hour, a[j].minute, 0, 0, time.UTC)
	return tj.After(ti)
}

func Stars() {
	inputs := strings.Split(input, "\n")
	records := make([]record, len(inputs))
	for i, v := range inputs {
		year, month, day, hour, minute := 0, 0, 0, 0, 0
		fmt.Sscanf(v, "[%d-%d-%d %d:%d]", &year, &month, &day, &hour, &minute)
		idx := strings.LastIndex(v, "] ") + 2
		whatHappens := v[idx:]
		r := record{year, month, day, hour, minute, event{0, beginsShift}}
		if strings.HasPrefix(whatHappens, "wakes up") {
			r.e.action = wakesUp
			r.e.guardId = -1
		} else if strings.HasPrefix(whatHappens, "Guard") {
			r.e.action = beginsShift
			fmt.Sscanf(v[strings.Index(v, "#")+1:], "%d", &(r.e.guardId))
		} else if strings.HasPrefix(whatHappens, "falls asleep") {
			r.e.action = fallsAsleep
			r.e.guardId = -1
		}
		records[i] = r
	}

	sortedRecords := ByTimeSorter(records)
	sort.Sort(sortedRecords)

	guardRecords := make(map[int][]int)
	currentGuardId := -1
	var currentGuardRecords []int = nil
	fellAsleepAt := -1
	for _, v := range sortedRecords {
		switch v.e.action {
		case beginsShift:
			currentGuardId = v.e.guardId
			ok := true
			currentGuardRecords, ok = guardRecords[currentGuardId]
			if !ok {
				currentGuardRecords = make([]int, 61)
				guardRecords[currentGuardId] = currentGuardRecords
			}

		case fallsAsleep:
			fellAsleepAt = v.minute

		case wakesUp:
			wokeUpAt := v.minute
			currentGuardRecords[60] += wokeUpAt - fellAsleepAt
			for i := fellAsleepAt; i < wokeUpAt; i++ {
				currentGuardRecords[i] += 1
			}
		}
	}

	mostAsleepGuardId := -1
	mostAsleepTime := -1
	for k, v := range guardRecords {
		if v[60] > mostAsleepTime {
			mostAsleepTime = v[60]
			mostAsleepGuardId = k
		}
	}

	mostSleptMinute := 0
	for k, v := range guardRecords[mostAsleepGuardId][:len(guardRecords[mostAsleepGuardId])-1] {
		if v > guardRecords[mostAsleepGuardId][mostSleptMinute] {
			mostSleptMinute = k
		}
	}

	fmt.Println("day4 star1:", mostAsleepGuardId*mostSleptMinute)

	mostFrequentSleeperId := -1
	mostFrequentSleeptMinute := -1
	mostFrequentSleeptTimes := -1
	for guardId, records := range guardRecords {
		for minute, timesAsleepAtMinute := range records[:len(records)-1] {
			if timesAsleepAtMinute > mostFrequentSleeptTimes {
				mostFrequentSleeptMinute = minute
				mostFrequentSleeptTimes = timesAsleepAtMinute
				mostFrequentSleeperId = guardId
			}
		}
	}

	fmt.Println("day4 star2:", mostFrequentSleeperId*mostFrequentSleeptMinute)
}
