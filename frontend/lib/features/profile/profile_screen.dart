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
        appBar: AppBar(title: const Text('내 정보')),
        body: Center(
          child: Column(
            mainAxisSize: MainAxisSize.min,
            children: [
              const Icon(Icons.person_outline, size: 64, color: AppTheme.textSecondary),
              const SizedBox(height: 16),
              const Text('로그인이 필요합니다', style: TextStyle(fontSize: 16)),
              const SizedBox(height: 16),
              ElevatedButton(
                onPressed: () => context.push('/login'),
                child: const Text('로그인'),
              ),
            ],
          ),
        ),
      );
    }

    final alignmentScore = user['alignmentScore'] as int? ?? 0;
    final alignmentGrade = user['alignmentGrade'] as String? ?? 'neutral';

    return Scaffold(
      appBar: AppBar(title: const Text('내 정보')),
      body: ListView(
        padding: const EdgeInsets.all(16),
        children: [
          // Profile header
          Center(
            child: Column(
              children: [
                CircleAvatar(
                  radius: 40,
                  backgroundColor: AppTheme.primary.withValues(alpha: 0.1),
                  child: Text(
                    (user['nickname'] ?? '?')[0],
                    style: const TextStyle(fontSize: 28, fontWeight: FontWeight.bold, color: AppTheme.primary),
                  ),
                ),
                const SizedBox(height: 12),
                Text(user['nickname'] ?? '', style: Theme.of(context).textTheme.titleLarge),
                const SizedBox(height: 4),
                // Alignment badge
                _AlignmentBadge(grade: alignmentGrade, score: alignmentScore),
              ],
            ),
          ),
          const SizedBox(height: 24),

          // Stats
          Card(
            child: Padding(
              padding: const EdgeInsets.all(16),
              child: Row(
                mainAxisAlignment: MainAxisAlignment.spaceAround,
                children: [
                  _statItem('거래', '${user['completedTradeCount'] ?? 0}'),
                  _statItem('후기', '${user['positiveReviewCount'] ?? 0}'),
                  _statItem('성향', '$alignmentScore'),
                ],
              ),
            ),
          ),
          const SizedBox(height: 16),

          // Menu items
          _menuItem(Icons.list_alt, '내 매물', () {
            Navigator.of(context).push(MaterialPageRoute(builder: (_) => const MyListingsScreen()));
          }),
          _menuItem(Icons.swap_horiz, '내 거래', () {
            Navigator.of(context).push(MaterialPageRoute(builder: (_) => const MyTradesScreen()));
          }),
          _menuItem(Icons.notifications_outlined, '알림', () {
            Navigator.of(context).push(MaterialPageRoute(builder: (_) => const NotificationsScreen()));
          }),
          const Divider(height: 32),
          _menuItem(Icons.logout, '로그아웃', () async {
            final api = ref.read(apiClientProvider);
            await api.clearTokens();
            ref.read(currentUserProvider.notifier).set(null);
          }, color: AppTheme.error),
        ],
      ),
    );
  }

  Widget _statItem(String label, String value) {
    return Column(
      children: [
        Text(value, style: const TextStyle(fontSize: 20, fontWeight: FontWeight.bold)),
        const SizedBox(height: 4),
        Text(label, style: const TextStyle(fontSize: 12, color: AppTheme.textSecondary)),
      ],
    );
  }

  Widget _menuItem(IconData icon, String label, VoidCallback onTap, {Color? color}) {
    return ListTile(
      leading: Icon(icon, color: color),
      title: Text(label, style: TextStyle(color: color)),
      trailing: const Icon(Icons.chevron_right),
      onTap: onTap,
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
        color: color.withValues(alpha: 0.1),
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
      'royal_knight' => ('로얄 나이트', Colors.blue, Icons.shield),
      'lawful' => ('라이풀', Colors.green, Icons.verified_user),
      'neutral' => ('뉴트럴', Colors.grey, Icons.person),
      'caution' => ('주의', Colors.orange, Icons.warning_amber),
      'chaotic' => ('카오틱', Colors.red, Icons.dangerous),
      _ => ('뉴트럴', Colors.grey, Icons.person),
    };
  }
}
