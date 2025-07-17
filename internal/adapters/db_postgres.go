package adapters

import "AzubiTool/internal/ports"

// DbPostgresAdapter implementiert Database-Interface.
type DbPostgresAdapter struct{}

var _ ports.Database = (*DbPostgresAdapter)(nil)
