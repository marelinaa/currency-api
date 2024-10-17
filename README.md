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
- **GET /v1/currency/date?date=date**: Retrieves currency rate for the given date.
- **GET /v1/currency/history?startDate=startDate&endDate=endDate**: Retrieves history of currency rates for the given date period.

### Example

1. To sign in and get token:
  ```
  GET localhost:8080/v1/sign-in
  ```
In body you need to write user login and password like that: { "login": "user1", "password": "12345" }.  
There are only 2 available users:
{ "login": "user1", "password": "12345" }, { "login": "user2", "password": "54321" }
<img width="845" alt="Снимок экрана 2024-10-17 в 17 31 32" src="https://github.com/user-attachments/assets/d4a6812d-8ecb-4cce-8e88-f5507d1f08e9">


2. To retrieve currency rate for the given date, you also need to use Bearer token:
  ```
  GET localhost:8080/v1/currency/date?date=2024-10-17
  ```
<img width="827" alt="Снимок экрана 2024-10-17 в 17 31 54" src="https://github.com/user-attachments/assets/941e4734-1738-48e1-886f-98eb72705a60">

  
3. To retrieve history of currency rates for the given date period, also using Bearer token:
  ```
  GET localhost:8080/v1/currency/history?startDate=2024-10-16&endDate=2024-10-17
  ```
<img width="841" alt="Снимок экрана 2024-10-17 в 17 32 07" src="https://github.com/user-attachments/assets/648f7a7f-4a19-4898-8be6-29446c3082bd">

