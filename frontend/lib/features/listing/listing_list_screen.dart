import 'package:flutter/material.dart';
import 'package:flutter_riverpod/flutter_riverpod.dart';
import 'package:go_router/go_router.dart';
import '../../shared/providers/app_providers.dart';
import '../../shared/theme/app_theme.dart';
import '../../shared/utils/listing_utils.dart' as utils;

class ListingListScreen extends ConsumerStatefulWidget {
  const ListingListScreen({super.key});

  @override
  ConsumerState<ListingListScreen> createState() => _ListingListScreenState();
}

class _ListingListScreenState extends ConsumerState<ListingListScreen> {
  List<dynamic> _listings = [];
  bool _loading = true;
  String? _selectedServer;
  String _searchQuery = '';
  final _searchController = TextEditingController();

  @override
  void initState() {
    super.initState();
    _loadListings();
  }

  @override
  void dispose() {
    _searchController.dispose();
    super.dispose();
  }

  Future<void> _loadListings() async {
    setState(() => _loading = true);
    try {
      final api = ref.read(apiClientProvider);
      final result = await api.getListings(
        serverId: _selectedServer,
        q: _searchQuery.isNotEmpty ? _searchQuery : null,
      );
      setState(() {
        _listings = result['data'] ?? [];
        _loading = false;
      });
    } catch (e) {
      setState(() => _loading = false);
    }
  }

  @override
  Widget build(BuildContext context) {
    final servers = ref.watch(serversProvider);

    return Scaffold(
      backgroundColor: AppColors.bg,
      appBar: AppBar(
        backgroundColor: AppColors.bg,
        elevation: 0,
        scrolledUnderElevation: 0,
        title: const Text(
          '기란장터',
          style: TextStyle(
            color: AppColors.gold,
            fontSize: 22,
            fontWeight: FontWeight.bold,
            letterSpacing: 1.2,
          ),
        ),
        actions: [
          IconButton(
            icon: const Icon(Icons.search, color: AppColors.gold),
            onPressed: () => _showSearchDialog(),
          ),
        ],
      ),
      body: Column(
        children: [
          // Server filter chips
          servers.when(
            data: (serverList) => Container(
              height: 52,
              decoration: const BoxDecoration(
                border: Border(
                  bottom: BorderSide(color: AppColors.border, width: 0.5),
                ),
              ),
              child: ListView(
                scrollDirection: Axis.horizontal,
                padding: const EdgeInsets.symmetric(horizontal: 12, vertical: 8),
                children: [
                  Padding(
                    padding: const EdgeInsets.only(right: 8),
                    child: _buildFilterChip(
                      label: '전체',
                      selected: _selectedServer == null,
                      onSelected: () {
                        setState(() => _selectedServer = null);
                        _loadListings();
                      },
                    ),
                  ),
                  ...serverList.map((s) => Padding(
                    padding: const EdgeInsets.only(right: 8),
                    child: _buildFilterChip(
                      label: s['serverName'],
                      selected: _selectedServer == s['serverId'],
                      onSelected: () {
                        setState(() => _selectedServer = s['serverId']);
                        _loadListings();
                      },
                    ),
                  )),
                ],
              ),
            ),
            loading: () => const SizedBox(height: 52),
            error: (_, __) => const SizedBox(height: 52),
          ),
          // Listing list
          Expanded(
            child: _loading
                ? Center(
                    child: CircularProgressIndicator(
                      color: AppColors.gold.withValues(alpha: 0.7),
                      strokeWidth: 2,
                    ),
                  )
                : _listings.isEmpty
                    ? _buildEmptyState()
                    : RefreshIndicator(
                        color: AppColors.gold,
                        backgroundColor: AppColors.bgCard,
                        onRefresh: _loadListings,
                        child: ListView.builder(
                          padding: const EdgeInsets.all(12),
                          itemCount: _listings.length,
                          itemBuilder: (context, index) =>
                              _buildListingCard(_listings[index]),
                        ),
                      ),
          ),
        ],
      ),
      floatingActionButton: FloatingActionButton.extended(
        onPressed: () async {
          await context.push('/create-listing');
          _loadListings();
        },
        backgroundColor: AppColors.gold,
        foregroundColor: AppColors.bg,
        icon: const Icon(Icons.add),
        label: const Text(
          '매물 등록',
          style: TextStyle(fontWeight: FontWeight.bold),
        ),
      ),
    );
  }

  Widget _buildFilterChip({
    required String label,
    required bool selected,
    required VoidCallback onSelected,
  }) {
    return GestureDetector(
      onTap: onSelected,
      child: Container(
        padding: const EdgeInsets.symmetric(horizontal: 14, vertical: 6),
        decoration: BoxDecoration(
          color: selected
              ? AppColors.gold.withValues(alpha: 0.15)
              : AppColors.bgCard,
          borderRadius: BorderRadius.circular(20),
          border: Border.all(
            color: selected ? AppColors.gold : AppColors.border,
            width: selected ? 1.5 : 1,
          ),
        ),
        child: Text(
          label,
          style: TextStyle(
            fontSize: 13,
            fontWeight: selected ? FontWeight.w600 : FontWeight.normal,
            color: selected ? AppColors.gold : AppColors.textSecondary,
          ),
        ),
      ),
    );
  }

