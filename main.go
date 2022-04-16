package main

import (
	"net/http"
	"time"

	"github.com/flosch/pongo2"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

const tmplPath = "src/template/"
var e = createMux()

func main() {
	e.GET("/", articleIndex)

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

func articleIndex(c echo.Context) error {
	data := map[string]interface{}{
		"message": "Hello World",
		"Now" : time.Now(),
	}
	return render(c, "article/index.html", data)
}

func htmlBlob(file string, data map[string]interface{}) ([]byte, error) {
	return pongo2.Must(pongo2.FromCache(tmplPath + file)).ExecuteBytes(data)
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
