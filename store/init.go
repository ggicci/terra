package store

import (
	"fmt"

	"github.com/io4io/terra/config"
	"github.com/io4io/terra/store/internal"
)

func Init() error {
	masterTerraDB, err := config.GetConfig().TerraDB.Connect()
	if err != nil {
		return fmt.Errorf("connect TerraDB: %w", err)
	}

	replaceGlobalStore(NewStore(internal.NewSqlStore(masterTerraDB)))
	return nil
}
