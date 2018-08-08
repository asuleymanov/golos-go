package transactions

//Chain structure storing ID blockchain
type Chain struct {
	ID string
}

//GolosChain ID GOLOS blockchain
var GolosChain = &Chain{
	ID: "782a3039b478c839e4cb0c941ff4eaeb7df40bdd68bd441afd444b9da763de12",
}

//TestChain ID TEST blockchain
var TestChain = &Chain{
	ID: "5876894a41e6361bde2e73278f07340f2eb8b41c2facd29099de9deef6cdb679",
}
