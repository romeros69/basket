package neo4j

import (
	"context"
	"fmt"

	"github.com/neo4j/neo4j-go-driver/v5/neo4j"
	"github.com/romeros69/basket/config"
)

type Neo4j struct {
	DB neo4j.SessionWithContext
}

func New(cfg *config.Config) (*Neo4j, error) {
	driver, err := neo4j.NewDriverWithContext(cfg.Neo4j.Neo4jURL, neo4j.BasicAuth(cfg.Neo4j.Neo4jLogin, cfg.Neo4j.Neo4jPassword, ""))
	if err != nil {
		return nil, err
	}

	err = driver.VerifyConnectivity(context.Background())
	if err != nil {
		return nil, fmt.Errorf("error in verify: %w", err)
	}

	// Открываем сессию
	session := driver.NewSession( context.Background(), neo4j.SessionConfig{AccessMode: neo4j.AccessModeWrite})

	return &Neo4j{
		DB: session,
	}, nil
}
