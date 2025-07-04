# vanillaAPIGo

Super simple CRUD API for GoLang.
Here as a practical reference.

## Features

- GoLang backend (REST API)
- In Memory Application
- Write Out to JSON file - for permanence
- Read In to Local Memory
- Full Crud Functionality
- MUX to avoid lockups/transaction errors

## Getting Started

### Prerequisites

- Go (1.22+)

### Backend Setup

```

go run main.go

```

backend will run on `http://localhost:8080`

Routes/Methods supported:

- `/products`
    - `GET`
    - `POST`
- `/product`
    - `GET`
    - `PUT`
    - `DELETE`

Data Body Shape

```azure
{
"id": "55",
"name": "Weed Wacker",
"price": 44.44
}
```