package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"log"
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

type IDResponse struct {
	ID string `json:"id"`
}

var (
	playerIDs   []string
	gameIDs     []string
	awardIDs    []string
	leagueIDs   []string
	playerStats []PlayerStat
)

const apiBase = "http://localhost:8080/v1"

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

// Функция для отправки POST-запроса
func sendPostRequest(url string, data interface{}, idField string, idList *[]string) {
	jsonData, err := json.Marshal(data)
	if err != nil {
		log.Printf("Ошибка маршалинга данных: %v", err)
		return
	}

	resp, err := http.Post(url, "application/json", bytes.NewBuffer(jsonData))
	if err != nil {
		log.Printf("Ошибка отправки запроса на %s: %v", url, err)
		return
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusCreated {
		log.Printf("Ошибка: неверный статус код %d при отправке запроса на %s", resp.StatusCode, url)
		return
	}

	var idResp IDResponse
	err = json.NewDecoder(resp.Body).Decode(&idResp)
	if err != nil {
		log.Printf("Ошибка декодирования ответа: %v", err)
		return
	}

	*idList = append(*idList, idResp.ID)
}

// Отправка запросов на создание сущностей и сбор ID
func createEntities() {
	// Создание игроков
	for i := 0; i < 100000; i++ {
		sendPostRequest(apiBase+"/player", generatePlayer(), "playerID", &playerIDs)
	}

	// Создание игр
	for i := 0; i < 100000; i++ {
		sendPostRequest(apiBase+"/game", generateGame(), "gameID", &gameIDs)
	}

	// Создание наград
	for i := 0; i < 100000; i++ {
		sendPostRequest(apiBase+"/award", generateAward(), "awardID", &awardIDs)
	}

	// Создание лиг
	for i := 0; i < 100000; i++ {
		sendPostRequest(apiBase+"/league", generateLeague(), "leagueID", &leagueIDs)
	}

	fmt.Println("Создание всех сущностей завершено")
}

// Генерация PlayerStat и RewardStat с использованием созданных ID
func generatePlayerStat() PlayerStat {
	return PlayerStat{
		Assists:       5,
		AvgGoals:      25.5,
		Goals:         30,
		Interceptions: 2,
		MatchID:       gameIDs[0],  // Использование первого gameID
		PlayerID:      playerIDs[0], // Использование первого playerID
		Rebounds:      12,
		TotalAvgStats: 20.5,
	}
}

func generateRewardStat() RewardStat {
	return RewardStat{
		Match:      gameIDs[0],  // Использование первого gameID
		Player:     playerIDs[0],  // Использование первого playerID
		Reward:     awardIDs[0],  // Использование первого awardID
		Tournament: leagueIDs[0], // Использование первого leagueID
	}
}

// Отправка PlayerStat и RewardStat
func createStats() {
	// Создание статистики игроков
	for i := 0; i < 100000; i++ {
		sendPostRequest(apiBase+"/stat_player", generatePlayerStat(), "statID", &playerIDs)
	}

	// Создание статистики наград
	for i := 0; i < 100000; i++ {
		sendPostRequest(apiBase+"/stat_awards", generateRewardStat(), "statID", &awardIDs)
	}

	fmt.Println("Создание статистики завершено")
}

func main() {
	start := time.Now()

	createEntities() // Создаём сущности и сохраняем их ID
	createStats()    // Создаём PlayerStat и RewardStat, используя созданные ID

	fmt.Printf("Все запросы выполнены за %v\n", time.Since(start))
}