import 'dart:async';

import 'package:flutter/widgets.dart';
import 'package:lincle/shared/api/auth_service.dart';

/// A mock [AuthService] for widget tests.
///
/// Does not depend on `google_sign_in` or any platform-specific packages,
/// so tests can run on any platform without `--platform chrome`.
///
/// Use [emitSignIn] / [emitSignOut] / [emitError] to simulate auth events.
class MockAuthService implements AuthService {
  final _controller = StreamController<AuthEvent>.broadcast();

  @override
  Future<void> initialize({String? clientId}) async {}

  @override
  Stream<AuthEvent> get authEvents => _controller.stream;

  @override
  Future<void> authenticate() async {}

  @override
  Future<void> signOut() async {}

  @override
  bool get supportsAuthenticate => false;

  @override
  Widget renderSignInButton({double minimumWidth = 300}) =>
      const SizedBox.shrink();

  // ── Test helpers ──

  /// Simulate a successful sign-in event.
  void emitSignIn({
    String idToken = 'mock_id_token',
    String? email = 'test@example.com',
    String? displayName = 'Test User',
  }) {
    _controller.add(AuthEventSignIn(AuthResult(
      idToken: idToken,
      email: email,
      displayName: displayName,
    )));
  }

  /// Simulate a sign-out event.
  void emitSignOut() {
    _controller.add(AuthEventSignOut());
  }

  /// Simulate an auth error.
  void emitError(Object error) {
    _controller.add(AuthEventError(error));
  }

  void dispose() => _controller.close();
}
