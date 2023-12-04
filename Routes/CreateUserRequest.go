package routes

import (
	"context"
	"fmt"
	"net/http"
	"os"
	"regexp"
	conns "run/question1/Connection"

	"github.com/gin-gonic/gin"
	"github.com/jackc/pgx/v5"
)

func CreateNewUser(ctx *gin.Context) {

	m_name := ctx.Query("name")
	m_phone_number := ctx.Query("phone_number")
	m_phone_number_RegEx := regexp.MustCompile(`[0-9]+`)

	m_phone_number = m_phone_number_RegEx.FindString(m_phone_number)

	if len(m_name) > 0 && len(m_phone_number) > 0 {
		//Connection
		conn, err := pgx.Connect(context.Background(), conns.GetConString())
		if err != nil {
			fmt.Fprintf(os.Stderr, "Unable to connect to database: %v\n", err)
			os.Exit(1)
		}
		defer conn.Close(context.Background())

		var returnedrows pgx.Rows
		//Get Total Count
		returnedrows, err = conn.Query(context.Background(), fmt.Sprintf("select count(*) from users where phone_number='%s'", m_phone_number))
		var rowcount int
		if err == nil {
			for returnedrows.Next() {
				err = returnedrows.Scan(&rowcount)
				returnedrows.Close()
				if err == nil {
					if rowcount > 0 {
						ctx.JSON(http.StatusBadRequest, gin.H{
							m_phone_number: "phone number Already Exists",
						})
					} else {

						//add values to table
						returnedrows, err = conn.Query(context.Background(), fmt.Sprintf("insert into users (name,phone_number) values ('%s','%s')", m_name, m_phone_number))
						if err == nil {

							ctx.JSON(http.StatusOK, gin.H{
								m_phone_number: "success",
							})
						} else {
							fmt.Println(err)
						}

					}

				} else {
					println("data mapping error %s", err)
				}
			}

		} else {
			println("query error %s", err)
		}

	} else {
		ctx.JSON(http.StatusConflict, gin.H{
			m_phone_number: "Enter Data Carefully",
			m_name:         "Enter Data Carefully",
		})
	}

}
