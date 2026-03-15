import 'package:flutter/material.dart';
import 'package:flutter_riverpod/flutter_riverpod.dart';
import 'package:go_router/go_router.dart';
import '../../shared/providers/app_providers.dart';
import '../../shared/theme/app_theme.dart';

class MyListingsScreen extends ConsumerStatefulWidget {
  const MyListingsScreen({super.key});

  @override
  ConsumerState<MyListingsScreen> createState() => _MyListingsScreenState();
}

class _MyListingsScreenState extends ConsumerState<MyListingsScreen> {
  List<dynamic> _listings = [];
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
      final res = await api.getMyListings();
      setState(() { _listings = res['data'] ?? []; _loading = false; });
    } catch (e) {
      setState(() { _error = e.toString(); _loading = false; });
    }
  }

  @override
  Widget build(BuildContext context) {
    return Scaffold(
      appBar: AppBar(title: const Text('내 매물')),
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
              : _listings.isEmpty
                  ? const Center(child: Column(
                      mainAxisSize: MainAxisSize.min,
                      children: [
                        Icon(Icons.inbox, size: 48, color: AppTheme.textSecondary),
                        SizedBox(height: 8),
                        Text('등록한 매물이 없습니다', style: TextStyle(color: AppTheme.textSecondary)),
                      ],
                    ))
                  : RefreshIndicator(
                      onRefresh: _load,
                      child: ListView.builder(
                        itemCount: _listings.length,
                        itemBuilder: (context, i) {
                          final item = _listings[i];
                          return ListTile(
                            title: Text(item['title'] ?? ''),
                            subtitle: Text('${item['itemName'] ?? ''} · ${_statusLabel(item['status'])}'),
                            trailing: Text(
                              item['priceAmount'] != null ? '${item['priceAmount']}원' : '가격제안',
                              style: const TextStyle(fontWeight: FontWeight.bold),
                            ),
                            onTap: () => context.push('/listings/${item['listingId']}'),
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
}
