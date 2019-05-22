package slack

type AppHomeOpened struct {
	Channel string `json: "channel"`
	EventTs string `json: "event_ts"`
	Type    string `json: "type"`
	User    string `json: "user"`
}

type ChannelRename struct {
	Channel struct {
		Created float64 `json: "created"`
		Id      string  `json: "id"`
		Name    string  `json: "name"`
	} `json: "channel"`
	Type string `json: "type"`
}

type GroupRename struct {
	Channel struct {
		Created float64 `json: "created"`
		Id      string  `json: "id"`
		Name    string  `json: "name"`
	} `json: "channel"`
	Type string `json: "type"`
}

type ImOpen struct {
	Channel string `json: "channel"`
	Type    string `json: "type"`
	User    string `json: "user"`
}

type PinAdded struct {
	ChannelId string                 `json: "channel_id"`
	EventTs   string                 `json: "event_ts"`
	Item      map[string]interface{} `json: "item"`
	Type      string                 `json: "type"`
	User      string                 `json: "user"`
}

type TeamProfileDelete struct {
	Profile struct {
		Fields []string `json: "fields"`
	} `json: "profile"`
	Type string `json: "type"`
}

type FilePublic struct {
	File struct {
		Id string `json: "id"`
	} `json: "file"`
	FileId string `json: "file_id"`
	Type   string `json: "type"`
}

type GridMigrationFinished struct {
	ApiAppId string `json: "api_app_id"`
	Event    struct {
		EnterpriseId string `json: "enterprise_id"`
		Type         string `json: "type"`
	} `json: "event"`
	EventId   string  `json: "event_id"`
	EventTime float64 `json: "event_time"`
	TeamId    string  `json: "team_id"`
	Token     string  `json: "token"`
	Type      string  `json: "type"`
}

type GroupDeleted struct {
	Channel string `json: "channel"`
	Type    string `json: "type"`
}

type MessageMpim struct {
	ApiAppId    string   `json: "api_app_id"`
	AuthedTeams []string `json: "authed_teams"`
	Event       struct {
		Channel     string `json: "channel"`
		ChannelType string `json: "channel_type"`
		EventTs     string `json: "event_ts"`
		Text        string `json: "text"`
		Ts          string `json: "ts"`
		Type        string `json: "type"`
		User        string `json: "user"`
	} `json: "event"`
	EventId   string  `json: "event_id"`
	EventTime float64 `json: "event_time"`
	TeamId    string  `json: "team_id"`
	Token     string  `json: "token"`
	Type      string  `json: "type"`
}

type SubteamSelfRemoved struct {
	SubteamId string `json: "subteam_id"`
	Type      string `json: "type"`
}

type TeamDomainChange struct {
	Domain string `json: "domain"`
	Type   string `json: "type"`
	Url    string `json: "url"`
}

type BotAdded struct {
	Bot struct {
		AppId string `json: "app_id"`
		Icons struct {
			Image48 string `json: "image_48"`
		} `json: "icons"`
		Id   string `json: "id"`
		Name string `json: "name"`
	} `json: "bot"`
	Type string `json: "type"`
}

type ChannelCreated struct {
	Channel struct {
		Created float64 `json: "created"`
		Creator string  `json: "creator"`
		Id      string  `json: "id"`
		Name    string  `json: "name"`
	} `json: "channel"`
	Type string `json: "type"`
}

type FileCommentAdded struct {
	Comment map[string]interface{} `json: "comment"`
	File    struct {
		Id string `json: "id"`
	} `json: "file"`
	FileId string `json: "file_id"`
	Type   string `json: "type"`
}

type ImMarked struct {
	Channel string `json: "channel"`
	Ts      string `json: "ts"`
	Type    string `json: "type"`
}

type ReactionRemoved struct {
	EventTs string `json: "event_ts"`
	Item    struct {
		Channel string `json: "channel"`
		Ts      string `json: "ts"`
		Type    string `json: "type"`
	} `json: "item"`
	ItemUser string `json: "item_user"`
	Reaction string `json: "reaction"`
	Type     string `json: "type"`
	User     string `json: "user"`
}

type SubteamSelfAdded struct {
	SubteamId string `json: "subteam_id"`
	Type      string `json: "type"`
}

