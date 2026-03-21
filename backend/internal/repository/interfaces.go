package repository

import (
	"context"
	"time"

	"github.com/jym/lincle/internal/domain"
)

// ─── Auth ────────────────────────────────────────────

// AuthRepo handles user accounts, profiles, and refresh tokens.
type AuthRepo interface {
	// FindUserByProvider looks up an existing user by login provider and key.
	// Returns (nil, nil) when no user matches.
	FindUserByProvider(ctx context.Context, provider, providerKey string) (*UserWithNickname, error)

	// CreateUserWithProfile inserts a new user row and a default profile in a single transaction.
	CreateUserWithProfile(ctx context.Context, userID, provider, providerKey, nickname string) error

	// UpdateLastLogin sets last_login_at to now for the given user.
	UpdateLastLogin(ctx context.Context, userID string, at time.Time) error

	// StoreRefreshToken persists a hashed refresh token.
	StoreRefreshToken(ctx context.Context, id, userID, tokenHash string, expiresAt time.Time) error

	// FindRefreshToken returns the token row ID if a valid (non-expired) token hash exists.
	// Returns ("", nil) when no matching token is found.
	FindRefreshToken(ctx context.Context, tokenHash string) (string, error)

	// GetAccountStatus returns the account_status for the given user.
	// Returns ("", nil) when the user is not found.
	GetAccountStatus(ctx context.Context, userID string) (string, error)

	// RotateRefreshToken deletes the old token and inserts a new one atomically.
	RotateRefreshToken(ctx context.Context, oldTokenID, newTokenID, userID, newTokenHash string, expiresAt time.Time) error

	// DeleteRefreshTokensByUser removes all refresh tokens for the user (logout).
	DeleteRefreshTokensByUser(ctx context.Context, userID string) error

	// GetUserProfile returns the full user + profile data needed by GET /me.
	// Returns (nil, nil) when the user is not found.
	GetUserProfile(ctx context.Context, userID string) (*FullUserProfile, error)

	// UpdateProfile dynamically updates the specified profile fields.
	UpdateProfile(ctx context.Context, userID string, fields ProfileUpdateFields) error
}

// UserWithNickname is the result of the login lookup query.
type UserWithNickname struct {
	UserID   string
	Role     string
	Nickname string
}

// FullUserProfile combines users + user_profiles for the GET /me endpoint.
type FullUserProfile struct {
	UserID          string
	Role            string
	AccountStatus   string
	Nickname        string
	AvatarURL       *string
	Introduction    *string
	PrimaryServerID *string
	TradeCount      int
	ReviewCount     int
	ResponseBadge   string
	TrustBadge      string
	AlignmentScore  int
	AlignmentGrade  string
}

// ProfileUpdateFields carries the optional fields for profile updates.
// nil means "don't change this field".
type ProfileUpdateFields struct {
	Nickname      *string
	Introduction  *string
	PrimaryServer *string
	AvatarURL     *string
}

// ─── Listing ─────────────────────────────────────────

// ListingRepo handles listings, favorites, and status history.
type ListingRepo interface {
	// CheckImageOwnership verifies an uploaded image belongs to the given user.
	// Returns false when the image does not exist or belongs to someone else.
	CheckImageOwnership(ctx context.Context, imageID, userID string) (bool, error)

	// InsertListing creates a new listing row.
	InsertListing(ctx context.Context, listing *InsertListingParams) error

	// InsertStatusHistory records a status transition event.
	InsertStatusHistory(ctx context.Context, params *InsertStatusHistoryParams) error

	// ListListings returns a paginated, filtered list of listings.
	ListListings(ctx context.Context, filter ListingFilter) ([]ListingListItem, error)

	// GetListing returns a single listing with all joined data (server, category, author, icon).
	// Returns (nil, nil) when the listing is not found.
	GetListing(ctx context.Context, listingID string) (*ListingDetail, error)

	// IncrementViewCount bumps the view counter for a listing.
	IncrementViewCount(ctx context.Context, listingID string) error

	// IsFavorited checks if the user has favorited the listing.
	IsFavorited(ctx context.Context, userID, listingID string) (bool, error)

	// GetListingOwnerAndStatus returns (authorUserID, status) for ownership/guard checks.
	// Returns (nil, nil) when the listing is not found.
	GetListingOwnerAndStatus(ctx context.Context, listingID string) (*ListingOwnerStatus, error)

	// UpdateListing dynamically updates the specified listing fields.
	UpdateListing(ctx context.Context, listingID string, fields ListingUpdateFields) error

	// UpdateListingStatus sets the status, updated_at, and last_activity_at.
	UpdateListingStatus(ctx context.Context, listingID string, status domain.ListingStatus, now time.Time) error

	// ListingExists returns true if the listing exists and is not soft-deleted.
	ListingExists(ctx context.Context, listingID string) (bool, error)

	// AddFavorite inserts a favorite (idempotent, ON CONFLICT DO NOTHING).
	AddFavorite(ctx context.Context, id, userID, listingID string) error

	// RemoveFavorite deletes a favorite by user+listing.
	RemoveFavorite(ctx context.Context, userID, listingID string) error

	// ListMyListings returns the caller's own listings, optionally filtered by status.
	ListMyListings(ctx context.Context, userID string, status *string) ([]MyListingItem, error)
}

