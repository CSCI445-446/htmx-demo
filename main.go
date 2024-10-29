package main

import (
	"database/sql"
	"fmt"
	"net/http"

	"github.com/a-h/templ"
	"github.com/go-sql-driver/mysql"
	"github.com/labstack/echo/v4"

	"htmx-demo/templates"
	"htmx-demo/types"
)

func connect() *sql.DB {
	config := mysql.Config{
		User: "htmx",
		Passwd: "password",
		DBName: "htmx_demo",
	}

	sql, _ := sql.Open("mysql", config.FormatDSN())

	return sql
}

func searchCustomers(conn *sql.DB, searchString string) []types.Customer {
	stmt, _ := conn.Prepare("SELECT * FROM customer WHERE last_name LIKE ?")
	rows, _ := stmt.Query(fmt.Sprintf("%%%s%%", searchString))

	var customer types.Customer
	var customers []types.Customer
	for rows.Next() {
		_ = rows.Scan(&customer.ID, &customer.FirstName, &customer.LastName, &customer.Email)
		customers = append(customers, customer)
	}

	return customers
}

func main() {
	e := echo.New()

	e.GET("/user_lookup", func(ctx echo.Context) error {
		return Render(ctx, http.StatusOK, templates.CustomerSearch())
	})

	e.GET("/user_search", func(ctx echo.Context) error {
		conn := connect()
		searchString := ctx.QueryParam("lastName")
		customers := searchCustomers(conn, searchString)
		return Render(ctx, http.StatusOK, templates.SearchResults(customers))
	})

	e.Logger.Fatal(e.Start(":8000"))
}

func Render(ctx echo.Context, statusCode int, t templ.Component) error {
	buf := templ.GetBuffer()
	defer templ.ReleaseBuffer(buf)

	if err := t.Render(ctx.Request().Context(), buf); err != nil {
		return err
	}

	return ctx.HTML(statusCode, buf.String())
}
