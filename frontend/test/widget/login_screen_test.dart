import 'package:flutter/material.dart';
import 'package:flutter_test/flutter_test.dart';
import 'package:flutter_riverpod/flutter_riverpod.dart';
import 'package:lincle/features/auth/login_screen.dart';
import 'package:lincle/shared/theme/app_theme.dart';

void main() {
  group('LoginScreen', () {
    Widget buildTestWidget() {
      // LoginScreen uses ConsumerStatefulWidget + GoRouter (context.go).
      // We wrap it in MaterialApp to provide Navigator context, and
      // ProviderScope for Riverpod. GoRouter is not needed for render tests
      // since we don't navigate in these tests.
      return ProviderScope(
        child: MaterialApp(
          theme: AppTheme.dark,
          home: const LoginScreen(),
        ),
      );
    }

    testWidgets('renders app title "기란장터"', (tester) async {
      await tester.pumpWidget(buildTestWidget());
      await tester.pump();

      expect(find.text('기란장터'), findsOneWidget);
    });

    testWidgets('renders subtitle "리니지 클래식 거래 플랫폼"', (tester) async {
      await tester.pumpWidget(buildTestWidget());
      await tester.pump();

      expect(find.text('리니지 클래식 거래 플랫폼'), findsOneWidget);
    });

    testWidgets('shows "둘러보기" button', (tester) async {
      await tester.pumpWidget(buildTestWidget());
      await tester.pump();

      expect(find.text('둘러보기'), findsOneWidget);
    });

    testWidgets('shows dev login button in debug mode', (tester) async {
      // LoginScreen conditionally shows dev login when kDebugMode is true.
      // In test environment, kDebugMode is true by default.
      await tester.pumpWidget(buildTestWidget());
      await tester.pump();

      expect(find.text('개발자 로그인 (테스트)'), findsOneWidget);
    });

    testWidgets('Google login button visibility depends on client ID', (tester) async {
      // The Google login button only appears when _googleInitialized or _clientId.isNotEmpty.
      // In test, GOOGLE_CLIENT_ID is empty (no --dart-define), so the Google button
      // will NOT appear unless _clientId is provided at compile time.
      await tester.pumpWidget(buildTestWidget());
      await tester.pump();

      // With empty GOOGLE_CLIENT_ID, Google button should not be rendered
      expect(find.text('Google로 시작하기'), findsNothing);
    });
  });
}