// InsertListingParams carries all columns for a new listing INSERT.
type InsertListingParams struct {
	ID              string
	ListingType     string
	AuthorUserID    string
	ServerID        string
	CategoryID      string
	ItemName        string
	Title           string
	Description     string
	PriceType       string
	PriceAmount     *int64
	Quantity        int
	Enhancement     *int
	OptionsText     *string
	TradeMethod     string
	MeetingArea     *string
	TimeText        *string
	Now             time.Time
}

// InsertStatusHistoryParams carries columns for a status_history INSERT.
type InsertStatusHistoryParams struct {
	ID            string
	EntityType    string
	EntityID      string
	FromStatus    *string // nil for initial creation
	ToStatus      string
	ChangedByUser string
	ReasonCode    *string
	CreatedAt     time.Time
}

// ListingFilter captures all query parameters for the listing list endpoint.
type ListingFilter struct {
	Query       string
	ServerID    string
	CategoryID  string
	ListingType string
	Status      string
	Sort        string
	Limit       int
	Cursor      string // created_at cursor for pagination
}

// ListingListItem is the shape returned by ListListings.
type ListingListItem struct {
	ListingID       string
	ListingType     string
	Title           string
	ItemName        string
	PriceType       string
	PriceAmount     *int64
	EnhancementLvl  *int
	ServerID        string
	ServerName      string
	Status          string
	TradeMethod     string
	ViewCount       int
	FavoriteCount   int
	ChatCount       int
	LastActivityAt  time.Time
	CreatedAt       time.Time
	AuthorID        string
	AuthorNickname  string
	TrustBadge      string
	ResponseBadge   string
	IconID          *string
}

// ListingDetail is the full result for a single listing page.
type ListingDetail struct {
	ID              string
	ListingType     string
	Title           string
	ItemName        string
	Description     string
	PriceType       string
	PriceAmount     *int64
	Quantity        int
	Enhancement     *int
	OptionsText     *string
	ServerID        string
	ServerName      string
	CategoryID      string
	CategoryName    string
	Status          string
	Visibility      string
	TradeMethod     string
	MeetingArea     *string
	TimeText        *string
	AuthorID        string
	AuthorNickname  string
	TrustBadge      string
	ResponseBadge   string
	TradeCount      int
	ViewCount       int
	FavoriteCount   int
	ChatCount       int
	ReservedChatID  *string
	LastActivityAt  time.Time
	CreatedAt       time.Time
	UpdatedAt       time.Time
	IconID          *string
}

// ListingOwnerStatus is used for ownership + guard checks.
type ListingOwnerStatus struct {
	AuthorUserID string
	Status       string
}

// ListingUpdateFields carries the optional fields for listing edits.
type ListingUpdateFields struct {
	Title       *string
	Description *string
	PriceType   *string
	PriceAmount *int64
	Quantity    *int
	Enhancement *int
	OptionsText *string
	TradeMethod *string
	MeetingArea *string
	TimeText    *string
}

// MyListingItem is the compact shape returned by ListMyListings.
type MyListingItem struct {
	ListingID       string
	ListingType     string
	Title           string
	ItemName        string
	PriceType       string
	PriceAmount     *int64
	EnhancementLvl  *int
	ServerName      string
	Status          string
	ViewCount       int
	FavoriteCount   int
	ChatCount       int
	CreatedAt       time.Time
	AuthorID        string
	AuthorNickname  string
	IconID          *string
}

// ─── Chat ────────────────────────────────────────────

