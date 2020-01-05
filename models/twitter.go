package models

// UserEntity holds data for a user that we followed
type UserEntity struct {
	ScreenName        string `json:"screenName"`
	UserID            int64  `json:"userID"`
	FollowedTimestamp int64  `json:"followedTimestamp"`
}
