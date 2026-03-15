import 'package:dio/dio.dart';
import 'package:shared_preferences/shared_preferences.dart';

class ApiClient {
  static const String baseUrl = 'http://localhost:8080/api/v1';

  late final Dio dio;
  String? _accessToken;
  String? _refreshToken;

  ApiClient() {
    dio = Dio(BaseOptions(
      baseUrl: baseUrl,
      connectTimeout: const Duration(seconds: 10),
      receiveTimeout: const Duration(seconds: 30),
      headers: {'Content-Type': 'application/json'},
    ));

    dio.interceptors.add(InterceptorsWrapper(
      onRequest: (options, handler) {
        if (_accessToken != null) {
          options.headers['Authorization'] = 'Bearer $_accessToken';
        }
        handler.next(options);
      },
      onError: (error, handler) async {
        if (error.response?.statusCode == 401 && _refreshToken != null) {
          try {
            final refreshed = await _doRefresh();
            if (refreshed) {
              final opts = error.requestOptions;
              opts.headers['Authorization'] = 'Bearer $_accessToken';
              final response = await dio.fetch(opts);
              handler.resolve(response);
              return;
            }
          } catch (_) {}
        }
        handler.next(error);
      },
    ));
  }

  Future<void> loadTokens() async {
    final prefs = await SharedPreferences.getInstance();
    _accessToken = prefs.getString('accessToken');
    _refreshToken = prefs.getString('refreshToken');
  }

  Future<void> saveTokens(String access, String refresh) async {
    _accessToken = access;
    _refreshToken = refresh;
    final prefs = await SharedPreferences.getInstance();
    await prefs.setString('accessToken', access);
    await prefs.setString('refreshToken', refresh);
  }

  Future<void> clearTokens() async {
    _accessToken = null;
    _refreshToken = null;
    final prefs = await SharedPreferences.getInstance();
    await prefs.remove('accessToken');
    await prefs.remove('refreshToken');
  }

  bool get isLoggedIn => _accessToken != null;

  Future<bool> _doRefresh() async {
    try {
      final response = await Dio(BaseOptions(baseUrl: baseUrl)).post(
        '/auth/refresh',
        data: {'refreshToken': _refreshToken},
      );
      if (response.statusCode == 200) {
        await saveTokens(
          response.data['accessToken'],
          response.data['refreshToken'],
        );
        return true;
      }
    } catch (_) {}
    return false;
  }

  // ── Auth ──
  Future<Map<String, dynamic>> login(String provider, String token) async {
    final res = await dio.post('/auth/login', data: {
      'provider': provider,
      'providerToken': token,
    });
    await saveTokens(res.data['accessToken'], res.data['refreshToken']);
    return res.data;
  }

  Future<Map<String, dynamic>> getMe() async {
    final res = await dio.get('/me');
    return res.data;
  }

  // ── Listings ──
  Future<Map<String, dynamic>> getListings({
    String? serverId,
    String? categoryId,
    String? q,
    String sort = 'recent',
    String? cursor,
  }) async {
    final params = <String, dynamic>{'sort': sort, 'limit': 20};
    if (serverId != null) params['serverId'] = serverId;
    if (categoryId != null) params['categoryId'] = categoryId;
    if (q != null) params['q'] = q;
    if (cursor != null) params['cursor'] = cursor;
    final res = await dio.get('/listings', queryParameters: params);
    return res.data;
  }

  Future<Map<String, dynamic>> getListing(String id) async {
    final res = await dio.get('/listings/$id');
    return res.data;
  }

  Future<Map<String, dynamic>> createListing(Map<String, dynamic> data) async {
    final res = await dio.post('/listings', data: data);
    return res.data;
  }

  Future<void> favoriteListing(String id) async {
    await dio.post('/listings/$id/favorite');
  }

  Future<void> unfavoriteListing(String id) async {
    await dio.delete('/listings/$id/favorite');
  }

  // ── Chat ──
  Future<Map<String, dynamic>> createChat(String listingId) async {
    final res = await dio.post(
      '/listings/$listingId/chats',
      options: Options(validateStatus: (status) => status == 201 || status == 409),
    );
    return res.data;
  }

  Future<Map<String, dynamic>> getChats() async {
    final res = await dio.get('/chats');
    return res.data;
  }

  Future<Map<String, dynamic>> getMessages(String chatId) async {
    final res = await dio.get('/chats/$chatId/messages');
    return res.data;
  }

  Future<Map<String, dynamic>> sendMessage(String chatId, String text) async {
    final res = await dio.post('/chats/$chatId/messages', data: {
      'messageType': 'text',
      'bodyText': text,
    });
    return res.data;
  }

  // ── My Data ──
  Future<Map<String, dynamic>> getMyListings({String? status}) async {
    final params = <String, dynamic>{};
    if (status != null) params['status'] = status;
    final res = await dio.get('/me/listings', queryParameters: params);
    return res.data;
  }

  Future<Map<String, dynamic>> getMyTrades() async {
    final res = await dio.get('/me/trades');
    return res.data;
  }

  Future<Map<String, dynamic>> getNotifications() async {
    final res = await dio.get('/notifications');
    return res.data;
  }

  Future<void> markNotificationsRead(List<String> ids) async {
    await dio.post('/notifications/read', data: {'notificationIds': ids});
  }

  Future<Map<String, dynamic>> getMyReports() async {
    final res = await dio.get('/me/reports');
    return res.data;
  }

  Future<Map<String, dynamic>> getUserReviews(String userId) async {
    final res = await dio.get('/users/$userId/reviews');
    return res.data;
  }

  // ── Reservations ──
  Future<Map<String, dynamic>> createReservation(String chatId, Map<String, dynamic> data) async {
    final res = await dio.post('/chats/$chatId/reservations', data: data);
    return res.data;
  }

  // ── Trade Completion ──
  Future<Map<String, dynamic>> completeTrade(String listingId, Map<String, dynamic> data) async {
    final res = await dio.post('/listings/$listingId/complete', data: data);
    return res.data;
  }

  // ── Reviews ──
  Future<Map<String, dynamic>> createReview(String completionId, Map<String, dynamic> data) async {
    final res = await dio.post('/trade-completions/$completionId/reviews', data: data);
    return res.data;
  }

  // ── Reports ──
  Future<Map<String, dynamic>> createReport(Map<String, dynamic> data) async {
    final res = await dio.post('/reports', data: data);
    return res.data;
  }

  // ── Servers & Categories ──
  Future<List<dynamic>> getServers() async {
    final res = await dio.get('/servers');
    return res.data['data'];
  }

  Future<List<dynamic>> getCategories() async {
    final res = await dio.get('/categories');
    return res.data['data'];
  }
}