type SubteamUpdated struct {
	Subteam struct {
		AutoType    string      `json: "auto_type"`
		CreatedBy   string      `json: "created_by"`
		DateCreate  float64     `json: "date_create"`
		DateDelete  float64     `json: "date_delete"`
		DateUpdate  float64     `json: "date_update"`
		DeletedBy   interface{} `json: "deleted_by"`
		Description string      `json: "description"`
		Handle      string      `json: "handle"`
		Id          string      `json: "id"`
		IsExternal  bool        `json: "is_external"`
		IsUsergroup bool        `json: "is_usergroup"`
		Name        string      `json: "name"`
		Prefs       struct {
			Channels []interface{} `json: "channels"`
			Groups   []interface{} `json: "groups"`
		} `json: "prefs"`
		TeamId    string   `json: "team_id"`
		UpdatedBy string   `json: "updated_by"`
		UserCount string   `json: "user_count"`
		Users     []string `json: "users"`
	} `json: "subteam"`
	Type string `json: "type"`
}

type ChannelJoined struct {
	Channel map[string]interface{} `json: "channel"`
	Type    string                 `json: "type"`
}

type FileCommentDeleted struct {
	Comment string `json: "comment"`
	File    struct {
		Id string `json: "id"`
	} `json: "file"`
	FileId string `json: "file_id"`
	Type   string `json: "type"`
}

type GroupArchive struct {
	Channel string `json: "channel"`
	Type    string `json: "type"`
}

type GroupLeft struct {
	Channel string `json: "channel"`
	Type    string `json: "type"`
}

type Hello struct {
	Type string `json: "type"`
}

type LinkShared struct {
	ApiAppId    string   `json: "api_app_id"`
	AuthedUsers []string `json: "authed_users"`
	Event       struct {
		Channel string `json: "channel"`
		Links   []struct {
			Domain string `json: "domain"`
			Url    string `json: "url"`
		} `json: "links"`
		MessageTs string `json: "message_ts"`
		ThreadTs  string `json: "thread_ts"`
		Type      string `json: "type"`
		User      string `json: "user"`
	} `json: "event"`
	EventId   string  `json: "event_id"`
	EventTime float64 `json: "event_time"`
	TeamId    string  `json: "team_id"`
	Token     string  `json: "token"`
	Type      string  `json: "type"`
}

type ScopeDenied struct {
	ApiAppId    string        `json: "api_app_id"`
	AuthedTeams []interface{} `json: "authed_teams"`
	Event       struct {
		Scopes    []string `json: "scopes"`
		TriggerId string   `json: "trigger_id"`
		Type      string   `json: "type"`
	} `json: "event"`
	EventId   string  `json: "event_id"`
	EventTime float64 `json: "event_time"`
	TeamId    string  `json: "team_id"`
	Token     string  `json: "token"`
	Type      string  `json: "type"`
}

type Conversation struct {
	Channel struct {
		Created            float64       `json: "created"`
		Creator            string        `json: "creator"`
		Id                 string        `json: "id"`
		IsArchived         bool          `json: "is_archived"`
		IsChannel          bool          `json: "is_channel"`
		IsExtShared        bool          `json: "is_ext_shared"`
		IsGeneral          bool          `json: "is_general"`
		IsGroup            bool          `json: "is_group"`
		IsIm               bool          `json: "is_im"`
		IsMember           bool          `json: "is_member"`
		IsMpim             bool          `json: "is_mpim"`
		IsOrgShared        bool          `json: "is_org_shared"`
		IsPendingExtShared bool          `json: "is_pending_ext_shared"`
		IsPrivate          bool          `json: "is_private"`
		IsReadOnly         bool          `json: "is_read_only"`
		IsShared           bool          `json: "is_shared"`
		LastRead           string        `json: "last_read"`
		Locale             string        `json: "locale"`
		Name               string        `json: "name"`
		NameNormalized     string        `json: "name_normalized"`
		NumMembers         float64       `json: "num_members"`
		PendingShared      []interface{} `json: "pending_shared"`
		PreviousNames      []string      `json: "previous_names"`
		Purpose            struct {
			Creator string  `json: "creator"`
			LastSet float64 `json: "last_set"`
			Value   string  `json: "value"`
		} `json: "purpose"`
		Topic struct {
			Creator string  `json: "creator"`
			LastSet float64 `json: "last_set"`
			Value   string  `json: "value"`
		} `json: "topic"`
		Unlinked float64 `json: "unlinked"`
	} `json: "channel"`
	Ok bool `json: "ok"`
}

type ChannelLeft struct {
	Channel string `json: "channel"`
	Type    string `json: "type"`
}

type GroupJoined struct {
	Channel map[string]interface{} `json: "channel"`
	Type    string                 `json: "type"`
}

type MemberLeftChannel struct {
	Channel     string `json: "channel"`
	ChannelType string `json: "channel_type"`
	Team        string `json: "team"`
	Type        string `json: "type"`
	User        string `json: "user"`
}

