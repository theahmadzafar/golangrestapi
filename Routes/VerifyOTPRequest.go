package routes

import (
	"context"
	"fmt"
	"net/http"
	"os"
	"regexp"
	conns "run/question1/Connection"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/jackc/pgx/v5"
)

func VerifyOTP(ctx *gin.Context) {
	m_otp := ctx.Query("otp")
	m_otp_RegEx := regexp.MustCompile(`[0-9]+`)
	m_otp = m_otp_RegEx.FindString(m_otp)

	m_phone_number := ctx.Query("phone_number")
	m_phone_number_RegEx := regexp.MustCompile(`[0-9]+`)

	m_phone_number = m_phone_number_RegEx.FindString(m_phone_number)

	if len(m_otp) > 0 && len(m_otp) < 5 && len(m_phone_number) > 0 {
		//Connection
		conn, err := pgx.Connect(context.Background(), conns.GetConString())
		if err != nil {
			fmt.Fprintf(os.Stderr, "Unable to connect to database: %v\n", err)
			os.Exit(1)
		} else {

			var returnedrows pgx.Rows
			//Get Total Count
			returnedrows, err = conn.Query(context.Background(), "select otp,otp_expiration_time from users where phone_number=$1", m_phone_number)
			var m_otp_db string
			var m_otp_expiration_time_db time.Time
			if err == nil {
				for returnedrows.Next() {
					err = returnedrows.Scan(&m_otp_db, &m_otp_expiration_time_db)
					returnedrows.Close()

					if err == nil {
						if m_otp_db == m_otp {

							//check for otp existance

							if m_otp_expiration_time_db.After(time.Now()) {
								ctx.JSON(http.StatusAccepted, gin.H{
									m_phone_number: "Verified ",
								})

							} else {
								ctx.JSON(http.StatusNotFound, gin.H{
									m_phone_number: "otp Expired ",
								})
							}

						} else {
							ctx.JSON(http.StatusNotFound, gin.H{
								m_phone_number: "otp Does not matched ",
							})
						}
					}
				}
			} else {
				ctx.JSON(http.StatusNotFound, gin.H{
					m_phone_number: "Enter Data carefully ",
				})
			}
		}
		defer conn.Close(context.Background())

	}
}
