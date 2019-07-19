# Noughts and Crosses in GoLang

The aim is to build a service that allows two users to play noughts and crosses remotely.

## MVP

As this is my first challenge in Go, I have decided that a MVP would be:
- The players would be able to select which box (or position) they wanted to place their "X" or "O", using indexes as coordinates.
- The board would be checked for winning combinations and when there is a winner the game will end and the winner declared.
- If board is full with no winning combinations, winner will be declared as "NONE".

To get the MVP working I created the program to be played in the terminal, then to add in routing once that was working. However this proved to be more time consuming because the functions were structured differently. Ideally I should have started with the routes as it would have been easier to visualise and save time.

## Running
```
go run main.go
```

## Routing
Ideally using postman.

#### Start new game
`POST http://localhost:9000/newgame`

#### See game state (board, current player, winner)
`GET http://localhost:9000/game`

#### Make move
`PUT http://localhost:9000/move`

Query parameters:<br>
I decided to create a multidimensional array depicting the board. The "coordinates" of the board correspond with the indexes of the multidimensional array. For example, to choose the top right corner box, coordinates would be:<br>
```row: 0, col: 2```
```
row     // between 0-2
col     // between 0-2
```

e.g. `PUT http://localhost:9000/move?row=0&col=2`

#### Notes

The json structure of the Game board is shown as a linear array, so it's not displayed as the actual visual of a noughts and crosses box. This can make it difficult to visualise the board however with frontend views this can be solved.

After making the full app work on one page, I decided to separate out the responsibilities into the `app` directory to have a cleaner layout. *There are differing opinions as to what file structure is best to use, so I implemented a simple structure separating out models, api routes, service and views.*
