package main

import (
	"context"
	"fmt"
	"os"
	"regexp"

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
		conn, err := pgx.Connect(context.Background(), GetConString())
		if err != nil {
			fmt.Fprintf(os.Stderr, "Unable to connect to database: %v\n", err)
			os.Exit(1)
		}
		defer conn.Close(context.Background())
	}
}
