package controllers

import (
	"reflect"
	"testing"
)

func Test_getGithubJobs(t *testing.T) {
	type args struct {
		key    string
		pageNo int
	}
	tests := []struct {
		name string
		args args
		want []jobs
	}{
		// TODO: Add test cases.
		// {"Checking retrieval from github API", args{"java", 1}, jobs{{}}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := GetGithubJobs(tt.args.key, tt.args.pageNo); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("getGithubJobs() = %v, want %v", got, tt.want)
			}
		})
	}
}
