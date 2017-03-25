package main

import ( "fmt" )

func header() string {
	return "gobug (Bug tracking system) by Philip Kavanagh";
}

func help() string {
	return	` 
	-a SUBJECT : adds a bug with the provided subject
	-l         : lists bugs
	-c BUGID   : closes a bug defined by the bugid
		`
}

func main(){
	fmt.Println(header());
	fmt.Println(help());
}
