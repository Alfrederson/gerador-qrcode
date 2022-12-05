package main

import (
	"errors"
	"os"
	"strconv"
	"fmt"
	"log"
	"net/http"
	"github.com/gin-gonic/gin"
	"github.com/skip2/go-qrcode"
)

type DadosBRCode = struct{
	nome string
	cidade string
	chave string
	valor string
	codigo string
}

func extraiBRCode(c *gin.Context) (DadosBRCode,error){
	resultado := DadosBRCode{codigo : "GeradorQR"}
	resultado.nome = c.Param("nome")
	if resultado.nome ==""{
		return resultado, errors.New("Sem nome")
	}

	resultado.cidade = c.Param("cidade")
	if resultado.cidade ==""{
		return resultado, errors.New("Sem cidade")
	}
	
	resultado.chave = c.Param("chave")
	if resultado.chave == ""{
		return resultado, errors.New("Sem chave")
	}

	valorf,err := strconv.ParseFloat(c.Param("valor"),64)

	if err != nil{
		valorf = 0.01
	}

	resultado.valor = fmt.Sprintf("%0.2f",valorf)

	return resultado,nil
}


// tabela gerada com polinomial 0x1021
var tabela_crc = [...]int {
	0x0000,0x1021,0x2042,0x3063,0x4084,0x50A5,0x60C6,0x70E7,0x8108,0x9129,0xA14A,0xB16B,0xC18C,0xD1AD,0xE1CE,0xF1EF,
	0x1231,0x0210,0x3273,0x2252,0x52B5,0x4294,0x72F7,0x62D6,0x9339,0x8318,0xB37B,0xA35A,0xD3BD,0xC39C,0xF3FF,0xE3DE,
	0x2462,0x3443,0x0420,0x1401,0x64E6,0x74C7,0x44A4,0x5485,0xA56A,0xB54B,0x8528,0x9509,0xE5EE,0xF5CF,0xC5AC,0xD58D,
	0x3653,0x2672,0x1611,0x0630,0x76D7,0x66F6,0x5695,0x46B4,0xB75B,0xA77A,0x9719,0x8738,0xF7DF,0xE7FE,0xD79D,0xC7BC,
	0x48C4,0x58E5,0x6886,0x78A7,0x0840,0x1861,0x2802,0x3823,0xC9CC,0xD9ED,0xE98E,0xF9AF,0x8948,0x9969,0xA90A,0xB92B,
	0x5AF5,0x4AD4,0x7AB7,0x6A96,0x1A71,0x0A50,0x3A33,0x2A12,0xDBFD,0xCBDC,0xFBBF,0xEB9E,0x9B79,0x8B58,0xBB3B,0xAB1A,
	0x6CA6,0x7C87,0x4CE4,0x5CC5,0x2C22,0x3C03,0x0C60,0x1C41,0xEDAE,0xFD8F,0xCDEC,0xDDCD,0xAD2A,0xBD0B,0x8D68,0x9D49,
	0x7E97,0x6EB6,0x5ED5,0x4EF4,0x3E13,0x2E32,0x1E51,0x0E70,0xFF9F,0xEFBE,0xDFDD,0xCFFC,0xBF1B,0xAF3A,0x9F59,0x8F78,
	0x9188,0x81A9,0xB1CA,0xA1EB,0xD10C,0xC12D,0xF14E,0xE16F,0x1080,0x00A1,0x30C2,0x20E3,0x5004,0x4025,0x7046,0x6067,
	0x83B9,0x9398,0xA3FB,0xB3DA,0xC33D,0xD31C,0xE37F,0xF35E,0x02B1,0x1290,0x22F3,0x32D2,0x4235,0x5214,0x6277,0x7256,
	0xB5EA,0xA5CB,0x95A8,0x8589,0xF56E,0xE54F,0xD52C,0xC50D,0x34E2,0x24C3,0x14A0,0x0481,0x7466,0x6447,0x5424,0x4405,
	0xA7DB,0xB7FA,0x8799,0x97B8,0xE75F,0xF77E,0xC71D,0xD73C,0x26D3,0x36F2,0x0691,0x16B0,0x6657,0x7676,0x4615,0x5634,
	0xD94C,0xC96D,0xF90E,0xE92F,0x99C8,0x89E9,0xB98A,0xA9AB,0x5844,0x4865,0x7806,0x6827,0x18C0,0x08E1,0x3882,0x28A3,
	0xCB7D,0xDB5C,0xEB3F,0xFB1E,0x8BF9,0x9BD8,0xABBB,0xBB9A,0x4A75,0x5A54,0x6A37,0x7A16,0x0AF1,0x1AD0,0x2AB3,0x3A92,
	0xFD2E,0xED0F,0xDD6C,0xCD4D,0xBDAA,0xAD8B,0x9DE8,0x8DC9,0x7C26,0x6C07,0x5C64,0x4C45,0x3CA2,0x2C83,0x1CE0,0x0CC1,
	0xEF1F,0xFF3E,0xCF5D,0xDF7C,0xAF9B,0xBFBA,0x8FD9,0x9FF8,0x6E17,0x7E36,0x4E55,0x5E74,0x2E93,0x3EB2,0x0ED1,0x1EF0,
}

