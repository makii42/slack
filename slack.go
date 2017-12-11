package slack

import (
	"context"
	"errors"
	"log"
	"net/url"
	"os"
)

var logger *log.Logger // A logger that can be set by consumers
/*
  Added as a var so that we can change this for testing purposes
*/
var SLACK_API string = "https://slack.com/api/"
var SLACK_WEB_API_FORMAT string = "https://%s.slack.com/api/users.admin.%s?t=%s"

type SlackResponse struct {
	Ok    bool   `json:"ok"`
	Error string `json:"error"`
}

type AuthTestResponse struct {
	URL    string `json:"url"`
	Team   string `json:"team"`
	User   string `json:"user"`
	TeamID string `json:"team_id"`
	UserID string `json:"user_id"`
}

type authTestResponseFull struct {
	SlackResponse
	AuthTestResponse
}

type Client struct {
	config struct {
		token string
	}
	info  Info
	debug bool
}

// Slack is the public interface to the Slack API Client.
type Slack interface {
	AddPin(channel string, item ItemRef) error
	AddPinContext(ctx context.Context, channel string, item ItemRef) error
	AddReaction(name string, item ItemRef) error
	AddReactionContext(ctx context.Context, name string, item ItemRef) error
	AddStar(channel string, item ItemRef) error
	AddStarContext(ctx context.Context, channel string, item ItemRef) error
	ArchiveChannel(channel string) error
	ArchiveChannelContext(ctx context.Context, channel string) error
	ArchiveGroup(group string) error
	ArchiveGroupContext(ctx context.Context, group string) error
	AuthTest() (response *AuthTestResponse, error error)
	AuthTestContext(ctx context.Context) (response *AuthTestResponse, error error)
	CloseGroup(group string) (bool, bool, error)
	CloseGroupContext(ctx context.Context, group string) (bool, bool, error)
	CloseIMChannel(channel string) (bool, bool, error)
	CloseIMChannelContext(ctx context.Context, channel string) (bool, bool, error)
	ConnectRTM() (info *Info, websocketURL string, err error)
	ConnectRTMContext(ctx context.Context) (info *Info, websocketURL string, err error)
	CreateChannel(channel string) (*Channel, error)
	CreateChannelContext(ctx context.Context, channel string) (*Channel, error)
	CreateChildGroup(group string) (*Group, error)
	CreateChildGroupContext(ctx context.Context, group string) (*Group, error)
	CreateGroup(group string) (*Group, error)
	CreateGroupContext(ctx context.Context, group string) (*Group, error)
	CreateUserGroup(userGroup UserGroup) (UserGroup, error)
	CreateUserGroupContext(ctx context.Context, userGroup UserGroup) (UserGroup, error)
	Debugf(format string, v ...interface{})
	Debugln(v ...interface{})
	DeleteFile(fileID string) error
	DeleteFileContext(ctx context.Context, fileID string) error
	DeleteMessage(channel, messageTimestamp string) (string, string, error)
	DeleteMessageContext(ctx context.Context, channel, messageTimestamp string) (string, string, error)
	DeleteUserPhoto() error
	DeleteUserPhotoContext(ctx context.Context) error
	DisableUser(teamName string, uid string) error
	DisableUserContext(ctx context.Context, teamName string, uid string) error
	DisableUserGroup(userGroup string) (UserGroup, error)
	DisableUserGroupContext(ctx context.Context, userGroup string) (UserGroup, error)
	EnableUserGroup(userGroup string) (UserGroup, error)
	EnableUserGroupContext(ctx context.Context, userGroup string) (UserGroup, error)
	EndDND() error
	EndDNDContext(ctx context.Context) error
	EndSnooze() (*DNDStatus, error)
	EndSnoozeContext(ctx context.Context) (*DNDStatus, error)
	GetAccessLogs(params AccessLogParameters) ([]Login, *Paging, error)
	GetAccessLogsContext(ctx context.Context, params AccessLogParameters) ([]Login, *Paging, error)
	GetBillableInfo(user string) (map[string]BillingActive, error)
	GetBillableInfoContext(ctx context.Context, user string) (map[string]BillingActive, error)
	GetBillableInfoForTeam() (map[string]BillingActive, error)
	GetBillableInfoForTeamContext(ctx context.Context) (map[string]BillingActive, error)
	GetBotInfo(bot string) (*Bot, error)
	GetBotInfoContext(ctx context.Context, bot string) (*Bot, error)
	GetChannelHistory(channel string, params HistoryParameters) (*History, error)
	GetChannelHistoryContext(ctx context.Context, channel string, params HistoryParameters) (*History, error)
	GetChannelInfo(channel string) (*Channel, error)
	GetChannelInfoContext(ctx context.Context, channel string) (*Channel, error)
	GetChannelReplies(channel, thread_ts string) ([]Message, error)
	GetChannelRepliesContext(ctx context.Context, channel, thread_ts string) ([]Message, error)
	GetChannels(excludeArchived bool) ([]Channel, error)
	GetChannelsContext(ctx context.Context, excludeArchived bool) ([]Channel, error)
	GetDNDInfo(user *string) (*DNDStatus, error)
	GetDNDInfoContext(ctx context.Context, user *string) (*DNDStatus, error)
	GetDNDTeamInfo(users []string) (map[string]DNDStatus, error)
	GetDNDTeamInfoContext(ctx context.Context, users []string) (map[string]DNDStatus, error)
	GetEmoji() (map[string]string, error)
	GetEmojiContext(ctx context.Context) (map[string]string, error)
	GetFileInfo(fileID string, count, page int) (*File, []Comment, *Paging, error)
	GetFileInfoContext(ctx context.Context, fileID string, count, page int) (*File, []Comment, *Paging, error)
	GetFiles(params GetFilesParameters) ([]File, *Paging, error)
	GetFilesContext(ctx context.Context, params GetFilesParameters) ([]File, *Paging, error)
	GetGroupHistory(group string, params HistoryParameters) (*History, error)
	GetGroupHistoryContext(ctx context.Context, group string, params HistoryParameters) (*History, error)
	GetGroupInfo(group string) (*Group, error)
	GetGroupInfoContext(ctx context.Context, group string) (*Group, error)
	GetGroups(excludeArchived bool) ([]Group, error)
	GetGroupsContext(ctx context.Context, excludeArchived bool) ([]Group, error)
	GetIMChannels() ([]IM, error)
	GetIMChannelsContext(ctx context.Context) ([]IM, error)
	GetIMHistory(channel string, params HistoryParameters) (*History, error)
	GetIMHistoryContext(ctx context.Context, channel string, params HistoryParameters) (*History, error)
	GetReactions(item ItemRef, params GetReactionsParameters) ([]ItemReaction, error)
	GetReactionsContext(ctx context.Context, item ItemRef, params GetReactionsParameters) ([]ItemReaction, error)
	GetStarred(params StarsParameters) ([]StarredItem, *Paging, error)
	GetStarredContext(ctx context.Context, params StarsParameters) ([]StarredItem, *Paging, error)
	GetTeamInfo() (*TeamInfo, error)
	GetTeamInfoContext(ctx context.Context) (*TeamInfo, error)
	GetUserGroupMembers(userGroup string) ([]string, error)
	GetUserGroupMembersContext(ctx context.Context, userGroup string) ([]string, error)
	GetUserGroups() ([]UserGroup, error)
	GetUserGroupsContext(ctx context.Context) ([]UserGroup, error)
	GetUserIdentity() (*UserIdentityResponse, error)
	GetUserIdentityContext(ctx context.Context) (*UserIdentityResponse, error)
	GetUserInfo(user string) (*User, error)
	GetUserInfoContext(ctx context.Context, user string) (*User, error)
	GetUserPresence(user string) (*UserPresence, error)
	GetUserPresenceContext(ctx context.Context, user string) (*UserPresence, error)
	GetUsers() ([]User, error)
	GetUsersContext(ctx context.Context) ([]User, error)
	InviteGuest(teamName, channel, firstName, lastName, emailAddress string) error
	InviteGuestContext(ctx context.Context, teamName, channel, firstName, lastName, emailAddress string) error
	InviteRestricted(teamName, channel, firstName, lastName, emailAddress string) error
	InviteRestrictedContext(ctx context.Context, teamName, channel, firstName, lastName, emailAddress string) error
	InviteToTeam(teamName, firstName, lastName, emailAddress string) error
	InviteToTeamContext(ctx context.Context, teamName, firstName, lastName, emailAddress string) error
	InviteUserToChannel(channel, user string) (*Channel, error)
	InviteUserToChannelContext(ctx context.Context, channel, user string) (*Channel, error)
	InviteUserToGroup(group, user string) (*Group, bool, error)
	InviteUserToGroupContext(ctx context.Context, group, user string) (*Group, bool, error)
	JoinChannel(channel string) (*Channel, error)
	JoinChannelContext(ctx context.Context, channel string) (*Channel, error)
	KickUserFromChannel(channel, user string) error
	KickUserFromChannelContext(ctx context.Context, channel, user string) error
	KickUserFromGroup(group, user string) error
	KickUserFromGroupContext(ctx context.Context, group, user string) error
	LeaveChannel(channel string) (bool, error)
	LeaveChannelContext(ctx context.Context, channel string) (bool, error)
	LeaveGroup(group string) error
	LeaveGroupContext(ctx context.Context, group string) error
	ListPins(channel string) ([]Item, *Paging, error)
	ListPinsContext(ctx context.Context, channel string) ([]Item, *Paging, error)
	ListReactions(params ListReactionsParameters) ([]ReactedItem, *Paging, error)
	ListReactionsContext(ctx context.Context, params ListReactionsParameters) ([]ReactedItem, *Paging, error)
	ListStars(params StarsParameters) ([]Item, *Paging, error)
	ListStarsContext(ctx context.Context, params StarsParameters) ([]Item, *Paging, error)
	MarkIMChannel(channel, ts string) (err error)
	MarkIMChannelContext(ctx context.Context, channel, ts string) (err error)
	NewRTM() Realtime
	NewRTMWithOptions(options *RTMOptions) Realtime
	OpenGroup(group string) (bool, bool, error)
	OpenGroupContext(ctx context.Context, group string) (bool, bool, error)
	OpenIMChannel(user string) (bool, bool, string, error)
	OpenIMChannelContext(ctx context.Context, user string) (bool, bool, string, error)
	PostMessage(channel, text string, params PostMessageParameters) (string, string, error)
	PostMessageContext(ctx context.Context, channel, text string, params PostMessageParameters) (string, string, error)
	RemovePin(channel string, item ItemRef) error
	RemovePinContext(ctx context.Context, channel string, item ItemRef) error
	RemoveReaction(name string, item ItemRef) error
	RemoveReactionContext(ctx context.Context, name string, item ItemRef) error
	RemoveStar(channel string, item ItemRef) error
	RemoveStarContext(ctx context.Context, channel string, item ItemRef) error
	RenameChannel(channel, name string) (*Channel, error)
	RenameChannelContext(ctx context.Context, channel, name string) (*Channel, error)
	RenameGroup(group, name string) (*Channel, error)
	RenameGroupContext(ctx context.Context, group, name string) (*Channel, error)
	RevokeFilePublicURL(fileID string) (*File, error)
	RevokeFilePublicURLContext(ctx context.Context, fileID string) (*File, error)
	Search(query string, params SearchParameters) (*SearchMessages, *SearchFiles, error)
	SearchContext(ctx context.Context, query string, params SearchParameters) (*SearchMessages, *SearchFiles, error)
	SearchFiles(query string, params SearchParameters) (*SearchFiles, error)
	SearchFilesContext(ctx context.Context, query string, params SearchParameters) (*SearchFiles, error)
	SearchMessages(query string, params SearchParameters) (*SearchMessages, error)
	SearchMessagesContext(ctx context.Context, query string, params SearchParameters) (*SearchMessages, error)
	SendMessage(channel string, options ...MsgOption) (string, string, string, error)
	SendMessageContext(ctx context.Context, channel string, options ...MsgOption) (string, string, string, error)
	SendSSOBindingEmail(teamName, user string) error
	SendSSOBindingEmailContext(ctx context.Context, teamName, user string) error
	SetChannelPurpose(channel, purpose string) (string, error)
	SetChannelPurposeContext(ctx context.Context, channel, purpose string) (string, error)
	SetChannelReadMark(channel, ts string) error
	SetChannelReadMarkContext(ctx context.Context, channel, ts string) error
	SetChannelTopic(channel, topic string) (string, error)
	SetChannelTopicContext(ctx context.Context, channel, topic string) (string, error)
	SetDebug(debug bool)
	SetGroupPurpose(group, purpose string) (string, error)
	SetGroupPurposeContext(ctx context.Context, group, purpose string) (string, error)
	SetGroupReadMark(group, ts string) error
	SetGroupReadMarkContext(ctx context.Context, group, ts string) error
	SetGroupTopic(group, topic string) (string, error)
	SetGroupTopicContext(ctx context.Context, group, topic string) (string, error)
	SetRegular(teamName, user string) error
	SetRegularContext(ctx context.Context, teamName, user string) error
	SetRestricted(teamName, uid string) error
	SetRestrictedContext(ctx context.Context, teamName, uid string) error
	SetSnooze(minutes int) (*DNDStatus, error)
	SetSnoozeContext(ctx context.Context, minutes int) (*DNDStatus, error)
	SetUltraRestricted(teamName, uid, channel string) error
	SetUltraRestrictedContext(ctx context.Context, teamName, uid, channel string) error
	SetUserAsActive() error
	SetUserAsActiveContext(ctx context.Context) error
	SetUserCustomStatus(statusText, statusEmoji string) error
	SetUserCustomStatusContext(ctx context.Context, statusText, statusEmoji string) error
	SetUserPhoto(ctx context.Context, image string, params UserSetPhotoParams) error
	SetUserPhotoContext(ctx context.Context, image string, params UserSetPhotoParams) error
	SetUserPresence(presence string) error
	SetUserPresenceContext(ctx context.Context, presence string) error
	ShareFilePublicURL(fileID string) (*File, []Comment, *Paging, error)
	ShareFilePublicURLContext(ctx context.Context, fileID string) (*File, []Comment, *Paging, error)
	StartRTM() (info *Info, websocketURL string, err error)
	StartRTMContext(ctx context.Context) (info *Info, websocketURL string, err error)
	UnarchiveChannel(channel string) error
	UnarchiveChannelContext(ctx context.Context, channel string) error
	UnarchiveGroup(group string) error
	UnarchiveGroupContext(ctx context.Context, group string) error
	UnsetUserCustomStatus() error
	UnsetUserCustomStatusContext(ctx context.Context) error
	UpdateMessage(channel, timestamp, text string) (string, string, string, error)
	UpdateMessageContext(ctx context.Context, channel, timestamp, text string) (string, string, string, error)
	UpdateUserGroup(userGroup UserGroup) (UserGroup, error)
	UpdateUserGroupContext(ctx context.Context, userGroup UserGroup) (UserGroup, error)
	UpdateUserGroupMembers(userGroup string, members string) (UserGroup, error)
	UpdateUserGroupMembersContext(ctx context.Context, userGroup string, members string) (UserGroup, error)
	UploadFile(params FileUploadParameters) (file *File, err error)
	UploadFileContext(ctx context.Context, params FileUploadParameters) (file *File, err error)
}

