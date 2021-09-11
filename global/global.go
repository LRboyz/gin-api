package global

import (
	"github.com/LRboyz/gin-api/config"
	ut "github.com/go-playground/universal-translator"
	"go.uber.org/zap"
	"gorm.io/gorm"
)

var (
	Settings config.ServerConfig
	Lg       *zap.Logger
	Trans    ut.Translator
	DB       *gorm.DB
)
