package cli

/*
	A mapping of OS to termial clearing functions.
*/
var clear map[string]func()
