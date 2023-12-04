package main

import (
	"context"
	"fmt"
	"os"

	"github.com/jackc/pgx/v5"
)

func main3() {
	//Connection
	conn, err := pgx.Connect(context.Background(), GetConString())

	if err != nil {
		fmt.Fprintf(os.Stderr, "Unable to connect to database: %v\n", err)
		os.Exit(1)
	}
	defer conn.Close(context.Background())

	var returnedrows pgx.Rows
	//Get Total Count
	returnedrows, err = conn.Query(context.Background(), "select count(*) from seat")
	var rowcount int
	if err == nil {
		for returnedrows.Next() {
			err = returnedrows.Scan(&rowcount)
			if err == nil {
				fmt.Printf("%d total rows \n", rowcount)
			} else {
				println("data mapping error %s", err)
			}
		}

	} else {
		println("query error %s", err)
	}
	//Swap in loop
	for i := 1; i < rowcount; i += 2 {
		var swapbufferi string
		var swapbufferiplus1 string
		returnedrows, err = conn.Query(context.Background(), fmt.Sprintf("select student from seat where id = %d", i))
		if err == nil {
			for returnedrows.Next() {
				err = returnedrows.Scan(&swapbufferi)
				if err == nil {
					returnedrows.Close()

					returnedrows, err = conn.Query(context.Background(), fmt.Sprintf("select student from seat where id = %d", i+1))
					if err == nil {
						for returnedrows.Next() {
							err = returnedrows.Scan(&swapbufferiplus1)
							if err == nil {

								fmt.Println(swapbufferi + swapbufferiplus1)
								returnedrows.Close()
								//Swapping in transaction
								m_tx, err := conn.BeginTx(context.Background(), pgx.TxOptions{})
								if err != nil {
									fmt.Println("begin tx error ", err)

								} else {
									_, err := m_tx.Exec(context.Background(), "update seat set student = $1  where id = $2", swapbufferi, i+1)
									if err != nil {

										fmt.Println("exec tx error ", err)
									}
									_, err = m_tx.Exec(context.Background(), "update seat set student = $1  where id = $2", swapbufferiplus1, i)
									if err != nil {

										fmt.Println("exec tx error ", err)
									}
								}
								defer m_tx.Rollback(context.Background())
								//swapping
								err = m_tx.Commit(context.Background())
								if err != nil {

									fmt.Println("commit tx error ", err)
								}
								//Swaping Transaction Ends
							} else {
								fmt.Println("data mapping error ", err)
							}
						}

					} else {
						fmt.Println("query error ", err)
					}

				} else {
					fmt.Println("data mapping error ", err)
				}
			}

		} else {
			fmt.Println("query error ", err)
		}
	}

	//Display
	returnedrows, err = conn.Query(context.Background(), "select * from seat order by id")
	var greeting string
	var id int
	if err == nil {
		for returnedrows.Next() {
			err = returnedrows.Scan(&id, &greeting)
			if err == nil {
				fmt.Printf("%d %s \n", id, greeting)
			} else {
				fmt.Println("data mapping error ", err)
			}
		}

	} else {
		fmt.Println("query error ", err)
	}

}
