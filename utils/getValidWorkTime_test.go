package utils

import (
	"reflect"
	"testing"
	"time"
)

func TestIsHoliday(t *testing.T) {
	type args struct {
		ti time.Time
	}
	// Time1 := time.Date(2022, 4, 1, 23, 30, 0, 0, time.Local)
	Time1 := time.Now()
	tests := []struct {
		name string
		args args
		want bool
	}{
		// TODO: Add test cases.
		{
			name: "test",
			args: args{
				ti: Time1,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := IsHoliday(tt.args.ti); got != tt.want {
				t.Errorf("IsHoliday() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestGetValidWorkingTimeDuration(t *testing.T) {
	type args struct {
		startTime time.Time
		endTime   time.Time
	}
	startTime1 := time.Date(2022, 4, 1, 23, 30, 0, 0, time.Local)
	endTime1 := time.Date(2022, 4, 5, 1, 0, 0, 0, time.Local)
	// // case2: 调休、周末、工作日： 4-23周六休息日，4-24周日调休工作日，4-25周一工作日。输出工作时长：1500min，4-23有效时长0min，4-24有效时长一天1440min，4-25有效时长60min
	// startTime1 := time.Date(2022, 4, 23, 23, 30, 0, 0, time.Local)
	// endTime1 := time.Date(2022, 4, 25, 1, 0, 0, 0, time.Local)
	tests := []struct {
		name                     string
		args                     args
		wantValidRepairedMinutes int
		wantHolidayList          []string
		wantErr                  bool
	}{
		// TODO: Add test cases.
		{
			name: "test",
			args: args{
				startTime: startTime1,
				endTime:   endTime1,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotValidRepairedMinutes, gotHolidayList, err := GetValidWorkingTimeDuration(tt.args.startTime, tt.args.endTime)
			if (err != nil) != tt.wantErr {
				t.Errorf("GetValidWorkingTimeDuration() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if gotValidRepairedMinutes != tt.wantValidRepairedMinutes {
				t.Errorf("GetValidWorkingTimeDuration() gotValidRepairedMinutes = %v, want %v", gotValidRepairedMinutes, tt.wantValidRepairedMinutes)
			}
			if !reflect.DeepEqual(gotHolidayList, tt.wantHolidayList) {
				t.Errorf("GetValidWorkingTimeDuration() gotHolidayList = %v, want %v", gotHolidayList, tt.wantHolidayList)
			}
		})
	}
}
