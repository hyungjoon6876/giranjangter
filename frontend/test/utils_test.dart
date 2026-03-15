import 'package:flutter/material.dart';
import 'package:flutter_test/flutter_test.dart';
import 'package:lincle/shared/utils/listing_utils.dart';
import 'package:lincle/shared/theme/app_theme.dart';

void main() {
  group('formatPrice', () {
    test('returns "0" for null', () {
      expect(formatPrice(null), '0');
    });

    test('returns plain number for values under 10,000', () {
      expect(formatPrice(0), '0');
      expect(formatPrice(1), '1');
      expect(formatPrice(999), '999');
      expect(formatPrice(9999), '9999');
    });

    test('formats values >= 10,000 as 만 (ten-thousands)', () {
      expect(formatPrice(10000), '1만');
      expect(formatPrice(50000), '5만');
      expect(formatPrice(99999), '10만'); // rounds to nearest 만
    });

    test('formats values >= 100,000,000 as 억 (hundred-millions)', () {
      expect(formatPrice(100000000), '1.0억');
      expect(formatPrice(250000000), '2.5억');
      expect(formatPrice(1000000000), '10.0억');
    });

    test('boundary: 10000 uses 만, 99999999 uses 만, 100000000 uses 억', () {
      expect(formatPrice(10000), '1만');
      expect(formatPrice(99999999), '10000만');
      expect(formatPrice(100000000), '1.0억');
    });
  });

  group('statusLabel', () {
    test('maps "available" to "판매중"', () {
      expect(statusLabel('available'), '판매중');
    });

    test('maps "reserved" to "예약중"', () {
      expect(statusLabel('reserved'), '예약중');
    });

    test('maps "pending_trade" to "거래중"', () {
      expect(statusLabel('pending_trade'), '거래중');
    });

    test('maps "completed" to "거래완료"', () {
      expect(statusLabel('completed'), '거래완료');
    });

    test('maps "cancelled" to "취소됨"', () {
      expect(statusLabel('cancelled'), '취소됨');
    });

    test('returns the status itself for unknown values', () {
      expect(statusLabel('unknown_status'), 'unknown_status');
    });

    test('returns empty string for null', () {
      expect(statusLabel(null), '');
    });
  });

  group('statusColor', () {
    test('returns success (green) for "available"', () {
      expect(statusColor('available'), AppColors.success);
    });

    test('returns warning (orange) for "reserved"', () {
      expect(statusColor('reserved'), AppColors.warning);
    });

    test('returns gold for "pending_trade"', () {
      expect(statusColor('pending_trade'), AppColors.gold);
    });

    test('returns textSecondary (grey) for "completed"', () {
      expect(statusColor('completed'), AppColors.textSecondary);
    });

    test('returns error (red) for "cancelled"', () {
      expect(statusColor('cancelled'), AppColors.error);
    });

    test('returns textSecondary for unknown status', () {
      expect(statusColor('something_else'), AppColors.textSecondary);
    });

    test('returns textSecondary for null', () {
      expect(statusColor(null), AppColors.textSecondary);
    });
  });

  group('chatStatusLabel', () {
    test('maps "open" to "채팅중"', () {
      expect(chatStatusLabel('open'), '채팅중');
    });

    test('maps "reservation_proposed" to "예약제안"', () {
      expect(chatStatusLabel('reservation_proposed'), '예약제안');
    });

    test('maps "reservation_confirmed" to "예약확정"', () {
      expect(chatStatusLabel('reservation_confirmed'), '예약확정');
    });

    test('maps "deal_completed" to "거래완료"', () {
      expect(chatStatusLabel('deal_completed'), '거래완료');
    });

    test('returns empty string for null', () {
      expect(chatStatusLabel(null), '');
    });
  });

  group('formatTimeAgo', () {
    test('returns empty string for null', () {
      expect(formatTimeAgo(null), '');
    });

    test('returns empty string for unparseable date', () {
      expect(formatTimeAgo('not-a-date'), '');
    });

    test('returns "방금" for times less than 1 minute ago', () {
      final now = DateTime.now().subtract(const Duration(seconds: 30));
      expect(formatTimeAgo(now.toIso8601String()), '방금');
    });

    test('returns "N분 전" for times less than 1 hour ago', () {
      final thirtyMinsAgo = DateTime.now().subtract(const Duration(minutes: 30));
      expect(formatTimeAgo(thirtyMinsAgo.toIso8601String()), '30분 전');
    });

    test('returns "N시간 전" for times less than 1 day ago', () {
      final fiveHoursAgo = DateTime.now().subtract(const Duration(hours: 5));
      expect(formatTimeAgo(fiveHoursAgo.toIso8601String()), '5시간 전');
    });

    test('returns "N일 전" for times less than 30 days ago', () {
      final tenDaysAgo = DateTime.now().subtract(const Duration(days: 10));
      expect(formatTimeAgo(tenDaysAgo.toIso8601String()), '10일 전');
    });

    test('returns "N개월 전" for times 30+ days ago', () {
      final sixtyDaysAgo = DateTime.now().subtract(const Duration(days: 60));
      expect(formatTimeAgo(sixtyDaysAgo.toIso8601String()), '2개월 전');
    });
  });
}
