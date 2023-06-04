# golang_food_log_reader

Golang program to read logs of a restaurant.

A restaurant keeps a log of (eater_id, foodmenu_id) for all the diners. The eater_id is a unique number for every diner and foodmenu_id is unique for every food item served on the menu. Write a program that reads this log file and returns the top 3 menu items consumed. If you find an eater_id with the same foodmenu_id more than once then show an error.

Expected output:

1.     Code that can handle the above problem statement

2.     Testcases (with example log files) that checks the possible conditions with unit tests.

3.     Code has to be submitted in your github repository (share the repo link).

4.     Containerize the application and host the image.

Output:

    $ go run main.go

    Log Parsing for file: ./logs/testcase1.log  
    Top Three Food Items Ids: [34 23 12]

    Log Parsing for file: ./logs/testcase2.log  
    An error occured: eater_id:10 with foodmenu_id:12 already exists

    Log Parsing for file: ./logs/testcase3.log  
    An error occured: food items are less than 3
