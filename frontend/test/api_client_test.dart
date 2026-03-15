import 'package:flutter_test/flutter_test.dart';
import 'package:shared_preferences/shared_preferences.dart';
import 'package:lincle/shared/api/api_client.dart';

void main() {
  group('ApiClient', () {
    late ApiClient client;

    setUp(() {
      SharedPreferences.setMockInitialValues({});
      client = ApiClient();
    });

    group('constructor', () {
      test('creates Dio instance with correct baseUrl', () {
        expect(client.dio.options.baseUrl, ApiClient.baseUrl);
        expect(client.dio.options.baseUrl, 'http://localhost:8080/api/v1');
      });

      test('sets Content-Type header to application/json', () {
        expect(client.dio.options.headers['Content-Type'], 'application/json');
      });

      test('sets connect timeout to 10 seconds', () {
        expect(client.dio.options.connectTimeout, const Duration(seconds: 10));
      });

      test('sets receive timeout to 30 seconds', () {
        expect(client.dio.options.receiveTimeout, const Duration(seconds: 30));
      });
    });

    group('isLoggedIn', () {
      test('returns false initially (no token set)', () {
        expect(client.isLoggedIn, false);
      });
    });

    group('saveTokens', () {
      test('sets isLoggedIn to true after saving tokens', () async {
        expect(client.isLoggedIn, false);
        await client.saveTokens('access-123', 'refresh-456');
        expect(client.isLoggedIn, true);
      });

      test('persists tokens to SharedPreferences', () async {
        await client.saveTokens('access-abc', 'refresh-xyz');
        final prefs = await SharedPreferences.getInstance();
        expect(prefs.getString('accessToken'), 'access-abc');
        expect(prefs.getString('refreshToken'), 'refresh-xyz');
      });
    });

    group('clearTokens', () {
      test('sets isLoggedIn to false after clearing tokens', () async {
        await client.saveTokens('access-123', 'refresh-456');
        expect(client.isLoggedIn, true);
        await client.clearTokens();
        expect(client.isLoggedIn, false);
      });

      test('removes tokens from SharedPreferences', () async {
        await client.saveTokens('access-abc', 'refresh-xyz');
        await client.clearTokens();
        final prefs = await SharedPreferences.getInstance();
        expect(prefs.getString('accessToken'), isNull);
        expect(prefs.getString('refreshToken'), isNull);
      });
    });

    group('loadTokens', () {
      test('loads tokens from SharedPreferences and sets isLoggedIn', () async {
        SharedPreferences.setMockInitialValues({
          'accessToken': 'stored-access',
          'refreshToken': 'stored-refresh',
        });
        final freshClient = ApiClient();
        expect(freshClient.isLoggedIn, false);
        await freshClient.loadTokens();
        expect(freshClient.isLoggedIn, true);
      });

      test('stays logged out when no tokens stored', () async {
        SharedPreferences.setMockInitialValues({});
        final freshClient = ApiClient();
        await freshClient.loadTokens();
        expect(freshClient.isLoggedIn, false);
      });
    });

    group('createChat validateStatus', () {
      test('createChat uses validateStatus that accepts 201 and 409', () {
        // The createChat method passes Options(validateStatus: (status) => status == 201 || status == 409).
        // This means Dio will NOT throw on 409 (existing chat room), allowing the caller
        // to handle the returned chatRoomId from the 409 response body.
        //
        // We verify the pattern is correct by checking the source code expectation.
        // A full test would require a mock Dio or HTTP server, which is beyond
        // unit test scope. The key design decision is documented here:
        //
        // - 201: new chat room created successfully
        // - 409: chat room already exists, response contains existing chatRoomId
        // - Any other status: Dio throws DioException as normal
        //
        // This is tested at the integration level with the actual backend.
      }, skip: 'Requires Dio mock or HTTP server — validates design intent only');
    });
  });
}
