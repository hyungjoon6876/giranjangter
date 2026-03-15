import 'package:flutter/material.dart';
import 'package:flutter_riverpod/flutter_riverpod.dart';
import '../../shared/providers/app_providers.dart';

class TradeCompleteSheet extends ConsumerStatefulWidget {
  final String listingId;
  final String reservationId;
  final VoidCallback onCompleted;

  const TradeCompleteSheet({
    super.key,
    required this.listingId,
    required this.reservationId,
    required this.onCompleted,
  });

  @override
  ConsumerState<TradeCompleteSheet> createState() => _TradeCompleteSheetState();
}

class _TradeCompleteSheetState extends ConsumerState<TradeCompleteSheet> {
  bool _submitting = false;

  @override
  Widget build(BuildContext context) {
    return Padding(
      padding: const EdgeInsets.all(24),
      child: Column(
        mainAxisSize: MainAxisSize.min,
        children: [
          const Icon(Icons.check_circle_outline, size: 64, color: Colors.green),
          const SizedBox(height: 16),
          const Text('거래 완료 요청', style: TextStyle(fontSize: 20, fontWeight: FontWeight.bold)),
          const SizedBox(height: 8),
          const Text(
            '거래가 정상적으로 완료되었나요?\n상대방에게 확인 요청이 전송됩니다.',
            textAlign: TextAlign.center,
            style: TextStyle(color: Colors.grey),
          ),
          const SizedBox(height: 24),
          Row(
            children: [
              Expanded(
                child: OutlinedButton(
                  onPressed: () => Navigator.pop(context),
                  child: const Text('취소'),
                ),
              ),
              const SizedBox(width: 12),
              Expanded(
                child: ElevatedButton(
                  onPressed: _submitting ? null : _submit,
                  child: _submitting
                      ? const SizedBox(width: 20, height: 20, child: CircularProgressIndicator(strokeWidth: 2, color: Colors.white))
                      : const Text('완료 요청'),
                ),
              ),
            ],
          ),
        ],
      ),
    );
  }

  Future<void> _submit() async {
    setState(() => _submitting = true);
    try {
      final api = ref.read(apiClientProvider);
      await api.dio.post('/listings/${widget.listingId}/complete', data: {
        'reservationId': widget.reservationId,
      });
      if (mounted) {
        Navigator.pop(context);
        widget.onCompleted();
        ScaffoldMessenger.of(context).showSnackBar(
          const SnackBar(content: Text('거래 완료가 요청되었습니다. 상대방 확인을 기다려주세요.')),
        );
      }
    } catch (e) {
      if (mounted) {
        ScaffoldMessenger.of(context).showSnackBar(SnackBar(content: Text('완료 요청 실패: $e')));
      }
    } finally {
      if (mounted) setState(() => _submitting = false);
    }
  }
}
