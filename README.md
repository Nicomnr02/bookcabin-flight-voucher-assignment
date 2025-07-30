
# Airline Voucher Seat Assignment App

A voucher seat assignment web application for an airline campaign. For each
flight, the airline wants to randomly assign 3 unique seat numbers, with specific seat maps
depending on aircraft type.
## Candidate

- [Nicolas Raja Oloan Manurung](https://github.com/Nicomnr02)


## Tech Stack

**Client:** React

**Server:** Go

**Database**: SQLite


## Run Locally 

1. Clone the project

```bash
  git clone https://github.com/Nicomnr02/bookcabin-flight-voucher-assignment.git
```

2. Go to the **BE** project directory

```bash
  cd backend
```

3. Install dependencies

```bash
  go mod tidy
```

4. Init SQLite GUI (Optional, Docker ON)

```bash
  make run-sqlite
```

5. Start the server

```bash
  make run-server
```


6. Go to the **FE** project directory

```bash
  cd frontend/voucher-assignments
```

7. Install dependencies

```bash
  npm install
```

8. Start the Web Application

```bash
  npm run dev
```

Go to the Web Application

-  [http://127.0.0.0:5173](http://127.0.0.0:5173)

See the SQLite GUI (Optional, Docker ON)

-  [http://127.0.0.0:8080](http://127.0.0.0:8080)

