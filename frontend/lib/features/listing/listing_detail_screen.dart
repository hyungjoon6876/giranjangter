import 'package:flutter/material.dart';
import 'package:flutter_riverpod/flutter_riverpod.dart';
import 'package:go_router/go_router.dart';
import '../../shared/providers/app_providers.dart';
import '../../shared/theme/app_theme.dart';
import '../../shared/utils/listing_utils.dart' as utils;

class ListingDetailScreen extends ConsumerStatefulWidget {
  final String listingId;
  const ListingDetailScreen({super.key, required this.listingId});

  @override
  ConsumerState<ListingDetailScreen> createState() => _ListingDetailScreenState();
}

class _ListingDetailScreenState extends ConsumerState<ListingDetailScreen> {
  Map<String, dynamic>? _listing;
  bool _loading = true;

  @override
  void initState() {
    super.initState();
    _load();
  }

  Future<void> _load() async {
    try {
      final api = ref.read(apiClientProvider);
      final data = await api.getListing(widget.listingId);
      setState(() { _listing = data; _loading = false; });
    } catch (e) {
      setState(() => _loading = false);
    }
  }

  @override
  Widget build(BuildContext context) {
    if (_loading) return const Scaffold(body: Center(child: CircularProgressIndicator()));
    if (_listing == null) return Scaffold(appBar: AppBar(), body: const Center(child: Text('매물을 찾을 수 없습니다')));

    final l = _listing!;
    final priceText = l['priceAmount'] != null ? '${utils.formatPrice(l['priceAmount'] is int ? l['priceAmount'] : int.tryParse(l['priceAmount'].toString()))}원' : '가격 제안';

    return Scaffold(
      appBar: AppBar(
        title: Text(l['serverName'] ?? ''),
        actions: [
          if (l['availableActions']?.contains('report') ?? false)
            IconButton(icon: const Icon(Icons.flag_outlined), onPressed: () {}),
        ],
      ),
      body: SingleChildScrollView(
        padding: const EdgeInsets.all(16),
        child: Column(
          crossAxisAlignment: CrossAxisAlignment.start,
          children: [
            // Status badge
            Row(
              children: [
                _badge(l['listingType'] == 'sell' ? '판매' : '구매',
                    l['listingType'] == 'sell' ? AppTheme.primary : AppTheme.secondary),
                const SizedBox(width: 6),
                _badge(utils.statusLabel(l['status']), utils.statusColor(l['status'])),
                const Spacer(),
                if (l['tradeMethod'] != null)
                  Text(_tradeMethodLabel(l['tradeMethod']),
                      style: const TextStyle(fontSize: 12, color: AppTheme.textSecondary)),
              ],
            ),
            const SizedBox(height: 12),
            // Title
            Text(l['title'] ?? '', style: Theme.of(context).textTheme.headlineMedium),
            const SizedBox(height: 8),
            // Item info
            Row(
              children: [
                Text(l['itemName'] ?? '', style: const TextStyle(fontSize: 16)),
                if (l['enhancementLevel'] != null)
                  Text(' +${l['enhancementLevel']}',
                      style: const TextStyle(fontSize: 16, color: AppTheme.primary, fontWeight: FontWeight.bold)),
              ],
            ),
            if (l['optionsText'] != null) ...[
              const SizedBox(height: 4),
              Text(l['optionsText'], style: const TextStyle(color: AppTheme.textSecondary)),
            ],
            const SizedBox(height: 16),
            // Price
            Text(priceText, style: const TextStyle(fontSize: 28, fontWeight: FontWeight.bold)),
            if (l['priceType'] == 'negotiable')
              const Text('협상 가능', style: TextStyle(color: AppTheme.textSecondary)),
            const Divider(height: 32),
            // Description
            Text(l['description'] ?? '', style: const TextStyle(fontSize: 15, height: 1.6)),
            const Divider(height: 32),
            // Trade info
            _infoRow('거래 방식', _tradeMethodLabel(l['tradeMethod'])),
            if (l['preferredMeetingAreaText'] != null)
              _infoRow('접선 장소', l['preferredMeetingAreaText']),
            if (l['availableTimeText'] != null)
              _infoRow('거래 가능 시간', l['availableTimeText']),
            _infoRow('수량', '${l['quantity']}개'),
            const Divider(height: 32),
            // Author
            _buildAuthorSection(l['author']),
            const SizedBox(height: 80),
          ],
        ),
      ),
      bottomNavigationBar: _buildBottomBar(l),
    );
  }

