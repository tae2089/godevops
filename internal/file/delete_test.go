package file

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_deleteOldFiles(t *testing.T) {
	t.Run("remove files - success", func(t *testing.T) {
		t.SkipNow()
		dirPath := "/tmv/var/file"
		daysOld := 0
		err := deleteOldFiles(dirPath, daysOld)
		assert.Nil(t, err)
	})
	t.Run("remove files - fail", func(t *testing.T) {
		dirPath := "./testing"
		daysOld := 30
		err := deleteOldFiles(dirPath, daysOld)
		assert.NotNil(t, err)
	})
}
