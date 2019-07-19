## Noughts and Crosses in GoLang

The aim is to build a service that allows two users to play noughts and crosses remotely.

#### MVP

As this is my first challenge in Go, I have decided that a MVP would be to be able to have two users play on the same computer instead of remotely. The players would be able to select which box (or position) they wanted to place their "X" or "O", using indexes as coordinates. The board would be checked for winning combinations and when there is a winner the game will end and the winner declared. If board is full with no winning combinations, winner will be declared as "NONE".

To get the MVP working I created the program to be played in the terminal, then to add in routing once that was working. However this proved to be more time consuming because the functions were structured differently. Ideally I should have started with the routes as it would have been easier to visualise and save time.

#### Notes

I decided to create a multidimensional array depicting the board. The "coordinates" of the board correspond with the indexes of the multidimensional array. For example, to choose the top right corner box, coordinates would be:
```row: 0, col: 2```
