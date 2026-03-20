export interface User {
  userId: string;
  nickname: string;
  avatarUrl?: string;
  introduction?: string;
  primaryServerId?: string;
  role: string;
  accountStatus: string;
  completedTradeCount: number;
  positiveReviewCount: number;
  responseBadge?: string;
  trustBadge?: string;
  lastActiveAt: string;
  createdAt: string;
}

export interface Author {
  userId: string;
  nickname: string;
  avatarUrl?: string;
  trustBadge?: string;
  responseBadge?: string;
  completedTradeCount?: number;
  lastActiveAt?: string;
}

export interface Listing {
  listingId: string;
  listingType: "sell" | "buy";
  title: string;
  itemName: string;
  description?: string;
  priceType: "fixed" | "negotiable" | "offer";
  priceAmount?: number;
  quantity: number;
  enhancementLevel?: number;
  optionsText?: string;
  serverId: string;
  serverName: string;
  categoryId: string;
  categoryName?: string;
  status: string;
  visibility: string;
  tradeMethod: string;
  preferredMeetingAreaText?: string;
  availableTimeText?: string;
  thumbnailUrl?: string;
  iconUrl?: string;
  images?: { imageId: string; url: string; order: number }[];
  author: Author;
  viewCount: number;
  favoriteCount: number;
  chatCount: number;
  isFavorited?: boolean;
  isOwner?: boolean;
  availableActions?: string[];
  lastActivityAt: string;
  createdAt: string;
  updatedAt?: string;
}

export interface ChatRoom {
  chatRoomId: string;
  listingId: string;
  listingTitle: string;
  listingThumbnail?: string;
  listingStatus: string;
  listingPrice?: number;
  listingServerId?: string;
  listingServerName?: string;
  counterparty: Author;
  chatStatus: string;
  lastMessage?: Message;
  unreadCount: number;
  myLastReadMessageId?: string;
  updatedAt: string;
}

export interface Message {
  messageId: string;
  chatRoomId: string;
  senderUserId?: string;
  messageType: "text" | "system" | "reservation_card";
  bodyText?: string;
  metadataJson?: Record<string, unknown>;
  sentAt: string;
  status?: "sending" | "sent" | "failed";  // client-only, not from server
}

export interface Server {
  serverId: string;
  serverName: string;
}

export interface Category {
  categoryId: string;
  categoryName: string;
  parentId?: string;
}

export interface PaginatedResponse<T> {
  data: T[];
  cursor: { next?: string; hasMore: boolean };
}

export interface AuthResponse {
  accessToken: string;
  refreshToken: string;
  expiresIn: number;
  user: User;
}

export interface ApiError {
  error: { code: string; message: string; details?: Record<string, unknown> };
}

export interface Notification {
  notificationId: string;
  message: string;
  readAt?: string;
  createdAt: string;
}

export interface ItemSearchResult {
  id: string;
  name: string;
  categoryId: string;
  iconUrl?: string;
}

export interface Review {
  reviewId: string;
  rating: "positive" | "negative";
  comment?: string;
  reviewerNickname: string;
  createdAt: string;
}

export interface UploadedImage {
  imageId: string;
  url: string;
  thumbnailUrl: string;
  sizeBytes?: number;
}
