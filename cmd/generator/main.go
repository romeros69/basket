package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"log"
	"math/rand"
	"net/http"
	"time"
)

type Award struct {
	Name    string `json:"name"`
	Surname string `json:"surname"`
}

type Game struct {
	Date       string `json:"date"`
	FirstTeam  string `json:"first_team"`
	SecondTeam string `json:"second_team"`
	League     string `json:"league"`
	Type       string `json:"type"`
}

type League struct {
	Name   string `json:"name"`
	Season string `json:"season"`
}

type Player struct {
	Age         int    `json:"age"`
	Citizenship string `json:"citizenship"`
	Height      int    `json:"height"`
	Name        string `json:"name"`
	Role        string `json:"role"`
	Surname     string `json:"surname"`
	Team        string `json:"team"`
	Weight      int    `json:"weight"`
}

type PlayerStat struct {
	Assists       int     `json:"assists"`
	AvgGoals      float64 `json:"avgGoals"`
	Goals         int     `json:"goals"`
	Interceptions int     `json:"interceptions"`
	MatchID       string  `json:"matchId"`
	PlayerID      string  `json:"playerId"`
	Rebounds      int     `json:"rebounds"`
	TotalAvgStats float64 `json:"totalAvgStats"`
}

type RewardStat struct {
	Match      string `json:"match"`
	Player     string `json:"player"`
	Reward     string `json:"reward"`
	Tournament string `json:"tournament"`
}

type CreatePlayerResp struct {
	PlayerID string `json:"playerID"`
}

type CreateGameResp struct {
	GameID string `json:"game_id"`
}

type CreateAwardResp struct {
	AwardID string `json:"award_id"`
}

type CreateLeagueResp struct {
	LeagueID string `json:"league_id"`
}

var (
	playerIDs   []string
	gameIDs     []string
	awardIDs    []string
	leagueIDs   []string
	playerStats []PlayerStat
)

const apiBase = "http://localhost:8080/v1"

// Генерация случайного текста для уникальности
func randomString(length int) string {
	const charset = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"
	seededRand := rand.New(rand.NewSource(time.Now().UnixNano()))
	b := make([]byte, length)
	for i := range b {
		b[i] = charset[seededRand.Intn(len(charset))]
	}
	return string(b)
}

// Генерация данных для сущностей
func generatePlayer() Player {
	return Player{
		Age:         30,
		Citizenship: "USA",
		Height:      198,
		Name:        "LeBron",
		Role:        "Forward",
		Surname:     "James",
		Team:        "LA Lakers",
		Weight:      113,
	}
}

func generateGame() Game {
	return Game{
		Date:       "12.03.24",
		FirstTeam:  "LA Lakers",
		SecondTeam: "Chicago Bulls",
		League:     "NBA",
		Type:       "Final",
	}
}

func generateAward() Award {
	return Award{
		Name:    "MVP of season 2024",
		Surname: "Best player of season 2024",
	}
}

func generateLeague() League {
	return League{
		Name:   "NBA",
		Season: "2023/2024",
	}
}

