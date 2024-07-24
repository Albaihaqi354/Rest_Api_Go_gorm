# Pustaka API

Pustaka API is a simple RESTful API for managing books in a library. It is built using Golang with the Gin framework and GORM for ORM.

## Getting Started

These instructions will help you set up and run the Pustaka API on your local machine.

### Prerequisites

- Go (1.16 or later)
- PostgreSQL
- Git

### Installation

1. **Clone the repository:**
    ```sh
    git clone https://github.com/Albaihaqi354/Rest_Api_Golang/
    cd pustaka-api
    ```

2. **Install dependencies:**
    ```sh
    go mod tidy
    ```

3. **Setup PostgreSQL:**

   Create a database named `pustaka-api` and update the connection details in the `connctDb` function in `main.go`:
    ```go
    var dsn = "host=localhost user=youruser password=yourpassword dbname=pustaka-api port=5432 sslmode=disable TimeZone=Asia/Shanghai"
    ```

4. **Run the application:**
    ```sh
    go run main.go
    ```

The server will start on `http://localhost:8080`.

## API Endpoints

### Get All Books

- **URL:** `/v1/books`
- **Method:** `GET`
- **Success Response:**
  - **Code:** 200
  - **Content:** 
    ```json
    {
      "data": [
        {
          "Id": 1,
          "Title": "Book Title",
          "Description": "Book Description",
          "Price": 100,
          "Rating": 5
        },
        ...
      ]
    }
    ```

### Get Book by ID

- **URL:** `/v1/books/:id`
- **Method:** `GET`
- **URL Params:**
  - `id` (int) - The ID of the book
- **Success Response:**
  - **Code:** 200
  - **Content:**
    ```json
    {
      "data": {
        "Id": 1,
        "Title": "Book Title",
        "Description": "Book Description",
        "Price": 100,
        "Rating": 5
      }
    }
    ```

### Create a New Book

- **URL:** `/v1/books`
- **Method:** `POST`
- **Data Params:**
  - `BookRequest` (json) - The book details
    ```json
    {
      "id": 1,
      "title": "New Book",
      "description": "New Description",
      "price": 100,
      "rating": 5
    }
    ```
- **Success Response:**
  - **Code:** 200
  - **Content:**
    ```json
    {
      "data": {
        "Id": 1,
        "Title": "New Book",
        "Description": "New Description",
        "Price": 100,
        "Rating": 5
      }
    }
    ```

## Project Structure

```plaintext
pustaka-api/
├── book/
│   ├── book.go
│   ├── book_request.go
│   ├── book_response.go
│   ├── repository.go
│   └── service.go
├── handler/
│   └── handler.go
├── main.go
├── go.mod
└── go.sum
