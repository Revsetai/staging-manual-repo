package retry

import "github.com/Revsetai/staging-manual-repo/internal/config"

// FromConfig builds a policy from the resolved upstream settings.
func FromConfig(u config.Upstream) Policy {
	p := DefaultPolicy()
	p.MaxAttempts = u.MaxAttempts
	p.BaseDelay = u.RetryBaseDelay
	return p
}
