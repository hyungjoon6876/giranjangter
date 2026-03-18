import 'package:flutter/material.dart';
import 'package:flutter_test/flutter_test.dart';
import 'package:flutter_riverpod/flutter_riverpod.dart';
import 'package:lincle/shared/providers/app_providers.dart';
import 'package:lincle/features/auth/login_screen.dart';
import 'helpers/mock_api_client.dart';
import 'helpers/mock_auth_service.dart';

void main() {
  testWidgets('App launches without error — LoginScreen smoke test',
      (tester) async {
    final mockAuthService = MockAuthService();
    addTearDown(() => mockAuthService.dispose());

    await tester.pumpWidget(
      ProviderScope(
        overrides: [
          apiClientProvider.overrideWithValue(MockApiClient()),
          authServiceProvider.overrideWithValue(mockAuthService),
        ],
        child: const MaterialApp(home: LoginScreen()),
      ),
    );
    await tester.pumpAndSettle();

    // Verify the app renders without throwing
    expect(find.byType(MaterialApp), findsOneWidget);
    expect(find.byType(LoginScreen), findsOneWidget);
  });
}
