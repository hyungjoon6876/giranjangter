/// Abstract interface for [ApiClient].
///
/// Enables dependency injection and testability — widget tests can
/// override [apiClientProvider] with a mock that implements this interface
/// without making real HTTP calls.
abstract class IApiClient {
  // ── Token Management ──
  Future<void> loadTokens();
  Future<void> saveTokens(String access, String refresh);
  Future<void> clearTokens();
  bool get isLoggedIn;

  // ── Auth ──
  Future<Map<String, dynamic>> login(String provider, String token);
  Future<Map<String, dynamic>> getMe();

  // ── Listings ──
  Future<Map<String, dynamic>> getListings({
    String? serverId,
    String? categoryId,
    String? q,
    String sort,
    String? cursor,
  });
  Future<Map<String, dynamic>> getListing(String id);
  Future<Map<String, dynamic>> createListing(Map<String, dynamic> data);
  Future<void> favoriteListing(String id);
  Future<void> unfavoriteListing(String id);

  // ── Chat ──
  Future<Map<String, dynamic>> createChat(String listingId);
  Future<Map<String, dynamic>> getChats();
  Future<Map<String, dynamic>> getMessages(String chatId, {String? cursor});
  Future<Map<String, dynamic>> sendMessage(String chatId, String text);

  // ── My Data ──
  Future<Map<String, dynamic>> getMyListings({String? status});
  Future<Map<String, dynamic>> getMyTrades();
  Future<Map<String, dynamic>> getNotifications();
  Future<void> markNotificationsRead(List<String> ids);
  Future<Map<String, dynamic>> getMyReports();
  Future<Map<String, dynamic>> getUserReviews(String userId);

  // ── Reservations ──
  Future<Map<String, dynamic>> createReservation(
      String chatId, Map<String, dynamic> data);

  // ── Trade Completion ──
  Future<Map<String, dynamic>> completeTrade(
      String listingId, Map<String, dynamic> data);

  // ── Reviews ──
  Future<Map<String, dynamic>> createReview(
      String completionId, Map<String, dynamic> data);

  // ── Reports ──
  Future<Map<String, dynamic>> createReport(Map<String, dynamic> data);

  // ── Servers & Categories ──
  Future<List<dynamic>> getServers();
  Future<List<dynamic>> getCategories();

  // ── Item Search ──
  Future<List<dynamic>> searchItems(String query, {String? categoryId});

  // ── Static Assets ──
  String get staticBaseUrl;
}
