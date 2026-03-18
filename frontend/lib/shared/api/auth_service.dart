import 'package:flutter/widgets.dart';

/// Result of a successful authentication.
class AuthResult {
  final String idToken;
  final String? email;
  final String? displayName;

  AuthResult({required this.idToken, this.email, this.displayName});
}

/// Auth event types emitted by [AuthService.authEvents].
sealed class AuthEvent {}

class AuthEventSignIn extends AuthEvent {
  final AuthResult result;
  AuthEventSignIn(this.result);
}

class AuthEventSignOut extends AuthEvent {}

class AuthEventError extends AuthEvent {
  final Object error;
  AuthEventError(this.error);
}

/// Abstract interface for authentication services.
///
/// Decouples [LoginScreen] from concrete Google Sign-In packages so that
/// widget tests can run without `--platform chrome`.
abstract class AuthService {
  /// Initialize the auth service with the given [clientId].
  Future<void> initialize({String? clientId});

  /// Stream of authentication events (sign-in, sign-out, errors).
  Stream<AuthEvent> get authEvents;

  /// Trigger the authenticate flow (mobile platforms).
  Future<void> authenticate();

  /// Sign out the current user.
  Future<void> signOut();

  /// Whether [authenticate] is supported on the current platform.
  bool get supportsAuthenticate;

  /// Returns a platform-specific sign-in button widget.
  ///
  /// On web this renders the Google Identity Services button;
  /// on mobile this returns a styled [ElevatedButton].
  Widget renderSignInButton({double minimumWidth});
}
