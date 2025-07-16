package adapters

import "yourproject/internal/ports"

// PlaywrightAdapter implementiert das WebAutomator-Interface.
type PlaywrightAdapter struct{}

var _ ports.WebAutomator = (*PlaywrightAdapter)(nil)
