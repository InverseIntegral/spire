//go:build !windows

package x509_test

var (
	x509RevokeUsage = `Usage of localauthority x509 revoke:
  -authorityID string
    	The authority ID of the X.509 authority to revoke
  -output value
    	Desired output format (pretty, json); default: pretty.
  -socketPath string
    	Path to the SPIRE Server API socket (default "/tmp/spire-server/private/api.sock")
`
)