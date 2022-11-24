package main

import (
	"os"
	"fmt"
	"net/http"
	"github.com/gin-gonic/gin"
	"github.com/skip2/go-qrcode"
)

var contador = 0

func getPix(c *gin.Context){
	mensagem := c.Param("mensagem")

	qr := gerarQR(mensagem)
	
	if qr != nil{
		c.Data(http.StatusOK,"image/png",qr)
	}else{
		c.String(http.StatusInternalServerError,"Ops! Alguma coisa não funcionou.")
	}
}

func gerarQR(txt string) ([]byte){
	var png []byte
	contador++
	png,err := qrcode.Encode( fmt.Sprintf("Prefixo qualquer\n%d\n%s",contador,txt),qrcode.Medium, 256)
	if err != nil{
		return nil
	}
	return png
}

func main(){

	router := gin.Default()
	router.GET("/pix/:mensagem",getPix)

	port := os.Getenv("PORT")
	if port == ""{
		port = "8080"
	}
	router.Run(":"+port)
}