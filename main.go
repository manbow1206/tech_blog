package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"strconv"
	"time"

	"github.com/flosch/pongo2"
	_ "github.com/go-sql-driver/mysql" // Using MySQL driver
	"github.com/jmoiron/sqlx"
	"github.com/joho/godotenv"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

const tmplPath = "src/template/"
var db *sqlx.DB
var e = createMux()

func main() {
	err := godotenv.Load(".env")
	if err != nil {
		// .env読めなかった場合の処理
}
	db = connectDB()

	e.GET("/", articleIndex)
	e.GET("/", articleIndex)
	e.GET("/new", articleNew)
	e.GET("/:id", articleShow)
	e.GET("/:id/edit", articleEdit)

	e.Logger.Fatal(e.Start(":8180"))
}

func createMux() *echo.Echo {
	// アプリケーションインスタンスの生成
	e := echo.New()

	// アプリケーションに各種ミドルウェアを設定
	e.Use(middleware.Recover())
	e.Use(middleware.Logger())
	e.Use(middleware.Gzip())

	e.Static("/css", "src/css")
	e.Static("/js", "src/js")
	// アプリケーションインスタンスを返却
	return e
}

func connectDB() *sqlx.DB {
	dsn := os.Getenv("DSN")
	fmt.Println(dsn)
	db, err := sqlx.Open("mysql", dsn)
	if err != nil {
		e.Logger.Fatal(err)
	}
	if err := db.Ping(); err != nil {
		e.Logger.Fatal(err)
}
log.Println("db connection succeeded")
return db
}

func articleIndex(c echo.Context) error {
	data := map[string]interface{}{
		"Message": "Hello World!",
		"Now":     time.Now(),
	}
	return render(c, "article/index.html", data)
}

func htmlBlob(file string, data map[string]interface{}) ([]byte, error) {
	return pongo2.Must(pongo2.FromCache(tmplPath + file)).ExecuteBytes(data)
}

func articleNew(c echo.Context) error {
	data := map[string]interface{}{
		"Message": "Article New",
		"Now":     time.Now(),
	}

	return render(c, "article/new.html", data)
}

func articleShow(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("id"))

	data := map[string]interface{}{
		"Message": "Article Show",
		"Now":     time.Now(),
		"ID":      id,
	}

	return render(c, "article/show.html", data)
}

func articleEdit(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("id"))

	data := map[string]interface{}{
		"Message": "Article Edit",
		"Now":     time.Now(),
		"ID":      id,
	}

	return render(c, "article/edit.html", data)
}

func render(c echo.Context, file string, data map[string]interface{}) error {
	// 定義したhtmlBlob()関数を呼び出し、生成されたHTMLをバイトデータとして受け取る
	b, err := htmlBlob(file, data)

	// エラーチェック
	if err != nil {
		return c.NoContent(http.StatusInternalServerError)
	}

	// ステータスコード200でHTMLデータをレスポンス
	return c.HTMLBlob(http.StatusOK, b)
}
