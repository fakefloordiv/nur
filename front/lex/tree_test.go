package lex

import (
	"github.com/stretchr/testify/require"
	"testing"
)

func TestNode_InsertAndMatch(t *testing.T) {
	node := new(treeNode)
	require.NoError(t, node.insert("+", OpPlus))
	require.NoError(t, node.insert("-", OpMinus))
	require.NoError(t, node.insert("*", OpStar))
	require.NoError(t, node.insert("**", OpStarStar))

	t.Run("Match single characters", func(t *testing.T) {
		require.Equal(t, OpPlus, node.Match("+"))
		require.Equal(t, OpMinus, node.Match("-"))
		require.Equal(t, OpStar, node.Match("*"))
	})

	t.Run("Match double characters", func(t *testing.T) {
		require.Equal(t, OpStarStar, node.Match("**"))
	})

	t.Run("Match unknown character", func(t *testing.T) {
		require.Equal(t, Unknown, node.Match(")"))
		require.Equal(t, Unknown, node.Match("++"))
		require.Equal(t, Unknown, node.Match("***"))
	})
}

func TestNode_HasPrefix(t *testing.T) {
	node := new(treeNode)
	require.NoError(t, node.insert("+", OpPlus))
	require.NoError(t, node.insert("*", OpStar))
	require.NoError(t, node.insert("**", OpStarStar))

	t.Run("Single characters", func(t *testing.T) {
		require.True(t, node.HasPrefix("+"))
		require.True(t, node.HasPrefix("*"))
	})

	t.Run("Double characters", func(t *testing.T) {
		require.Equal(t, OpStarStar, node.Match("**"))
	})

	t.Run("Unknown character sequences", func(t *testing.T) {
		require.Equal(t, Unknown, node.Match(")"))
		require.Equal(t, Unknown, node.Match("++"))
		require.Equal(t, Unknown, node.Match("***"))
	})
}
