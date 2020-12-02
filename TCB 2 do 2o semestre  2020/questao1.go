package main

import (
	"database/sql"
	"fmt"
	"log"
	"math"

	_ "github.com/go-sql-driver/mysql"
)

func exec(db *sql.DB, sql string) sql.Result {
	result, err := db.Exec(sql)
	if err != nil {
		panic(err)
	}
	return result
}

func main() {

	db, err := sql.Open("mysql", "root:abcd1234@/")
	if err != nil {
		panic(err)
	}

	defer db.Close()

	var a float64
	var b float64
	var c float64
	var x1 float64
	var x2 float64
	codigo := 0
	var valoresa []float64
	var valoresb []float64
	var valoresc []float64
	var valoresxpos []float64
	var valoresxneg []float64
	totalregistros := 0

	fmt.Println("\nCalculando uma equação de segundo grau")
	fmt.Printf("\n")

	fmt.Printf("Digite um valor A: ")
	fmt.Scanln(&a)

	fmt.Printf("Digite um valor B: ")
	fmt.Scanln(&b)

	fmt.Printf("Digite um valor C: ")
	fmt.Scanln(&c)

	x1 = (-b + math.Sqrt(b*b-4*a*c)) / (2 * a)

	x2 = (-b - math.Sqrt(b*b-4*a*c)) / (2 * a)

	fmt.Println("X+ = ", x1)
	fmt.Println("X- = ", x2)

	fmt.Println("\n\nCriando banco de dados Tabela.")
	exec(db, "CREATE DATABASE IF NOT EXISTS tabela")

	exec(db, "USE tabela")

	fmt.Println("Criando tabela a.")
	exec(db, `CREATE TABLE IF NOT EXISTS a(
		valor_ida INT AUTO_INCREMENT PRIMARY KEY,
		valor_a DECIMAL(7,2)
	)`)

	fmt.Println("Criando tabela b.")
	exec(db, `CREATE TABLE IF NOT EXISTS b(
		valor_idb INT AUTO_INCREMENT PRIMARY KEY,
		valor_b DECIMAL(7,2)
	)`)

	fmt.Println("Criando tabela c.")
	exec(db, `CREATE TABLE IF NOT EXISTS c(
		valor_idc INT AUTO_INCREMENT PRIMARY KEY,
		valor_c DECIMAL(7,2)
	)`)

	fmt.Println("Criando tabela xpos.")
	exec(db, `CREATE TABLE IF NOT EXISTS xpos(
		valor_idxpos INT AUTO_INCREMENT PRIMARY KEY,
		valor_xpos DECIMAL(7,2)
	)`)

	fmt.Println("Criando tabela xneg.")
	exec(db, `CREATE TABLE IF NOT EXISTS xneg(
		valor_idxneg INT AUTO_INCREMENT PRIMARY KEY,
		valor_xneg DECIMAL(7,2)
	)`)

	fmt.Println("Adicionando valor a")
	valoresa = append(valoresa, a)

	rows, _ := db.Query("SELECT MAX(valor_ida) AS ultimo FROM a")
	defer rows.Close()
	for rows.Next() {
		rows.Scan(&codigo)
		codigo = codigo + 1
	}
	totalregistros = len(valoresa)
	tx, _ := db.Begin()
	stmt, _ := tx.Prepare("INSERT INTO a(valor_ida, valor_a) VALUES(?,?)")

	for i := 0; i < totalregistros; i++ {
		stmt.Exec(codigo, valoresa[i])
		codigo = codigo + 1
	}

	if err != nil {
		tx.Rollback()
		log.Fatal(err)
	}
	tx.Commit()

	fmt.Println("Adicionando valor b")
	valoresb = append(valoresb, b)

	rows, _ = db.Query("SELECT MAX(valor_idb) AS ultimo FROM b")
	defer rows.Close()
	for rows.Next() {
		rows.Scan(&codigo)
		codigo = codigo + 1
	}
	totalregistros = len(valoresb)
	tx, _ = db.Begin()
	stmt, _ = tx.Prepare("INSERT INTO b(valor_idb, valor_b) values(?,?)")

	for i := 0; i < totalregistros; i++ {
		stmt.Exec(codigo, valoresb[i])
		codigo = codigo + 1
	}

	if err != nil {
		tx.Rollback()
		log.Fatal(err)
	}
	tx.Commit()

	fmt.Println("Adicionando valor c")
	valoresc = append(valoresc, c)

	rows, _ = db.Query("SELECT MAX(valor_idc) AS ultimo FROM c")
	defer rows.Close()
	for rows.Next() {
		rows.Scan(&codigo)
		codigo = codigo + 1
	}
	totalregistros = len(valoresc)
	tx, _ = db.Begin()
	stmt, _ = tx.Prepare("INSERT INTO c(valor_idc, valor_c) values(?,?)")

	for i := 0; i < totalregistros; i++ {
		stmt.Exec(codigo, valoresc[i])
		codigo = codigo + 1
	}

	if err != nil {
		tx.Rollback()
		log.Fatal(err)
	}
	tx.Commit()

	fmt.Println("Adicionando valor x+")
	valoresxpos = append(valoresxpos, x1)

	rows, _ = db.Query("SELECT MAX(valor_idxpos) AS ultimo FROM xpos")
	defer rows.Close()
	for rows.Next() {
		rows.Scan(&codigo)
		codigo = codigo + 1
	}
	totalregistros = len(valoresxpos)
	tx, _ = db.Begin()
	stmt, _ = tx.Prepare("INSERT INTO xpos(valor_idxpos, valor_xpos) VALUES(?,?)")

	for i := 0; i < totalregistros; i++ {
		stmt.Exec(codigo, valoresxpos[i])
		codigo = codigo + 1
	}

	if err != nil {
		tx.Rollback()
		log.Fatal(err)
	}
	tx.Commit()

	fmt.Println("Adicionando valor x-")
	valoresxneg = append(valoresxneg, x2)

	rows, _ = db.Query("SELECT MAX(valor_idxneg) AS ultimo FROM xneg")
	defer rows.Close()
	for rows.Next() {
		rows.Scan(&codigo)
		codigo = codigo + 1
	}
	totalregistros = len(valoresxneg)
	tx, _ = db.Begin()
	stmt, _ = tx.Prepare("INSERT INTO xneg(valor_idxneg, valor_xneg) VALUES(?,?)")

	for i := 0; i < totalregistros; i++ {
		stmt.Exec(codigo, valoresxneg[i])
		codigo = codigo + 1
	}

	if err != nil {
		tx.Rollback()
		log.Fatal(err)
	}
	tx.Commit()
}
