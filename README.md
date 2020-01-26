**SWAPI API**

Api em Go com as seguintes funcionalidades:

- Adicionar um planeta (com nome, clima e terreno)

- Listar planetas

- Buscar por nome

- Buscar por ID

- Remover planeta

Sendo que para cada planeta buscado era necessário retorna a quantidade de aparições em filmes, dado obtido na API pública do Star Wars: https://swapi.co/

O banco de dado usado utilizado foi o Cassandra. Para conexão é necessário passar as seguintes variáveis de ambiente:

- CASSANDRA_HOST localhost
- CASSANDRA_USERNAME cassandra
- CASSANDRA_PASSWORD cassandra

Sendo que esses valores foram colocados como padrão no docker file.

A porta da API deve ser passada como variáveis de ambiente:
- API_PORT 8080

Sendo 8080 o valor padrão colocado no DockerFile

Para rodar o código é necessário ter um cassandra iniciado, para isso pode ser usado o comando do docker a seguir:
- docker run --name cassandra --network="host" -d cassandra:3.11

Após isso é necessário aplicar o o schema que esta no arquivo ./cassandra/cassandra.sql.

Para rodar a aplicação com docker é necessário os seguintes comandos:
- docker build -t swapi-api .
- docker run --name swapi -p 8080:8080 --network="host" swapi-ap

