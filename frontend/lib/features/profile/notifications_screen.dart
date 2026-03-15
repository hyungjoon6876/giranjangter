import 'package:flutter/material.dart';
import 'package:flutter_riverpod/flutter_riverpod.dart';
import '../../shared/providers/app_providers.dart';
import '../../shared/theme/app_theme.dart';

class NotificationsScreen extends ConsumerStatefulWidget {
  const NotificationsScreen({super.key});

  @override
  ConsumerState<NotificationsScreen> createState() => _NotificationsScreenState();
}

class _NotificationsScreenState extends ConsumerState<NotificationsScreen> {
  List<dynamic> _notifications = [];
  bool _loading = true;
  String? _error;

  @override
  void initState() {
    super.initState();
    _load();
  }

  Future<void> _load() async {
    setState(() { _loading = true; _error = null; });
    try {
      final api = ref.read(apiClientProvider);
      final res = await api.getNotifications();
      setState(() { _notifications = res['data'] ?? []; _loading = false; });
    } catch (e) {
      setState(() { _error = e.toString(); _loading = false; });
    }
  }

  Future<void> _markAllRead() async {
    final unread = _notifications
        .where((n) => n['isRead'] != true)
        .map<String>((n) => n['notificationId'] as String)
        .toList();
    if (unread.isEmpty) return;

    try {
      final api = ref.read(apiClientProvider);
      await api.markNotificationsRead(unread);
      await _load();
    } catch (_) {}
  }

  @override
  Widget build(BuildContext context) {
    return Scaffold(
      backgroundColor: AppColors.bg,
      appBar: AppBar(
        title: const Text('알림'),
        backgroundColor: AppColors.bgSurface,
        actions: [
          if (_notifications.any((n) => n['isRead'] != true))
            TextButton(
              onPressed: _markAllRead,
              child: const Text('모두 읽음', style: TextStyle(color: AppColors.gold)),
            ),
        ],
      ),
      body: _loading
          ? const Center(child: CircularProgressIndicator(color: AppColors.gold))
          : _error != null
              ? Center(child: Column(
                  mainAxisSize: MainAxisSize.min,
                  children: [
                    const Icon(Icons.error_outline, size: 48, color: AppColors.textMuted),
                    const SizedBox(height: 8),
                    const Text('불러오기 실패', style: TextStyle(color: AppColors.textSecondary)),
                    const SizedBox(height: 8),
                    TextButton(
                      onPressed: _load,
                      child: const Text('다시 시도', style: TextStyle(color: AppColors.gold)),
                    ),
                  ],
                ))
              : _notifications.isEmpty
                  ? const Center(child: Column(
                      mainAxisSize: MainAxisSize.min,
                      children: [
                        Icon(Icons.notifications_none, size: 48, color: AppColors.textMuted),
                        SizedBox(height: 8),
                        Text('알림이 없습니다', style: TextStyle(color: AppColors.textSecondary)),
                      ],
                    ))
                  : RefreshIndicator(
                      color: AppColors.gold,
                      backgroundColor: AppColors.bgCard,
                      onRefresh: _load,
                      child: ListView.builder(
                        padding: const EdgeInsets.symmetric(vertical: 8, horizontal: 12),
                        itemCount: _notifications.length,
                        itemBuilder: (context, i) {
                          final notif = _notifications[i];
                          final isRead = notif['isRead'] == true;
                          return Padding(
                            padding: const EdgeInsets.only(bottom: 8),
                            child: Container(
                              padding: const EdgeInsets.all(14),
                              decoration: BoxDecoration(
                                color: isRead ? AppColors.bgCard : AppColors.bgSurface,
                                borderRadius: BorderRadius.circular(12),
                                border: Border.all(
                                  color: isRead ? AppColors.border : AppColors.gold.withValues(alpha: 0.3),
                                ),
                              ),
                              child: Row(
                                crossAxisAlignment: CrossAxisAlignment.start,
                                children: [
                                  Container(
                                    width: 36,
                                    height: 36,
                                    decoration: BoxDecoration(
                                      color: (isRead ? AppColors.textMuted : AppColors.gold).withValues(alpha: 0.15),
                                      shape: BoxShape.circle,
                                    ),
                                    child: Icon(
                                      _notifIcon(notif['type']),
                                      color: isRead ? AppColors.textMuted : AppColors.gold,
                                      size: 18,
                                    ),
                                  ),
                                  const SizedBox(width: 12),
                                  Expanded(
                                    child: Column(
                                      crossAxisAlignment: CrossAxisAlignment.start,
                                      children: [
                                        Text(
                                          notif['title'] ?? '',
                                          style: TextStyle(
                                            fontWeight: isRead ? FontWeight.normal : FontWeight.bold,
                                            color: AppColors.textPrimary,
                                            fontSize: 14,
                                          ),
                                        ),
                                        const SizedBox(height: 4),
                                        Text(
                                          notif['body'] ?? '',
                                          maxLines: 2,
                                          overflow: TextOverflow.ellipsis,
                                          style: const TextStyle(
                                            color: AppColors.textSecondary,
                                            fontSize: 13,
                                          ),
                                        ),
                                      ],
                                    ),
                                  ),
                                  if (!isRead)
                                    Container(
                                      width: 8,
                                      height: 8,
                                      margin: const EdgeInsets.only(top: 4),
                                      decoration: const BoxDecoration(
                                        color: AppColors.gold,
                                        shape: BoxShape.circle,
                                      ),
                                    ),
                                ],
                              ),
                            ),
                          );
                        },
                      ),
                    ),
    );
  }

  IconData _notifIcon(String? type) {
    return switch (type) {
      'chat' => Icons.chat_bubble_outline,
      'reservation' => Icons.schedule,
      'trade' => Icons.swap_horiz,
      'review' => Icons.rate_review,
      'report' => Icons.flag,
      'system' => Icons.info_outline,
      _ => Icons.notifications_outlined,
    };
  }
}
