package fs

import (
	"fmt"
	"testing"
	"time"

	"github.com/spf13/pflag"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

// Check it satisfies the interface
var _ pflag.Value = (*Duration)(nil)

func TestParseDuration(t *testing.T) {
	for _, test := range []struct {
		in   string
		want time.Duration
		err  bool
	}{
		{"0", 0, false},
		{"", 0, true},
		{"1ms", time.Millisecond, false},
		{"1s", time.Second, false},
		{"1m", time.Minute, false},
		{"1.5m", (3 * time.Minute) / 2, false},
		{"1h", time.Hour, false},
		{"1d", time.Hour * 24, false},
		{"1w", time.Hour * 24 * 7, false},
		{"1M", time.Hour * 24 * 30, false},
		{"1y", time.Hour * 24 * 365, false},
		{"1.5y", time.Hour * 24 * 365 * 3 / 2, false},
		{"-1s", -time.Second, false},
		{"1.s", time.Second, false},
		{"1x", 0, true},
		{"off", time.Duration(DurationOff), false},
		{"1h2m3s", time.Hour + 2*time.Minute + 3*time.Second, false},
	} {
		duration, err := ParseDuration(test.in)
		if test.err {
			require.Error(t, err)
		} else {
			require.NoError(t, err)
		}
		assert.Equal(t, test.want, duration)
	}
}

func TestDurationString(t *testing.T) {
	for _, test := range []struct {
		in   time.Duration
		want string
	}{
		{time.Duration(0), "0s"},
		{time.Second, "1s"},
		{time.Minute, "1m0s"},
		{time.Millisecond, "1ms"},
		{time.Second, "1s"},
		{(3 * time.Minute) / 2, "1m30s"},
		{time.Hour, "1h0m0s"},
		{time.Hour * 24, "1d"},
		{time.Hour * 24 * 7, "1w"},
		{time.Hour * 24 * 30, "1M"},
		{time.Hour * 24 * 365, "1y"},
		{time.Hour * 24 * 365 * 3 / 2, "1.5y"},
		{-time.Second, "-1s"},
		{time.Second, "1s"},
		{time.Duration(DurationOff), "off"},
		{time.Hour + 2*time.Minute + 3*time.Second, "1h2m3s"},
		{time.Hour * 24, "1d"},
		{time.Hour * 24 * 7, "1w"},
		{time.Hour * 24 * 30, "1M"},
		{time.Hour * 24 * 365, "1y"},
		{time.Hour * 24 * 365 * 3 / 2, "1.5y"},
		{-time.Hour * 24 * 365 * 3 / 2, "-1.5y"},
	} {
		got := Duration(test.in).String()
		assert.Equal(t, test.want, got)
		// Test the reverse
		reverse, err := ParseDuration(test.want)
		assert.NoError(t, err)
		assert.Equal(t, test.in, reverse)
	}
}

func TestDurationReadableString(t *testing.T) {
	for _, test := range []struct {
		negative bool
		in       time.Duration
		want     string
	}{
		// Edge Cases
		{false, time.Duration(DurationOff), "off"},
		// Base Cases
		{false, time.Duration(0), "0s"},
		{true, time.Millisecond, "1ms"},
		{true, time.Second, "1s"},
		{true, time.Minute, "1m"},
		{true, (3 * time.Minute) / 2, "1m30s"},
		{true, time.Hour, "1h"},
		{true, time.Hour * 24, "1d"},
		{true, time.Hour * 24 * 7, "1w"},
		{true, time.Hour * 24 * 365, "1y"},
		// Composite Cases
		{true, time.Hour + 2*time.Minute + 3*time.Second, "1h2m3s"},
		{true, time.Hour * 24 * (365 + 14), "1y2w"},
		{true, time.Hour*24*4 + time.Hour*3 + time.Minute*2 + time.Second, "4d3h2m1s"},
		{true, time.Hour * 24 * (365*3 + 7*2 + 1), "3y2w1d"},
		{true, time.Hour*24*(365*3+7*2+1) + time.Hour*2 + time.Second, "3y2w1d2h1s"},
		{true, time.Hour*24*(365*3+7*2+1) + time.Second, "3y2w1d1s"},
		{true, time.Hour*24*(365+7*2+3) + time.Hour*4 + time.Minute*5 + time.Second*6 + time.Millisecond*7, "1y2w3d4h5m6s7ms"},
	} {
		got := Duration(test.in).ReadableString()
		assert.Equal(t, test.want, got)

		// Test Negative Case
		if test.negative {
			got = Duration(-test.in).ReadableString()
			assert.Equal(t, "-"+test.want, got)
		}
	}
}

func TestDurationScan(t *testing.T) {
	var v Duration
	n, err := fmt.Sscan(" 17m ", &v)
	require.NoError(t, err)
	assert.Equal(t, 1, n)
	assert.Equal(t, Duration(17*60*time.Second), v)
}
