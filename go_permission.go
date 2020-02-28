package main
import ("fmt")
const (
	isAdmin = 1 << iota
	isHeadquarters
	canSeeFinancials
	canSeeAfrica
	canSeeEurope
	canSeeNorthAmerica
	canSeeSouthAmerica)

func main(){
	var roles byte = isAdmin | canSeeFinancials | canSeeEurope
	fmt.Printf("%v\n", 1|3|4)
	fmt.Printf("%v\n",isAdmin)
	fmt.Printf("%b\n",roles)
	fmt.Printf("Is Admin? %v\n", isAdmin & roles ==isAdmin)
}