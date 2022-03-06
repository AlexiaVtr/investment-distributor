package main

// Realiza la repartición de la inversión a los diferentes créditos:
func GetCredit(inv, credit300, credit500, credit700 int32) (int32, int32, int32, error) {
	n := inv

	// Se realizan todas las restas posibles en 6 casos:
	switch {

	case n > 0:
		credit700, credit500, credit300 = getCredit(n, 700, 500, 300)

		//Si los contadores son ">0" entonces se finalizó la distribución correctamente:
		if credit300 != 0 || credit500 != 0 || credit700 != 0 {
			return credit300, credit500, credit700, nil
		}

		//Permite que no se saltee los demás casos:
		fallthrough

	case n > 0:

		// El orden de las restas cambia en cada caso:
		credit300, credit500, credit700 = getCredit(n, 300, 500, 700)
		if credit300 != 0 || credit500 != 0 || credit700 != 0 {
			return credit300, credit500, credit700, nil
		}
		fallthrough

	case n > 0:
		credit500, credit300, credit700 = getCredit(n, 500, 300, 700)
		if credit300 != 0 || credit500 != 0 || credit700 != 0 {
			return credit300, credit500, credit700, nil
		}
		fallthrough

	case n > 0:
		credit500, credit700, credit300 = getCredit(n, 500, 700, 300)
		if credit300 != 0 || credit500 != 0 || credit700 != 0 {
			return credit300, credit500, credit700, nil
		}
		fallthrough

	case n > 0:
		credit700, credit300, credit500 = getCredit(n, 700, 300, 500)
		if credit300 != 0 || credit500 != 0 || credit700 != 0 {
			return credit300, credit500, credit700, nil
		}
		fallthrough
	case n > 0:
		credit300, credit700, credit500 = getCredit(n, 300, 700, 500)
		if credit300 != 0 || credit500 != 0 || credit700 != 0 {
			return credit300, credit500, credit700, nil
		}

	}

	// Si ningúno puede dar resto "0" entonces devuelve los contadores y el error creado:
	return credit300, credit500, credit700, &MyError{}
}

// Reutiiza el ciclo for con montos definidos en los parámetros:
func getCredit(i, a, b, c int32) (int32, int32, int32) {

	// Las variables son contadores para definir la cantidad de créditos a asignar:
	var creditA, creditB, creditC int32

	// "i"     -> inversión
	// "a,b,c" -> montos
	for i > 0 {
		if i >= a {
			i -= a
			creditA += 1
		}
		if i >= b {
			i -= b
			creditB += 1
		}
		if i >= c {
			i -= c
			creditC += 1
		}

		// Si la inversión recibida no se acredito al 100% finaliza el ciclo:

		if i < 300 && i != 0 {
			creditA, creditB, creditC = 0, 0, 0
			break
		}
	}

	//Retorna los contadores sin importar los resultados:
	return creditA, creditB, creditC

}

// Retorna el promedio de lo enviado:
func GetAverage(a int64, b int64) float32 {
	var c float32
	c = float32(a) / float32(b)
	return c
}
