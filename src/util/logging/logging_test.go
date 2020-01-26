package logging

import (
	"os"
	"testing"

	"github.com/sirupsen/logrus"

	"github.com/stretchr/testify/require"
)

func TestGetOutputWriter(t *testing.T) {
	val, err := GetOutputWriter("stdout")
	require.NoError(t, err)
	require.Equal(t, os.Stdout, val)

	val, err = GetOutputWriter("stderr")
	require.NoError(t, err)
	require.Equal(t, os.Stderr, val)

	val, err = GetOutputWriter("none")
	require.NoError(t, err)
	require.Equal(t, NoWriter{}, val)

	_, err = GetOutputWriter("log")
	require.NoError(t, err)

	_, err = GetOutputWriter(".....*/*/a**%$&#$@log")
	require.Error(t, err)
}

func TestLevelFromString(t *testing.T) {
	val, err := LevelFromString("debug")
	require.NoError(t, err)
	require.Equal(t, logrus.DebugLevel, val)

	val, err = LevelFromString("Debug")
	require.NoError(t, err)
	require.Equal(t, logrus.DebugLevel, val)

	val, err = LevelFromString("info")
	require.NoError(t, err)
	require.Equal(t, logrus.InfoLevel, val)

	val, err = LevelFromString("warn")
	require.NoError(t, err)
	require.Equal(t, logrus.WarnLevel, val)

	val, err = LevelFromString("error")
	require.NoError(t, err)
	require.Equal(t, logrus.ErrorLevel, val)

	val, err = LevelFromString("fatal")
	require.NoError(t, err)
	require.Equal(t, logrus.FatalLevel, val)

	val, err = LevelFromString("panic")
	require.NoError(t, err)
	require.Equal(t, logrus.PanicLevel, val)

	_, err = LevelFromString("information")
	require.Error(t, err)
}
