package projects

import (
	"fmt"
)

const (
	isAdmin = 1 << iota //1
	isHeadquarters
	canSeeFinancials //100

	canSeeAfrica //1000
	canSeeAsia   //10000
	canSeeEurope //100000
	canSeeNorthAmerica
	canSeeSouthAmerica
)

func AdminAccess() {
	//assigned Roles
	var roles byte = isAdmin | canSeeFinancials | canSeeEurope
	fmt.Printf("%b \n", roles)
	fmt.Printf("Is a Admin ? %v \n", isAdmin&roles == isAdmin)
	fmt.Printf("Can You See Asia ? %v \n", canSeeAsia&roles == canSeeAsia)
	fmt.Printf("Can You See Europe ? %v \n", canSeeEurope&roles == canSeeEurope)
}
