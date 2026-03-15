import 'package:flutter/material.dart';
import 'package:flutter_riverpod/flutter_riverpod.dart';
import '../../shared/providers/app_providers.dart';
import '../../shared/theme/app_theme.dart';
import '../../shared/utils/listing_utils.dart' as utils;

class MyTradesScreen extends ConsumerStatefulWidget {
  const MyTradesScreen({super.key});

  @override
  ConsumerState<MyTradesScreen> createState() => _MyTradesScreenState();
}

class _MyTradesScreenState extends ConsumerState<MyTradesScreen> {
  List<dynamic> _trades = [];
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
      final res = await api.getMyTrades();
      setState(() { _trades = res['data'] ?? []; _loading = false; });
    } catch (e) {
      setState(() { _error = e.toString(); _loading = false; });
    }
  }

  @override
  Widget build(BuildContext context) {
    return Scaffold(
      backgroundColor: AppColors.bg,
      appBar: AppBar(
        title: const Text('내 거래'),
        backgroundColor: AppColors.bgSurface,
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
              : _trades.isEmpty
                  ? const Center(child: Column(
                      mainAxisSize: MainAxisSize.min,
                      children: [
                        Icon(Icons.swap_horiz, size: 48, color: AppColors.textMuted),
                        SizedBox(height: 8),
                        Text('거래 내역이 없습니다', style: TextStyle(color: AppColors.textSecondary)),
                      ],
                    ))
                  : RefreshIndicator(
                      color: AppColors.gold,
                      backgroundColor: AppColors.bgCard,
                      onRefresh: _load,
                      child: ListView.builder(
                        padding: const EdgeInsets.symmetric(vertical: 8, horizontal: 12),
                        itemCount: _trades.length,
                        itemBuilder: (context, i) {
                          final trade = _trades[i];
                          final cp = trade['counterparty'] as Map<String, dynamic>? ?? {};
                          final statusColor = _tradeStatusColor(trade['tradeStatus']);
                          return Padding(
                            padding: const EdgeInsets.only(bottom: 8),
                            child: Container(
                              padding: const EdgeInsets.all(14),
                              decoration: BoxDecoration(
                                color: AppColors.bgCard,
                                borderRadius: BorderRadius.circular(12),
                                border: Border.all(color: AppColors.border),
                              ),
                              child: Row(
                                children: [
                                  // Status icon
                                  Container(
                                    width: 40,
                                    height: 40,
                                    decoration: BoxDecoration(
                                      color: statusColor.withValues(alpha: 0.15),
                                      shape: BoxShape.circle,
                                    ),
                                    child: Icon(
                                      _statusIcon(trade['tradeStatus']),
                                      color: statusColor,
                                      size: 20,
                                    ),
                                  ),
                                  const SizedBox(width: 12),

                                  // Content
                                  Expanded(
                                    child: Column(
                                      crossAxisAlignment: CrossAxisAlignment.start,
                                      children: [
                                        Text(
                                          trade['listingTitle'] ?? '',
                                          style: const TextStyle(
                                            color: AppColors.textPrimary,
                                            fontWeight: FontWeight.w600,
                                            fontSize: 15,
                                          ),
                                        ),
                                        const SizedBox(height: 4),
                                        Text(
                                          '${cp['nickname'] ?? '상대방'} \u00B7 ${utils.statusLabel(trade['tradeStatus'])}',
                                          style: const TextStyle(
                                            color: AppColors.textSecondary,
                                            fontSize: 13,
                                          ),
                                        ),
                                      ],
                                    ),
                                  ),

                                  // Chat status
                                  Container(
                                    padding: const EdgeInsets.symmetric(horizontal: 8, vertical: 4),
                                    decoration: BoxDecoration(
                                      color: AppColors.bgElevated,
                                      borderRadius: BorderRadius.circular(8),
                                    ),
                                    child: Text(
                                      utils.chatStatusLabel(trade['chatStatus']),
                                      style: const TextStyle(
                                        fontSize: 11,
                                        color: AppColors.textSecondary,
                                        fontWeight: FontWeight.w500,
                                      ),
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

  Color _tradeStatusColor(String? status) {
    return switch (status) {
      'available' => AppColors.success,
      'reserved' => AppColors.warning,
      'completed' => AppColors.textSecondary,
      'cancelled' => AppColors.error,
      _ => AppColors.textSecondary,
    };
  }

  IconData _statusIcon(String? status) {
    return switch (status) {
      'available' => Icons.storefront,
      'reserved' => Icons.schedule,
      'completed' => Icons.check_circle,
      'cancelled' => Icons.cancel,
      _ => Icons.swap_horiz,
    };
  }
}
