package train

import (
	"github.com/athunlal/bookNow-Api-Gateway/pkg/auth"
	"github.com/athunlal/bookNow-Api-Gateway/pkg/config"
	"github.com/athunlal/bookNow-Api-Gateway/pkg/train/routes"
	"github.com/gin-gonic/gin"
)

func UserTrainSvc(r *gin.Engine, cfg *config.Config, userSVC *auth.ServiceAuth) {

	svc := &TraiService{
		client: InitTrainService(cfg),
	}

	authorize := auth.InitAuthMiddleware(userSVC)
	r.Use(authorize.AuthRequired)
	train := r.Group("/train")
	{

		train.GET("/search", svc.SearchTrain)
	}
}

func (svc *TraiService) SearchTrain(ctx *gin.Context) {

	routes.SeatchTrainRoute(ctx, svc.client)
}
