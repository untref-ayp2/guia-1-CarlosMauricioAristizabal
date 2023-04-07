package ordenamientos

func Burbujeo(datos []int) {
	lendatos := len(datos)
	for i := 0; i < lendatos-1; i++ {
		for j := 0; j < lendatos-i-1; j++ {
			swap(datos, j)
		}
	}

}
func swap(dato []int, j int) {
	if dato[j] > dato[j+1] {
		dato[j], dato[j+1] = dato[j+1], dato[j]
	}
}
