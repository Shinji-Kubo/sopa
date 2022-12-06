package server

import (
	"bufio"
	"os"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
)

func RootPath(ctx *gin.Context) {
	line := ""
	f, _ := os.Open("data/data.csv")
	defer f.Close()

	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		line = scanner.Text()
	}

	values := strings.Split(line, ",")

	time, _ := time.Parse(time.RFC3339, values[0])

	ctx.Set("time", time.String())
	ctx.Set("value", values[1])

	ctx.HTML(200, "index.html", ctx.Keys)
}
