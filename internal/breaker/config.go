package breaker

import "github.com/Revsetai/staging-manual-repo/internal/config"

// FromConfig builds a breaker from the resolved upstream settings, so the
// tuning lives in one place instead of being re-derived at each call site.
func FromConfig(u config.Upstream) *Breaker {
	return New(Settings{
		FailureThreshold: u.BreakerThreshold,
		Window:           u.BreakerWindow,
		Cooldown:         u.BreakerCooldown,
	})
}