// ChatRepo handles chat rooms, messages, read cursors, and block relations.
type ChatRepo interface {
	// GetListingAuthor returns the author_user_id of a listing.
	// Returns ("", nil) when the listing is not found.
	GetListingAuthor(ctx context.Context, listingID string) (string, error)

	// FindExistingChatRoom checks if a chat room already exists between two users for a listing.
	// Returns ("", nil) when no room exists.
	FindExistingChatRoom(ctx context.Context, listingID, sellerID, buyerID string) (string, error)

	// CreateChatRoom inserts a new chat_rooms row and increments the listing's chat_count.
	CreateChatRoom(ctx context.Context, params *CreateChatRoomParams) error

	// ListChatRooms returns all chat rooms for the user with counterpart info.
	ListChatRooms(ctx context.Context, userID string) ([]ChatRoomListItem, error)

	// IsChatParticipant returns true if the user is a seller or buyer in the chat room.
	IsChatParticipant(ctx context.Context, chatRoomID, userID string) (bool, error)

	// ListMessages returns paginated messages for a chat room.
	ListMessages(ctx context.Context, chatRoomID string, limit int, cursor string) ([]ChatMessageItem, error)

	// GetChatRoomParticipants returns (sellerID, buyerID) for the given chat room,
	// verifying the caller is a participant.
	// Returns (nil, nil) when the room is not found or user is not a participant.
	GetChatRoomParticipants(ctx context.Context, chatRoomID, userID string) (*ChatParticipants, error)

	// CheckDuplicateMessage returns true if a message with the given client_message_id exists.
	CheckDuplicateMessage(ctx context.Context, clientMessageID string) (bool, error)

	// InsertMessage creates a new chat_messages row and updates the room's last_message_at.
	InsertMessage(ctx context.Context, params *InsertMessageParams) error

	// UpsertReadCursor updates or inserts the user's read cursor for a chat room.
	UpsertReadCursor(ctx context.Context, chatRoomID, userID, lastReadMessageID string) error

	// BlockUser inserts a block_relations row (idempotent, ON CONFLICT DO NOTHING).
	BlockUser(ctx context.Context, id, blockerID, blockedID string) error

	// UnblockUser removes a block_relations row.
	UnblockUser(ctx context.Context, blockerID, blockedID string) error
}

// CreateChatRoomParams carries columns for a new chat room INSERT.
type CreateChatRoomParams struct {
	ID        string
	ListingID string
	SellerID  string
	BuyerID   string
	Now       time.Time
}

// ChatRoomListItem is the shape returned by ListChatRooms.
type ChatRoomListItem struct {
	ChatRoomID          string
	ListingID           string
	ListingTitle        string
	ChatStatus          string
	LastMessageAt       *time.Time
	UpdatedAt           *time.Time
	CounterpartID       string
	CounterpartNick     string
	CounterpartBadge    string
	ListingThumbnail    *string    // thumbnail URL from listing_images
	ListingStatus       string     // listing status (available, reserved, sold, deleted)
	ListingPrice        *int64     // listing price_amount
	ListingServerID     *string    // listing server_id
	ListingServerName   *string    // server name from servers table
	UnreadCount         int        // count of unread messages for current user
	LastMessageBody     *string    // body text of the most recent message
	LastMessageSentAt   *time.Time // sent_at of the most recent message
	MyLastReadMsgID     *string    // current user's last read message ID
}

// ChatParticipants holds the seller and buyer for a chat room.
type ChatParticipants struct {
	SellerID string
	BuyerID  string
}

// ChatMessageItem is the shape returned by ListMessages.
type ChatMessageItem struct {
	MessageID    string
	SenderUserID *string
	MessageType  string
	BodyText     *string
	MetadataJSON *string
	SentAt       time.Time
}

// InsertMessageParams carries columns for a new chat message INSERT.
type InsertMessageParams struct {
	ID              string
	ChatRoomID      string
	SenderUserID    string
	MessageType     string
	BodyText        *string
	ClientMessageID *string
	Now             time.Time
}

// ─── Reservation ─────────────────────────────────────

