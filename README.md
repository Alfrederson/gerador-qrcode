## Gerador de QR Code de Pix

Funciona assim:

host:porta/pix/nome/cidade/chave/valor

Nome é o nome do beneficiário.
Cidade é a cidade em que a transação ocorre.
Chave é a chave (de qualquer tipo)
Valor é o valor da transferência. É arredondado para 1 centavo. Não tente botar valores negativos.

Exemplo:


GET localhost:8080/pix/Fulanildo/Aracemopolis/fulanildo@bol.com/10.00

Isso gera um QR code para fazer um PIX de R$ 10.00 para a chave fulanildo@bol.com.

Exemplo ao vivo e à cores:

https://gerador-qrcode-r2pk3ukt5q-rj.a.run.app/pix/Fulano/Santos/fulano@bol.com/1.00