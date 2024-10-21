Para rodar o usecase de listagem das ordens:

Setup:
1. rodar docker compose up -d -> inicia os containers mysql e rabbitMQ
2. rodar docker ps -> pega o container ID que está rodando o mysql
3. rodar docker exec -it a4c34d47aab1 mysql -uroot -p orders -> acessa o mysql do container na database orders
4. rodar CREATE TABLE `orders` (`id` varchar(255) NOT NULL, `price` float NOT NULL, `tax` float NOT NULL, `final_price` float NOT NULL, PRIMARY KEY (`id`))
5. rodar go run main.go wire_gen.go

Endpoint REST (GET /order)
1. acessar o arquivo 20-CleanArch/api/list_orders.http
2. rodar o GET http://localhost:8000/order -> retorna todas as ordens do banco em json

Service ListOrders com GRPC
1. rodar evans -r repl
2. rodar call ListOrders -> retorna todas as ordens do banco em json

Query ListOrders GraphQL
1. acessar o graphql http://localhost:8080/
2. rodar a query query {
  listOrders {
    id
    Price
    Tax
    FinalPrice
  } -> retorna todas as ordens do banco em json
