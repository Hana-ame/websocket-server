package main

import (
	"fmt"
	"log"
	"net/http"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
)

const DST = "/var/www/moonchan/upload/"

// for gin
func upload(c *gin.Context) {
	// single file
	file, err := c.FormFile("file")
	if err != nil {
		c.String(http.StatusInternalServerError, err.Error())
		return
	}
	log.Println(file.Filename)

	// Upload the file to specific dst.
	fn := strconv.Itoa(int(time.Now().UnixMilli())) + "-" + file.Filename
	err = c.SaveUploadedFile(file, DST+fn)
	if err != nil {
		c.String(http.StatusInternalServerError, err.Error())
		return
	}

	// c.Writer.Header().Set("Location", "/location")
	c.String(http.StatusOK, fmt.Sprintf("'%s' uploaded!", fn))

}
