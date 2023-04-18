package p2p

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestToPeerAddrInfo(t *testing.T) {
	/// Given
	addresses := []string{
		"/ip4/10.20.30.40/tcp/443/p2p/QmYyQSo1c1Ym7orWxLYvCrM2EmxFTANf8wXmmE7DWjhx5N",
		"/ip4/192.168.0.13/tcp/80/p2p/QmbLHAnMoJPWSCR5Zhtx6BHJX9KiKNN6tpvbUcqanj75Nb",
	}

	/// When
	peersAddrInfo := toAddrInfo(addresses)

	/// Then
	assert.Equal(t, peersAddrInfo[0].ID.String(), "QmYyQSo1c1Ym7orWxLYvCrM2EmxFTANf8wXmmE7DWjhx5N")
	assert.Equal(t, peersAddrInfo[0].Addrs[0].String(), "/ip4/10.20.30.40/tcp/443")

	assert.Equal(t, peersAddrInfo[1].ID.String(), "QmbLHAnMoJPWSCR5Zhtx6BHJX9KiKNN6tpvbUcqanj75Nb")
	assert.Equal(t, peersAddrInfo[1].Addrs[0].String(), "/ip4/192.168.0.13/tcp/80")
}
