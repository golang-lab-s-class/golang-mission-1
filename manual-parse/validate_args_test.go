package main

import (
	"errors"
	"testing"
)

//validateArgs 메소드 유닛 테스트 코드

func TestValidateArgs(t *testing.T) {
	//익명 구조체 및 테스트 케이스 작성
	tests := []struct {
		c   config
		err error
	}{
		{
			c:   config{},
			err: errors.New("Must specify a number greater than 0"),
		},
		{
			c:   config{numTimes: -1},
			err: errors.New("Must specify a number greater than 0"),
		},
		{
			c:   config{numTimes: 10},
			err: nil,
		},
	}

	//테스트 검증 과정
	for _, tc := range tests {
		err := validateArgs(tc.c)
		if tc.err != nil && err.Error() != tc.err.Error() {
			t.Errorf("Expected error to be: %v, got: %v\n", tc.err, err)
		}
		if tc.err == nil && err != nil {
			t.Errorf("Expected nil error, got: %v\n", err)
		}
	}
}
