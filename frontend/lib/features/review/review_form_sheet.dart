import 'package:flutter/material.dart';
import 'package:flutter_riverpod/flutter_riverpod.dart';
import '../../shared/providers/app_providers.dart';
import '../../shared/theme/app_theme.dart';

class ReviewFormSheet extends ConsumerStatefulWidget {
  final String completionId;
  final VoidCallback onSubmitted;

  const ReviewFormSheet({super.key, required this.completionId, required this.onSubmitted});

  @override
  ConsumerState<ReviewFormSheet> createState() => _ReviewFormSheetState();
}

class _ReviewFormSheetState extends ConsumerState<ReviewFormSheet> {
  String _rating = 'positive';
  final _commentCtrl = TextEditingController();
  bool _submitting = false;

  @override
  Widget build(BuildContext context) {
    return Padding(
      padding: EdgeInsets.only(
        left: 24, right: 24, top: 24,
        bottom: MediaQuery.of(context).viewInsets.bottom + 24,
      ),
      child: Column(
        mainAxisSize: MainAxisSize.min,
        crossAxisAlignment: CrossAxisAlignment.start,
        children: [
          const Text('후기 작성', style: TextStyle(fontSize: 20, fontWeight: FontWeight.bold)),
          const SizedBox(height: 16),

          const Text('이 거래는 어땠나요?', style: TextStyle(fontSize: 15)),
          const SizedBox(height: 12),

          Row(
            children: [
              Expanded(
                child: _ratingButton(
                  icon: Icons.thumb_up,
                  label: '추천',
                  value: 'positive',
                  color: AppTheme.secondary,
                ),
              ),
              const SizedBox(width: 12),
              Expanded(
                child: _ratingButton(
                  icon: Icons.thumb_down,
                  label: '비추천',
                  value: 'negative',
                  color: AppTheme.error,
                ),
              ),
            ],
          ),
          const SizedBox(height: 16),

          TextField(
            controller: _commentCtrl,
            decoration: const InputDecoration(
              labelText: '한줄 후기 (선택)',
              hintText: '예: 빠른 거래 감사합니다!',
            ),
            maxLines: 2,
            maxLength: 500,
          ),
          const SizedBox(height: 16),

          SizedBox(
            width: double.infinity,
            child: ElevatedButton(
              onPressed: _submitting ? null : _submit,
              child: _submitting
                  ? const SizedBox(width: 20, height: 20, child: CircularProgressIndicator(strokeWidth: 2, color: Colors.white))
                  : const Text('후기 등록'),
            ),
          ),
        ],
      ),
    );
  }

  Widget _ratingButton({
    required IconData icon,
    required String label,
    required String value,
    required Color color,
  }) {
    final selected = _rating == value;
    return OutlinedButton.icon(
      onPressed: () => setState(() => _rating = value),
      icon: Icon(icon, color: selected ? Colors.white : color),
      label: Text(label, style: TextStyle(color: selected ? Colors.white : color)),
      style: OutlinedButton.styleFrom(
        backgroundColor: selected ? color : null,
        side: BorderSide(color: color),
        padding: const EdgeInsets.symmetric(vertical: 14),
      ),
    );
  }

  Future<void> _submit() async {
    setState(() => _submitting = true);
    try {
      final api = ref.read(apiClientProvider);
      await api.createReview(widget.completionId, {
        'rating': _rating,
        'comment': _commentCtrl.text.isNotEmpty ? _commentCtrl.text : null,
      });
      if (mounted) {
        Navigator.pop(context);
        widget.onSubmitted();
        ScaffoldMessenger.of(context).showSnackBar(const SnackBar(content: Text('후기가 등록되었습니다!')));
      }
    } catch (e) {
      if (mounted) {
        ScaffoldMessenger.of(context).showSnackBar(SnackBar(content: Text('후기 등록 실패: $e')));
      }
    } finally {
      if (mounted) setState(() => _submitting = false);
    }
  }
}
