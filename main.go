package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

// Todo構造体の定義
type Todo struct {
	ID        int    `json:"id"`
	Title     string `json:"title"`
	Completed bool   `json:"completed"`
}

// 新しいTodoを作成するためのリクエスト構造体
// IDは自動で付与
type CreateTodoRequest struct {
	Title     string `json: "title"`
	Completed bool   `json: "completed"`
}

// エラーレスポンス用の構造体
type ErrorResponse struct {
	Message string `json: "message"`
}

// パッケージレベルで変数を定義（テストケース）
var todos = []Todo{
	{ID: 1, Title: "買い物に行く", Completed: false},
	{ID: 2, Title: "Goの勉強をする", Completed: true},
}

const todoPath = "/api/todos"

// パッケージレベルのハンドラ関数
func handleTodos(w http.ResponseWriter, r *http.Request) {

	// レスポンスヘッダーの設定
	w.Header().Set("Content-Type", "application/json")

	switch r.Method {
	case http.MethodGet:
		// todosをjsonとしてエンコードしてレスポンスに書き込む
		json.NewEncoder(w).Encode(todos)
	case http.MethodPost:
		var createRequest CreateTodoRequest
		if err := json.NewDecoder(r.Body).Decode(&createRequest); err != nil {
			w.WriteHeader(http.StatusBadRequest)
			json.NewEncoder(w).Encode(ErrorResponse{
				Message: "Invalid request body",
			})
			return
		}
	default:
		w.WriteHeader(http.StatusMethodNotAllowed)
		json.NewEncoder(w).Encode(ErrorResponse{
			Message: "Method not allowed. User GET or POST instead",
		})
	}
}

func main() {
	fmt.Printf("Registering handler for path: %s\n", todoPath)
	// ルーティングの登録
	http.HandleFunc(todoPath, handleTodos)

	fmt.Println("Starting server on port 8080")

	http.ListenAndServe(":8080", nil)
}
