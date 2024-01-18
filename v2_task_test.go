package amclient

import (
	"encoding/json"
	"fmt"
	"net/http"
	"testing"
	"time"

	"gotest.tools/v3/assert"
)

func TestTask_Read(t *testing.T) {
	setup()
	defer teardown()

	mux.HandleFunc("/api/v2beta/task/96acb0a1-525c-456a-9060-51bb84f5f708/", func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, "GET")
		fmt.Fprint(w, `{
	"uuid": "96acb0a1-525c-456a-9060-51bb84f5f708",
	"exit_code": 1,
	"file_uuid": "e502c0d9-becf-455a-8b20-091526947a09",
	"file_name": "foobar.txt",
	"time_created": "2019-06-18T00:00:00",
	"time_started": "2019-07-18T00:00:00",
	"time_ended": "2019-08-18T00:00:00",
	"duration": 4294967295
}`)
	})

	payload, _, err := client.Task.Read(ctx, "96acb0a1-525c-456a-9060-51bb84f5f708")

	assert.NilError(t, err)
	assert.DeepEqual(t, &Task{
		ID:          "96acb0a1-525c-456a-9060-51bb84f5f708",
		ExitCode:    1,
		FileID:      "e502c0d9-becf-455a-8b20-091526947a09",
		CreatedAt:   TaskDateTime{Time: time.Date(2019, time.June, 18, 0, 0, 0, 0, time.UTC)},
		StartedAt:   TaskDateTime{Time: time.Date(2019, time.July, 18, 0, 0, 0, 0, time.UTC)},
		CompletedAt: TaskDateTime{Time: time.Date(2019, time.August, 18, 0, 0, 0, 0, time.UTC)},
		Filename:    "foobar.txt",
		Duration:    TaskDuration(4294967295 * time.Second),
	}, payload)
}

func TestTaskDateTimeUnmarshalJSON(t *testing.T) {
	tests := map[string]struct {
		input     string
		expected  TaskDateTime
		expectErr bool
	}{
		"Expected format": {
			input:    `"2006-01-02T15:04:05"`,
			expected: TaskDateTime{Time: time.Date(2006, time.January, 2, 15, 4, 5, 0, time.UTC)},
		},
		"Unexpected format": {
			input:     `"2006-01-02T15:04:05Z"`,
			expectErr: true,
		},
		"Empty string": {
			input:     `"2006-01-02T15:04:05Z"`,
			expectErr: true,
		},
	}

	for name, tt := range tests {
		t.Run(name, func(t *testing.T) {
			var dt TaskDateTime

			err := json.Unmarshal([]byte(tt.input), &dt)

			if tt.expectErr {
				assert.Assert(t, err != nil)
			} else {
				assert.NilError(t, err)
				assert.Equal(t, dt, tt.expected)
			}
		})
	}
}

func TestTaskDurationUnmarshalJSON(t *testing.T) {
	tests := map[string]struct {
		input     string
		expected  TaskDuration
		expectErr bool
	}{
		"Integer literal": {
			input:    `30`,
			expected: TaskDuration(30 * time.Second),
		},
		"Negative integer": {
			input:    `-10`,
			expected: TaskDuration(-10 * time.Second),
		},
		"Minus one": {
			input:    `"< 1"`,
			expected: TaskDuration(time.Second / 2),
		},
		"Empty string": {
			input:    `""`,
			expected: TaskDuration(0),
		},
		"Invalid string": {
			input:     `"invalid"`,
			expectErr: true,
		},
		"Invalid string with < symbol": {
			input:     `"inv < alid"`,
			expectErr: true,
		},
		"Invalid integer (too big)": {
			input:     `123456789012345678901234567890123456789`,
			expectErr: true,
		},
	}

	for name, tt := range tests {
		t.Run(name, func(t *testing.T) {
			var dur TaskDuration

			err := json.Unmarshal([]byte(tt.input), &dur)

			if tt.expectErr {
				assert.Assert(t, err != nil)
			} else {
				assert.NilError(t, err)
				assert.Equal(t, dur, tt.expected)
			}
		})
	}
}
