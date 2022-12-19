package task

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestTask(t *testing.T) {
	t.Run("Parse Toml config file", func(t *testing.T) {
		task, err := ParseFile("../../testdata/task.toml")
		assert.NoError(t, err)
		assert.Equal(t, "https://www.google.com", task.URL)
		assert.Equal(t, "GET", task.Method)
		assert.Equal(t, []uint32{200}, task.AcceptedStatusCodes)
	})
}
