# Currency API

This repository contains a Currency API project that provides users a daily currency exchange rate.

## How to Run

To run the Currency API, follow these steps:

1. Clone the repository:
   ```bash
   git clone https://github.com/marelinaa/currency-api.git
   ```

2. Navigate to the project directory:
   ```bash
   cd currency-api
   ```

3. You need to check if Docker and Docker Compose are installed on your device. You can do this by running the commands "docker" and "docker compose". If installed, run:
   ```bash
   docker-compose up --build
   ```

This application consists only of the server component without a user-friendly interface. 
To interact with it, you will need to utilize tools like Postman or any other API interaction tool. 
### API Endpoints

- **GET /v1/sign-in**: Generates token.
- **GET /v1/currency/date/:date**: Retrieves currency rate for the given date.
- **GET /v1/currency/history/:startDate/:endDate**: Retrieves history of currency rates for the given date period.

### Example

1. To sign in and get token:
  ```
  GET /v1/sign-in
  ```
  In body you need to write user login and password like that: { "login": "Elya", "password": "54321" } There are only 2 available users:
{ "login": "Elya", "password": "54321" }, { "login": "Artyom", "password": "12345" }
<img width="836" alt="Снимок экрана 2024-10-17 в 15 23 27" src="https://github.com/user-attachments/assets/1fafb7d0-750c-4ea8-a5cf-ab412138fca0">


2. To retrieve currency rate for the given date, you also need to use Bearer token:
  ```
  GET /v1/currency/date/2024-10-16
  ```
  <img width="840" alt="Снимок экрана 2024-10-17 в 15 23 51" src="https://github.com/user-attachments/assets/d0c3acc6-d9e6-4bed-83e6-a56380540c9e">

  
3. To retrieve history of currency rates for the given date period, also using Bearer token:
  ```
  GET /v1/currency/history/2024-10-15/2024-10-16
  ```
<img width="845" alt="Снимок экрана 2024-10-17 в 15 25 54" src="https://github.com/user-attachments/assets/be54e089-a28c-4a5f-815f-7bfa3e506280">
