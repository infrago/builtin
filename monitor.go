package builtin

import "github.com/infrago/infra"

func (m *builtinModule) Ready() bool { return m.loaded }

func (m *builtinModule) Health() infra.ModuleHealth {
	return infra.NewModuleHealth("builtin", m.Ready(), nil)
}

func (m *builtinModule) Stats() infra.ModuleStats {
	return infra.NewModuleStats("builtin", m.Ready(), nil)
}
