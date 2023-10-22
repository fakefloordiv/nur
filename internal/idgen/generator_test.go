package idgen

import (
	"github.com/stretchr/testify/require"
	"testing"
)

func TestGenerator(t *testing.T) {
	t.Run("determined sequence", func(t *testing.T) {
		gen := NewGenerator(5)
		var ids []string

		for i := 0; i < 5; i++ {
			ids = append(ids, gen.ID())
		}

		require.Equal(t, []string{"BCsWm", "uwqoE", "DKQpV", "viZIe", "zexQb"}, ids)
	})
}
