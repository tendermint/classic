package abci

type SimpleEvent string

func (_ SimpleEvent) AssertABCIEvent() {}
