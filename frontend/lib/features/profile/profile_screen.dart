import 'package:flutter/material.dart';
import 'package:flutter_riverpod/flutter_riverpod.dart';
import 'package:go_router/go_router.dart';
import '../../shared/providers/app_providers.dart';
import '../../shared/theme/app_theme.dart';
import 'my_listings_screen.dart';
import 'my_trades_screen.dart';
import 'notifications_screen.dart';

class ProfileScreen extends ConsumerWidget {
  const ProfileScreen({super.key});

  @override
  Widget build(BuildContext context, WidgetRef ref) {
    final user = ref.watch(currentUserProvider);

    if (user == null) {
      return Scaffold(
        backgroundColor: AppColors.bg,
        appBar: AppBar(
          title: const Text('내 정보'),
          backgroundColor: AppColors.bgSurface,
        ),
        body: Center(
          child: Column(
            mainAxisSize: MainAxisSize.min,
            children: [
              Icon(Icons.person_outline, size: 64, color: AppColors.textMuted),
              const SizedBox(height: 16),
              const Text(
                '로그인이 필요합니다',
                style: TextStyle(fontSize: 16, color: AppColors.textSecondary),
              ),
              const SizedBox(height: 20),
              ElevatedButton(
                onPressed: () => context.push('/login'),
                style: ElevatedButton.styleFrom(
                  backgroundColor: AppColors.gold,
                  foregroundColor: AppColors.bg,
                  padding: const EdgeInsets.symmetric(horizontal: 32, vertical: 14),
                  shape: RoundedRectangleBorder(
                    borderRadius: BorderRadius.circular(10),
                  ),
                ),
                child: const Text(
                  '로그인',
                  style: TextStyle(fontWeight: FontWeight.w700, fontSize: 15),
                ),
              ),
            ],
          ),
        ),
      );
    }

    final alignmentScore = user['alignmentScore'] as int? ?? 0;
    final alignmentGrade = user['alignmentGrade'] as String? ?? 'neutral';

    return Scaffold(
      backgroundColor: AppColors.bg,
      appBar: AppBar(
        title: const Text('내 정보'),
        backgroundColor: AppColors.bgSurface,
      ),
      body: ListView(
        padding: const EdgeInsets.all(16),
        children: [
          // Profile header
          Center(
            child: Column(
              children: [
                // Avatar with gold border glow
                Container(
                  padding: const EdgeInsets.all(3),
                  decoration: BoxDecoration(
                    shape: BoxShape.circle,
                    border: Border.all(color: AppColors.gold, width: 2),
                    boxShadow: [
                      BoxShadow(
                        color: AppColors.gold.withValues(alpha: 0.3),
                        blurRadius: 16,
                        spreadRadius: 2,
                      ),
                    ],
                  ),
                  child: CircleAvatar(
                    radius: 44,
                    backgroundColor: AppColors.bgElevated,
                    child: Text(
                      (user['nickname'] ?? '?')[0],
                      style: const TextStyle(
                        fontSize: 32,
                        fontWeight: FontWeight.bold,
                        color: AppColors.gold,
                      ),
                    ),
                  ),
                ),
                const SizedBox(height: 14),
                Text(
                  user['nickname'] ?? '',
                  style: const TextStyle(
                    fontSize: 22,
                    fontWeight: FontWeight.w700,
                    color: AppColors.textPrimary,
                  ),
                ),
                const SizedBox(height: 8),
                // Alignment badge
                _AlignmentBadge(grade: alignmentGrade, score: alignmentScore),
              ],
            ),
          ),
          const SizedBox(height: 24),

          // Stats card
          Container(
            padding: const EdgeInsets.symmetric(vertical: 20, horizontal: 16),
            decoration: BoxDecoration(
              color: AppColors.bgCard,
              borderRadius: BorderRadius.circular(12),
              border: Border.all(color: AppColors.border),
            ),
            child: Row(
              mainAxisAlignment: MainAxisAlignment.spaceAround,
              children: [
                _statItem('거래', '${user['completedTradeCount'] ?? 0}'),
                Container(width: 1, height: 36, color: AppColors.border),
                _statItem('후기', '${user['positiveReviewCount'] ?? 0}'),
                Container(width: 1, height: 36, color: AppColors.border),
                _statItem('성향', '$alignmentScore'),
              ],
            ),
          ),
          const SizedBox(height: 20),

          // Menu items
          _menuItem(Icons.list_alt, '내 매물', () {
            Navigator.of(context).push(MaterialPageRoute(builder: (_) => const MyListingsScreen()));
          }),
          const SizedBox(height: 8),
          _menuItem(Icons.swap_horiz, '내 거래', () {
            Navigator.of(context).push(MaterialPageRoute(builder: (_) => const MyTradesScreen()));
          }),
          const SizedBox(height: 8),
          _menuItem(Icons.notifications_outlined, '알림', () {
            Navigator.of(context).push(MaterialPageRoute(builder: (_) => const NotificationsScreen()));
          }),
          const SizedBox(height: 16),
          const Divider(color: AppColors.border),
          const SizedBox(height: 8),
          _menuItem(Icons.logout, '로그아웃', () async {
            final api = ref.read(apiClientProvider);
            await api.clearTokens();
            ref.read(currentUserProvider.notifier).set(null);
          }, isDestructive: true),
        ],
      ),
    );
  }

