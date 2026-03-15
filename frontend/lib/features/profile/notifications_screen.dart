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
      appBar: AppBar(
        title: const Text('알림'),
        actions: [
          if (_notifications.any((n) => n['isRead'] != true))
            TextButton(
              onPressed: _markAllRead,
              child: const Text('모두 읽음'),
            ),
        ],
      ),
      body: _loading
          ? const Center(child: CircularProgressIndicator())
          : _error != null
              ? Center(child: Column(
                  mainAxisSize: MainAxisSize.min,
                  children: [
                    const Icon(Icons.error_outline, size: 48, color: AppTheme.textSecondary),
                    const SizedBox(height: 8),
                    Text('불러오기 실패', style: TextStyle(color: AppTheme.textSecondary)),
                    const SizedBox(height: 8),
                    TextButton(onPressed: _load, child: const Text('다시 시도')),
                  ],
                ))
              : _notifications.isEmpty
                  ? const Center(child: Column(
                      mainAxisSize: MainAxisSize.min,
                      children: [
                        Icon(Icons.notifications_none, size: 48, color: AppTheme.textSecondary),
                        SizedBox(height: 8),
                        Text('알림이 없습니다', style: TextStyle(color: AppTheme.textSecondary)),
                      ],
                    ))
                  : RefreshIndicator(
                      onRefresh: _load,
                      child: ListView.builder(
                        itemCount: _notifications.length,
                        itemBuilder: (context, i) {
                          final notif = _notifications[i];
                          final isRead = notif['isRead'] == true;
                          return ListTile(
                            leading: Icon(
                              _notifIcon(notif['type']),
                              color: isRead ? AppTheme.textSecondary : AppTheme.primary,
                            ),
                            title: Text(
                              notif['title'] ?? '',
                              style: TextStyle(fontWeight: isRead ? FontWeight.normal : FontWeight.bold),
                            ),
                            subtitle: Text(notif['body'] ?? '', maxLines: 2, overflow: TextOverflow.ellipsis),
                            tileColor: isRead ? null : AppTheme.primary.withValues(alpha: 0.03),
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
