import 'auth_service.dart';
import 'google_auth_service.dart';

/// Web implementation — safe to import google_sign_in_web here because
/// this file is only resolved when compiling for the web platform.
AuthService createAuthService() => GoogleAuthService();
