package shop

//AppConfig allows to configure the app
type AppConfig struct {
	DBDialect string
	DBArgs    []interface{}
}
