package main

import "fmt"

type pair struct {
	timestamp int
	value     string
}

type TimeMap struct {
	values map[string][]pair
}

func Constructor() TimeMap {
	return TimeMap{
		values: map[string][]pair{},
	}
}

func (this *TimeMap) Set(key string, value string, timestamp int) {
	if times, ok := this.values[key]; ok {
		this.values[key] = append(times, pair{value: value, timestamp: timestamp})
	} else {
		newTimes := []pair{{value: value, timestamp: timestamp}}
		this.values[key] = newTimes
	}
}

func (this *TimeMap) Get(key string, timestamp int) string {
	if _, ok := this.values[key]; !ok {
		return ""
	}

	if timestamp < this.values[key][0].timestamp {
		return ""
	}

	times := this.values[key]
	l, r := 0, len(times)

	for l < r {
		mid := (l + r) / 2
		if timestamp < times[mid].timestamp {
			r = mid
		} else {
			l = mid + 1
		}
	}
	return times[r-1].value
}

func main() {
	timeMap := Constructor()
	timeMap.Set("love", "high", 10)
	timeMap.Set("love", "low", 20)
	fmt.Println(timeMap.Get("love", 5))
	fmt.Println(timeMap.Get("love", 10))
	fmt.Println(timeMap.Get("love", 15))
	fmt.Println(timeMap.Get("love", 20))
	fmt.Println(timeMap.Get("love", 25))
}
