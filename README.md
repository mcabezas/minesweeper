# minesweeper

Step 1:
 I have decided start by the very beginning modeling first the Game entity 
 and a very basic in-memory persistence
 
Step 2:
  Setup of Game endpoints in order to have a quickly piece of operational code
  REST API covered by tests respecting the standard status codes for each scenario

Step 3:
  Swagger documentation setup

Step 4: I have identified a new entity -> The GameBoard (Board)
 * The Board has attached all the board cells
 * I have been working on the Board creation scenario
 * I have decided to use a map for the board to hold all the Cells
   This decision will give me the change to have a lineal complexity when I have to manipulate them.
   
Step 5: Attaching board creation to game creation API
  * In this case when the create_game endpoint is called I'm creating first the board and then the game
  * In case of any issue I have to keep in mind the rollback
  * Because I'm doing a big abstraction on regards of the Storage there are no guarantees of having
    the change to use relational transactions for rollback.
    Because of that I will do a very tiny implementation of saga patter for rollback
      