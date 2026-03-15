import 'dart:async';
import 'package:flutter/foundation.dart';
import 'package:flutter/material.dart';
import 'package:flutter_riverpod/flutter_riverpod.dart';
import 'package:go_router/go_router.dart';
import 'package:google_sign_in/google_sign_in.dart';
import 'package:google_sign_in_web/web_only.dart' as web;
import '../../shared/providers/app_providers.dart';
import '../../shared/theme/app_theme.dart';

class LoginScreen extends ConsumerStatefulWidget {
  const LoginScreen({super.key});

  @override
  ConsumerState<LoginScreen> createState() => _LoginScreenState();
}

class _LoginScreenState extends ConsumerState<LoginScreen> {
  bool _loading = false;
  bool _googleInitialized = false;
  StreamSubscription? _authSub;

  static const _clientId = String.fromEnvironment('GOOGLE_CLIENT_ID', defaultValue: '');

  @override
  void initState() {
    super.initState();
    _initGoogle();
  }

  @override
  void dispose() {
    _authSub?.cancel();
    super.dispose();
  }

  Future<void> _initGoogle() async {
    if (_clientId.isEmpty) return;
    try {
      final signIn = GoogleSignIn.instance;
      await signIn.initialize(clientId: _clientId);

      // Listen for authentication events (Web uses renderButton flow)
      _authSub = signIn.authenticationEvents.listen(
        (event) {
          if (event is GoogleSignInAuthenticationEventSignIn) {
            _handleSignIn(event.user);
          }
        },
        onError: (error) {
          if (mounted) {
            ScaffoldMessenger.of(context).showSnackBar(
              SnackBar(
                content: Text('Google 로그인 실패: $error'),
                backgroundColor: AppColors.error,
              ),
            );
          }
        },
      );

      setState(() => _googleInitialized = true);
    } catch (e) {
      debugPrint('Google Sign-In init failed: $e');
    }
  }

  Future<void> _handleSignIn(GoogleSignInAccount user) async {
    setState(() => _loading = true);
    try {
      final idToken = user.authentication.idToken;
      if (idToken == null) {
        throw Exception('Google ID Token을 받지 못했습니다.');
      }

      final api = ref.read(apiClientProvider);
      final result = await api.login('google', idToken);
      ref.read(currentUserProvider.notifier).set(result['user'] as Map<String, dynamic>);
      if (mounted) context.go('/');
    } catch (e) {
      if (mounted) {
        ScaffoldMessenger.of(context).showSnackBar(
          SnackBar(
            content: Text('로그인 실패: $e'),
            backgroundColor: AppColors.error,
          ),
        );
      }
    } finally {
      if (mounted) setState(() => _loading = false);
    }
  }

  Future<void> _devLogin() async {
    setState(() => _loading = true);
    try {
      final api = ref.read(apiClientProvider);
      final result = await api.login('google', 'dev_user_${DateTime.now().millisecondsSinceEpoch}');
      ref.read(currentUserProvider.notifier).set(result['user'] as Map<String, dynamic>);
      if (mounted) context.go('/');
    } catch (e) {
      if (mounted) {
        ScaffoldMessenger.of(context).showSnackBar(SnackBar(content: Text('로그인 실패: $e')));
      }
    } finally {
      if (mounted) setState(() => _loading = false);
    }
  }

  @override
  Widget build(BuildContext context) {
    return Scaffold(
      backgroundColor: AppColors.bg,
      body: SafeArea(
        child: Padding(
          padding: const EdgeInsets.symmetric(horizontal: 32),
          child: Column(
            mainAxisAlignment: MainAxisAlignment.center,
            children: [
              const Spacer(flex: 2),

              // Logo
              Image.asset(
                'assets/images/logo.png',
                height: 80,
                fit: BoxFit.contain,
              ),
              const SizedBox(height: 12),
              const Text(
                '리니지 클래식 거래 플랫폼',
                style: TextStyle(
                  color: AppColors.textSecondary,
                  fontSize: 15,
                  letterSpacing: 1,
                ),
              ),

              const Spacer(flex: 2),

              // Google Sign-In: official rendered button (Web) or custom button (mobile)
              if (_googleInitialized) ...[
                if (kIsWeb)
                  // Google official sign-in button for Web (full width, large)
                  web.renderButton(
                    configuration: web.GSIButtonConfiguration(
                      type: web.GSIButtonType.standard,
                      theme: web.GSIButtonTheme.filledBlue,
                      size: web.GSIButtonSize.large,
                      text: web.GSIButtonText.signinWith,
                      shape: web.GSIButtonShape.rectangular,
                      minimumWidth: 320,
                    ),
                  )
                else if (GoogleSignIn.instance.supportsAuthenticate())
                  SizedBox(
                    width: double.infinity,
                    child: ElevatedButton.icon(
                      onPressed: _loading ? null : () async {
                        try {
                          await GoogleSignIn.instance.authenticate();
                        } catch (e) {
                          if (mounted) {
                            ScaffoldMessenger.of(context).showSnackBar(
                              SnackBar(content: Text('Google 로그인 실패: $e')),
                            );
                          }
                        }
                      },
                      icon: const Icon(Icons.g_mobiledata, size: 24),
                      label: const Text('Google로 시작하기'),
                      style: ElevatedButton.styleFrom(
                        backgroundColor: Colors.white,
                        foregroundColor: Colors.black87,
                        padding: const EdgeInsets.symmetric(vertical: 16),
                        shape: RoundedRectangleBorder(
                          borderRadius: BorderRadius.circular(12),
                        ),
                      ),
                    ),
                  ),
              ],

              // Dev login (debug mode only)
              if (kDebugMode) ...[
                const SizedBox(height: 12),
                SizedBox(
                  width: 320,
                  height: 44,
                  child: OutlinedButton.icon(
                    onPressed: _loading ? null : _devLogin,
                    icon: const Icon(Icons.developer_mode, size: 18),
                    label: const Text('개발자 로그인 (테스트)', style: TextStyle(fontSize: 14)),
                    style: OutlinedButton.styleFrom(
                      foregroundColor: AppColors.goldLight,
                      side: const BorderSide(color: AppColors.gold),
                      shape: RoundedRectangleBorder(
                        borderRadius: BorderRadius.circular(4),
                      ),
                    ),
                  ),
                ),
              ],

              const SizedBox(height: 16),
              TextButton(
                onPressed: () => context.go('/'),
                child: const Text(
                  '둘러보기',
                  style: TextStyle(color: AppColors.textSecondary, fontSize: 14),
                ),
              ),

              const Spacer(flex: 1),
            ],
          ),
        ),
      ),
    );
  }
}