// ReservationRepo handles reservations, trade completions, reviews, and related transactions.
type ReservationRepo interface {
	// GetChatRoomForReservation returns listing_id, seller_id, buyer_id from the chat room,
	// verifying the caller is a participant.
	// Returns (nil, nil) when the room is not found or user is not a participant.
	GetChatRoomForReservation(ctx context.Context, chatRoomID, userID string) (*ChatRoomReservationInfo, error)

	// CountActiveReservations returns the number of proposed/confirmed reservations for a listing.
	CountActiveReservations(ctx context.Context, listingID string) (int, error)

	// CreateReservation inserts a reservation, a system chat message, and updates the
	// chat room status to 'reservation_proposed' -- all in a single transaction.
	CreateReservation(ctx context.Context, params *CreateReservationParams) error

	// GetReservationForConfirm returns the counterpart_user_id, listing_id, and chat_room_id
	// for a 'proposed' reservation.
	// Returns (nil, nil) when not found.
	GetReservationForConfirm(ctx context.Context, reservationID string) (*ReservationConfirmInfo, error)

	// ConfirmReservation sets the reservation to 'confirmed', updates listing to 'reserved',
	// updates chat_room status, and records status_history -- all in a single transaction.
	ConfirmReservation(ctx context.Context, params *ConfirmReservationParams) error

	// GetReservationForCancel returns listing_id, chat_room_id, proposer_user_id, counterpart_user_id
	// for a reservation in 'proposed' or 'confirmed' status.
	// Returns (nil, nil) when not found.
	GetReservationForCancel(ctx context.Context, reservationID string) (*ReservationCancelInfo, error)

	// CancelReservation sets the reservation to 'cancelled', reverts listing to 'available',
	// and updates chat_room status to 'open' -- all in a single transaction.
	CancelReservation(ctx context.Context, params *CancelReservationParams) error

	// GetConfirmedReservation returns proposer and counterpart for a confirmed reservation on a listing.
	// Returns (nil, nil) when not found.
	GetConfirmedReservation(ctx context.Context, reservationID, listingID string) (*ReservationParticipants, error)

	// CreateTradeCompletion inserts a trade_completions row.
	CreateTradeCompletion(ctx context.Context, params *CreateTradeCompletionParams) error

	// GetPendingCompletion returns the pending_confirm completion for confirmation.
	// Returns (nil, nil) when not found.
	GetPendingCompletion(ctx context.Context, completionID string) (*PendingCompletionInfo, error)

	// ConfirmCompletion confirms the trade: updates trade_completions, fulfills the reservation,
	// completes the listing, and adjusts alignment scores -- all in a single transaction.
	ConfirmCompletion(ctx context.Context, params *ConfirmCompletionParams) error

	// ListMyTrades returns the user's trade history (chat rooms with listing info).
	ListMyTrades(ctx context.Context, userID string) ([]MyTradeItem, error)

	// GetCompletionForReview returns status, requested_by, counterpart for a completion.
	// Returns (nil, nil) when not found.
	GetCompletionForReview(ctx context.Context, completionID string) (*CompletionReviewInfo, error)

	// CreateReview inserts a review, updates profile counters, and adjusts alignment scores
	// in a single transaction.
	CreateReview(ctx context.Context, params *CreateReviewParams) error

	// ListUserReviews returns reviews received by a target user.
	ListUserReviews(ctx context.Context, targetUserID string) ([]UserReviewItem, error)

	// CreateReport inserts a reports row.
	CreateReport(ctx context.Context, params *CreateReportParams) error

	// ListMyReports returns reports submitted by the user.
	ListMyReports(ctx context.Context, userID string) ([]MyReportItem, error)

	// ListNotifications returns the user's notifications.
	ListNotifications(ctx context.Context, userID string) ([]NotificationItem, error)

	// MarkNotificationsRead marks the specified notification IDs as read.
	MarkNotificationsRead(ctx context.Context, userID string, notificationIDs []string) error
}

// ChatRoomReservationInfo is the chat room data needed to create a reservation.
type ChatRoomReservationInfo struct {
	ListingID string
	SellerID  string
	BuyerID   string
}

// CreateReservationParams carries all data for the reservation creation transaction.
type CreateReservationParams struct {
	ReservationID    string
	ListingID        string
	ChatRoomID       string
	ProposerUserID   string
	CounterpartID    string
	ScheduledAt      string
	MeetingType      string
	ServerID         *string
	MeetingPoint     *string
	Note             *string
	ExpiresAt        *string
	SystemMessageID  string
	SystemMetaJSON   string
	Now              time.Time
}

// ReservationConfirmInfo is the data needed to confirm a reservation.
type ReservationConfirmInfo struct {
	CounterpartUserID string
	ListingID         string
	ChatRoomID        string
}

// ConfirmReservationParams carries all data for the reservation confirmation transaction.
type ConfirmReservationParams struct {
	ReservationID    string
	ListingID        string
	ChatRoomID       string
	UserID           string
	StatusHistoryID  string
	Now              time.Time
}

// ReservationCancelInfo is the data needed to cancel a reservation.
type ReservationCancelInfo struct {
	ListingID    string
	ChatRoomID   string
	ProposerID   string
	CounterpartID string
}

