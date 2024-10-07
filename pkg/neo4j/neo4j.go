package neo4j

import (
	"context"
	"fmt"

	"github.com/neo4j/neo4j-go-driver/v5/neo4j"
	"github.com/neo4j/neo4j-go-driver/v5/neo4j/log"
	"github.com/romeros69/basket/config"
)

type Neo4j struct {
	DB neo4j.SessionWithContext
}


// Construct a new driver

func New(cfg *config.Config) (*Neo4j, error) {
    
    useConsoleLogger := func(level neo4j.LogLevel) func(config *neo4j.Config) {
	    return func(config *neo4j.Config) {
	    	config.Log = log.ToConsole(level)
	    }
    }

    driver, err := neo4j.NewDriverWithContext("neo4j://127.0.0.1:7687", neo4j.NoAuth(), useConsoleLogger(log.DEBUG))

	//driver, err := neo4j.NewDriverWithContext(cfg.Neo4j.Neo4jURL, neo4j.BasicAuth(cfg.Neo4j.Neo4jLogin, cfg.Neo4j.Neo4jPassword, ""))
	//driver, err := neo4j.NewDriverWithContext("neo4j://127.0.0.1:7687", neo4j.BasicAuth(cfg.Neo4j.Neo4jLogin, cfg.Neo4j.Neo4jPassword, ""))
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
