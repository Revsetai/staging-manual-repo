package breaker

import (
	"testing"
	"time"

	"github.com/Revsetai/staging-manual-repo/internal/config"
)

func TestFromConfigCarriesTheTuning(t *testing.T) {
	b := FromConfig(config.Upstream{
		BreakerThreshold: 7,
		BreakerWindow:    2 * time.Minute,
		BreakerCooldown:  90 * time.Second,
	})

	if b.settings.FailureThreshold != 7 {
		t.Errorf("failure threshold = %d, want 7", b.settings.FailureThreshold)
	}
	if b.settings.Window != 2*time.Minute {
		t.Errorf("window = %s, want 2m", b.settings.Window)
	}
	if b.settings.Cooldown != 90*time.Second {
		t.Errorf("cooldown = %s, want 90s", b.settings.Cooldown)
	}
}

func TestFromConfigFallsBackToDefaults(t *testing.T) {
	b := FromConfig(config.Upstream{})
	if b.settings.FailureThreshold != DefaultSettings().FailureThreshold {
		t.Errorf("threshold = %d, want the default", b.settings.FailureThreshold)
	}
}
