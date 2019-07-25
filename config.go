package spider

type Device struct{
	Register[] Register
}
type Register struct{
	Base uint32
	Type string 
	Cmd string 
	Tag string 
	Name string
}
