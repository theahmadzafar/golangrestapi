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

func GenerateOTP(ctx *gin.Context) {

	m_phone_number := ctx.Query("phone_number")
	m_phone_number_RegEx := regexp.MustCompile(`[0-9]+`)

	m_phone_number = m_phone_number_RegEx.FindString(m_phone_number)

	if len(m_phone_number) > 0 {
		//Connection
		conn, err := pgx.Connect(context.Background(), conns.GetConString())
		if err != nil {
			fmt.Fprintf(os.Stderr, "Unable to connect to database: %v\n", err)
			os.Exit(1)
		} else {

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

							//Genrate otp
							otp := OTPGen()

							_, err = conn.Query(context.Background(), "update users  set otp=$1,otp_expiration_time=$2 where phone_number=$3", otp, time.Now().Add(time.Minute), m_phone_number)
							if err == nil {
								ctx.JSON(http.StatusOK, gin.H{
									"OTP": otp,
								})
							} else {
								fmt.Println("adsfasfa $1", err)
							}

						} else {
							ctx.JSON(http.StatusNotFound, gin.H{
								m_phone_number: "phone number Does not matched ",
							})
						}
					}
				}
			}
		}
		defer conn.Close(context.Background())
	} else {
		ctx.JSON(http.StatusConflict, gin.H{
			m_phone_number: "Enter Data Carefully",
		})
	}
}

func OTPGen() string {
	return fmt.Sprint(time.Now().Nanosecond())[:4]
}
