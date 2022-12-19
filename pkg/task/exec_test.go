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
		_, err := task.exec()
		assert.NoError(t, err)
	})
	t.Run("404 URL", func(t *testing.T) {
		task := Task{
			URL:                 "https://www.google.com/404",
			Headers:             nil,
			Method:              "GET",
			Timeout:             1,
			AcceptedStatusCodes: []uint32{404},
		}
		_, err := task.exec()
		assert.NoError(t, err)
	})
	t.Run("Non-existent site", func(t *testing.T) {
		task := Task{
			URL:                 "https://www.some-non-existent.site",
			Headers:             nil,
			Method:              "GET",
			Timeout:             1,
			AcceptedStatusCodes: []uint32{404},
		}
		_, err := task.exec()
		assert.Error(t, err)
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
		assert.Len(t, r.responses, 30)
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

	t.Run("Invalid config: concurrency > total", func(t *testing.T) {
		task := Task{
			URL:                 "https://www.google.com/404",
			Headers:             nil,
			Method:              "GET",
			Timeout:             1,
			AcceptedStatusCodes: []uint32{200},
		}
		_, err := task.Execute(&Config{
			Concurrency: 100,
			Total:       30,
		})
		assert.Error(t, err)
		assert.Equal(t, ErrConcurrencyGreaterThanTotal, err)
	})

	t.Run("Invalid config: concurrency number is 0", func(t *testing.T) {
		task := Task{
			URL:                 "https://www.google.com/404",
			Headers:             nil,
			Method:              "GET",
			Timeout:             1,
			AcceptedStatusCodes: []uint32{200},
		}
		_, err := task.Execute(&Config{
			Concurrency: 0,
			Total:       30,
		})
		assert.Error(t, err)
		assert.Equal(t, ErrInvalidConcurrencyNumber, err)
	})

	t.Run("Invalid config: total number is 0", func(t *testing.T) {
		task := Task{
			URL:                 "https://www.google.com/404",
			Headers:             nil,
			Method:              "GET",
			Timeout:             1,
			AcceptedStatusCodes: []uint32{200},
		}
		_, err := task.Execute(&Config{
			Concurrency: 10,
			Total:       0,
		})
		assert.Error(t, err)
		assert.Equal(t, ErrInvalidTotalNumber, err)
	})
}