type ReconnectUrl struct {
	Type string `json: "type"`
}

type ScopeGranted struct {
	ApiAppId    string        `json: "api_app_id"`
	AuthedTeams []interface{} `json: "authed_teams"`
	Event       struct {
		Scopes    []string `json: "scopes"`
		TriggerId string   `json: "trigger_id"`
		Type      string   `json: "type"`
	} `json: "event"`
	EventId   string  `json: "event_id"`
	EventTime float64 `json: "event_time"`
	TeamId    string  `json: "team_id"`
	Token     string  `json: "token"`
	Type      string  `json: "type"`
}

type SubteamCreated struct {
	Subteam struct {
		AutoType    interface{} `json: "auto_type"`
		CreatedBy   string      `json: "created_by"`
		DateCreate  float64     `json: "date_create"`
		DateDelete  float64     `json: "date_delete"`
		DateUpdate  float64     `json: "date_update"`
		DeletedBy   interface{} `json: "deleted_by"`
		Description string      `json: "description"`
		Handle      string      `json: "handle"`
		Id          string      `json: "id"`
		IsExternal  bool        `json: "is_external"`
		IsUsergroup bool        `json: "is_usergroup"`
		Name        string      `json: "name"`
		Prefs       struct {
			Channels []interface{} `json: "channels"`
			Groups   []interface{} `json: "groups"`
		} `json: "prefs"`
		TeamId    string `json: "team_id"`
		UpdatedBy string `json: "updated_by"`
		UserCount string `json: "user_count"`
	} `json: "subteam"`
	Type string `json: "type"`
}

type TeamPlanChange struct {
	CanAddUra    bool     `json: "can_add_ura"`
	PaidFeatures []string `json: "paid_features"`
	Plan         string   `json: "plan"`
	Type         string   `json: "type"`
}

type ManualPresenceChange struct {
	Presence string `json: "presence"`
	Type     string `json: "type"`
}

type ReactionAdded struct {
	EventTs string `json: "event_ts"`
	Item    struct {
		Channel string `json: "channel"`
		Ts      string `json: "ts"`
		Type    string `json: "type"`
	} `json: "item"`
	ItemUser string `json: "item_user"`
	Reaction string `json: "reaction"`
	Type     string `json: "type"`
	User     string `json: "user"`
}

type SubteamMembersChanged struct {
	AddedUsers         []string `json: "added_users"`
	AddedUsersCount    string   `json: "added_users_count"`
	DatePreviousUpdate float64  `json: "date_previous_update"`
	DateUpdate         float64  `json: "date_update"`
	RemovedUsers       []string `json: "removed_users"`
	RemovedUsersCount  string   `json: "removed_users_count"`
	SubteamId          string   `json: "subteam_id"`
	TeamId             string   `json: "team_id"`
	Type               string   `json: "type"`
}

type TeamProfileReorder struct {
	Profile struct {
		Fields []struct {
			Id       string  `json: "id"`
			Ordering float64 `json: "ordering"`
		} `json: "fields"`
	} `json: "profile"`
	Type string `json: "type"`
}

type MessageAppHome struct {
	ApiAppId    string   `json: "api_app_id"`
	AuthedTeams []string `json: "authed_teams"`
	Event       struct {
		Channel     string `json: "channel"`
		ChannelType string `json: "channel_type"`
		EventTs     string `json: "event_ts"`
		Text        string `json: "text"`
		Ts          string `json: "ts"`
		Type        string `json: "type"`
		User        string `json: "user"`
	} `json: "event"`
	EventId   string  `json: "event_id"`
	EventTime float64 `json: "event_time"`
	TeamId    string  `json: "team_id"`
	Token     string  `json: "token"`
	Type      string  `json: "type"`
}

type AccountsChanged struct {
	Type string `json: "type"`
}

type AppMention struct {
	Channel string `json: "channel"`
	EventTs string `json: "event_ts"`
	Text    string `json: "text"`
	Ts      string `json: "ts"`
	Type    string `json: "type"`
	User    string `json: "user"`
}

type DndUpdatedUser struct {
	DndStatus struct {
		DndEnabled     bool    `json: "dnd_enabled"`
		NextDndEndTs   float64 `json: "next_dnd_end_ts"`
		NextDndStartTs float64 `json: "next_dnd_start_ts"`
	} `json: "dnd_status"`
	Type string `json: "type"`
	User string `json: "user"`
}

type EmojiChanged struct {
	EventTs string   `json: "event_ts"`
	Names   []string `json: "names"`
	Subtype string   `json: "subtype"`
	Type    string   `json: "type"`
}

