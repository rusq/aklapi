package aklapi

import (
	"fmt"
	"reflect"
	"testing"
	"time"
)

func mustTime(t time.Time, err error) time.Time {
	if err != nil {
		panic(err)
	}
	return t
}

func Test_adjustYear(t *testing.T) {
	const dttmFmt = "2006-01-02 15:04:05"
	type args struct {
		t time.Time
	}
	tests := []struct {
		name string
		now  func() time.Time
		args args
		want time.Time
	}{
		{"same day",
			func() time.Time {
				t, _ := time.Parse(dttmFmt, fmt.Sprintf("%d-09-16 10:00:00", time.Now().Year()))
				return t
			},
			args{mustTime(time.Parse(dttmFmt, "0000-09-16 00:00:00"))},
			mustTime(time.Parse(dttmFmt, fmt.Sprintf("%d-09-16 00:00:00", time.Now().Year()))),
		},
		{"same day",
			func() time.Time {
				t, _ := time.Parse(dttmFmt, fmt.Sprintf("%d-04-07 15:04:05", time.Now().Year()))
				return t
			},
			args{mustTime(time.Parse(dttmFmt, "0000-04-07 00:00:00"))},
			mustTime(time.Parse(dttmFmt, fmt.Sprintf("%d-04-07 00:00:00", time.Now().Year()))),
		},
		{"differnt days",
			func() time.Time {
				t, _ := time.Parse(dttmFmt, fmt.Sprintf("%d-09-15 20:00:00", time.Now().Year()))
				return t
			},
			args{mustTime(time.Parse(dttmFmt, "0000-09-16 00:00:00"))},
			mustTime(time.Parse(dttmFmt, fmt.Sprintf("%d-09-16 00:00:00", time.Now().Year()))),
		},
		{"new year",
			func() time.Time {
				t, _ := time.Parse(dttmFmt, fmt.Sprintf("%d-12-31 20:00:00", time.Now().Year()))
				return t
			},
			args{mustTime(time.Parse(dttmFmt, "0000-01-20 00:00:00"))},
			mustTime(time.Parse(dttmFmt, fmt.Sprintf("%d-01-20 00:00:00", time.Now().Year()+1))),
		},
	}
	for _, tt := range tests {
		oldNow := now
		defer func() {
			now = oldNow
		}()
		now = tt.now
		t.Run(tt.name, func(t *testing.T) {
			if got := adjustYear(tt.args.t); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("adjustYear() = %v, want %v", got, tt.want)
			}
		})
	}
}