  Widget _buildListingCard(Map<String, dynamic> listing) {
    final priceText = listing['priceAmount'] != null
        ? '${utils.formatPrice(listing['priceAmount'] is int ? listing['priceAmount'] : int.tryParse(listing['priceAmount'].toString()))}원'
        : '가격 제안';

    final statusText = utils.statusLabel(listing['status']);
    final statusClr = _statusColorDark(listing['status']);
    final hasEnhancement = listing['enhancementLevel'] != null &&
        listing['enhancementLevel'] != 0;

    return Padding(
      padding: const EdgeInsets.only(bottom: 10),
      child: Material(
        color: AppColors.bgCard,
        borderRadius: BorderRadius.circular(12),
        child: InkWell(
          onTap: () => context.push('/listings/${listing['listingId']}'),
          borderRadius: BorderRadius.circular(12),
          splashColor: AppColors.gold.withValues(alpha: 0.05),
          highlightColor: AppColors.gold.withValues(alpha: 0.03),
          child: Container(
            padding: const EdgeInsets.all(14),
            decoration: BoxDecoration(
              borderRadius: BorderRadius.circular(12),
              border: Border.all(color: AppColors.border, width: 0.5),
            ),
            child: Row(
              crossAxisAlignment: CrossAxisAlignment.start,
              children: [
                // LEFT: Item icon with glow
                _buildItemIcon(listing['iconUrl']),
                const SizedBox(width: 12),
                // CENTER: Item info
                Expanded(
                  child: Column(
                    crossAxisAlignment: CrossAxisAlignment.start,
                    children: [
                      // Item name + enhancement badge
                      Row(
                        children: [
                          Expanded(
                            child: Row(
                              children: [
                                Flexible(
                                  child: Text(
                                    listing['itemName'] ?? listing['title'] ?? '',
                                    style: const TextStyle(
                                      fontSize: 15,
                                      fontWeight: FontWeight.bold,
                                      color: AppColors.textPrimary,
                                    ),
                                    maxLines: 1,
                                    overflow: TextOverflow.ellipsis,
                                  ),
                                ),
                                if (hasEnhancement) ...[
                                  const SizedBox(width: 4),
                                  Container(
                                    padding: const EdgeInsets.symmetric(
                                        horizontal: 5, vertical: 1),
                                    decoration: BoxDecoration(
                                      color: AppColors.gold.withValues(alpha: 0.15),
                                      borderRadius: BorderRadius.circular(4),
                                    ),
                                    child: Text(
                                      '+${listing['enhancementLevel']}',
                                      style: const TextStyle(
                                        fontSize: 12,
                                        fontWeight: FontWeight.bold,
                                        color: AppColors.gold,
                                      ),
                                    ),
                                  ),
                                ],
                              ],
                            ),
                          ),
                          // Status chip
                          Container(
                            padding: const EdgeInsets.symmetric(
                                horizontal: 6, vertical: 2),
                            decoration: BoxDecoration(
                              color: statusClr.withValues(alpha: 0.15),
                              borderRadius: BorderRadius.circular(4),
                            ),
                            child: Text(
                              statusText,
                              style: TextStyle(
                                fontSize: 10,
                                fontWeight: FontWeight.w600,
                                color: statusClr,
                              ),
                            ),
                          ),
                        ],
                      ),
                      const SizedBox(height: 4),
                      // Category + server subtitle
                      Text(
                        [
                          if (listing['categoryName'] != null)
                            listing['categoryName'],
                          if (listing['serverName'] != null)
                            listing['serverName'],
                        ].join(' / '),
                        style: const TextStyle(
                          fontSize: 12,
                          color: AppColors.textSecondary,
                        ),
                        maxLines: 1,
                        overflow: TextOverflow.ellipsis,
                      ),
                      const SizedBox(height: 8),
                      // Price row + meta
                      Row(
                        children: [
                          Text(
                            priceText,
                            style: const TextStyle(
                              fontSize: 17,
                              fontWeight: FontWeight.bold,
                              color: AppColors.gold,
                            ),
                          ),
                          if (listing['priceType'] == 'negotiable')
                            const Text(
                              ' (협상)',
                              style: TextStyle(
                                fontSize: 11,
                                color: AppColors.textSecondary,
                              ),
                            ),
                          const Spacer(),
                          _metaChip(Icons.remove_red_eye_outlined,
                              '${listing['viewCount'] ?? 0}'),
                          const SizedBox(width: 6),
                          _metaChip(Icons.favorite_border,
                              '${listing['favoriteCount'] ?? 0}'),
                          const SizedBox(width: 6),
                          _metaChip(Icons.chat_bubble_outline,
                              '${listing['chatCount'] ?? 0}'),
                        ],
                      ),
                      const SizedBox(height: 6),
                      // Author + time
                      Row(
                        children: [
                          Text(
                            listing['author']?['nickname'] ?? '',
                            style: const TextStyle(
                              fontSize: 11,
                              color: AppColors.textSecondary,
                            ),
                          ),
                          const Spacer(),
                          Text(
                            utils.formatTimeAgo(listing['createdAt']),
                            style: TextStyle(
                              fontSize: 11,
                              color: AppColors.textSecondary.withValues(alpha: 0.7),
                            ),
                          ),
                        ],
                      ),
                    ],
                  ),
                ),
              ],
            ),
          ),
        ),
      ),
    );
  }