type FileUnshared struct {
	File struct {
		Id string `json: "id"`
	} `json: "file"`
	FileId string `json: "file_id"`
	Type   string `json: "type"`
}

type GroupUnarchive struct {
	Channel string `json: "channel"`
	Type    string `json: "type"`
}

type ImHistoryChanged struct {
	EventTs string `json: "event_ts"`
	Latest  string `json: "latest"`
	Ts      string `json: "ts"`
	Type    string `json: "type"`
}

type UrlVerification struct {
	Challenge string `json: "challenge"`
	Token     string `json: "token"`
	Type      string `json: "type"`
}

type Usergroup struct {
	AutoType    string      `json: "auto_type"`
	CreatedBy   string      `json: "created_by"`
	DateCreate  float64     `json: "date_create"`
	DateDelete  float64     `json: "date_delete"`
	DateUpdate  float64     `json: "date_update"`
	DeletedBy   interface{} `json: "deleted_by"`
	Description string      `json: "description"`
	Handle      string      `json: "handle"`
	Id          string      `json: "id"`
	IsExternal  bool        `json: "is_external"`
	IsUsergroup bool        `json: "is_usergroup"`
	Name        string      `json: "name"`
	Prefs       struct {
		Channels []interface{} `json: "channels"`
		Groups   []interface{} `json: "groups"`
	} `json: "prefs"`
	TeamId    string   `json: "team_id"`
	UpdatedBy string   `json: "updated_by"`
	UserCount string   `json: "user_count"`
	Users     []string `json: "users"`
}

type AppRateLimited struct {
	ApiAppId          string  `json: "api_app_id"`
	MinuteRateLimited float64 `json: "minute_rate_limited"`
	TeamId            string  `json: "team_id"`
	Token             string  `json: "token"`
	Type              string  `json: "type"`
}

type ChannelDeleted struct {
	Channel string `json: "channel"`
	Type    string `json: "type"`
}

type PresenceSub struct {
	Ids  []string `json: "ids"`
	Type string   `json: "type"`
}

type UserResourceRemoved struct {
	ApiAppId    string        `json: "api_app_id"`
	AuthedTeams []interface{} `json: "authed_teams"`
	Event       struct {
		TriggerId string `json: "trigger_id"`
		Type      string `json: "type"`
		User      string `json: "user"`
	} `json: "event"`
	EventId   string  `json: "event_id"`
	EventTime float64 `json: "event_time"`
	TeamId    string  `json: "team_id"`
	Token     string  `json: "token"`
	Type      string  `json: "type"`
}

type Event struct {
	ApiAppId    string   `json: "api_app_id"`
	AuthedUsers []string `json: "authed_users"`
	Event       struct {
		EventTs string `json: "event_ts"`
		Type    string `json: "type"`
		User    string `json: "user"`
	} `json: "event"`
	EventId   string  `json: "event_id"`
	EventTime float64 `json: "event_time"`
	TeamId    string  `json: "team_id"`
	Token     string  `json: "token"`
	Type      string  `json: "type"`
}

type ChannelArchive struct {
	Channel string `json: "channel"`
	Type    string `json: "type"`
	User    string `json: "user"`
}

type ChannelMarked struct {
	Channel string `json: "channel"`
	Ts      string `json: "ts"`
	Type    string `json: "type"`
}

type GroupClose struct {
	Channel string `json: "channel"`
	Type    string `json: "type"`
	User    string `json: "user"`
}

type GroupOpen struct {
	Channel string `json: "channel"`
	Type    string `json: "type"`
	User    string `json: "user"`
}

type TeamPrefChange struct {
	Name  string `json: "name"`
	Type  string `json: "type"`
	Value bool   `json: "value"`
}

type DndUpdated struct {
	DndStatus struct {
		DndEnabled     bool    `json: "dnd_enabled"`
		NextDndEndTs   float64 `json: "next_dnd_end_ts"`
		NextDndStartTs float64 `json: "next_dnd_start_ts"`
		SnoozeEnabled  bool    `json: "snooze_enabled"`
		SnoozeEndtime  float64 `json: "snooze_endtime"`
	} `json: "dnd_status"`
	Type string `json: "type"`
	User string `json: "user"`
}

type TeamJoin struct {
	Type string                 `json: "type"`
	User map[string]interface{} `json: "user"`
}

type FileChange struct {
	File struct {
		Id string `json: "id"`
	} `json: "file"`
	FileId string `json: "file_id"`
	Type   string `json: "type"`
}

type PrefChange struct {
	Name  string `json: "name"`
	Type  string `json: "type"`
	Value string `json: "value"`
}

