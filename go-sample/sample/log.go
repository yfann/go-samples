package sample


func Testlog(){
	log := mainLogger.WithName("reconcilers").WithValues("target-type", "Foo")
}