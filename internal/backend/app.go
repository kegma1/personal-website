package backend

import (
	"context"
	"personal_website/repositories"
)


type App struct {
	Ctx context.Context
	Q *repositories.Queries
}	
