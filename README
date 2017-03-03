# This is a Board Game called "华容道" and its resolution

I first try to implement this by a question of interview

Usage:
  go build -o hrd *.go
  ./hrd ./board

File format:
5 4               // board height, width
1 2 1 0 0         // chessman name number, chessman height, chessman width, chessman location[0], chessman location[1]
2 2 2 0 1         // same as above
3 2 1 0 3
4 2 1 2 0
5 1 2 2 1
6 2 1 2 3
7 1 1 3 1
8 1 1 3 2
9 1 1 4 0
10 1 1 4 3




## Project schedule:
* About 4 PM 1 March 2017, I receive this question from the email.
* Before implement this, I google "华容道" to make it clear.
* About 5 PM, I try to implement this
  * For this program, I just use INIT then REPL
  * And I decide to init the starting map, and load to the program
* I design the most important instruction "move"
* I try to abstract the board and chessman, to make all things as number
* I number all the chessman from 1 to 10, and 2 is "曹操"
* After 6 hours， about 11 PM， 1 March, I implement the main feature
  and still need debug the checkCollision function and move function

* At 11 AM, 2 March 2017, the program finally work well, and can receive my command
  "move" to rerender the board
* From 12 AM, I decide to implement some details,
  for example: save current status to a file
* About 12:30 AM, the save feature is done, and do some tests to make sure It will not crash

* About 2 PM 3 March 2017, I start to implement the most complex part, "auto solve" part.
* About 4 PM 3 March 2017, almost done, but met the step loop issue.
* After lunch, about 11:30 PM 3 March 2017, I am checking the loop bug.
  Then I add all passed status into a hash map.
  And make possible steps if the step will not cause a passed status.
* About 00:30 4 March 2017, Finally it work.
  In the following 1 hour, I test the program to avoid panic.
  I also make the pick next step randomization.
  And pprof enabled.

## Issues met:
* c.location can not be update:
    since I pass a copy of c, so the origin c.location can not be update
  solution: pass the pointer of the parameter chessMan
* Out of index:
    when checkCollision and move, met this
  solution: take care of index to make this impossible
* loop for some steps pattern:
    Add all passed status into a hash map.
    Make possible steps if the step will not cause a passed status.
