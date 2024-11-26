package main

import (
	"github.com/stretchr/testify/assert"
	"math"
	"testing"
	"time"
)

func TestSolution(t *testing.T) {
	t.Run("success", func(t *testing.T) {
		type testCase struct {
			test         string
			partsCount   int
			chunkSize    int
			expectedTime float64
		}
		testCases := []testCase{
			{
				test:         "100 - 10",
				partsCount:   100,
				chunkSize:    10,
				expectedTime: 2.5,
			},
			{
				test:         "100 - 8",
				partsCount:   100,
				chunkSize:    8,
				expectedTime: 3.3,
			},
			{
				test:         "100 - 16",
				partsCount:   100,
				chunkSize:    16,
				expectedTime: 1.8,
			},
		}

		for _, tc := range testCases {
			t.Run(tc.test, func(t *testing.T) {
				now := time.Now()
				errs := DownloadFile(tc.partsCount, tc.chunkSize)

				after := time.Since(now) // time in seconds
				formattedAfter := math.Round(float64(after)/float64(time.Second)*10) / 10

				assert.Equal(t, tc.expectedTime, formattedAfter)
				assert.Len(t, errs, tc.partsCount/20)
			})
		}
	})
}
