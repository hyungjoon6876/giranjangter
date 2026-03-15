// ---- Dashboard ----
export interface DashboardStats {
  totalUsers: number;
  newUsersToday: number;
  activeListings: number;
  newListingsToday: number;
  pendingReports: number;
  tradesToday: number;
  activeChats: number;
  restrictedUsers: number;
}

// ---- Users ----
export interface AdminUser {
  userId: string;
  nickname: string;
  accountStatus: string;
  role: string;
  introduction?: string | null;
  primaryServerId?: string | null;
  completedTradeCount: number;
  positiveReviewCount?: number;
  alignmentScore: number;
  alignmentGrade: string;
  trustBadge?: string | null;
  lastLoginAt?: string;
  createdAt: string;
}

// ---- Moderation ----
export interface ModerationAction {
  actionId: string;
  reportId?: string | null;
  actorUserId: string;
  actionCode: string;
  restrictionScope?: string | null;
  durationDays?: number;
  memo?: string | null;
  createdAt: string;
}

// ---- Audit ----
export interface AuditLog {
  logId: string;
  actorId?: string | null;
  actorRole?: string | null;
  action: string;
  targetType: string;
  targetId: string;
  details?: string | null;
  ipAddress?: string | null;
  createdAt: string;
}

// ---- Reports ----
export interface AdminReport {
  reportId: string;
  reporterUserId: string;
  targetType: string;
  targetId: string;
  reportType: string;
  description?: string;
  status: string;
  createdAt: string;
}

// ---- Listings ----
export interface AdminListing {
  listingId: string;
  title: string;
  itemName: string;
  status: string;
  visibility: string;
  listingType: string;
  authorNickname: string;
  createdAt: string;
}

// ---- Trades ----
export interface TradeCompletion {
  completionId: string;
  listingId: string;
  listingTitle: string;
  status: string;
  requestedByUserId: string;
  counterpartUserId: string;
  autoConfirmAt?: string;
  createdAt: string;
}

// ---- Chat Messages (admin view) ----
export interface AdminChatMessage {
  messageId: string;
  senderUserId?: string | null;
  messageType: string;
  bodyText?: string | null;
  metadataJson?: string | null;
  sentAt: string;
}

// ---- API response wrappers ----
export interface DataResponse<T> {
  data: T[];
}

export interface ApiError {
  error: { code: string; message: string; details?: Record<string, unknown> };
}
