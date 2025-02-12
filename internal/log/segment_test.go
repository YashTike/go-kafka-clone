package log

import (
	"io"
	"os"
	"testing"

	api "github.com/YashTike/proglog/api/v1"
	"github.com/stretchr/testify/require"
)

func TestSegment(t *testing.T) {

	dir, err := os.MkdirTemp("", "segment_test")
	require.NoError(t, err)
	defer os.RemoveAll(dir)

	want := &api.Record{
		Value: []byte("hello, world"),
	}

	c := Config{}
	c.Segment.MaxStoreBytes = 1024
	c.Segment.MaxIndexBytes = entWidth * 3

	s, err := newSegment(dir, 16, c)
	require.NoError(t, err)
	require.Equal(t, uint64(16), s.nextOffset, s.baseOffset)
	require.False(t, s.IsMaxed())

	for i := uint64(0); i < 3; i++ {
		offset, err := s.Append(want)
		require.NoError(t, err)
		require.Equal(t, 16+i, offset)

		got, err := s.Read(offset)
		require.NoError(t, err)
		require.Equal(t, want.Value, got.Value)
	}

	_, err = s.Append(want)
	require.Equal(t, io.EOF, err)
	require.True(t, s.IsMaxed())
	
	c.Segment.MaxStoreBytes = uint64(len(want.Value) * 3)
	c.Segment.MaxIndexBytes = 1024

	s, err = newSegment(dir, 16, c)
	require.NoError(t, err)
	require.True(t, s.IsMaxed())

	err = s.Remove()
	require.NoError(t, err)
	s, err = newSegment(dir, 16, c)
	require.NoError(t, err)
	require.False(t, s.IsMaxed())

}