type ResourcesAdded struct {
	ApiAppId    string        `json: "api_app_id"`
	AuthedTeams []interface{} `json: "authed_teams"`
	Event       struct {
		Resources []struct {
			Resource struct {
				Grant struct {
					ResourceId string `json: "resource_id"`
					Type       string `json: "type"`
				} `json: "grant"`
				Type string `json: "type"`
			} `json: "resource"`
			Scopes []string `json: "scopes"`
		} `json: "resources"`
		Type string `json: "type"`
	} `json: "event"`
	EventId   string  `json: "event_id"`
	EventTime float64 `json: "event_time"`
	TeamId    string  `json: "team_id"`
	Token     string  `json: "token"`
	Type      string  `json: "type"`
}

type TeamProfileChange struct {
	Profile struct {
		Fields []map[string]interface{} `json: "fields"`
	} `json: "profile"`
	Type string `json: "type"`
}

type ChannelUnarchive struct {
	Channel string `json: "channel"`
	Type    string `json: "type"`
	User    string `json: "user"`
}

type FileCreated struct {
	File struct {
		Id string `json: "id"`
	} `json: "file"`
	FileId string `json: "file_id"`
	Type   string `json: "type"`
}

type FileShared struct {
	File struct {
		Id string `json: "id"`
	} `json: "file"`
	FileId string `json: "file_id"`
	Type   string `json: "type"`
}

type ImClose struct {
	Channel string `json: "channel"`
	Type    string `json: "type"`
	User    string `json: "user"`
}

type ImCreated struct {
	Channel map[string]interface{} `json: "channel"`
	Type    string                 `json: "type"`
	User    string                 `json: "user"`
}

type StarRemoved struct {
	EventTs string                 `json: "event_ts"`
	Item    map[string]interface{} `json: "item"`
	Type    string                 `json: "type"`
	User    string                 `json: "user"`
}

type TeamMigrationStarted struct {
	Type string `json: "type"`
}

type UserChange struct {
	Type string                 `json: "type"`
	User map[string]interface{} `json: "user"`
}

type UserResourceDenied struct {
	ApiAppId    string        `json: "api_app_id"`
	AuthedTeams []interface{} `json: "authed_teams"`
	Event       struct {
		Scopes    []string `json: "scopes"`
		TriggerId string   `json: "trigger_id"`
		Type      string   `json: "type"`
		User      string   `json: "user"`
	} `json: "event"`
	EventId   string  `json: "event_id"`
	EventTime float64 `json: "event_time"`
	TeamId    string  `json: "team_id"`
	Token     string  `json: "token"`
	Type      string  `json: "type"`
}

type FileCommentEdited struct {
	Comment map[string]interface{} `json: "comment"`
	File    struct {
		Id string `json: "id"`
	} `json: "file"`
	FileId string `json: "file_id"`
	Type   string `json: "type"`
}

type FileDeleted struct {
	EventTs string `json: "event_ts"`
	FileId  string `json: "file_id"`
	Type    string `json: "type"`
}

type GroupMarked struct {
	Channel string `json: "channel"`
	Ts      string `json: "ts"`
	Type    string `json: "type"`
}

type TeamRename struct {
	Name string `json: "name"`
	Type string `json: "type"`
}

type Im struct {
	Created       float64 `json: "created"`
	Id            string  `json: "id"`
	IsIm          bool    `json: "is_im"`
	IsUserDeleted bool    `json: "is_user_deleted"`
	User          string  `json: "user"`
}

type Goodbye struct {
	Type string `json: "type"`
}

type MessageGroups struct {
	ApiAppId    string   `json: "api_app_id"`
	AuthedTeams []string `json: "authed_teams"`
	Event       struct {
		Channel     string `json: "channel"`
		ChannelType string `json: "channel_type"`
		EventTs     string `json: "event_ts"`
		Text        string `json: "text"`
		Ts          string `json: "ts"`
		Type        string `json: "type"`
		User        string `json: "user"`
	} `json: "event"`
	EventId   string  `json: "event_id"`
	EventTime float64 `json: "event_time"`
	TeamId    string  `json: "team_id"`
	Token     string  `json: "token"`
	Type      string  `json: "type"`
}

type PinRemoved struct {
	ChannelId string                 `json: "channel_id"`
	EventTs   string                 `json: "event_ts"`
	HasPins   bool                   `json: "has_pins"`
	Item      map[string]interface{} `json: "item"`
	Type      string                 `json: "type"`
	User      string                 `json: "user"`
}

