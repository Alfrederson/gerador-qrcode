## Gerador de QR Code de Pix

# Funciona assim:

```
/copicola/nome/cidade/chave/valor
/pix/nome/cidade/chave/valor
/pix/nome/cidade/chave/valor/tamanho

```

Nome é o nome do beneficiário.
Cidade é a cidade em que a transação ocorre.
Chave é a chave (de qualquer tipo)
Valor é o valor da transferência. É arredondado para 1 centavo. Não tente botar valores negativos.
Tamanho é o tamanho da imagem, limitado à faixa de 128 a 1024 pixels.

É possível gerar QRCodes inválidos com isso porque não ~~quero fazer~~ fiz toda a checagem necessária, mas o aplicativo do banco que for pagar vai simplesmente se recusar a fazer o pagamento. Por exemplo: o aplicativo do NUBank ignora um QR Code para uma transferência de 1.5 milhão de reais.

# Exemplo:

```
GET localhost:8080/pix/Fulanildo/Aracemopolis/fulanildo@bol.com/10.00
```
Ou:
```
GET localhost:8080/pix/Fulanildo/Aracemopolis/fulanildo@bol.com/10.00/128
```

Pra gerar o código em formato de texto, é só fazer assim:
```
GET localhost:8080/copicola/Fulanildo/Aracemopolis/fulanildo@bol.com/10.00
```
Isso gera um br code para fazer um PIX de R$ 10.00 para a chave fulanildo@bol.com.

# Exemplo ao vivo e à cores:

https://gerador-qrcode-r2pk3ukt5q-rj.a.run.app/pix/Fulano/Santos/fulano@bol.com/1.00

https://gerador-qrcode-r2pk3ukt5q-rj.a.run.app/copicola/Fulano/Santos/fulano@bol.com/1.00

# Rodar localmente

Basta clonar o repositório e usar ./fazer para dar build no Dockerfile ou ./rodar para rodar com o go run.

