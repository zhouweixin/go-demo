package main

import (
	"encoding/json"
	"fmt"
	"strconv"
)

func main() {
	str := `aaa`
	query, err := parseQueryFromString(str)
	if err != nil {
		panic(err)
	}
	fmt.Println(query)
}

type SearchQuery struct {
	CityIdStr     string `json:"city_id"`
	CityId        int64
	SequenceId    int64 `json:"sequence_id"`
	SubSequenceId int64 `json:"sub_sequence_id"`
	LevelId       int64 `json:"level_id"`
	Score         int16 `json:"score"`
	Status        int16 `json:"status"`
}

func parseQueryFromString(v string) (*SearchQuery, error) {

	var res SearchQuery

	if v == "" {
		return &res, nil
	}

	err := json.Unmarshal([]byte(v), &res)
	if err != nil {
		return nil, err
	}

	if res.CityIdStr != "" {
		res.CityId, err = strconv.ParseInt(res.CityIdStr, 10, 64)
		if err != nil {
			return nil, err
		}
	}

	return &res, nil
}