  Widget _statItem(String label, String value) {
    return Column(
      children: [
        Text(
          value,
          style: const TextStyle(
            fontSize: 22,
            fontWeight: FontWeight.bold,
            color: AppColors.gold,
          ),
        ),
        const SizedBox(height: 4),
        Text(
          label,
          style: const TextStyle(fontSize: 12, color: AppColors.textSecondary),
        ),
      ],
    );
  }

  Widget _menuItem(IconData icon, String label, VoidCallback onTap, {bool isDestructive = false}) {
    return Material(
      color: Colors.transparent,
      child: InkWell(
        onTap: onTap,
        borderRadius: BorderRadius.circular(12),
        child: Container(
          padding: const EdgeInsets.symmetric(horizontal: 16, vertical: 14),
          decoration: BoxDecoration(
            color: AppColors.bgCard,
            borderRadius: BorderRadius.circular(12),
            border: Border.all(color: AppColors.border),
          ),
          child: Row(
            children: [
              Icon(icon, color: isDestructive ? AppColors.error : AppColors.textSecondary, size: 22),
              const SizedBox(width: 12),
              Expanded(
                child: Text(
                  label,
                  style: TextStyle(
                    color: isDestructive ? AppColors.error : AppColors.textPrimary,
                    fontSize: 15,
                    fontWeight: FontWeight.w500,
                  ),
                ),
              ),
              Icon(
                Icons.chevron_right,
                color: isDestructive ? AppColors.error.withValues(alpha: 0.5) : AppColors.goldLight,
                size: 20,
              ),
            ],
          ),
        ),
      ),
    );
  }
}

class _AlignmentBadge extends StatelessWidget {
  final String grade;
  final int score;

  const _AlignmentBadge({required this.grade, required this.score});

  @override
  Widget build(BuildContext context) {
    final (label, color, icon) = _gradeInfo(grade);
    return Container(
      padding: const EdgeInsets.symmetric(horizontal: 12, vertical: 6),
      decoration: BoxDecoration(
        color: color.withValues(alpha: 0.15),
        borderRadius: BorderRadius.circular(16),
        border: Border.all(color: color.withValues(alpha: 0.3)),
      ),
      child: Row(
        mainAxisSize: MainAxisSize.min,
        children: [
          Icon(icon, size: 16, color: color),
          const SizedBox(width: 4),
          Text(label, style: TextStyle(color: color, fontWeight: FontWeight.w600, fontSize: 13)),
        ],
      ),
    );
  }

  (String, Color, IconData) _gradeInfo(String grade) {
    return switch (grade) {
      'royal_knight' => ('로얄 나이트', AppColors.royalKnight, Icons.shield),
      'lawful' => ('라이풀', AppColors.lawful, Icons.verified_user),
      'neutral' => ('뉴트럴', AppColors.neutral, Icons.person),
      'caution' => ('주의', AppColors.caution, Icons.warning_amber),
      'chaotic' => ('카오틱', AppColors.chaotic, Icons.dangerous),
      _ => ('뉴트럴', AppColors.neutral, Icons.person),
    };
  }
}