type TokensRevoked struct {
	ApiAppId string `json: "api_app_id"`
	Event    struct {
		Tokens struct {
			Bot   []string `json: "bot"`
			Oauth []string `json: "oauth"`
		} `json: "tokens"`
		Type string `json: "type"`
	} `json: "event"`
	EventId   string  `json: "event_id"`
	EventTime float64 `json: "event_time"`
	TeamId    string  `json: "team_id"`
	Token     string  `json: "token"`
	Type      string  `json: "type"`
}

type UserResourceGranted struct {
	ApiAppId    string        `json: "api_app_id"`
	AuthedTeams []interface{} `json: "authed_teams"`
	Event       struct {
		Scopes    []string `json: "scopes"`
		TriggerId string   `json: "trigger_id"`
		Type      string   `json: "type"`
		User      string   `json: "user"`
	} `json: "event"`
	EventId   string  `json: "event_id"`
	EventTime float64 `json: "event_time"`
	TeamId    string  `json: "team_id"`
	Token     string  `json: "token"`
	Type      string  `json: "type"`
}

type Channel struct {
	Channel struct {
		Created     float64 `json: "created"`
		Creator     string  `json: "creator"`
		Id          string  `json: "id"`
		IsArchived  bool    `json: "is_archived"`
		IsChannel   bool    `json: "is_channel"`
		IsGeneral   bool    `json: "is_general"`
		IsMember    bool    `json: "is_member"`
		IsMpim      bool    `json: "is_mpim"`
		IsOrgShared bool    `json: "is_org_shared"`
		IsPrivate   bool    `json: "is_private"`
		IsShared    bool    `json: "is_shared"`
		LastRead    string  `json: "last_read"`
		Latest      struct {
			Attachments []struct {
				Fallback string  `json: "fallback"`
				Id       float64 `json: "id"`
				Text     string  `json: "text"`
			} `json: "attachments"`
			BotId    string `json: "bot_id"`
			Subtype  string `json: "subtype"`
			Text     string `json: "text"`
			Ts       string `json: "ts"`
			Type     string `json: "type"`
			Username string `json: "username"`
		} `json: "latest"`
		Members        []string `json: "members"`
		Name           string   `json: "name"`
		NameNormalized string   `json: "name_normalized"`
		PreviousNames  []string `json: "previous_names"`
		Purpose        struct {
			Creator string  `json: "creator"`
			LastSet float64 `json: "last_set"`
			Value   string  `json: "value"`
		} `json: "purpose"`
		Topic struct {
			Creator string  `json: "creator"`
			LastSet float64 `json: "last_set"`
			Value   string  `json: "value"`
		} `json: "topic"`
		UnreadCount        float64 `json: "unread_count"`
		UnreadCountDisplay float64 `json: "unread_count_display"`
	} `json: "channel"`
	Ok bool `json: "ok"`
}

type File struct {
	Channels          []string      `json: "channels"`
	CommentsCount     float64       `json: "comments_count"`
	Created           float64       `json: "created"`
	DeanimateGif      string        `json: "deanimate_gif"`
	DisplayAsBot      bool          `json: "display_as_bot"`
	Editable          bool          `json: "editable"`
	ExternalType      string        `json: "external_type"`
	Filetype          string        `json: "filetype"`
	Groups            []interface{} `json: "groups"`
	HasRichPreview    bool          `json: "has_rich_preview"`
	Id                string        `json: "id"`
	ImageExifRotation float64       `json: "image_exif_rotation"`
	Ims               []interface{} `json: "ims"`
	IsExternal        bool          `json: "is_external"`
	IsPublic          bool          `json: "is_public"`
	IsStarred         bool          `json: "is_starred"`
	Mimetype          string        `json: "mimetype"`
	Mode              string        `json: "mode"`
	Name              string        `json: "name"`
	OriginalH         float64       `json: "original_h"`
	OriginalW         float64       `json: "original_w"`
	Permalink         string        `json: "permalink"`
	PermalinkPublic   string        `json: "permalink_public"`
	Pjpeg             string        `json: "pjpeg"`
	PrettyType        string        `json: "pretty_type"`
	PublicUrlShared   bool          `json: "public_url_shared"`
	Shares            struct {
		Public struct {
			C0T8SE4AU []struct {
				ChannelName     string   `json: "channel_name"`
				LatestReply     string   `json: "latest_reply"`
				ReplyCount      float64  `json: "reply_count"`
				ReplyUsers      []string `json: "reply_users"`
				ReplyUsersCount float64  `json: "reply_users_count"`
				TeamId          string   `json: "team_id"`
				ThreadTs        string   `json: "thread_ts"`
				Ts              string   `json: "ts"`
			} `json: "C0T8SE4AU"`
		} `json: "public"`
	} `json: "shares"`
	Size               float64 `json: "size"`
	Thumb160           string  `json: "thumb_160"`
	Thumb360           string  `json: "thumb_360"`
	Thumb360Gif        string  `json: "thumb_360_gif"`
	Thumb360H          float64 `json: "thumb_360_h"`
	Thumb360W          float64 `json: "thumb_360_w"`
	Thumb64            string  `json: "thumb_64"`
	Thumb80            string  `json: "thumb_80"`
	Timestamp          float64 `json: "timestamp"`
	Title              string  `json: "title"`
	UrlPrivate         string  `json: "url_private"`
	UrlPrivateDownload string  `json: "url_private_download"`
	User               string  `json: "user"`
	Username           string  `json: "username"`
}

