package main

import (
	"bufio"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"sort"
	"strconv"
	"strings"
)

func findTopFood(f *os.File, topValue int) ([]int, error) {

	var topFood []int
	foodCountMap := make(map[int]int, 0)
	checkMap := make(map[int][]int)
	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		str := scanner.Text()
		eater_id, _ := strconv.Atoi(strings.Split(strings.Split(str, " ")[0], ":")[1])
		foodmenu_id, _ := strconv.Atoi(strings.Split(strings.Split(str, " ")[1], ":")[1])
		foodItems := checkMap[eater_id]
		// fmt.Print("MAP",checkMap)
		for _, v := range foodItems {
			if v == foodmenu_id {
				err := fmt.Errorf("eater_id:%d with foodmenu_id:%d already exists", eater_id, foodmenu_id)
				return topFood, err
			}
		}
		checkMap[eater_id] = append(checkMap[eater_id], foodmenu_id)
		foodCountMap[foodmenu_id] += 1
	}

	type kv struct {
		Key   int
		Value int
	}

	var ss []kv
	for k, v := range foodCountMap {
		ss = append(ss, kv{k, v})
	}

	sort.Slice(ss, func(i, j int) bool {
		return ss[i].Value > ss[j].Value
	})

	if topValue > len(foodCountMap) {
		err := fmt.Errorf("food items are less than %d", topValue)
		return topFood, err
	}

	for i := 0; i < topValue; i++ {
		topFood = append(topFood, ss[i].Key)
	}

	return topFood, nil
}

func main() {

	files, err := ioutil.ReadDir("./logs/")
	if err != nil {
		log.Fatal(err)
	}

	var paths []string
	for _, file := range files {
        paths = append(paths,fmt.Sprintf("./logs/%s",file.Name()))
    }
	
	for _,path := range paths{
		fmt.Println("\nLog Parsing for file: ",path)

		f, err := os.Open(path)
		if err != nil {
			fmt.Print("There has been an error opening log: ", err)
		} else {
			topFood, err := findTopFood(f, 3)
			if err != nil {
				fmt.Println("An error occured:", err)
			} else {
				fmt.Println("Top Three Food Items Ids:", topFood)
			}
		}
		defer f.Close()
	}
	

}
