package operations


var (
	TypeFollow         = "follow"
	TypeReblog         = "reblog"
  TypeDeleteReblog = "delete_reblog"
)

var customJSONDataObjects = map[string]interface{}{
	TypeFollow:         &FollowOperation{},
	TypeReblog:         &ReblogOperation{},
  TypeDeleteReblog:         &DeleteReblogOperation{},
}

//FollowOperation the structure for the operation CustomJSONOperation.
type FollowOperation struct {
	Follower  string   `json:"follower"`
	Following string   `json:"following"`
	What      []string `json:"what"`
}

//ReblogOperation the structure for the operation CustomJSONOperation.
type ReblogOperation struct {
	Account  string `json:"account"`
	Author   string `json:"author"`
	Permlink string `json:"permlink"`
 Title string `json:"title"`
Body   string `json:"body"`
JSONMetadata   string `json:"json_metadata"`
}

//DeleteReblogOperation the structure for the operation CustomJSONOperation.
type DeleteReblogOperation struct {
	Account  string `json:"account"`
	Author   string `json:"author"`
	Permlink string `json:"permlink"`
}