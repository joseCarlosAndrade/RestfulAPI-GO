# Criando uma API rest para armazenamento de informações relativas a compras de uma feira

## Restful

A API precisa ser stateless e fornecer requisições com os verbos http:

- GET
- POST
- PUT
- DELETE
- PATCH

## Sugestão de uso

Pacote [gorilla mux](github.com/gorilla/mux) do Golang.

```shell
    go get -u github.com/gorilla/mux
```

---

## Documentação

Para entender melhor os endpoints da API, basta acessar a página estática na raíz do servidor `localhost:3000`, a qual explica melhor como funciona as chamadas.

---

## Compilando e executando

```shell
    make install # instala as dependencias 
    make all # compila
    make run # executa
```
