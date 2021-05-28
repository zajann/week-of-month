package main

/*
특정 날짜 x가 해당 월의 몇주차인지 구하는 함수를 TDD로 구현
- 1주일의 기준은 월요일부터 일요일
- 1주차는 해당 월 첫번째 목요일을 포함하는 주차
예시)
- 2021년 5월 1일 = 2021년 4월 5주차
- 2021년 5월 28일 = 2021년 5월 4주차
*/

import (
	"reflect"
	"testing"
	"time"
)

type weekOfMonth struct {
	year  int
	month int
	week  int
}

func makeWeekOfMonth(t time.Time) weekOfMonth {
	m := weekOfMonth{
		year:  t.Year(),
		month: int(t.Month()),
	}

	fisrtWeekday := getFisrtWeekday(t)
	firstDayOfWeeks := getFirstDayOfWeeks(fisrtWeekday)

	inDay := t.Day()
	if inDay < firstDayOfWeeks[0] {
		m.month--
		m.week = 5
		return m
	}

	var week int
	for _, day := range firstDayOfWeeks {
		if inDay >= day {
			week++
		} else {
			break
		}
	}
	m.week = week
	return m
}

func TestWeekOfMonth(t *testing.T) {
	tests := []struct {
		input time.Time
		want  weekOfMonth
	}{
		{
			time.Date(2021, 5, 1, 0, 0, 0, 0, time.UTC),
			weekOfMonth{2021, 4, 5},
		},
		{
			time.Date(2021, 5, 28, 0, 0, 0, 0, time.UTC),
			weekOfMonth{2021, 5, 4},
		},
		{
			time.Date(2021, 5, 7, 0, 0, 0, 0, time.UTC),
			weekOfMonth{2021, 5, 1},
		},
		{
			time.Date(2021, 7, 1, 0, 0, 0, 0, time.UTC),
			weekOfMonth{2021, 7, 1},
		},
		{
			time.Date(2021, 7, 5, 0, 0, 0, 0, time.UTC),
			weekOfMonth{2021, 7, 2},
		},
	}

	for _, tt := range tests {
		if got := makeWeekOfMonth(tt.input); !reflect.DeepEqual(got, tt.want) {
			t.Errorf("FAILD, want %v, but got %v", got, tt.want)
		}
	}
}

// getFirstDayOfWeeks returns firsrt day of week. fwd is day of first(1) of month.
func getFirstDayOfWeeks(fwd int) []int {
	var days []int
	if fwd <= 3 {
		days = append(days, 1)
	}
	for i := 1 + (6 - fwd) + 1; i < 32; i += 7 {
		days = append(days, i)
	}
	return days
}

func TestFirstDayOfWeeks(t *testing.T) {
	tests := []struct {
		input int
		want  []int
	}{
		{
			5,
			[]int{3, 10, 17, 24, 31},
		},
		{
			1,
			[]int{1, 7, 14, 21, 28},
		},
	}

	for _, tt := range tests {
		if got := getFirstDayOfWeeks(tt.input); !reflect.DeepEqual(got, tt.want) {
			t.Errorf("FAILD, want %v, but got %v", got, tt.want)
		}
	}
}

// getFisrtWeekday returns day of first(1) of month.
// Mon: 0, Tue: 1, Wed: 2, Thu: 3, Fri: 4, Sat: 5, Sun: 6
func getFisrtWeekday(t time.Time) int {
	d := time.Date(t.Year(), t.Month(), 1, 0, 0, 0, 0, time.UTC).Weekday() - 1
	// Sunday
	if d == -1 {
		d = 6
	}
	return int(d)
}

func TestFirstWeekay(t *testing.T) {
	tests := []struct {
		input time.Time
		want  int
	}{
		{
			time.Date(2021, 6, 28, 0, 0, 0, 0, time.UTC),
			1,
		},
		{
			time.Date(2021, 5, 28, 0, 0, 0, 0, time.UTC),
			5,
		},
	}

	for _, tt := range tests {
		if got := getFisrtWeekday(tt.input); got != tt.want {
			t.Errorf("FAILD, want %v, but got %v", got, tt.want)
		}
	}
}
