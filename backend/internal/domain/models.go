package domain

import "time"

// ─── Enums ───────────────────────────────────────────

type ListingType string

const (
	ListingTypeSell ListingType = "sell"
	ListingTypeBuy  ListingType = "buy"
)

type ListingStatus string

const (
	ListingAvailable    ListingStatus = "available"
	ListingReserved     ListingStatus = "reserved"
	ListingPendingTrade ListingStatus = "pending_trade"
	ListingCompleted    ListingStatus = "completed"
	ListingCancelled    ListingStatus = "cancelled"
)

type Visibility string

const (
	VisibilityPublic  Visibility = "public"
	VisibilityHidden  Visibility = "hidden"
	VisibilityBlocked Visibility = "blocked"
)

type PriceType string

const (
	PriceFixed      PriceType = "fixed"
	PriceNegotiable PriceType = "negotiable"
	PriceOffer      PriceType = "offer"
)

type TradeMethod string

const (
	TradeInGame    TradeMethod = "in_game"
	TradeOfflinePC TradeMethod = "offline_pc_bang"
	TradeEither    TradeMethod = "either"
)

type ChatStatus string

const (
	ChatOpen                 ChatStatus = "open"
	ChatReservationProposed  ChatStatus = "reservation_proposed"
	ChatReservationConfirmed ChatStatus = "reservation_confirmed"
	ChatTradeDue             ChatStatus = "trade_due"
	ChatDealCompleted        ChatStatus = "deal_completed"
	ChatDealCancelled        ChatStatus = "deal_cancelled"
	ChatReportLocked         ChatStatus = "report_locked"
)

type ReservationStatus string

const (
	ReservationProposed    ReservationStatus = "proposed"
	ReservationConfirmed   ReservationStatus = "confirmed"
	ReservationExpired     ReservationStatus = "expired"
	ReservationCancelled   ReservationStatus = "cancelled"
	ReservationFulfilled   ReservationStatus = "fulfilled"
	ReservationNoShowReported ReservationStatus = "no_show_reported"
)

type CompletionStatus string

const (
	CompletionPendingConfirm CompletionStatus = "pending_confirm"
	CompletionConfirmed      CompletionStatus = "confirmed"
	CompletionExpired        CompletionStatus = "expired"
	CompletionDisputed       CompletionStatus = "disputed"
)

// ─── Alignment (성향치) ─────────────────────────────

type AlignmentGrade string

const (
	GradeRoyalKnight AlignmentGrade = "royal_knight" // 100+
	GradeLawful      AlignmentGrade = "lawful"       // 50~99
	GradeNeutral     AlignmentGrade = "neutral"      // 0~49
	GradeCaution     AlignmentGrade = "caution"      // -1~-30
	GradeChaotic     AlignmentGrade = "chaotic"       // -31~
)

// Alignment score deltas
const (
	AlignmentTradeConfirmed   = 5   // 거래 완료 (양측 확인)
	AlignmentPositiveReview   = 3   // 긍정 후기 받음
	AlignmentReviewWritten    = 1   // 후기 작성 (작성자 인센티브)
	AlignmentNegativeReview   = -10 // 부정 후기 받음
	AlignmentTradeExpired     = -3  // 거래완료 미응답 (expired)
	AlignmentReportConfirmed  = -20 // 신고 처리 확정
)

// CalcAlignmentGrade returns the grade for a given score.
func CalcAlignmentGrade(score int) AlignmentGrade {
	switch {
	case score >= 100:
		return GradeRoyalKnight
	case score >= 50:
		return GradeLawful
	case score >= 0:
		return GradeNeutral
	case score >= -30:
		return GradeCaution
	default:
		return GradeChaotic
	}
}

type ReviewRating string

const (
	ReviewPositive ReviewRating = "positive"
	ReviewNegative ReviewRating = "negative"
)

type ReportType string

const (
	ReportFakeListing     ReportType = "fake_listing"
	ReportScamSuspicion   ReportType = "scam_suspicion"
	ReportNoShow          ReportType = "no_show"
	ReportHarassment      ReportType = "harassment"
	ReportSpam            ReportType = "spam"
	ReportProhibitedItem  ReportType = "prohibited_item"
	ReportPrivacyExposure ReportType = "privacy_exposure"
	ReportOther           ReportType = "other"
)

type ReportStatus string

const (
	ReportSubmitted     ReportStatus = "submitted"
	ReportAssigned      ReportStatus = "assigned"
	ReportInReview      ReportStatus = "in_review"
	ReportResolved      ReportStatus = "resolved"
	ReportDismissed     ReportStatus = "dismissed"
)

type AccountStatus string

const (
	AccountActive     AccountStatus = "active"
	AccountRestricted AccountStatus = "restricted"
	AccountSuspended  AccountStatus = "suspended"
	AccountWithdrawn  AccountStatus = "withdrawn"
)

