# Battleship

## Architecture
Below we are presenting architecture diagram of the project.
Dotted line denote flow during starting new game from waiting for opponent until received by frontend game id.
* Players
    
    Service responsible for handling user accounts.
    Has access to the data base to store user names and theirs password's hash.

* FrontEnd

    Service share GUI with game.

* MatchMaker

    Service handling games. If any user is starting new game, then the service wait until another player will also be waiting for new game. After this service notify Frontend that new game has started.

* GameCore

    Proxy service intermediary between Frontend and

    * Engine

        responsible for notify about win, loss and any validation error, like forbidden ship place.
    
    * Board

        Share necessary information about current state on the board to visualize it.
```mermaid
flowchart TB
    FrontEnd <-->|user, hash| /sign_up
    FrontEnd <-->|user, hash| /login
    FrontEnd -.->|user| /new_game
    /new_game -.->|game_id| /give_game
    /run_battle <-.->|user1, user2| MatchMaker
    /put_ship <-->|user, ship| FrontEnd
    /shot <-->|user, field| FrontEnd
    /get_board <==>|user, game_idp| FrontEnd
    /mark_shot <--> /shot
    Players === db

    subgraph FrontEnd
        /give_game
    end

    subgraph Players
        /sign_up
        /login
    end

    subgraph MatchMaker
        /new_game
    end
    subgraph GameCore
        subgraph Board
            /mark_shot
            /get_board
            /put_ship
        end

        subgraph Engine
            /run_battle
            /shot
        end
    end

    db[(DataBase)]

    style GameCore stroke:#bbf,stroke-dasharray: 5 5
```