func crc(do_que string) int {
	// valor inicial 0xFFFF
	var v = 0xFFFF
	for i := 0 ; i < len(do_que) ; i++{
		v = v << 8 ^ tabela_crc[ byte(v>>8)^do_que[i]]
	}
	return v & 0xFFFF
}

/*
	acho que isso pode ser transformado em um método do DadosBRCode.
*/
func gerarPixCopiaECola(b DadosBRCode) string {
	nome := b.nome
	cidade := b.cidade
	chave := b.chave
	valor := b.valor
	
	codigo := b.codigo // código da transferência sem espaços
	
	seq := "000201" // 00 02 01 payload format indicator
					// 01 02 12 point of initiation não é usado
					// 26 xx    merchant account information  0014BR.GOV.BCB.PIX e depois 01(tamanho da chave)(chave)

	seq += fmt.Sprintf("26%02d", 14 + len(chave) + 8)
	seq += "0014BR.GOV.BCB.PIX"
	seq += fmt.Sprintf("01%02d%s", len(chave), chave)

	seq += "52040000" // 52 04 0000 merchant category code
	seq += "5303986"  // 53 03 986 transaction currency, 986 = real brasileiro                
    seq += fmt.Sprintf("54%02d%s", len(valor), valor) // 54 (tamanho do valor) (valor) => 04 0.01 pix de 1 centavo
	seq += "5802BR"   // 58 02 xx => código do país (é BR)     
	seq += fmt.Sprintf("59%02d%s", len(nome), nome)   // 59 (tamanho do nome) (nome) = nome do beneficiário
	seq += fmt.Sprintf("60%02d%s", len(cidade), cidade) // 60 (tamanho da cidade) (cidade) = cidade onde está acontecendo
	
	var esse_tamanho = 4
	if len(codigo) > 0{
		esse_tamanho += len(codigo)
	}
	seq += fmt.Sprintf("62%02d", esse_tamanho) // 62 (não entendi que tamanho é esse) (05) (tamanho do id) (id)
	if len(codigo) > 0{
		seq += fmt.Sprintf("05%02d%s", len(codigo), codigo)
	}
	// se não tiver isso, o app do banco não realiza a transferência.
	seq += fmt.Sprintf("6304") //%04d",crc)
	// calcular o CRC16 com polinomial 0x1021 e valor inicial 0xFFFF
	// ele tem que ter sempre 4 digitos, então tem o %04 na frente do X
	seq += fmt.Sprintf("%04X", crc(seq))

	log.Println("Código PIX gerado:"+seq)
	return seq
}

func enviarImagem(c *gin.Context, nome string, imagem []byte){
	c.Header("Content-Disposition" , "inline; filename=\""+nome+".png\"")
	c.Data(http.StatusOK,"image/png",imagem)
}

func getPix(c *gin.Context){
	brCode, error := extraiBRCode(c)

	if error !=nil{
		c.String(http.StatusInternalServerError,"Ops! Alguma coisa não funcionou:\n"+error.Error())
		return 
	}
	
	qr := gerarQR(gerarPixCopiaECola(brCode),512)
	enviarImagem(c, "qr_code_para_"+brCode.nome , qr)
}

// acho que dá pra eliminar essa duplicação com uma lambda ou alguma outra
// coisa que eu não quero fazer agora.
func getPixTamanho(c *gin.Context){
	brCode,error := extraiBRCode(c)

	if error != nil{
		c.String(http.StatusInternalServerError, "Ops! Alguma coisa deu errado:\n"+error.Error())
		return
	}

	tamanho,error := strconv.ParseInt(c.Param("tamanho"),10,32)

	if error != nil{
		tamanho = 256
	}

	qr := gerarQR(gerarPixCopiaECola(brCode), int(tamanho))
	enviarImagem(c, "qr_code_para_"+brCode.nome , qr)
}

// violando DRY
func getCopicola(c *gin.Context){
	brCode, error := extraiBRCode(c)

	if error != nil{
		c.String(http.StatusInternalServerError,"Ops! Alguma coisa não funcionou:\n"+error.Error())
		return
	}

	c.String(http.StatusOK,gerarPixCopiaECola(brCode))
}

func gerarQR(txt string, tamanho int) ([]byte){
	var png []byte
	if tamanho > 1024{
		tamanho = 1024
	}
	if tamanho < 128{
		tamanho = 128
	}

	png,err := qrcode.Encode( txt ,qrcode.Medium, tamanho)
	if err != nil{
		return nil
	}
	return png
}

func main(){
	// forçando o negócio do google a dar build!
	router := gin.Default()
	router.GET("/pix/:nome/:cidade/:chave/:valor",getPix)
	router.GET("/pix/:nome/:cidade/:chave/:valor/:tamanho",getPixTamanho)

	router.GET("/copicola/:nome/:cidade/:chave/:valor",getCopicola)


	port := os.Getenv("PORT")
	if port == ""{
		port = "8080"
	}
	router.Run(":"+port)
}