type UserRole string

const (
	RoleUser      UserRole = "user"
	RoleModerator UserRole = "moderator"
	RoleAdmin     UserRole = "admin"
)

type MessageType string

const (
	MsgText            MessageType = "text"
	MsgSystem          MessageType = "system"
	MsgReservationCard MessageType = "reservation_card"
	MsgImage           MessageType = "image"
)

// ─── Entities ────────────────────────────────────────

type User struct {
	ID            string        `json:"userId" db:"id"`
	LoginProvider string        `json:"loginProvider" db:"login_provider"`
	ProviderKey   string        `json:"-" db:"login_provider_user_key"`
	AccountStatus AccountStatus `json:"accountStatus" db:"account_status"`
	Role          UserRole      `json:"role" db:"role"`
	LastLoginAt   *time.Time    `json:"lastLoginAt,omitempty" db:"last_login_at"`
	CreatedAt     time.Time     `json:"createdAt" db:"created_at"`
}

type UserProfile struct {
	UserID              string         `json:"userId" db:"user_id"`
	Nickname            string         `json:"nickname" db:"nickname"`
	AvatarURL           *string        `json:"avatarUrl,omitempty" db:"avatar_url"`
	Introduction        *string        `json:"introduction,omitempty" db:"introduction"`
	PrimaryServerID     *string        `json:"primaryServerId,omitempty" db:"primary_server_id"`
	CompletedTradeCount int            `json:"completedTradeCount" db:"completed_trade_count"`
	PositiveReviewCount int            `json:"positiveReviewCount" db:"positive_review_count"`
	ResponseBadge       string         `json:"responseBadge" db:"response_badge"`
	TrustBadge          string         `json:"trustBadge" db:"trust_badge"`
	AlignmentScore      int            `json:"alignmentScore" db:"alignment_score"`
	AlignmentGrade      AlignmentGrade `json:"alignmentGrade" db:"alignment_grade"`
	LastActiveAt        *time.Time     `json:"lastActiveAt,omitempty" db:"last_active_at"`
}

type Listing struct {
	ID                     string        `json:"listingId" db:"id"`
	ListingType            ListingType   `json:"listingType" db:"listing_type"`
	AuthorUserID           string        `json:"authorUserId" db:"author_user_id"`
	ServerID               string        `json:"serverId" db:"server_id"`
	CategoryID             string        `json:"categoryId" db:"category_id"`
	ItemName               string        `json:"itemName" db:"item_name"`
	Title                  string        `json:"title" db:"title"`
	Description            string        `json:"description" db:"description"`
	PriceType              PriceType     `json:"priceType" db:"price_type"`
	PriceAmount            *int64        `json:"priceAmount,omitempty" db:"price_amount"`
	Quantity               int           `json:"quantity" db:"quantity"`
	EnhancementLevel       *int          `json:"enhancementLevel,omitempty" db:"enhancement_level"`
	OptionsText            *string       `json:"optionsText,omitempty" db:"options_text"`
	TradeMethod            TradeMethod   `json:"tradeMethod" db:"trade_method"`
	PreferredMeetingArea   *string       `json:"preferredMeetingAreaText,omitempty" db:"preferred_meeting_area_text"`
	AvailableTimeText      *string       `json:"availableTimeText,omitempty" db:"available_time_text"`
	Status                 ListingStatus `json:"status" db:"status"`
	Visibility             Visibility    `json:"visibility" db:"visibility"`
	ReservedChatRoomID     *string       `json:"reservedChatRoomId,omitempty" db:"reserved_chat_room_id"`
	ViewCount              int           `json:"viewCount" db:"view_count"`
	FavoriteCount          int           `json:"favoriteCount" db:"favorite_count"`
	ChatCount              int           `json:"chatCount" db:"chat_count"`
	LastActivityAt         time.Time     `json:"lastActivityAt" db:"last_activity_at"`
	CreatedAt              time.Time     `json:"createdAt" db:"created_at"`
	UpdatedAt              time.Time     `json:"updatedAt" db:"updated_at"`
	DeletedAt              *time.Time    `json:"-" db:"deleted_at"`
}

type ChatRoom struct {
	ID           string     `json:"chatRoomId" db:"id"`
	ListingID    string     `json:"listingId" db:"listing_id"`
	SellerUserID string     `json:"sellerUserId" db:"seller_user_id"`
	BuyerUserID  string     `json:"buyerUserId" db:"buyer_user_id"`
	ChatStatus   ChatStatus `json:"chatStatus" db:"chat_status"`
	LastMessageAt *time.Time `json:"lastMessageAt,omitempty" db:"last_message_at"`
	CreatedAt    time.Time  `json:"createdAt" db:"created_at"`
	UpdatedAt    time.Time  `json:"updatedAt" db:"updated_at"`
}

