import 'package:flutter/material.dart';
import 'package:flutter_riverpod/flutter_riverpod.dart';
import '../../shared/providers/app_providers.dart';
import '../../shared/theme/app_theme.dart';

class ReservationFormSheet extends ConsumerStatefulWidget {
  final String chatId;
  final VoidCallback onCreated;

  const ReservationFormSheet({super.key, required this.chatId, required this.onCreated});

  @override
  ConsumerState<ReservationFormSheet> createState() => _ReservationFormSheetState();
}

class _ReservationFormSheetState extends ConsumerState<ReservationFormSheet> {
  DateTime _selectedDate = DateTime.now().add(const Duration(days: 1));
  TimeOfDay _selectedTime = const TimeOfDay(hour: 14, minute: 0);
  String _meetingType = 'in_game';
  final _meetingPointCtrl = TextEditingController();
  final _noteCtrl = TextEditingController();
  bool _submitting = false;

  @override
  Widget build(BuildContext context) {
    return Padding(
      padding: EdgeInsets.only(
        left: 16, right: 16, top: 16,
        bottom: MediaQuery.of(context).viewInsets.bottom + 16,
      ),
      child: Column(
        mainAxisSize: MainAxisSize.min,
        crossAxisAlignment: CrossAxisAlignment.start,
        children: [
          Row(
            children: [
              const Text('예약 제안', style: TextStyle(fontSize: 18, fontWeight: FontWeight.bold)),
              const Spacer(),
              IconButton(icon: const Icon(Icons.close), onPressed: () => Navigator.pop(context)),
            ],
          ),
          const SizedBox(height: 12),

          // Date picker
          ListTile(
            contentPadding: EdgeInsets.zero,
            leading: const Icon(Icons.calendar_today, color: AppColors.gold),
            title: Text('${_selectedDate.year}-${_selectedDate.month.toString().padLeft(2, '0')}-${_selectedDate.day.toString().padLeft(2, '0')}'),
            subtitle: const Text('거래 날짜'),
            onTap: () async {
              final date = await showDatePicker(
                context: context,
                initialDate: _selectedDate,
                firstDate: DateTime.now(),
                lastDate: DateTime.now().add(const Duration(days: 30)),
              );
              if (date != null) setState(() => _selectedDate = date);
            },
          ),

          // Time picker
          ListTile(
            contentPadding: EdgeInsets.zero,
            leading: const Icon(Icons.access_time, color: AppColors.gold),
            title: Text('${_selectedTime.hour.toString().padLeft(2, '0')}:${_selectedTime.minute.toString().padLeft(2, '0')}'),
            subtitle: const Text('거래 시간'),
            onTap: () async {
              final time = await showTimePicker(context: context, initialTime: _selectedTime);
              if (time != null) setState(() => _selectedTime = time);
            },
          ),

          const SizedBox(height: 8),

          // Meeting type
          DropdownButtonFormField<String>(
            decoration: const InputDecoration(labelText: '거래 방식'),
            value: _meetingType,
            items: const [
              DropdownMenuItem(value: 'in_game', child: Text('인게임')),
              DropdownMenuItem(value: 'offline_pc_bang', child: Text('PC방/오프라인')),
              DropdownMenuItem(value: 'either', child: Text('무관')),
            ],
            onChanged: (v) => setState(() => _meetingType = v ?? 'in_game'),
          ),
          const SizedBox(height: 12),

          TextField(
            controller: _meetingPointCtrl,
            decoration: const InputDecoration(
              labelText: '접선 장소',
              hintText: '예: 기란마을 분수대',
            ),
          ),
          const SizedBox(height: 12),

          TextField(
            controller: _noteCtrl,
            decoration: const InputDecoration(
              labelText: '메모 (선택)',
              hintText: '상대방에게 전할 메시지',
            ),
          ),
          const SizedBox(height: 16),

          SizedBox(
            width: double.infinity,
            child: ElevatedButton(
              onPressed: _submitting ? null : _submit,
              child: _submitting
                  ? const SizedBox(width: 20, height: 20, child: CircularProgressIndicator(strokeWidth: 2, color: Colors.white))
                  : const Text('예약 제안하기'),
            ),
          ),
        ],
      ),
    );
  }

  Future<void> _submit() async {
    setState(() => _submitting = true);
    try {
      final api = ref.read(apiClientProvider);
      final scheduledAt = DateTime(
        _selectedDate.year, _selectedDate.month, _selectedDate.day,
        _selectedTime.hour, _selectedTime.minute,
      ).toUtc().toIso8601String();

      await api.createReservation(widget.chatId, {
        'scheduledAt': scheduledAt,
        'meetingType': _meetingType,
        'meetingPointText': _meetingPointCtrl.text.isNotEmpty ? _meetingPointCtrl.text : null,
        'noteToCounterparty': _noteCtrl.text.isNotEmpty ? _noteCtrl.text : null,
      });

      if (mounted) {
        Navigator.pop(context);
        widget.onCreated();
        ScaffoldMessenger.of(context).showSnackBar(
          const SnackBar(content: Text('예약이 제안되었습니다!')),
        );
      }
    } catch (e) {
      if (mounted) {
        ScaffoldMessenger.of(context).showSnackBar(
          SnackBar(content: Text('예약 제안 실패: $e')),
        );
      }
    } finally {
      if (mounted) setState(() => _submitting = false);
    }
  }
}