// CancelReservationParams carries all data for the reservation cancellation transaction.
type CancelReservationParams struct {
	ReservationID string
	ListingID     string
	ChatRoomID    string
	ReasonCode    string
	Now           time.Time
}

// ReservationParticipants holds proposer and counterpart for a confirmed reservation.
type ReservationParticipants struct {
	ProposerID    string
	CounterpartID string
}

// CreateTradeCompletionParams carries columns for a new trade_completions INSERT.
type CreateTradeCompletionParams struct {
	ID              string
	ListingID       string
	ReservationID   string
	RequestedByID   string
	CounterpartID   string
	CompletionNote  *string
	AutoConfirmAt   time.Time
	Now             time.Time
}

// PendingCompletionInfo is the data needed to confirm a trade completion.
type PendingCompletionInfo struct {
	CounterpartUserID string
	ListingID         string
	ReservationID     string
	RequestedByUserID string
}

// ConfirmCompletionParams carries all data for the trade confirmation transaction.
type ConfirmCompletionParams struct {
	CompletionID  string
	ReservationID string
	ListingID     string
	RequestedByID string
	CounterpartID string
	Now           time.Time
}

// MyTradeItem is the shape returned by ListMyTrades.
type MyTradeItem struct {
	ChatRoomID     string
	ListingID      string
	ListingTitle   string
	ListingStatus  string
	CounterpartID  string
	CounterpartNick string
	ChatStatus     string
	UpdatedAt      time.Time
}

// CompletionReviewInfo is the data needed to create a review.
type CompletionReviewInfo struct {
	Status            string
	RequestedByUserID string
	CounterpartUserID string
}

// CreateReviewParams carries all data for the review creation transaction.
type CreateReviewParams struct {
	ReviewID     string
	CompletionID string
	ReviewerID   string
	TargetUserID string
	Rating       string
	Comment      *string
}

// UserReviewItem is the shape returned by ListUserReviews.
type UserReviewItem struct {
	ReviewID         string
	Rating           string
	Comment          *string
	CreatedAt        time.Time
	ReviewerNickname string
}

// CreateReportParams carries columns for a new reports INSERT.
type CreateReportParams struct {
	ID          string
	ReporterID  string
	TargetType  string
	TargetID    string
	ReportType  string
	Description string
}

// MyReportItem is the shape returned by ListMyReports.
type MyReportItem struct {
	ReportID   string
	TargetType string
	TargetID   string
	ReportType string
	Status     string
	CreatedAt  time.Time
}

// NotificationItem is the shape returned by ListNotifications.
type NotificationItem struct {
	NotificationID string
	Type           string
	Title          string
	Body           string
	ReferenceType  *string
	ReferenceID    *string
	DeepLink       *string
	IsRead         bool
	CreatedAt      time.Time
}

// ─── Master Data ─────────────────────────────────────

// MasterRepo handles read-only reference data (servers, categories, item search).
type MasterRepo interface {
	// ListServers returns all active game servers.
	ListServers(ctx context.Context) ([]ServerItem, error)

	// ListCategories returns all item categories.
	ListCategories(ctx context.Context) ([]CategoryItem, error)

	// SearchItems searches the item_master table by name, optionally filtered by category.
	SearchItems(ctx context.Context, query string, categoryID *string) ([]ItemSearchResult, error)
}

// ServerItem is the shape returned by ListServers.
type ServerItem struct {
	ServerID   string
	ServerName string
}

// CategoryItem is the shape returned by ListCategories.
type CategoryItem struct {
	CategoryID   string
	CategoryName string
	ParentID     *string
}

// ItemSearchResult is the shape returned by SearchItems.
type ItemSearchResult struct {
	ID             string
	Name           string
	CategoryID     string
	IconID         *string
	SubCategory    string  // NOT NULL in DB
	OptionText     *string // nullable
	IsEnchantable  int     // 0 or 1, NOT NULL
	SafeEnchantLvl int     // NOT NULL DEFAULT 0
	MaxEnchantLvl  int     // NOT NULL DEFAULT 0
}

// ─── Upload ──────────────────────────────────────────

// UploadRepo handles image upload metadata.
type UploadRepo interface {
	// InsertImage records an uploaded image's metadata.
	InsertImage(ctx context.Context, params *InsertImageParams) error
}

// InsertImageParams carries columns for an uploaded_images INSERT.
type InsertImageParams struct {
	ID          string
	UserID      string
	Filename    string
	URL         string
	ContentType string
	SizeBytes   int64
}
