import 'package:flutter/material.dart';
import 'package:flutter_riverpod/flutter_riverpod.dart';
import 'package:go_router/go_router.dart';
import '../../shared/providers/app_providers.dart';
import '../../shared/theme/app_theme.dart';

class ChatListScreen extends ConsumerStatefulWidget {
  const ChatListScreen({super.key});

  @override
  ConsumerState<ChatListScreen> createState() => _ChatListScreenState();
}

class _ChatListScreenState extends ConsumerState<ChatListScreen> {
  List<dynamic> _chats = [];
  bool _loading = true;

  @override
  void initState() {
    super.initState();
    _load();
  }

  Future<void> _load() async {
    try {
      final api = ref.read(apiClientProvider);
      final data = await api.getChats();
      setState(() { _chats = data['data'] ?? []; _loading = false; });
    } catch (e) {
      setState(() => _loading = false);
    }
  }

  @override
  Widget build(BuildContext context) {
    return Scaffold(
      appBar: AppBar(title: const Text('채팅')),
      body: _loading
          ? const Center(child: CircularProgressIndicator())
          : _chats.isEmpty
              ? const Center(child: Text('채팅이 없습니다', style: TextStyle(color: AppTheme.textSecondary)))
              : RefreshIndicator(
                  onRefresh: _load,
                  child: ListView.separated(
                    itemCount: _chats.length,
                    separatorBuilder: (_, __) => const Divider(height: 1),
                    itemBuilder: (context, i) => _buildChatTile(_chats[i]),
                  ),
                ),
    );
  }

  Widget _buildChatTile(Map<String, dynamic> chat) {
    final cp = chat['counterparty'] as Map<String, dynamic>?;

    return ListTile(
      leading: CircleAvatar(
        backgroundColor: AppTheme.primary.withValues(alpha: 0.1),
        child: Text(
          (cp?['nickname'] ?? '?')[0],
          style: const TextStyle(fontWeight: FontWeight.bold, color: AppTheme.primary),
        ),
      ),
      title: Text(cp?['nickname'] ?? '', style: const TextStyle(fontWeight: FontWeight.w600)),
      subtitle: Text(chat['listingTitle'] ?? '', maxLines: 1, overflow: TextOverflow.ellipsis),
      trailing: Column(
        mainAxisAlignment: MainAxisAlignment.center,
        crossAxisAlignment: CrossAxisAlignment.end,
        children: [
          Text(
            _statusLabel(chat['chatStatus']),
            style: TextStyle(fontSize: 11, color: _statusColor(chat['chatStatus'])),
          ),
        ],
      ),
      onTap: () => context.push('/chats/${chat['chatRoomId']}'),
    );
  }

  String _statusLabel(String? s) => {
    'open': '대화중', 'reservation_proposed': '예약제안',
    'reservation_confirmed': '예약확정', 'deal_completed': '거래완료',
  }[s] ?? s ?? '';

  Color _statusColor(String? s) => {
    'open': AppTheme.secondary, 'reservation_proposed': AppTheme.warning,
    'reservation_confirmed': AppTheme.primary, 'deal_completed': AppTheme.textSecondary,
  }[s] ?? AppTheme.textSecondary;
}