type ResourcesRemoved struct {
	ApiAppId    string        `json: "api_app_id"`
	AuthedTeams []interface{} `json: "authed_teams"`
	Event       struct {
		Resources []struct {
			Resource struct {
				Grant struct {
					ResourceId string `json: "resource_id"`
					Type       string `json: "type"`
				} `json: "grant"`
				Type string `json: "type"`
			} `json: "resource"`
			Scopes []string `json: "scopes"`
		} `json: "resources"`
		Type string `json: "type"`
	} `json: "event"`
	EventId   string  `json: "event_id"`
	EventTime float64 `json: "event_time"`
	TeamId    string  `json: "team_id"`
	Token     string  `json: "token"`
	Type      string  `json: "type"`
}

type AppUninstalled struct {
	ApiAppId string `json: "api_app_id"`
	Event    struct {
		Type string `json: "type"`
	} `json: "event"`
	EventId   string  `json: "event_id"`
	EventTime float64 `json: "event_time"`
	TeamId    string  `json: "team_id"`
	Token     string  `json: "token"`
	Type      string  `json: "type"`
}

type EmailDomainChanged struct {
	EmailDomain string `json: "email_domain"`
	EventTs     string `json: "event_ts"`
	Type        string `json: "type"`
}

type GridMigrationStarted struct {
	ApiAppId string `json: "api_app_id"`
	Event    struct {
		EnterpriseId string `json: "enterprise_id"`
		Type         string `json: "type"`
	} `json: "event"`
	EventId   string  `json: "event_id"`
	EventTime float64 `json: "event_time"`
	TeamId    string  `json: "team_id"`
	Token     string  `json: "token"`
	Type      string  `json: "type"`
}

type GroupHistoryChanged struct {
	EventTs string `json: "event_ts"`
	Latest  string `json: "latest"`
	Ts      string `json: "ts"`
	Type    string `json: "type"`
}

type MemberJoinedChannel struct {
	Channel     string `json: "channel"`
	ChannelType string `json: "channel_type"`
	Inviter     string `json: "inviter"`
	Team        string `json: "team"`
	Type        string `json: "type"`
	User        string `json: "user"`
}

type MessageChannels struct {
	ApiAppId    string   `json: "api_app_id"`
	AuthedTeams []string `json: "authed_teams"`
	Event       struct {
		Channel     string `json: "channel"`
		ChannelType string `json: "channel_type"`
		EventTs     string `json: "event_ts"`
		Text        string `json: "text"`
		Ts          string `json: "ts"`
		Type        string `json: "type"`
		User        string `json: "user"`
	} `json: "event"`
	EventId   string  `json: "event_id"`
	EventTime float64 `json: "event_time"`
	TeamId    string  `json: "team_id"`
	Token     string  `json: "token"`
	Type      string  `json: "type"`
}

type PresenceChange struct {
	Presence string `json: "presence"`
	Type     string `json: "type"`
	User     string `json: "user"`
}

type StarAdded struct {
	EventTs string                 `json: "event_ts"`
	Item    map[string]interface{} `json: "item"`
	Type    string                 `json: "type"`
	User    string                 `json: "user"`
}

type Mpim struct {
	Created            float64                `json: "created"`
	Creator            string                 `json: "creator"`
	Id                 string                 `json: "id"`
	IsGroup            bool                   `json: "is_group"`
	IsMpim             bool                   `json: "is_mpim"`
	LastRead           string                 `json: "last_read"`
	Latest             map[string]interface{} `json: "latest"`
	Members            []string               `json: "members"`
	Name               string                 `json: "name"`
	UnreadCount        float64                `json: "unread_count"`
	UnreadCountDisplay float64                `json: "unread_count_display"`
}

