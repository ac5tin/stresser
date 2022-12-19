package task

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestExec(t *testing.T) {
	t.Run("Basic Get", func(t *testing.T) {
		task := Task{
			URL:                 "https://www.google.com",
			Headers:             nil,
			Method:              "GET",
			Timeout:             1,
			AcceptedStatusCodes: []uint32{200},
		}
		assert.NoError(t, task.exec())
	})
	t.Run("404 URL", func(t *testing.T) {
		task := Task{
			URL:                 "https://www.google.com/404",
			Headers:             nil,
			Method:              "GET",
			Timeout:             1,
			AcceptedStatusCodes: []uint32{404},
		}
		assert.NoError(t, task.exec())
	})
	t.Run("Non-existent site", func(t *testing.T) {
		task := Task{
			URL:                 "https://www.some-non-existent.site",
			Headers:             nil,
			Method:              "GET",
			Timeout:             1,
			AcceptedStatusCodes: []uint32{404},
		}
		assert.Error(t, task.exec())
	})

	t.Run("Successful Results", func(t *testing.T) {
		task := Task{
			URL:                 "https://www.google.com",
			Headers:             nil,
			Method:              "GET",
			Timeout:             1,
			AcceptedStatusCodes: []uint32{200},
		}
		r, err := task.Execute(&Config{
			Concurrency: 10,
			Total:       30,
		})
		assert.NoError(t, err)
		assert.Equal(t, uint32(30), r.SuccessCount)
	})

	t.Run("Failed Results", func(t *testing.T) {
		task := Task{
			URL:                 "https://www.google.com/404",
			Headers:             nil,
			Method:              "GET",
			Timeout:             1,
			AcceptedStatusCodes: []uint32{200},
		}
		r, err := task.Execute(&Config{
			Concurrency: 10,
			Total:       30,
		})
		assert.NoError(t, err)
		assert.Equal(t, uint32(30), r.FailedCount)
	})
}