type ChatMessage struct {
	ID           string      `json:"messageId" db:"id"`
	ChatRoomID   string      `json:"chatRoomId" db:"chat_room_id"`
	SenderUserID *string     `json:"senderUserId,omitempty" db:"sender_user_id"`
	MessageType  MessageType `json:"messageType" db:"message_type"`
	BodyText     *string     `json:"bodyText,omitempty" db:"body_text"`
	MetadataJSON *string     `json:"metadataJson,omitempty" db:"metadata_json"`
	SentAt       time.Time   `json:"sentAt" db:"sent_at"`
}

type Reservation struct {
	ID                string            `json:"reservationId" db:"id"`
	ListingID         string            `json:"listingId" db:"listing_id"`
	ChatRoomID        string            `json:"chatRoomId" db:"chat_room_id"`
	ProposerUserID    string            `json:"proposerUserId" db:"proposer_user_id"`
	CounterpartUserID string            `json:"counterpartUserId" db:"counterpart_user_id"`
	Status            ReservationStatus `json:"reservationStatus" db:"status"`
	ScheduledAt       time.Time         `json:"scheduledAt" db:"scheduled_at"`
	MeetingType       TradeMethod       `json:"meetingType" db:"meeting_type"`
	ServerID          *string           `json:"serverId,omitempty" db:"server_id"`
	MeetingPointText  *string           `json:"meetingPointText,omitempty" db:"meeting_point_text"`
	NoteToCounterparty *string          `json:"noteToCounterparty,omitempty" db:"note_to_counterparty"`
	ExpiresAt         *time.Time        `json:"expiresAt,omitempty" db:"expires_at"`
	ConfirmedAt       *time.Time        `json:"confirmedAt,omitempty" db:"confirmed_at"`
	CancelledAt       *time.Time        `json:"cancelledAt,omitempty" db:"cancelled_at"`
	CreatedAt         time.Time         `json:"createdAt" db:"created_at"`
}

type TradeCompletion struct {
	ID                string           `json:"completionId" db:"id"`
	ListingID         string           `json:"listingId" db:"listing_id"`
	ReservationID     string           `json:"reservationId" db:"reservation_id"`
	RequestedByUserID string           `json:"requestedByUserId" db:"requested_by_user_id"`
	CounterpartUserID string           `json:"counterpartUserId" db:"counterpart_user_id"`
	Status            CompletionStatus `json:"completionStatus" db:"status"`
	AutoConfirmAt     *time.Time       `json:"autoConfirmAt,omitempty" db:"auto_confirm_at"`
	ConfirmedAt       *time.Time       `json:"confirmedAt,omitempty" db:"confirmed_at"`
	CreatedAt         time.Time        `json:"createdAt" db:"created_at"`
}

type Review struct {
	ID             string       `json:"reviewId" db:"id"`
	CompletionID   string       `json:"completionId" db:"completion_id"`
	ReviewerUserID string       `json:"reviewerUserId" db:"reviewer_user_id"`
	TargetUserID   string       `json:"targetUserId" db:"target_user_id"`
	Rating         ReviewRating `json:"rating" db:"rating"`
	Comment        *string      `json:"comment,omitempty" db:"comment"`
	CreatedAt      time.Time    `json:"createdAt" db:"created_at"`
}

type Report struct {
	ID             string       `json:"reportId" db:"id"`
	ReporterUserID string       `json:"reporterUserId" db:"reporter_user_id"`
	TargetType     string       `json:"targetType" db:"target_type"`
	TargetID       string       `json:"targetId" db:"target_id"`
	ReportType     ReportType   `json:"reportType" db:"report_type"`
	Description    string       `json:"description" db:"description"`
	Status         ReportStatus `json:"status" db:"status"`
	CreatedAt      time.Time    `json:"createdAt" db:"created_at"`
}

type Notification struct {
	ID            string     `json:"notificationId" db:"id"`
	UserID        string     `json:"userId" db:"user_id"`
	Type          string     `json:"type" db:"type"`
	Title         string     `json:"title" db:"title"`
	Body          string     `json:"body" db:"body"`
	ReferenceType *string    `json:"referenceType,omitempty" db:"reference_type"`
	ReferenceID   *string    `json:"referenceId,omitempty" db:"reference_id"`
	DeepLink      *string    `json:"deepLink,omitempty" db:"deep_link"`
	IsRead        bool       `json:"isRead" db:"is_read"`
	CreatedAt     time.Time  `json:"createdAt" db:"created_at"`
}

type Server struct {
	ID   string `json:"serverId" db:"id"`
	Name string `json:"serverName" db:"name"`
}

type Category struct {
	ID       string  `json:"categoryId" db:"id"`
	Name     string  `json:"categoryName" db:"name"`
	ParentID *string `json:"parentId,omitempty" db:"parent_id"`
}
