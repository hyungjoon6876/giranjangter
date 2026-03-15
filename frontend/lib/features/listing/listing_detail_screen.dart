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
    if (_loading) {
      return Scaffold(
        backgroundColor: AppColors.bg,
        body: Center(
          child: CircularProgressIndicator(
            color: AppColors.gold.withValues(alpha: 0.7),
            strokeWidth: 2,
          ),
        ),
      );
    }
    if (_listing == null) {
      return Scaffold(
        backgroundColor: AppColors.bg,
        appBar: AppBar(
          backgroundColor: AppColors.bg,
          foregroundColor: AppColors.textPrimary,
        ),
        body: const Center(
          child: Text(
            '매물을 찾을 수 없습니다',
            style: TextStyle(color: AppColors.textSecondary),
          ),
        ),
      );
    }

    final l = _listing!;
    final priceText = l['priceAmount'] != null
        ? '${utils.formatPrice(l['priceAmount'] is int ? l['priceAmount'] : int.tryParse(l['priceAmount'].toString()))}원'
        : '가격 제안';
    final hasEnhancement = l['enhancementLevel'] != null && l['enhancementLevel'] != 0;

    return Scaffold(
      backgroundColor: AppColors.bg,
      appBar: AppBar(
        backgroundColor: AppColors.bg,
        elevation: 0,
        scrolledUnderElevation: 0,
        foregroundColor: AppColors.textPrimary,
        title: Text(
          l['serverName'] ?? '',
          style: const TextStyle(
            color: AppColors.textSecondary,
            fontSize: 16,
          ),
        ),
        actions: [
          if (l['availableActions']?.contains('report') ?? false)
            IconButton(
              icon: const Icon(Icons.flag_outlined, color: AppColors.textSecondary),
              onPressed: () {},
            ),
        ],
      ),
      body: SingleChildScrollView(
        padding: const EdgeInsets.all(16),
        child: Column(
          crossAxisAlignment: CrossAxisAlignment.start,
          children: [
            // Item header section
            _buildItemHeader(l, hasEnhancement),
            const SizedBox(height: 20),

            // Info section (game tooltip style)
            _buildInfoCard(l, priceText, hasEnhancement),
            const SizedBox(height: 12),

            // Description section
            if (l['description'] != null && (l['description'] as String).isNotEmpty)
              _buildDescriptionCard(l['description']),
            if (l['description'] != null && (l['description'] as String).isNotEmpty)
              const SizedBox(height: 12),

            // Author section
            _buildAuthorCard(l['author']),
            const SizedBox(height: 80),
          ],
        ),
      ),
      bottomNavigationBar: _buildBottomBar(l),
    );
  }

  Widget _buildItemHeader(Map<String, dynamic> l, bool hasEnhancement) {
    return Container(
      padding: const EdgeInsets.all(16),
      decoration: BoxDecoration(
        color: AppColors.bgCard,
        borderRadius: BorderRadius.circular(12),
        border: Border.all(color: AppColors.border, width: 0.5),
      ),
      child: Row(
        crossAxisAlignment: CrossAxisAlignment.start,
        children: [
          // Large item icon with gold glow
          Container(
            width: 64,
            height: 64,
            decoration: BoxDecoration(
              color: AppColors.bgSurface,
              borderRadius: BorderRadius.circular(12),
              boxShadow: [
                BoxShadow(
                  color: AppColors.gold.withValues(alpha: 0.2),
                  blurRadius: 12,
                  spreadRadius: 2,
                ),
              ],
            ),
            child: l['iconUrl'] != null
                ? ClipRRect(
                    borderRadius: BorderRadius.circular(12),
                    child: Image.network(
                      'http://localhost:8080${l['iconUrl']}',
                      width: 64,
                      height: 64,
                      fit: BoxFit.cover,
                      errorBuilder: (_, __, ___) => const Icon(
                        Icons.shield_outlined,
                        size: 32,
                        color: AppColors.textSecondary,
                      ),
                    ),
                  )
                : const Icon(
                    Icons.shield_outlined,
                    size: 32,
                    color: AppColors.textSecondary,
                  ),
          ),
          const SizedBox(width: 16),
          // Item name + meta
          Expanded(
            child: Column(
              crossAxisAlignment: CrossAxisAlignment.start,
              children: [
                Row(
                  children: [
                    Flexible(
                      child: Text(
                        l['itemName'] ?? l['title'] ?? '',
                        style: const TextStyle(
                          fontSize: 20,
                          fontWeight: FontWeight.bold,
                          color: AppColors.textPrimary,
                        ),
                      ),
                    ),
                    if (hasEnhancement) ...[
                      const SizedBox(width: 6),
                      Container(
                        padding: const EdgeInsets.symmetric(
                            horizontal: 6, vertical: 2),
                        decoration: BoxDecoration(
                          color: AppColors.gold.withValues(alpha: 0.15),
                          borderRadius: BorderRadius.circular(4),
                          border: Border.all(
                            color: AppColors.gold.withValues(alpha: 0.3),
                          ),
                        ),
                        child: Text(
                          '+${l['enhancementLevel']}',
                          style: const TextStyle(
                            fontSize: 14,
                            fontWeight: FontWeight.bold,
                            color: AppColors.gold,
                          ),
                        ),
                      ),
                    ],
                  ],
                ),
                const SizedBox(height: 6),
                // Category + server
                Text(
                  [
                    if (l['categoryName'] != null) l['categoryName'],
                    if (l['serverName'] != null) l['serverName'],
                  ].join(' / '),
                  style: const TextStyle(
                    fontSize: 13,
                    color: AppColors.textSecondary,
                  ),
                ),
                const SizedBox(height: 8),
                // Status badges
                Row(
                  children: [
                    _badge(
                      l['listingType'] == 'sell' ? '판매' : '구매',
                      l['listingType'] == 'sell'
                          ? AppColors.gold
                          : AppColors.blue,
                    ),
                    const SizedBox(width: 6),
                    _badge(
                      utils.statusLabel(l['status']),
                      _statusColorDark(l['status']),
                    ),
                  ],
                ),
              ],
            ),
          ),
        ],
      ),
    );
  }

  Widget _buildInfoCard(Map<String, dynamic> l, String priceText, bool hasEnhancement) {
    return Container(
      decoration: BoxDecoration(
        color: AppColors.bgCard,
        borderRadius: BorderRadius.circular(12),
        border: Border.all(color: AppColors.border, width: 0.5),
      ),
      child: Column(
        children: [
          _tooltipRow('가격', priceText, valueColor: AppColors.gold, isFirst: true),
          if (l['priceType'] == 'negotiable')
            _tooltipRow('가격유형', '협상 가능'),
          _tooltipRow('거래방식', _tradeMethodLabel(l['tradeMethod'])),
          if (hasEnhancement)
            _tooltipRow('강화수치', '+${l['enhancementLevel']}',
                valueColor: AppColors.gold),
          if (l['optionsText'] != null)
            _tooltipRow('옵션', l['optionsText']),
          _tooltipRow('수량', '${l['quantity'] ?? 1}개'),
          if (l['preferredMeetingAreaText'] != null)
            _tooltipRow('접선 장소', l['preferredMeetingAreaText']),
          if (l['availableTimeText'] != null)
            _tooltipRow('거래 시간', l['availableTimeText'], isLast: true),
        ],
      ),
    );
  }

  Widget _tooltipRow(
    String label,
    String value, {
    Color? valueColor,
    bool isFirst = false,
    bool isLast = false,
  }) {
    return Container(
      padding: const EdgeInsets.symmetric(horizontal: 16, vertical: 12),
      decoration: BoxDecoration(
        border: isLast
            ? null
            : const Border(
                bottom: BorderSide(
                  color: AppColors.border,
                  width: 0.5,
                ),
              ),
      ),
      child: Row(
        children: [
          SizedBox(
            width: 80,
            child: Text(
              label,
              style: const TextStyle(
                fontSize: 13,
                color: AppColors.textSecondary,
              ),
            ),
          ),
          Expanded(
            child: Text(
              value,
              style: TextStyle(
                fontSize: 14,
                fontWeight: valueColor != null ? FontWeight.bold : FontWeight.normal,
                color: valueColor ?? AppColors.textPrimary,
              ),
            ),
          ),
        ],
      ),
    );
  }

  Widget _buildDescriptionCard(String description) {
    return Container(
      width: double.infinity,
      padding: const EdgeInsets.all(16),
      decoration: BoxDecoration(
        color: AppColors.bgCard,
        borderRadius: BorderRadius.circular(12),
        border: Border.all(color: AppColors.border, width: 0.5),
      ),
      child: Column(
        crossAxisAlignment: CrossAxisAlignment.start,
        children: [
          const Text(
            '설명',
            style: TextStyle(
              fontSize: 13,
              fontWeight: FontWeight.w600,
              color: AppColors.textSecondary,
            ),
          ),
          const SizedBox(height: 8),
          Text(
            description,
            style: const TextStyle(
              fontSize: 14,
              height: 1.6,
              color: AppColors.textPrimary,
            ),
          ),
        ],
      ),
    );
  }

  Widget _buildAuthorCard(Map<String, dynamic>? author) {
    if (author == null) return const SizedBox();
    final tradeCount = author['completedTradeCount'] ?? 0;

    return Container(
      padding: const EdgeInsets.all(16),
      decoration: BoxDecoration(
        color: AppColors.bgCard,
        borderRadius: BorderRadius.circular(12),
        border: Border.all(color: AppColors.border, width: 0.5),
      ),
      child: Row(
        children: [
          CircleAvatar(
            radius: 22,
            backgroundColor: AppColors.bgSurface,
            child: Text(
              (author['nickname'] ?? '?')[0],
              style: const TextStyle(
                fontWeight: FontWeight.bold,
                color: AppColors.gold,
                fontSize: 16,
              ),
            ),
          ),
          const SizedBox(width: 12),
          Expanded(
            child: Column(
              crossAxisAlignment: CrossAxisAlignment.start,
              children: [
                Row(
                  children: [
                    Text(
                      author['nickname'] ?? '',
                      style: const TextStyle(
                        fontWeight: FontWeight.w600,
                        fontSize: 15,
                        color: AppColors.textPrimary,
                      ),
                    ),
                    if (tradeCount >= 5) ...[
                      const SizedBox(width: 6),
                      Container(
                        padding: const EdgeInsets.symmetric(
                            horizontal: 6, vertical: 2),
                        decoration: BoxDecoration(
                          color: AppColors.success.withValues(alpha: 0.15),
                          borderRadius: BorderRadius.circular(4),
                        ),
                        child: Row(
                          mainAxisSize: MainAxisSize.min,
                          children: [
                            Icon(
                              Icons.verified,
                              size: 12,
                              color: AppColors.success.withValues(alpha: 0.8),
                            ),
                            const SizedBox(width: 2),
                            Text(
                              '신뢰',
                              style: TextStyle(
                                fontSize: 10,
                                fontWeight: FontWeight.w600,
                                color: AppColors.success.withValues(alpha: 0.8),
                              ),
                            ),
                          ],
                        ),
                      ),
                    ],
                  ],
                ),
                const SizedBox(height: 2),
                Text(
                  '거래 ${tradeCount}회 완료',
                  style: const TextStyle(
                    fontSize: 12,
                    color: AppColors.textSecondary,
                  ),
                ),
              ],
            ),
          ),
        ],
      ),
    );
  }

  Widget _buildBottomBar(Map<String, dynamic> l) {
    final actions = (l['availableActions'] as List?)?.cast<String>() ?? [];

    return Container(
      padding: const EdgeInsets.all(16),
      decoration: const BoxDecoration(
        color: AppColors.bgCard,
        border: Border(
          top: BorderSide(color: AppColors.border, width: 0.5),
        ),
      ),
      child: SafeArea(
        child: Row(
          children: [
            if (actions.contains('favorite'))
              Container(
                decoration: BoxDecoration(
                  border: Border.all(color: AppColors.border),
                  borderRadius: BorderRadius.circular(10),
                ),
                child: IconButton(
                  icon: Icon(
                    l['isFavorited'] == true
                        ? Icons.favorite
                        : Icons.favorite_border,
                    color: l['isFavorited'] == true
                        ? AppColors.error
                        : AppColors.textSecondary,
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
              ),
            const SizedBox(width: 10),
            if (actions.contains('start_chat'))
              Expanded(
                child: SizedBox(
                  height: 48,
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
                            SnackBar(
                              content: const Text('채팅을 시작할 수 없습니다'),
                              backgroundColor: AppColors.error,
                            ),
                          );
                        }
                      }
                    },
                    style: ElevatedButton.styleFrom(
                      backgroundColor: AppColors.gold,
                      foregroundColor: AppColors.bg,
                      shape: RoundedRectangleBorder(
                        borderRadius: BorderRadius.circular(10),
                      ),
                      elevation: 0,
                    ),
                    child: const Row(
                      mainAxisAlignment: MainAxisAlignment.center,
                      children: [
                        Icon(Icons.chat_bubble_outline, size: 18),
                        SizedBox(width: 6),
                        Text(
                          '채팅 시작',
                          style: TextStyle(
                            fontWeight: FontWeight.bold,
                            fontSize: 15,
                          ),
                        ),
                      ],
                    ),
                  ),
                ),
              ),
          ],
        ),
      ),
    );
  }

  Widget _badge(String text, Color color) {
    return Container(
      padding: const EdgeInsets.symmetric(horizontal: 8, vertical: 3),
      decoration: BoxDecoration(
        color: color.withValues(alpha: 0.15),
        borderRadius: BorderRadius.circular(4),
      ),
      child: Text(
        text,
        style: TextStyle(
          fontSize: 12,
          fontWeight: FontWeight.w600,
          color: color,
        ),
      ),
    );
  }

  Color _statusColorDark(String? status) {
    return switch (status) {
      'available' => AppColors.success,
      'reserved' => const Color(0xFFE67E22),
      'pending_trade' => AppColors.blue,
      'completed' => AppColors.textSecondary,
      'cancelled' => AppColors.error,
      _ => AppColors.textSecondary,
    };
  }

  String _tradeMethodLabel(String? m) =>
      {'in_game': '인게임', 'offline_pc_bang': 'PC방/오프라인', 'either': '무관'}[m] ??
      m ??
      '';
}
