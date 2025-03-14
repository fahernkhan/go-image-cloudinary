package logger

import (
	"go.uber.org/zap"
)

var Logger *zap.Logger

func InitLogger() {
	var err error
	Logger, err = zap.NewProduction() // Atau gunakan zap.NewDevelopment() untuk mode development
	if err != nil {
		panic("Failed to initialize logger: " + err.Error())
	}

	defer Logger.Sync() // Flush buffer log sebelum aplikasi berhenti
}
