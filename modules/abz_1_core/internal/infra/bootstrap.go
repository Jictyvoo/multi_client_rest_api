package infra

import (
	"database/sql"
	"github.com/jictyvoo/multi_client_rest_api/modules/abz_1_core/internal/utils"
	"github.com/wrapped-owls/goremy-di/remy"
)

func init() {
	remy.Register(utils.Injector, remy.Singleton(func(retriever remy.DependencyRetriever) *sql.DB {
		return nil
	}))
}
