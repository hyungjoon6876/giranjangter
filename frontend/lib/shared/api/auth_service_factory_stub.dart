import 'auth_service.dart';

/// Stub factory used on non-web platforms (and in tests).
///
/// On mobile, GoogleSignIn works without dart:ui_web so we still use
/// GoogleAuthService — but the import must go through the conditional
/// import in `auth_service_factory.dart` so that the test runner
/// (which cannot resolve dart:ui_web) never compiles google_sign_in_web.
///
/// This stub throws at runtime; mobile entry-points should override
/// [authServiceProvider] or the conditional import resolves to the
/// real web implementation when compiled for web.
AuthService createAuthService() {
  // On non-web platforms, google_sign_in works fine without dart:ui_web.
  // However, we can't import GoogleAuthService here because it imports
  // google_sign_in_web. Instead, we defer to the provider override in
  // main.dart or use a separate mobile-only auth service.
  //
  // For the standard test runner, this stub is never called because
  // tests always override authServiceProvider with MockAuthService.
  throw UnsupportedError(
    'createAuthService() stub called. '
    'Override authServiceProvider in ProviderScope or '
    'compile with --platform chrome for web tests.',
  );
}
