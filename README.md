# Gator
### This is a CLI RSS program that will allow you to create aggregates of feeds from your favorite websites
- Database connection and active user are tracked in a *config file* so that multiple users can interact with the program on the same machine
  - Users can follow each other's feeds once established in the database
- **Postgres** serves as the database for the program, allowing for posts to be retrieved after scraping the websites
  - **Goose** for database migrations
  - **SQLC** to generate Go code from SQL queries for interacting wih the tables
  - Leverages numerous best techniques in database design, such as *ON DELETE CASCADE* to prevent orphaned records
- Parses *XML* into a user-friendly format
- Leverages *middleware* to keep code **DRY**
- Designed to be a live program, it continues running at a user-specified (or default) rate, scraping feeds at a regular interval
- Once scraped, the user can *browse* their followed posts

### In order to run the program on your *own* machine:
- Requires both **Go** and **Postgres** installed on your machine
- Create a config file in your home directory, `~/.gatorconfig.json` with the following content
  - `{
     "db_url":"postgres://postgres:postgres@localhost:5432/gator",
     "current_user_name":""
    }`
- Install the progarm using *go install*
  - `go install github.com/brettlazarine/gator@latest`
- Interact with Gator CLI via one of the following commands:
  - *register* ==> `gator register <username>`
    - registers a user and sets them as the active user
  - *login* ==> `gator login <username>`
    - sets the entered user as the active user
  - *users* ==> `gator users`
    - lists all registerd users
  - *addfeed* ==> `gator addfeed <name> <url>`
    - adds a feed by url with a given feed name to the database, the adding user becomes a follower
  - *follow* ==> `gator follow <url>`
    - adds the active user as a follower of the entered feed by url
  - *following* ==> `gator following`
    - lists all feeds the active user is following
  - *unfollow* ==> `gator unfollow <url>`
    - unfollows active user from an entered feed by url
  - *agg* ==> `gator agg <time interval>`
    - scrapes feeds from following list at the specified interval
    - interval syntax: 30s = 30 seconds, 1m = 1 minute
      - using 30s to 1m as the argument is suggested
    - ***TO CANCEL/END WEB SCRAPING*** ==> ctrl+c
  - *browse* ==> `gator browse <limit **optional**>`
      - prints posts from the feeds that have been scraped by the *agg* command
      - the `<limit>` tag specifies the amount of posts displayed, default of 2
  - *reset* ==> `gator reset`
    - **THIS IS PRIMARILY A DEVELOPMENT COMMAND, IT WILL RESET THE DATABASE**
