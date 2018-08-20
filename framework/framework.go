package framework


type BoYu struct {
	Banner string
}


func NewAapp() *BoYu{
	app := new(BoYu)
	app.Banner = banner()


	return  app
}


func (by *BoYu) Start(){
	println(by.Banner)
}