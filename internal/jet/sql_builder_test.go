package jet

import (
	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
	"testing"
	"time"
)

func TestArgToString(t *testing.T) {
	assert.Equal(t, argToString(true), "TRUE")
	assert.Equal(t, argToString(false), "FALSE")

	assert.Equal(t, argToString(int(-32)), "-32")
	assert.Equal(t, argToString(uint(32)), "32")
	assert.Equal(t, argToString(int8(-43)), "-43")
	assert.Equal(t, argToString(uint8(43)), "43")
	assert.Equal(t, argToString(int16(-54)), "-54")
	assert.Equal(t, argToString(uint16(54)), "54")
	assert.Equal(t, argToString(int32(-65)), "-65")
	assert.Equal(t, argToString(uint32(65)), "65")
	assert.Equal(t, argToString(int64(-64)), "-64")
	assert.Equal(t, argToString(uint64(64)), "64")
	assert.Equal(t, argToString(float32(2.0)), "2")
	assert.Equal(t, argToString(float64(1.11)), "1.11")

	assert.Equal(t, argToString("john"), "'john'")
	assert.Equal(t, argToString("It's text"), "'It''s text'")
	assert.Equal(t, argToString([]byte("john")), "'john'")
	assert.Equal(t, argToString(uuid.MustParse("b68dbff4-a87d-11e9-a7f2-98ded00c39c6")), "'b68dbff4-a87d-11e9-a7f2-98ded00c39c6'")

	time, err := time.Parse("Mon Jan 2 15:04:05 -0700 MST 2006", "Mon Jan 2 15:04:05 -0700 MST 2006")
	assert.NoError(t, err)
	assert.Equal(t, argToString(time), "'2006-01-02 15:04:05-07:00'")

	func() {
		defer func() {
			assert.Equal(t, recover().(string), "jet: map[string]bool type can not be used as SQL query parameter")
		}()

		argToString(map[string]bool{})
	}()
}

func TestFallTrough(t *testing.T) {
	assert.Equal(t, FallTrough([]SerializeOption{ShortName}), []SerializeOption{ShortName})
	assert.Equal(t, FallTrough([]SerializeOption{SkipNewLine}), []SerializeOption(nil))
	assert.Equal(t, FallTrough([]SerializeOption{ShortName, SkipNewLine}), []SerializeOption{ShortName})
}
