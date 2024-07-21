package main

import (
	"errors"
	"testing"
)

func TestParseArgs(t *testing.T) {
	//테스트 케이스를 추상하기 위한 testConfig 구조체 정의
	type testConfig struct {
		args []string
		err  error
		config
	}

	//테스트 케이스 작성
	tests := []testConfig{
		{
			args:   []string{"-h"},
			err:    nil,
			config: config{printUsage: true, numTimes: 0},
		},
		{
			args:   []string{"10"},
			err:    nil,
			config: config{printUsage: false, numTimes: 10},
		},
		{
			args:   []string{"abc"},
			err:    errors.New("strconv.Atoi: parsing \"abc\": invalid syntax"),
			config: config{printUsage: false, numTimes: 0},
		},
		{
			args:   []string{"1", "foo"},
			err:    errors.New("Invalid number of arguments"),
			config: config{printUsage: false, numTimes: 0},
		},
	}

	for _, tc := range tests {
		c, err := parseArgs(tc.args)
		if tc.err != nil && err.Error() != tc.err.Error() {
			t.Fatalf("Expected error to be: %v, got: %v\n", tc.err, err)
		}
		if tc.err == nil && err != nil {
			t.Errorf("Expected nil error, got: %v\n", err)
		}
		if c.printUsage != tc.config.printUsage {
			t.Errorf("Expected printUsage to be: %v, got: %v\n",
				tc.config.printUsage, c.printUsage)
		}
		if c.numTimes != tc.config.numTimes {
			t.Errorf("Expected numTimes to be: %v, got: %v\n",
				tc.config.numTimes, c.numTimes)
		}
	}
}
