
# Car Shop Go

This is a refactor of a previous Node.js version. I created this API with the intent of learning more about Golang and its features.
Authough it only performs a simple CRUD operation there is more going on. In this repo i didn't apply OOP or SOLID principals, as Mnetioned above
the purpose of this project is learning !!!


## Documentação da API

#### Returns all cars

```
  GET /car
```

| Parâmetro   | Tipo       | Descrição                           |
| :---------- | :--------- | :---------------------------------- |
| `none` | `null` | Returns all cars |

#### Returns one car

```
  GET /car/${id}
```

| Parâmetro   | Tipo       | Descrição                                   |
| :---------- | :--------- | :------------------------------------------ |
| `id`      | `string{24}` | **Obrigatório**. ID of the Car |

#### Creates a Car

```
  POST /car
```

| Parâmetro   | Tipo       | Descrição                                   |
| :---------- | :--------- | :------------------------------------------ |
| `body`      | `Car json Representation` | **Obrigatório**. Fields to update |

#### Updates a Car

```
  PUT /car/${id} 
```

| Parâmetro   | Tipo       | Descrição                                   |
| :---------- | :--------- | :------------------------------------------ |
| `id`      | `string{24}` | **Obrigatório**. ID of the Car |
| `body`      | `Car json Representation` | **Obrigatório**. Fields to update |

#### Delete a Car

```
  DELETE /car/${id} 
```

| Parâmetro   | Tipo       | Descrição                                   |
| :---------- | :--------- | :------------------------------------------ |
| `id`      | `string{24}` | **Obrigatório**. ID of the Car |


## Etiquetas

[![MIT License](https://img.shields.io/badge/License-MIT-green.svg)](https://choosealicense.com/licenses/mit/)
[![GPLv3 License](https://img.shields.io/badge/License-GPL%20v3-yellow.svg)](https://opensource.org/licenses/)
[![AGPL License](https://img.shields.io/badge/license-AGPL-blue.svg)](http://www.gnu.org/licenses/agpl-3.0)


## Running locally

Clone project

```bash
  git clone https://this-project
```

Enter root directory

```bash
  cd my-project
```

Install dependencies

```bash
  go get -d ./...
```

Start a MongoDB Container with no user/password

```bash
  docker run --name <name-of-your-container> -p 27017:27017 -d mongo
```

Run the application

```bash
  go run .
```

