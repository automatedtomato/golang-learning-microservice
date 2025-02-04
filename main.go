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

// パッケージレベルで変数を定義（テストケース）
var todos = []Todo{
	{ID: 1, Title: "買い物に行く", Completed: false},
	{ID: 2, Title: "Goの勉強をする", Completed: true},
}

const todoPath = "/api/todos"

// パッケージレベルのハンドラ関数
func getTodos(w http.ResponseWriter, r *http.Request) {
	// レスポンスヘッダーの設定
	w.Header().Set("Content-Type", "application/json")

	// todosをjsonとしてエンコードしてレスポンスに書き込む
	json.NewEncoder(w).Encode(todos)
}

func main() {
	fmt.Printf("Registering handler for path: %s\n", todoPath)
	// ルーティングの登録
	http.HandleFunc(todoPath, getTodos)

	fmt.Println("Starting server on port 8080")

	http.ListenAndServe(":8080", nil)
}