// SetLogger let's library users supply a logger, so that api debugging
// can be logged along with the application's debugging info.
func SetLogger(l *log.Logger) {
	logger = l
}

func New(token string) Slack {
	s := &Client{}
	s.config.token = token
	return s
}

// AuthTest tests if the user is able to do authenticated requests or not
func (api *Client) AuthTest() (response *AuthTestResponse, error error) {
	return api.AuthTestContext(context.Background())
}

// AuthTestContext tests if the user is able to do authenticated requests or not with a custom context
func (api *Client) AuthTestContext(ctx context.Context) (response *AuthTestResponse, error error) {
	responseFull := &authTestResponseFull{}
	err := post(ctx, "auth.test", url.Values{"token": {api.config.token}}, responseFull, api.debug)
	if err != nil {
		return nil, err
	}
	if !responseFull.Ok {
		return nil, errors.New(responseFull.Error)
	}
	return &responseFull.AuthTestResponse, nil
}

// SetDebug switches the api into debug mode
// When in debug mode, it logs various info about what its doing
// If you ever use this in production, don't call SetDebug(true)
func (api *Client) SetDebug(debug bool) {
	api.debug = debug
	if debug && logger == nil {
		logger = log.New(os.Stdout, "nlopes/slack", log.LstdFlags|log.Lshortfile)
	}
}

func (api *Client) Debugf(format string, v ...interface{}) {
	if api.debug {
		logger.Printf(format, v...)
	}
}

func (api *Client) Debugln(v ...interface{}) {
	if api.debug {
		logger.Println(v...)
	}
}
