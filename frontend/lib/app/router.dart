import 'package:flutter/material.dart';
import 'package:go_router/go_router.dart';
import '../features/auth/login_screen.dart';
import '../features/listing/listing_list_screen.dart';
import '../features/listing/listing_detail_screen.dart';
import '../features/listing/listing_create_screen.dart';
import '../features/chat/chat_list_screen.dart';
import '../features/chat/chat_detail_screen.dart';
import '../features/profile/profile_screen.dart';
import '../shared/theme/app_theme.dart';

final appRouter = GoRouter(
  initialLocation: '/',
  routes: [
    StatefulShellRoute.indexedStack(
      builder: (context, state, navigationShell) {
        return ScaffoldWithNavBar(navigationShell: navigationShell);
      },
      branches: [
        // Home / Listings
        StatefulShellBranch(routes: [
          GoRoute(
            path: '/',
            builder: (context, state) => const ListingListScreen(),
          ),
        ]),
        // Chat
        StatefulShellBranch(routes: [
          GoRoute(
            path: '/chats',
            builder: (context, state) => const ChatListScreen(),
          ),
        ]),
        // Profile
        StatefulShellBranch(routes: [
          GoRoute(
            path: '/profile',
            builder: (context, state) => const ProfileScreen(),
          ),
        ]),
      ],
    ),
    // Detail routes (no bottom nav)
    // IMPORTANT: /listings/create must come before /listings/:id
    GoRoute(
      path: '/create-listing',
      builder: (context, state) => const ListingCreateScreen(),
    ),
    GoRoute(
      path: '/listings/:id',
      builder: (context, state) => ListingDetailScreen(
        listingId: state.pathParameters['id']!,
      ),
    ),
    GoRoute(
      path: '/chats/:chatId',
      builder: (context, state) => ChatDetailScreen(
        chatId: state.pathParameters['chatId']!,
      ),
    ),
    GoRoute(
      path: '/login',
      builder: (context, state) => const LoginScreen(),
    ),
  ],
);

class ScaffoldWithNavBar extends StatelessWidget {
  const ScaffoldWithNavBar({super.key, required this.navigationShell});

  final StatefulNavigationShell navigationShell;

  @override
  Widget build(BuildContext context) {
    return Scaffold(
      body: navigationShell,
      bottomNavigationBar: NavigationBar(
        backgroundColor: AppColors.bgCard,
        indicatorColor: AppColors.gold.withValues(alpha: 0.15),
        selectedIndex: navigationShell.currentIndex,
        onDestinationSelected: (index) {
          navigationShell.goBranch(index, initialLocation: index == navigationShell.currentIndex);
        },
        destinations: const [
          NavigationDestination(
            icon: Icon(Icons.store, color: AppColors.textSecondary),
            selectedIcon: Icon(Icons.store, color: AppColors.gold),
            label: '거래소',
          ),
          NavigationDestination(
            icon: Icon(Icons.chat_bubble_outline, color: AppColors.textSecondary),
            selectedIcon: Icon(Icons.chat_bubble, color: AppColors.gold),
            label: '채팅',
          ),
          NavigationDestination(
            icon: Icon(Icons.person_outline, color: AppColors.textSecondary),
            selectedIcon: Icon(Icons.person, color: AppColors.gold),
            label: '내 정보',
          ),
        ],
        labelBehavior: NavigationDestinationLabelBehavior.alwaysShow,
      ),
    );
  }
}
