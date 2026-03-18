import 'package:flutter/material.dart';
import 'package:flutter_test/flutter_test.dart';
import 'package:flutter_riverpod/flutter_riverpod.dart';
import 'package:lincle/shared/providers/app_providers.dart';
import 'package:lincle/features/listing/listing_list_screen.dart';
import '../helpers/mock_api_client.dart';

void main() {
  group('ListingListScreen', () {
    Widget buildTestWidget() {
      return ProviderScope(
        overrides: [
          apiClientProvider.overrideWithValue(MockApiClient()),
        ],
        child: const MaterialApp(home: ListingListScreen()),
      );
    }

    testWidgets('renders app bar with logo image', (tester) async {
      await tester.pumpWidget(buildTestWidget());
      await tester.pumpAndSettle();

      // AppBar contains an Image.asset with height 32
      final images = find.byType(Image);
      expect(images, findsOneWidget);
    });

    testWidgets('shows FAB with "매물 등록" label', (tester) async {
      await tester.pumpWidget(buildTestWidget());
      await tester.pumpAndSettle();

      expect(find.byType(FloatingActionButton), findsOneWidget);
      expect(find.text('매물 등록'), findsOneWidget);
      expect(find.byIcon(Icons.add), findsOneWidget);
    });

    testWidgets('shows search icon in app bar', (tester) async {
      await tester.pumpWidget(buildTestWidget());
      await tester.pumpAndSettle();

      expect(find.byIcon(Icons.search), findsOneWidget);
    });

    testWidgets('shows empty state when no listings', (tester) async {
      await tester.pumpWidget(buildTestWidget());
      // Pump once to build, then settle for async _loadListings to complete
      await tester.pumpAndSettle();

      // MockApiClient returns empty list, so empty state is shown
      expect(find.text('매물이 없습니다'), findsOneWidget);
      expect(find.text('첫 매물을 등록해보세요!'), findsOneWidget);
    });
  });
}
