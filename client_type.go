package client

//The type is used as a keystroke for a specific user.
//Only a few keys can be set.
type Keys struct {
	PKey string
	AKey string
	OKey string
	MKey string
}

//Type of response when using operations.
type BResp struct {
	ID       string
	BlockNum uint32
	TrxNum   uint32
	Expired  bool
}

//Type for the Comment and Post functions.
//Sets the water to receive payment for a comment or post.
type PCOptions struct {
	Percent uint16
}

type PCVote struct {
	Weight int
}

//Type for MultiTransfer function
type ArrTransfer struct {
	To      string
	Memo    string
	Ammount string
}

//Type for MultiVote function
type ArrVote struct {
	User   string
	Weight int
}