type User struct {
	Ok   bool `json: "ok"`
	User struct {
		Color             string `json: "color"`
		Deleted           bool   `json: "deleted"`
		Has2fa            bool   `json: "has_2fa"`
		Id                string `json: "id"`
		IsAdmin           bool   `json: "is_admin"`
		IsAppUser         bool   `json: "is_app_user"`
		IsBot             bool   `json: "is_bot"`
		IsOwner           bool   `json: "is_owner"`
		IsPrimaryOwner    bool   `json: "is_primary_owner"`
		IsRestricted      bool   `json: "is_restricted"`
		IsStranger        bool   `json: "is_stranger"`
		IsUltraRestricted bool   `json: "is_ultra_restricted"`
		Locale            string `json: "locale"`
		Name              string `json: "name"`
		Profile           struct {
			AvatarHash            string  `json: "avatar_hash"`
			DisplayName           string  `json: "display_name"`
			DisplayNameNormalized string  `json: "display_name_normalized"`
			Email                 string  `json: "email"`
			Image192              string  `json: "image_192"`
			Image24               string  `json: "image_24"`
			Image32               string  `json: "image_32"`
			Image48               string  `json: "image_48"`
			Image512              string  `json: "image_512"`
			Image72               string  `json: "image_72"`
			ImageOriginal         string  `json: "image_original"`
			RealName              string  `json: "real_name"`
			RealNameNormalized    string  `json: "real_name_normalized"`
			StatusEmoji           string  `json: "status_emoji"`
			StatusExpiration      float64 `json: "status_expiration"`
			StatusText            string  `json: "status_text"`
			Team                  string  `json: "team"`
		} `json: "profile"`
		RealName string  `json: "real_name"`
		TeamId   string  `json: "team_id"`
		Tz       string  `json: "tz"`
		TzLabel  string  `json: "tz_label"`
		TzOffset float64 `json: "tz_offset"`
		Updated  float64 `json: "updated"`
	} `json: "user"`
}

type Group struct {
	Created    float64                `json: "created"`
	Creator    string                 `json: "creator"`
	Id         string                 `json: "id"`
	IsArchived bool                   `json: "is_archived"`
	IsGroup    string                 `json: "is_group"`
	IsMpim     bool                   `json: "is_mpim"`
	LastRead   string                 `json: "last_read"`
	Latest     map[string]interface{} `json: "latest"`
	Members    []string               `json: "members"`
	Name       string                 `json: "name"`
	Purpose    struct {
		Creator string  `json: "creator"`
		LastSet float64 `json: "last_set"`
		Value   string  `json: "value"`
	} `json: "purpose"`
	Topic struct {
		Creator string  `json: "creator"`
		LastSet float64 `json: "last_set"`
		Value   string  `json: "value"`
	} `json: "topic"`
	UnreadCount        float64 `json: "unread_count"`
	UnreadCountDisplay float64 `json: "unread_count_display"`
}

type BotChanged struct {
	Bot struct {
		AppId string `json: "app_id"`
		Icons struct {
			Image48 string `json: "image_48"`
		} `json: "icons"`
		Id   string `json: "id"`
		Name string `json: "name"`
	} `json: "bot"`
	Type string `json: "type"`
}

type ChannelHistoryChanged struct {
	EventTs string `json: "event_ts"`
	Latest  string `json: "latest"`
	Ts      string `json: "ts"`
	Type    string `json: "type"`
}

type CommandsChanged struct {
	EventTs string `json: "event_ts"`
	Type    string `json: "type"`
}

type Message struct {
	Channel string `json: "channel"`
	Text    string `json: "text"`
	Ts      string `json: "ts"`
	Type    string `json: "type"`
	User    string `json: "user"`
}

type MessageIm struct {
	ApiAppId    string   `json: "api_app_id"`
	AuthedTeams []string `json: "authed_teams"`
	Event       struct {
		Channel     string `json: "channel"`
		ChannelType string `json: "channel_type"`
		EventTs     string `json: "event_ts"`
		Text        string `json: "text"`
		Ts          string `json: "ts"`
		Type        string `json: "type"`
		User        string `json: "user"`
	} `json: "event"`
	EventId   string  `json: "event_id"`
	EventTime float64 `json: "event_time"`
	TeamId    string  `json: "team_id"`
	Token     string  `json: "token"`
	Type      string  `json: "type"`
}

type PresenceQuery struct {
	Ids  []string `json: "ids"`
	Type string   `json: "type"`
}

type UserTyping struct {
	Channel string `json: "channel"`
	Type    string `json: "type"`
	User    string `json: "user"`
}
