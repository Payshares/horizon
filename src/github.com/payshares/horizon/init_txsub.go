package horizon

import (
	"net/http"

	"github.com/payshares/horizon/db2/core"
	"github.com/payshares/horizon/db2/history"
	"github.com/payshares/horizon/txsub"
	results "github.com/payshares/horizon/txsub/results/db"
	"github.com/payshares/horizon/txsub/sequence"
)

func initSubmissionSystem(app *App) {
	cq := &core.Q{Session: app.CoreSession(nil)}

	app.submitter = &txsub.System{
		Pending:         txsub.NewDefaultSubmissionList(),
		Submitter:       txsub.NewDefaultSubmitter(http.DefaultClient, app.config.PaysharesCoreURL),
		SubmissionQueue: sequence.NewManager(),
		Results: &results.DB{
			Core:    cq,
			History: &history.Q{Session: app.HorizonSession(nil)},
		},
		Sequences:         cq.SequenceProvider(),
		NetworkPassphrase: app.networkPassphrase,
	}
}

func init() {
	appInit.Add("txsub", initSubmissionSystem, "app-context", "log", "horizon-db", "core-db")
}
