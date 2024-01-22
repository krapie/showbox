package server

import (
	"context"

	v1 "github.com/krapie/showbox/api/v1"
)

func Run() error {
	router, err := v1.Endpoints(context.Background())
	if err != nil {
		return err
	}

	if err = router.Run(":8080"); err != nil {
		return err
	}

	return nil
}
