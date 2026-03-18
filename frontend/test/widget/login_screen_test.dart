import 'package:flutter/material.dart';
import 'package:flutter_test/flutter_test.dart';
import 'package:flutter_riverpod/flutter_riverpod.dart';
import 'package:lincle/shared/providers/app_providers.dart';
import 'package:lincle/features/auth/login_screen.dart';
import '../helpers/mock_api_client.dart';
import '../helpers/mock_auth_service.dart';

void main() {
  group('LoginScreen', () {
    late MockAuthService mockAuthService;

    setUp(() {
      mockAuthService = MockAuthService();
    });

    tearDown(() {
      mockAuthService.dispose();
    });

    Widget buildTestWidget() {
      return ProviderScope(
        overrides: [
          apiClientProvider.overrideWithValue(MockApiClient()),
          authServiceProvider.overrideWithValue(mockAuthService),
        ],
        child: const MaterialApp(home: LoginScreen()),
      );
    }

    testWidgets('renders logo image', (tester) async {
      await tester.pumpWidget(buildTestWidget());
      await tester.pumpAndSettle();

      // LoginScreen shows the logo via Image.asset
      final images = find.byType(Image);
      expect(images, findsOneWidget);
    });

    testWidgets('renders subtitle "리니지 클래식 거래 플랫폼"', (tester) async {
      await tester.pumpWidget(buildTestWidget());
      await tester.pumpAndSettle();

      expect(find.text('리니지 클래식 거래 플랫폼'), findsOneWidget);
    });

    testWidgets('shows "둘러보기" button', (tester) async {
      await tester.pumpWidget(buildTestWidget());
      await tester.pumpAndSettle();

      expect(find.text('둘러보기'), findsOneWidget);
      // It's a TextButton
      final button = find.widgetWithText(TextButton, '둘러보기');
      expect(button, findsOneWidget);
    });

    testWidgets('Google login button hidden when no client ID', (tester) async {
      await tester.pumpWidget(buildTestWidget());
      await tester.pumpAndSettle();

      // GOOGLE_CLIENT_ID is empty in test env, so _initGoogle() returns early
      // and _googleInitialized stays false → no Google sign-in button rendered.
      // Only "둘러보기" button should be visible, no ElevatedButton for Google.
      expect(find.byType(ElevatedButton), findsNothing);
    });
  });
}
