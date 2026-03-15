import 'package:flutter/foundation.dart';
import 'package:flutter/material.dart';
import 'package:flutter_riverpod/flutter_riverpod.dart';
import 'package:go_router/go_router.dart';
import 'package:google_sign_in/google_sign_in.dart';
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

  // TODO: Google Cloud Console에서 발급받은 클라이언트 ID를 입력
  // flutter run --dart-define=GOOGLE_CLIENT_ID=xxx --dart-define=GOOGLE_SERVER_CLIENT_ID=xxx
  static const _clientId = String.fromEnvironment('GOOGLE_CLIENT_ID', defaultValue: '');
  static const _serverClientId = String.fromEnvironment('GOOGLE_SERVER_CLIENT_ID', defaultValue: '');

  @override
  void initState() {
    super.initState();
    _initGoogle();
  }

  Future<void> _initGoogle() async {
    if (_clientId.isEmpty) return;
    try {
      await GoogleSignIn.instance.initialize(
        clientId: _clientId,
        serverClientId: _serverClientId.isNotEmpty ? _serverClientId : null,
      );
      setState(() => _googleInitialized = true);
    } catch (e) {
      debugPrint('Google Sign-In init failed: $e');
    }
  }

  Future<void> _loginWithGoogle() async {
    setState(() => _loading = true);
    try {
      final account = await GoogleSignIn.instance.authenticate(
        scopeHint: ['email', 'profile'],
      );

      final idToken = account.authentication.idToken;
      if (idToken == null) {
        throw Exception('Google ID Token을 받지 못했습니다.');
      }

      final api = ref.read(apiClientProvider);
      final result = await api.login('google', idToken);
      ref.read(currentUserProvider.notifier).set(result['user'] as Map<String, dynamic>);
      if (mounted) context.go('/');
    } on GoogleSignInException catch (e) {
      if (e.code == GoogleSignInExceptionCode.canceled) {
        // User cancelled — do nothing
      } else if (mounted) {
        ScaffoldMessenger.of(context).showSnackBar(
          SnackBar(content: Text('Google 로그인 실패: ${e.description ?? e.code}')),
        );
      }
    } catch (e) {
      if (mounted) {
        ScaffoldMessenger.of(context).showSnackBar(
          SnackBar(content: Text('로그인 실패: $e')),
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
              Text(
                '기란장터',
                style: TextStyle(
                  fontSize: 40,
                  fontWeight: FontWeight.w900,
                  color: AppColors.gold,
                  letterSpacing: 4,
                  shadows: [
                    Shadow(
                      color: AppColors.gold.withValues(alpha: 0.4),
                      blurRadius: 20,
                    ),
                  ],
                ),
              ),
              const SizedBox(height: 8),
              const Text(
                '리니지 클래식 거래 플랫폼',
                style: TextStyle(
                  color: AppColors.textSecondary,
                  fontSize: 15,
                  letterSpacing: 1,
                ),
              ),

              const Spacer(flex: 2),

              // Google Sign-In button
              if (_googleInitialized || _clientId.isNotEmpty)
                SizedBox(
                  width: double.infinity,
                  child: ElevatedButton.icon(
                    onPressed: _loading ? null : _loginWithGoogle,
                    icon: _loading
                        ? const SizedBox(
                            width: 18,
                            height: 18,
                            child: CircularProgressIndicator(strokeWidth: 2, color: Colors.black54),
                          )
                        : const Icon(Icons.g_mobiledata, size: 24),
                    label: const Text('Google로 시작하기'),
                    style: ElevatedButton.styleFrom(
                      backgroundColor: Colors.white,
                      foregroundColor: Colors.black87,
                      padding: const EdgeInsets.symmetric(vertical: 16),
                      side: const BorderSide(color: Color(0xFFDDDDDD)),
                      elevation: 0,
                      shape: RoundedRectangleBorder(
                        borderRadius: BorderRadius.circular(12),
                      ),
                    ),
                  ),
                ),

              // Dev login (debug mode only)
              if (kDebugMode) ...[
                const SizedBox(height: 12),
                SizedBox(
                  width: double.infinity,
                  child: OutlinedButton.icon(
                    onPressed: _loading ? null : _devLogin,
                    icon: const Icon(Icons.developer_mode, size: 20),
                    label: const Text('개발자 로그인 (테스트)'),
                    style: OutlinedButton.styleFrom(
                      foregroundColor: AppColors.goldLight,
                      side: const BorderSide(color: AppColors.gold),
                      padding: const EdgeInsets.symmetric(vertical: 14),
                      shape: RoundedRectangleBorder(
                        borderRadius: BorderRadius.circular(12),
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