  Widget _buildItemIcon(String? iconUrl) {
    return Container(
      width: 40,
      height: 40,
      decoration: BoxDecoration(
        color: AppColors.bgSurface,
        borderRadius: BorderRadius.circular(8),
        boxShadow: [
          BoxShadow(
            color: AppColors.gold.withValues(alpha: 0.15),
            blurRadius: 8,
            spreadRadius: 1,
          ),
        ],
      ),
      child: iconUrl != null
          ? ClipRRect(
              borderRadius: BorderRadius.circular(8),
              child: Image.network(
                'http://localhost:8080$iconUrl',
                width: 40,
                height: 40,
                fit: BoxFit.cover,
                errorBuilder: (_, __, ___) => const Icon(
                  Icons.shield_outlined,
                  size: 20,
                  color: AppColors.textSecondary,
                ),
              ),
            )
          : const Icon(
              Icons.shield_outlined,
              size: 20,
              color: AppColors.textSecondary,
            ),
    );
  }

  Widget _metaChip(IconData icon, String count) {
    return Row(
      mainAxisSize: MainAxisSize.min,
      children: [
        Icon(icon, size: 12, color: AppColors.textSecondary.withValues(alpha: 0.6)),
        const SizedBox(width: 2),
        Text(
          count,
          style: TextStyle(
            fontSize: 11,
            color: AppColors.textSecondary.withValues(alpha: 0.6),
          ),
        ),
      ],
    );
  }

  Widget _buildEmptyState() {
    return Center(
      child: Column(
        mainAxisSize: MainAxisSize.min,
        children: [
          Icon(
            Icons.inventory_2_outlined,
            size: 64,
            color: AppColors.textSecondary.withValues(alpha: 0.3),
          ),
          const SizedBox(height: 16),
          const Text(
            '매물이 없습니다',
            style: TextStyle(
              fontSize: 16,
              fontWeight: FontWeight.w500,
              color: AppColors.textSecondary,
            ),
          ),
          const SizedBox(height: 8),
          Text(
            '첫 매물을 등록해보세요!',
            style: TextStyle(
              fontSize: 14,
              color: AppColors.textSecondary.withValues(alpha: 0.6),
            ),
          ),
        ],
      ),
    );
  }

  void _showSearchDialog() {
    showDialog(
      context: context,
      builder: (ctx) => AlertDialog(
        backgroundColor: AppColors.bgSurface,
        shape: RoundedRectangleBorder(
          borderRadius: BorderRadius.circular(16),
          side: const BorderSide(color: AppColors.border),
        ),
        title: const Text(
          '매물 검색',
          style: TextStyle(color: AppColors.textPrimary),
        ),
        content: TextField(
          controller: _searchController,
          style: const TextStyle(color: AppColors.textPrimary),
          cursorColor: AppColors.gold,
          decoration: InputDecoration(
            hintText: '아이템명 또는 제목',
            hintStyle: TextStyle(
              color: AppColors.textSecondary.withValues(alpha: 0.5),
            ),
            filled: true,
            fillColor: AppColors.bgCard,
            border: OutlineInputBorder(
              borderRadius: BorderRadius.circular(10),
              borderSide: const BorderSide(color: AppColors.border),
            ),
            enabledBorder: OutlineInputBorder(
              borderRadius: BorderRadius.circular(10),
              borderSide: const BorderSide(color: AppColors.border),
            ),
            focusedBorder: OutlineInputBorder(
              borderRadius: BorderRadius.circular(10),
              borderSide: const BorderSide(color: AppColors.gold, width: 1.5),
            ),
          ),
          autofocus: true,
          onSubmitted: (_) {
            _searchQuery = _searchController.text;
            Navigator.pop(ctx);
            _loadListings();
          },
        ),
        actions: [
          TextButton(
            onPressed: () {
              _searchController.clear();
              _searchQuery = '';
              Navigator.pop(ctx);
              _loadListings();
            },
            child: const Text(
              '초기화',
              style: TextStyle(color: AppColors.textSecondary),
            ),
          ),
          TextButton(
            onPressed: () {
              _searchQuery = _searchController.text;
              Navigator.pop(ctx);
              _loadListings();
            },
            child: const Text(
              '검색',
              style: TextStyle(color: AppColors.gold),
            ),
          ),
        ],
      ),
    );
  }

  Color _statusColorDark(String? status) {
    return switch (status) {
      'available' => AppColors.success,
      'reserved' => const Color(0xFFE67E22),
      'pending_trade' => AppColors.blue,
      'completed' => AppColors.textSecondary,
      'cancelled' => AppColors.error,
      _ => AppColors.textSecondary,
    };
  }
}
