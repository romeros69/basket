package chouse

import (
	"context"
	"database/sql"
	"fmt"

	// Добавляем драйвер для ClickHouse
	_ "github.com/ClickHouse/clickhouse-go/v2"
	"github.com/romeros69/basket/config"
)

type Chouse struct {
	DB *sql.DB
}

func New(cfg *config.Config) (*Chouse, error) {
	// Используем DSN для ClickHouse
	db, err := sql.Open("clickhouse", cfg.ClickHouse.ClickHouseURL) // Указываем драйвер "clickhouse"
	if err != nil {
		return nil, fmt.Errorf("ошибка при подключении к ClickHouse: %w", err)
	}

	// Проверяем подключение
	if err = db.Ping(); err != nil {
		return nil, fmt.Errorf("ошибка при проверке соединения с ClickHouse: %w", err)
	}

	// Создание таблицы
	err = createTable(context.Background(), db)
	if err != nil {
		return nil, fmt.Errorf("ошибка при создании таблицы: %w", err)
	}

	return &Chouse{DB: db}, nil
}

// Функция для создания таблицы в ClickHouse
func createTable(ctx context.Context, db *sql.DB) error {
	createTableQuery := `
		CREATE TABLE IF NOT EXISTS player_stats
		(
			player_id        String,  -- Идентификатор игрока
			match_id         Int,     -- Идентификатор матча
			goals            Int,     -- Количество забитых голов
			assists          Int,     -- Количество передач
			interceptions    Int,     -- Количество перехватов
			rebounds         Int      -- Количество подборов
		) 
		ENGINE = MergeTree()
		ORDER BY (player_id, match_id)
	`
	_, err := db.ExecContext(ctx, createTableQuery)
	if err != nil {
		return fmt.Errorf("ошибка при создании таблицы: %w", err)
	}
	return nil
}
