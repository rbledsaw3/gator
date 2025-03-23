# Blog Aggregator

This repository contains a command-line blog aggregator application written in Go. It fetches RSS feeds, stores posts in a PostgreSQL database, and provides convenient CLI commands for interacting with blog posts and feeds.

## Features

- **RSS Feed Aggregation:** Automatically fetches and parses RSS feeds at configurable intervals.
- **Post Storage:** Stores blog posts in a PostgreSQL database with graceful handling of duplicates.
- **CLI Commands:** Provides commands to browse, follow, unfollow, and list feeds and posts.
- **User Management:** Supports user registration, login, and feed-following functionality.

## Installation

Clone the repository:

```bash
git clone https://github.com/rbledsaw3/blog_aggregator.git
cd blog_aggregator
```

Ensure Go and PostgreSQL are installed and properly configured.

## Configuration

Create a configuration file at `~/.gatorconfig.json` with the following structure:

```json
{
  "db_url": "postgres://user:password@localhost:5432/blog_aggregator?sslmode=disable",
  "current_user_name": ""
}
```

## Database Setup

Migrate the database schema using Goose:

```bash
goose -dir sql/schema postgres "your_db_url_here" up
```

## Usage

### Running the Aggregator

Start aggregating RSS feeds periodically:

```bash
go run . agg 30m
```

### CLI Commands

- **Register a new user:**
  ```bash
  go run . register username
  ```

- **Login as an existing user:**
  ```bash
  go run . login username
  ```

- **Add and follow a new feed:**
  ```bash
  go run . addfeed "Feed Name" "https://feedurl.com/rss"
  ```

- **Browse recent posts:**
  ```bash
  go run . browse [limit]
  ```
  _Default limit is 2 posts._

- **List feeds you are following:**
  ```bash
  go run . following
  ```

- **Unfollow a feed:**
  ```bash
  go run . unfollow "https://feedurl.com/rss"
  ```

### Resetting the Database

To clear all users and feeds:

```bash
go run . reset
```

## Project Structure

```bash
.
├── internal
│   ├── config      # Application configuration management
│   └── database    # Database schema, migrations, and queries
├── sql
│   ├── queries     # SQL queries managed by sqlc
│   └── schema      # Database migrations managed by Goose
├── main.go         # Application entry point
└── handlers        # CLI command handlers
```

## Development

### Dependencies

This project uses Go modules. Dependencies are managed automatically.

### Generating SQL code

After modifying SQL queries, regenerate Go bindings:

```bash
sqlc generate
```


## License

This project is licensed under the MIT License.


