import 'package:flutter/material.dart';
import 'package:flutter_riverpod/flutter_riverpod.dart';
import '../../shared/providers/app_providers.dart';
import '../../shared/theme/app_theme.dart';

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
      appBar: AppBar(title: const Text('내 거래')),
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
              : _trades.isEmpty
                  ? const Center(child: Column(
                      mainAxisSize: MainAxisSize.min,
                      children: [
                        Icon(Icons.swap_horiz, size: 48, color: AppTheme.textSecondary),
                        SizedBox(height: 8),
                        Text('거래 내역이 없습니다', style: TextStyle(color: AppTheme.textSecondary)),
                      ],
                    ))
                  : RefreshIndicator(
                      onRefresh: _load,
                      child: ListView.builder(
                        itemCount: _trades.length,
                        itemBuilder: (context, i) {
                          final trade = _trades[i];
                          final cp = trade['counterparty'] as Map<String, dynamic>? ?? {};
                          return ListTile(
                            leading: CircleAvatar(
                              backgroundColor: _statusColor(trade['tradeStatus']).withValues(alpha: 0.1),
                              child: Icon(_statusIcon(trade['tradeStatus']), color: _statusColor(trade['tradeStatus']), size: 20),
                            ),
                            title: Text(trade['listingTitle'] ?? ''),
                            subtitle: Text('${cp['nickname'] ?? '상대방'} · ${_statusLabel(trade['tradeStatus'])}'),
                            trailing: Text(_chatStatusLabel(trade['chatStatus']), style: TextStyle(fontSize: 12, color: AppTheme.textSecondary)),
                          );
                        },
                      ),
                    ),
    );
  }

  String _statusLabel(String? status) {
    return switch (status) {
      'available' => '판매중',
      'reserved' => '예약중',
      'completed' => '거래완료',
      'cancelled' => '취소됨',
      _ => status ?? '',
    };
  }

  String _chatStatusLabel(String? status) {
    return switch (status) {
      'open' => '채팅중',
      'reservation_proposed' => '예약제안',
      'reservation_confirmed' => '예약확정',
      'deal_completed' => '완료',
      _ => status ?? '',
    };
  }

  Color _statusColor(String? status) {
    return switch (status) {
      'available' => Colors.green,
      'reserved' => Colors.orange,
      'completed' => Colors.blue,
      'cancelled' => Colors.grey,
      _ => Colors.grey,
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
