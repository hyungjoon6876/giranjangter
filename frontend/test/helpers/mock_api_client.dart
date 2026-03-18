import 'package:lincle/shared/api/api_client_interface.dart';

/// A mock implementation of [IApiClient] for widget tests.
///
/// Returns canned data without making any HTTP calls.
/// Override individual methods via constructor callbacks when you need
/// specific return values for a test case.
class MockApiClient implements IApiClient {
  bool _isLoggedIn = false;

  // ── Token Management ──
  @override
  Future<void> loadTokens() async {}

  @override
  Future<void> saveTokens(String access, String refresh) async {
    _isLoggedIn = true;
  }

  @override
  Future<void> clearTokens() async {
    _isLoggedIn = false;
  }

  @override
  bool get isLoggedIn => _isLoggedIn;

  // ── Auth ──
  @override
  Future<Map<String, dynamic>> login(String provider, String token) async {
    _isLoggedIn = true;
    return {
      'accessToken': 'mock_access_token',
      'refreshToken': 'mock_refresh_token',
      'user': {'id': 'mock_user', 'nickname': '테스트유저'},
    };
  }

  @override
  Future<Map<String, dynamic>> getMe() async {
    return {'id': 'mock_user', 'nickname': '테스트유저'};
  }

  // ── Listings ──
  @override
  Future<Map<String, dynamic>> getListings({
    String? serverId,
    String? categoryId,
    String? q,
    String sort = 'recent',
    String? cursor,
  }) async {
    return {'data': [], 'next_cursor': null};
  }

  @override
  Future<Map<String, dynamic>> getListing(String id) async {
    return {'id': id, 'title': '테스트 매물', 'status': 'available'};
  }

  @override
  Future<Map<String, dynamic>> createListing(
      Map<String, dynamic> data) async {
    return {'id': 'new_listing', ...data};
  }

  @override
  Future<void> favoriteListing(String id) async {}

  @override
  Future<void> unfavoriteListing(String id) async {}

  // ── Chat ──
  @override
  Future<Map<String, dynamic>> createChat(String listingId) async {
    return {'id': 'mock_chat', 'listingId': listingId};
  }

  @override
  Future<Map<String, dynamic>> getChats() async {
    return {'data': []};
  }

  @override
  Future<Map<String, dynamic>> getMessages(String chatId,
      {String? cursor}) async {
    return {'data': []};
  }

  @override
  Future<Map<String, dynamic>> sendMessage(
      String chatId, String text) async {
    return {'id': 'mock_msg', 'bodyText': text};
  }

  // ── My Data ──
  @override
  Future<Map<String, dynamic>> getMyListings({String? status}) async {
    return {'data': []};
  }

  @override
  Future<Map<String, dynamic>> getMyTrades() async {
    return {'data': []};
  }

  @override
  Future<Map<String, dynamic>> getNotifications() async {
    return {'data': []};
  }

  @override
  Future<void> markNotificationsRead(List<String> ids) async {}

  @override
  Future<Map<String, dynamic>> getMyReports() async {
    return {'data': []};
  }

  @override
  Future<Map<String, dynamic>> getUserReviews(String userId) async {
    return {'data': []};
  }

  // ── Reservations ──
  @override
  Future<Map<String, dynamic>> createReservation(
      String chatId, Map<String, dynamic> data) async {
    return {'id': 'mock_reservation', 'chatId': chatId};
  }

  // ── Trade Completion ──
  @override
  Future<Map<String, dynamic>> completeTrade(
      String listingId, Map<String, dynamic> data) async {
    return {'id': 'mock_completion', 'listingId': listingId};
  }

  // ── Reviews ──
  @override
  Future<Map<String, dynamic>> createReview(
      String completionId, Map<String, dynamic> data) async {
    return {'id': 'mock_review', 'completionId': completionId};
  }

  // ── Reports ──
  @override
  Future<Map<String, dynamic>> createReport(
      Map<String, dynamic> data) async {
    return {'id': 'mock_report'};
  }

  // ── Servers & Categories ──
  @override
  Future<List<dynamic>> getServers() async {
    return [
      {'id': 'test_server', 'name': '테스트서버'},
    ];
  }

  @override
  Future<List<dynamic>> getCategories() async {
    return [
      {'id': 'weapon', 'name': '무기'},
    ];
  }

  // ── Item Search ──
  @override
  Future<List<dynamic>> searchItems(String query,
      {String? categoryId}) async {
    return [];
  }

  // ── Static Assets ──
  @override
  String get staticBaseUrl => 'http://localhost:8080';
}
