package images

import (
	"main/config"
	"os"
)

func InitContainerImage() error {
	cfg, _ := config.Config()
	if _, err := os.Stat(cfg.PathImages); err == nil {

	} else {
		os.MkdirAll(cfg.PathImages, 0755)
	}

	return nil
}
