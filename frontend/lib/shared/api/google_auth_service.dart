import 'dart:async';

import 'package:flutter/foundation.dart' show kIsWeb;
import 'package:flutter/material.dart';
import 'package:google_sign_in/google_sign_in.dart';
import 'package:google_sign_in_web/web_only.dart' as web;

import 'auth_service.dart';

/// Concrete [AuthService] backed by the `google_sign_in` package.
///
/// This is the **only** file in the project that imports `google_sign_in` or
/// `google_sign_in_web`. All other code accesses auth through the abstract
/// [AuthService] interface.
class GoogleAuthService implements AuthService {
  final _controller = StreamController<AuthEvent>.broadcast();
  StreamSubscription? _eventSub;

  @override
  Future<void> initialize({String? clientId}) async {
    final signIn = GoogleSignIn.instance;
    await signIn.initialize(clientId: clientId);

    _eventSub = signIn.authenticationEvents.listen(
      (event) {
        if (event is GoogleSignInAuthenticationEventSignIn) {
          final user = event.user;
          final idToken = user.authentication.idToken;
          if (idToken != null) {
            _controller.add(AuthEventSignIn(AuthResult(
              idToken: idToken,
              email: user.email,
              displayName: user.displayName,
            )));
          } else {
            _controller.add(AuthEventError(
              Exception('Google ID Token을 받지 못했습니다.'),
            ));
          }
        } else if (event is GoogleSignInAuthenticationEventSignOut) {
          _controller.add(AuthEventSignOut());
        }
      },
      onError: (error) {
        _controller.add(AuthEventError(error));
      },
    );
  }

  @override
  Stream<AuthEvent> get authEvents => _controller.stream;

  @override
  Future<void> authenticate() async {
    await GoogleSignIn.instance.authenticate();
  }

  @override
  Future<void> signOut() async {
    await GoogleSignIn.instance.signOut();
  }

  @override
  bool get supportsAuthenticate =>
      GoogleSignIn.instance.supportsAuthenticate();

  @override
  Widget renderSignInButton({double minimumWidth = 300}) {
    if (kIsWeb) {
      return web.renderButton(
        configuration: web.GSIButtonConfiguration(
          type: web.GSIButtonType.standard,
          theme: web.GSIButtonTheme.filledBlue,
          size: web.GSIButtonSize.large,
          text: web.GSIButtonText.signinWith,
          shape: web.GSIButtonShape.rectangular,
          minimumWidth: minimumWidth,
        ),
      );
    }
    // Mobile: return a placeholder — LoginScreen builds its own button
    // because it needs access to _loading state and onPressed callback.
    return const SizedBox.shrink();
  }

  /// Release resources. Call when the service is no longer needed.
  void dispose() {
    _eventSub?.cancel();
    _controller.close();
  }
}
