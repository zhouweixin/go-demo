package main

import (
	"fmt"
	"github.com/pkg/errors"
	"golang.org/x/sync/errgroup"
)

func output(num int) (int, error) {
	if num < 0 {
		return 0, errors.New("num < 0")
	}
	return num, nil
}

func main() {
	group := new(errgroup.Group)
	nums := []int{3, 2, -1}

	for _, num := range nums {
		num := num
		group.Go(func() error {
			res, err := output(num)
			fmt.Println(res)
			return err
		})
	}

	if err := group.Wait(); err != nil {
		panic(err)
	}
	fmt.Println("完成")
}