  Widget _buildAuthorSection(Map<String, dynamic>? author) {
    if (author == null) return const SizedBox();
    return Row(
      children: [
        CircleAvatar(
          radius: 20,
          backgroundColor: AppTheme.border,
          child: Text(
            (author['nickname'] ?? '?')[0],
            style: const TextStyle(fontWeight: FontWeight.bold),
          ),
        ),
        const SizedBox(width: 12),
        Column(
          crossAxisAlignment: CrossAxisAlignment.start,
          children: [
            Text(author['nickname'] ?? '', style: const TextStyle(fontWeight: FontWeight.w600)),
            Text(
              '거래 ${author['completedTradeCount'] ?? 0}회',
              style: const TextStyle(fontSize: 12, color: AppTheme.textSecondary),
            ),
          ],
        ),
      ],
    );
  }

  Widget _buildBottomBar(Map<String, dynamic> l) {
    final actions = (l['availableActions'] as List?)?.cast<String>() ?? [];
    return Container(
      padding: const EdgeInsets.all(16),
      decoration: const BoxDecoration(
        color: Colors.white,
        border: Border(top: BorderSide(color: AppTheme.border)),
      ),
      child: Row(
        children: [
          if (actions.contains('favorite'))
            IconButton(
              icon: Icon(
                l['isFavorited'] == true ? Icons.favorite : Icons.favorite_border,
                color: l['isFavorited'] == true ? AppTheme.error : null,
              ),
              onPressed: () async {
                final api = ref.read(apiClientProvider);
                if (l['isFavorited'] == true) {
                  await api.unfavoriteListing(widget.listingId);
                } else {
                  await api.favoriteListing(widget.listingId);
                }
                _load();
              },
            ),
          const SizedBox(width: 8),
          if (actions.contains('start_chat'))
            Expanded(
              child: ElevatedButton(
                onPressed: () async {
                  final api = ref.read(apiClientProvider);
                  try {
                    final chat = await api.createChat(widget.listingId);
                    if (context.mounted) {
                      context.push('/chats/${chat['chatRoomId']}');
                    }
                  } catch (e) {
                    if (context.mounted) {
                      ScaffoldMessenger.of(context).showSnackBar(
                        const SnackBar(content: Text('채팅을 시작할 수 없습니다')),
                      );
                    }
                  }
                },
                child: const Text('채팅하기'),
              ),
            ),
        ],
      ),
    );
  }

  Widget _badge(String text, Color color) {
    return Container(
      padding: const EdgeInsets.symmetric(horizontal: 8, vertical: 3),
      decoration: BoxDecoration(
        color: color.withValues(alpha: 0.1),
        borderRadius: BorderRadius.circular(4),
      ),
      child: Text(text, style: TextStyle(fontSize: 12, fontWeight: FontWeight.w600, color: color)),
    );
  }

  Widget _infoRow(String label, String value) {
    return Padding(
      padding: const EdgeInsets.symmetric(vertical: 6),
      child: Row(
        children: [
          SizedBox(width: 100, child: Text(label, style: const TextStyle(color: AppTheme.textSecondary))),
          Expanded(child: Text(value)),
        ],
      ),
    );
  }

  String _tradeMethodLabel(String? m) => {'in_game': '인게임', 'offline_pc_bang': 'PC방/오프라인', 'either': '무관'}[m] ?? m ?? '';
}
