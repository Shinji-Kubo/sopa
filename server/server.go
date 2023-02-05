package server

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
)

func RootPath(ctx *gin.Context) {
	path := os.Getenv("HOME")

	line := ""
	f, _ := os.Open(fmt.Sprintf("%s/record/data.csv", path))
	defer f.Close()

	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		line = scanner.Text()
	}

	values := strings.Split(line, ",")

	time, err := time.Parse(time.RFC3339, fmt.Sprintf("%sZ", values[0]))
	if err != nil {
		log.Println(err)
	}

	ctx.Set("time", time.String())
	ctx.Set("value", values[1])

	ctx.HTML(200, "index.html", ctx.Keys)
}
