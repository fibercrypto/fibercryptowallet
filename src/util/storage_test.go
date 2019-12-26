package util

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestKeyValuesNoDefaults(t *testing.T) {
	m := NewKeyValueMap()
	require.Equal(t, m, NewKeyValuesWithDefaults(m, nil))
}

func TestKeyValueiWithDefaultsGetValue(t *testing.T) {
	// Values override defaults
	m1 := NewMapWithSingleKey("k1", "v1")
	m2 := NewMapWithSingleKey("k1", "v2")
	s := NewKeyValuesWithDefaults(m1, m2)
	require.Equal(t, "v1", m1.GetValue("k1"))
	require.Equal(t, "v2", m2.GetValue("k1"))
	require.Equal(t, "v1", s.GetValue("k1"))

	// Not in values and not in defaults
	require.Nil(t, m1.GetValue("z"))
	require.Nil(t, m2.GetValue("z"))
	require.Nil(t, s.GetValue("z"))

	// Defaults returned if no value found
	m2.SetValue("k2", "v3")
	require.Nil(t, m1.GetValue("k2"))
	require.Equal(t, "v3", m2.GetValue("k2"))
	require.Equal(t, "v1", m1.GetValue("k1"))
	require.Equal(t, "v2", m2.GetValue("k1"))
	require.Nil(t, m1.GetValue("z"))
	require.Nil(t, m2.GetValue("z"))
	require.Equal(t, "v3", s.GetValue("k2"))
	require.Equal(t, "v1", s.GetValue("k1"))
	require.Nil(t, s.GetValue("z"))

	// Values returned even if no defaults found
	m1.SetValue("k3", "v4")
	require.Equal(t, "v4", m1.GetValue("k3"))
	require.Nil(t, m2.GetValue("k3"))
	require.Nil(t, m1.GetValue("k2"))
	require.Equal(t, "v3", m2.GetValue("k2"))
	require.Equal(t, "v1", m1.GetValue("k1"))
	require.Equal(t, "v2", m2.GetValue("k1"))
	require.Nil(t, m1.GetValue("z"))
	require.Nil(t, m2.GetValue("z"))
	require.Equal(t, "v4", s.GetValue("k3"))
	require.Equal(t, "v3", s.GetValue("k2"))
	require.Equal(t, "v1", s.GetValue("k1"))
	require.Nil(t, s.GetValue("z"))
}
