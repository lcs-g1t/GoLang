rows, _ := db.Query("select max(marca_id) as ultimo from marcas")
				defer rows.Close()
				for rows.Next() {
					rows.Scan(&codigo)
					codigo = codigo + 1
				}