import 'package:flutter/material.dart';
import '../theme/app_theme.dart';

String statusLabel(String? status) {
  return switch (status) {
    'available' => '판매중',
    'reserved' => '예약중',
    'pending_trade' => '거래중',
    'completed' => '거래완료',
    'cancelled' => '취소됨',
    _ => status ?? '',
  };
}

Color statusColor(String? status) {
  return switch (status) {
    'available' => AppTheme.secondary,
    'reserved' => AppTheme.warning,
    'pending_trade' => AppTheme.primary,
    'completed' => AppTheme.textSecondary,
    'cancelled' => AppTheme.error,
    _ => AppTheme.textSecondary,
  };
}

String chatStatusLabel(String? status) {
  return switch (status) {
    'open' => '채팅중',
    'reservation_proposed' => '예약제안',
    'reservation_confirmed' => '예약확정',
    'deal_completed' => '거래완료',
    _ => status ?? '',
  };
}

String formatPrice(int? amount) {
  if (amount == null) return '0';
  if (amount >= 100000000) return '${(amount / 100000000).toStringAsFixed(1)}억';
  if (amount >= 10000) return '${(amount / 10000).toStringAsFixed(0)}만';
  return amount.toString();
}

String formatTimeAgo(String? dateStr) {
  if (dateStr == null) return '';
  final date = DateTime.tryParse(dateStr);
  if (date == null) return '';
  final diff = DateTime.now().difference(date);
  if (diff.inMinutes < 1) return '방금';
  if (diff.inHours < 1) return '${diff.inMinutes}분 전';
  if (diff.inDays < 1) return '${diff.inHours}시간 전';
  if (diff.inDays < 30) return '${diff.inDays}일 전';
  return '${(diff.inDays / 30).floor()}개월 전';
}
