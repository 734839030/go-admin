package database

import (
	"testing"
)

func Test_getConnection(t *testing.T) {
	t.Run("connect", func(t *testing.T) {
		if got := getDb(); nil == got {
			t.Errorf("getConnection() = %v", got)
		}
	})
}