// Функция для отправки POST-запроса с корректной обработкой ID в ответе
func sendPostRequest(url string, data interface{}, responseStruct interface{}) error {
	jsonData, err := json.Marshal(data)
	if err != nil {
		log.Printf("Ошибка маршалинга данных: %v", err)
		return err
	}

	resp, err := http.Post(url, "application/json", bytes.NewBuffer(jsonData))
	if err != nil {
		log.Printf("Ошибка отправки запроса на %s: %v", url, err)
		return err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusCreated {
		log.Printf("Ошибка: неверный статус код %d при отправке запроса на %s", resp.StatusCode, url)
		return fmt.Errorf("статус код: %d", resp.StatusCode)
	}

	err = json.NewDecoder(resp.Body).Decode(&responseStruct)
	if err != nil {
		log.Printf("Ошибка декодирования ответа: %v", err)
		return err
	}

	return nil
}

// Отправка запросов на создание сущностей и сбор ID
func createEntities() {
	// Создание игроков
	for i := 0; i < 100000; i++ {
		var resp CreatePlayerResp
		err := sendPostRequest(apiBase+"/player", generatePlayer(), &resp)
		if err == nil {
			playerIDs = append(playerIDs, resp.PlayerID)
		}
	}

	// Создание игр
	for i := 0; i < 100000; i++ {
		var resp CreateGameResp
		err := sendPostRequest(apiBase+"/game", generateGame(), &resp)
		if err == nil {
			gameIDs = append(gameIDs, resp.GameID)
		}
	}

	// Создание наград
	for i := 0; i < 100000; i++ {
		var resp CreateAwardResp
		err := sendPostRequest(apiBase+"/award", generateAward(), &resp)
		if err == nil {
			awardIDs = append(awardIDs, resp.AwardID)
		}
	}

	// Создание лиг
	for i := 0; i < 100000; i++ {
		var resp CreateLeagueResp
		err := sendPostRequest(apiBase+"/league", generateLeague(), &resp)
		if err == nil {
			leagueIDs = append(leagueIDs, resp.LeagueID)
		}
	}

	fmt.Println("Создание всех сущностей завершено")
}

// Генерация PlayerStat с использованием созданных playerID и gameID
func generatePlayerStat() PlayerStat {
	// Проверяем, что playerIDs и gameIDs содержат хотя бы одну запись
	if len(playerIDs) == 0 || len(gameIDs) == 0 {
		log.Println("Ошибка: playerIDs или gameIDs пусты!")
		return PlayerStat{}
	}
	return PlayerStat{
		Assists:       rand.Intn(10),
		AvgGoals:      rand.Float64() * 30,
		Goals:         rand.Intn(50),
		Interceptions: rand.Intn(10),
		MatchID:       gameIDs[rand.Intn(len(gameIDs))],   // Переиспользование случайного gameID
		PlayerID:      playerIDs[rand.Intn(len(playerIDs))], // Переиспользование случайного playerID
		Rebounds:      rand.Intn(20),
		TotalAvgStats: rand.Float64() * 100,
	}
}

// Генерация RewardStat с использованием созданных playerID и gameID
func generateRewardStat() RewardStat {
	// Проверяем, что playerIDs и gameIDs содержат хотя бы одну запись
	if len(playerIDs) == 0 || len(gameIDs) == 0 {
		log.Println("Ошибка: playerIDs или gameIDs пусты!")
		return RewardStat{}
	}
	return RewardStat{
		Match:      gameIDs[rand.Intn(len(gameIDs))],   // Переиспользование случайного gameID
		Player:     playerIDs[rand.Intn(len(playerIDs))], // Переиспользование случайного playerID
		Reward:     "Reward " + randomString(8),         // Случайная строка для награды
		Tournament: "Tournament " + randomString(8),     // Случайная строка для турнира
	}
}

// Отправка PlayerStat и RewardStat
func createStats() {
	// Проверка, есть ли данные для генерации статистики
	if len(playerIDs) == 0 || len(gameIDs) == 0 {
		log.Println("Ошибка: Недостаточно данных для создания статистики")
		return
	}

	// Создание статистики игроков
	for i := 0; i < 100000; i++ {
		playerStat := generatePlayerStat()
		if playerStat.PlayerID != "" && playerStat.MatchID != "" {
			sendPostRequest(apiBase+"/stat_player", playerStat, nil)
		}
	}

	// Создание статистики наград
	for i := 0; i < 100000; i++ {
		rewardStat := generateRewardStat()
		if rewardStat.Player != "" && rewardStat.Match != "" {
			sendPostRequest(apiBase+"/stat_awards", rewardStat, nil)
		}
	}

	fmt.Println("Создание статистики завершено")
}


func main() {
	start := time.Now()

	rand.Seed(time.Now().UnixNano()) // Инициализация генератора случайных чисел

	createEntities() // Создаём сущности и сохраняем их ID
	createStats()    // Создаём PlayerStat и RewardStat, используя созданные ID

	fmt.Printf("Все запросы выполнены за %v\n", time.Since(start))
}