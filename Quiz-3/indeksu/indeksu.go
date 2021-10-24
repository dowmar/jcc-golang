package indeksu

func GetKetebalan(hal string) (tebal string){
	if hal <= "100"{
		tebal = "tipis"
	}else{
		tebal = "tebal"
	}
	return
}