package horizon

func initPaysharesCoreInfo(app *App) {
	app.UpdatePaysharesCoreInfo()
	return
}

func init() {
	appInit.Add("paysharesCoreInfo", initPaysharesCoreInfo, "app-context", "log")
}
