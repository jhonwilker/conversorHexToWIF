# conversorHexToWIF
Converte a chave privada em hexadecimal para o formato WIF

#Como Usar
Supondo que a chave privada hexadecimal seja 0x268329e2e8fb78d0a, basta remover o 0x e colocar a frente do parametro -key como mostra o código abaixo.

```
go run main.go -key 268329e2e8fb78d0a
```
