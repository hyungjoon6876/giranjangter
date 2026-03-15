import 'package:flutter/material.dart';
import 'package:flutter_riverpod/flutter_riverpod.dart';
import '../../shared/providers/app_providers.dart';
import '../../shared/theme/app_theme.dart';

class ReportFormSheet extends ConsumerStatefulWidget {
  final String targetType; // listing, user, chat_room, message, review
  final String targetId;

  const ReportFormSheet({super.key, required this.targetType, required this.targetId});

  @override
  ConsumerState<ReportFormSheet> createState() => _ReportFormSheetState();
}

class _ReportFormSheetState extends ConsumerState<ReportFormSheet> {
  String? _reportType;
  final _descCtrl = TextEditingController();
  bool _submitting = false;

  static const _reportTypes = [
    {'value': 'fake_listing', 'label': '허위매물', 'icon': Icons.warning_amber},
    {'value': 'scam_suspicion', 'label': '사기 의심', 'icon': Icons.security},
    {'value': 'no_show', 'label': '노쇼', 'icon': Icons.person_off},
    {'value': 'harassment', 'label': '욕설/괴롭힘', 'icon': Icons.mood_bad},
    {'value': 'spam', 'label': '스팸/광고', 'icon': Icons.mark_email_unread},
    {'value': 'prohibited_item', 'label': '금지 품목', 'icon': Icons.block},
    {'value': 'privacy_exposure', 'label': '개인정보 노출', 'icon': Icons.privacy_tip},
    {'value': 'other', 'label': '기타', 'icon': Icons.more_horiz},
  ];

  @override
  Widget build(BuildContext context) {
    return DraggableScrollableSheet(
      initialChildSize: 0.7,
      maxChildSize: 0.9,
      minChildSize: 0.5,
      builder: (context, scrollCtrl) => Container(
        decoration: const BoxDecoration(
          color: Colors.white,
          borderRadius: BorderRadius.vertical(top: Radius.circular(16)),
        ),
        child: ListView(
          controller: scrollCtrl,
          padding: const EdgeInsets.all(24),
          children: [
            Row(
              children: [
                const Text('신고하기', style: TextStyle(fontSize: 20, fontWeight: FontWeight.bold)),
                const Spacer(),
                IconButton(icon: const Icon(Icons.close), onPressed: () => Navigator.pop(context)),
              ],
            ),
            const SizedBox(height: 8),
            const Text('신고 사유를 선택해주세요', style: TextStyle(color: AppTheme.textSecondary)),
            const SizedBox(height: 16),

            // Report type grid
            ...(_reportTypes.map((rt) => RadioListTile<String>(
              title: Row(
                children: [
                  Icon(rt['icon'] as IconData, size: 20, color: AppTheme.textSecondary),
                  const SizedBox(width: 8),
                  Text(rt['label'] as String),
                ],
              ),
              value: rt['value'] as String,
              groupValue: _reportType,
              onChanged: (v) => setState(() => _reportType = v),
              contentPadding: EdgeInsets.zero,
              dense: true,
            ))),

            const SizedBox(height: 16),
            TextField(
              controller: _descCtrl,
              decoration: const InputDecoration(
                labelText: '상세 설명 *',
                hintText: '신고 내용을 자세히 설명해주세요',
              ),
              maxLines: 4,
              maxLength: 1000,
            ),
            const SizedBox(height: 16),

            SizedBox(
              width: double.infinity,
              child: ElevatedButton(
                onPressed: (_reportType == null || _descCtrl.text.length < 10 || _submitting) ? null : _submit,
                style: ElevatedButton.styleFrom(backgroundColor: AppTheme.error),
                child: _submitting
                    ? const SizedBox(width: 20, height: 20, child: CircularProgressIndicator(strokeWidth: 2, color: Colors.white))
                    : const Text('신고 접수'),
              ),
            ),
            const SizedBox(height: 8),
            const Text(
              '허위 신고는 제재 대상이 될 수 있습니다.',
              textAlign: TextAlign.center,
              style: TextStyle(fontSize: 12, color: AppTheme.textSecondary),
            ),
          ],
        ),
      ),
    );
  }

  Future<void> _submit() async {
    if (_reportType == null || _descCtrl.text.length < 10) return;
    setState(() => _submitting = true);
    try {
      final api = ref.read(apiClientProvider);
      await api.createReport({
        'targetType': widget.targetType,
        'targetId': widget.targetId,
        'reportType': _reportType,
        'description': _descCtrl.text,
      });
      if (mounted) {
        Navigator.pop(context);
        ScaffoldMessenger.of(context).showSnackBar(
          const SnackBar(content: Text('신고가 접수되었습니다. 검토 후 처리됩니다.')),
        );
      }
    } catch (e) {
      if (mounted) {
        ScaffoldMessenger.of(context).showSnackBar(SnackBar(content: Text('신고 실패: $e')));
      }
    } finally {
      if (mounted) setState(() => _submitting = false);
    }
  }
}
