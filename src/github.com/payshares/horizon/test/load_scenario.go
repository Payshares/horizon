package test

import (
	"github.com/payshares/horizon/test/scenarios"
)

func loadScenario(scenarioName string, includeHorizon bool) {
	paysharesCorePath := scenarioName + "-core.sql"
	horizonPath := scenarioName + "-horizon.sql"

	if !includeHorizon {
		horizonPath = "blank-horizon.sql"
	}

	scenarios.Load(PaysharesCoreDatabaseURL(), paysharesCorePath)
	scenarios.Load(DatabaseURL(), horizonPath)
}
