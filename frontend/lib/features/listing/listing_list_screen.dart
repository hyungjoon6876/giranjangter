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
      appBar: AppBar(
        title: const Text('린클 거래소'),
        actions: [
          IconButton(
            icon: const Icon(Icons.search),
            onPressed: () => _showSearchDialog(),
          ),
        ],
      ),
      body: Column(
        children: [
          // Server filter chips
          servers.when(
            data: (serverList) => SizedBox(
              height: 48,
              child: ListView(
                scrollDirection: Axis.horizontal,
                padding: const EdgeInsets.symmetric(horizontal: 12),
                children: [
                  Padding(
                    padding: const EdgeInsets.only(right: 8),
                    child: FilterChip(
                      label: const Text('전체'),
                      selected: _selectedServer == null,
                      onSelected: (_) {
                        setState(() => _selectedServer = null);
                        _loadListings();
                      },
                    ),
                  ),
                  ...serverList.map((s) => Padding(
                    padding: const EdgeInsets.only(right: 8),
                    child: FilterChip(
                      label: Text(s['serverName']),
                      selected: _selectedServer == s['serverId'],
                      onSelected: (_) {
                        setState(() => _selectedServer = s['serverId']);
                        _loadListings();
                      },
                    ),
                  )),
                ],
              ),
            ),
            loading: () => const SizedBox(height: 48),
            error: (_, __) => const SizedBox(height: 48),
          ),
          // Listing list
          Expanded(
            child: _loading
                ? const Center(child: CircularProgressIndicator())
                : _listings.isEmpty
                    ? _buildEmptyState()
                    : RefreshIndicator(
                        onRefresh: _loadListings,
                        child: ListView.builder(
                          padding: const EdgeInsets.all(12),
                          itemCount: _listings.length,
                          itemBuilder: (context, index) => _buildListingCard(_listings[index]),
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
        icon: const Icon(Icons.add),
        label: const Text('매물 등록'),
      ),
    );
  }

  Widget _buildListingCard(Map<String, dynamic> listing) {
    final priceText = listing['priceAmount'] != null
        ? '${utils.formatPrice(listing['priceAmount'] is int ? listing['priceAmount'] : int.tryParse(listing['priceAmount'].toString()))}원'
        : '가격 제안';
    final sc = utils.statusColor(listing['status']);

    return Card(
      margin: const EdgeInsets.only(bottom: 10),
      child: InkWell(
        onTap: () => context.push('/listings/${listing['listingId']}'),
        borderRadius: BorderRadius.circular(12),
        child: Padding(
          padding: const EdgeInsets.all(14),
          child: Column(
            crossAxisAlignment: CrossAxisAlignment.start,
            children: [
              Row(
                children: [
                  Container(
                    padding: const EdgeInsets.symmetric(horizontal: 8, vertical: 2),
                    decoration: BoxDecoration(
                      color: listing['listingType'] == 'sell'
                          ? AppTheme.primary.withValues(alpha: 0.1)
                          : AppTheme.secondary.withValues(alpha: 0.1),
                      borderRadius: BorderRadius.circular(4),
                    ),
                    child: Text(
                      listing['listingType'] == 'sell' ? '판매' : '구매',
                      style: TextStyle(
                        fontSize: 11,
                        fontWeight: FontWeight.w600,
                        color: listing['listingType'] == 'sell'
                            ? AppTheme.primary
                            : AppTheme.secondary,
                      ),
                    ),
                  ),
                  const SizedBox(width: 6),
                  Container(
                    padding: const EdgeInsets.symmetric(horizontal: 6, vertical: 2),
                    decoration: BoxDecoration(
                      color: sc.withValues(alpha: 0.1),
                      borderRadius: BorderRadius.circular(4),
                    ),
                    child: Text(
                      utils.statusLabel(listing['status']),
                      style: TextStyle(fontSize: 10, color: sc, fontWeight: FontWeight.w600),
                    ),
                  ),
                  const Spacer(),
                  Text(
                    listing['serverName'] ?? '',
                    style: const TextStyle(fontSize: 12, color: AppTheme.textSecondary),
                  ),
                ],
              ),
              const SizedBox(height: 8),
              Text(
                listing['title'] ?? '',
                style: const TextStyle(fontSize: 16, fontWeight: FontWeight.w600),
                maxLines: 1,
                overflow: TextOverflow.ellipsis,
              ),
              const SizedBox(height: 4),
              Row(
                children: [
                  if (listing['iconUrl'] != null) ...[
                    Image.network(
                      'http://localhost:8080${listing['iconUrl']}',
                      width: 20,
                      height: 20,
                      errorBuilder: (_, __, ___) => const SizedBox(width: 20, height: 20),
                    ),
                    const SizedBox(width: 4),
                  ],
                  Text(
                    listing['itemName'] ?? '',
                    style: const TextStyle(fontSize: 13, color: AppTheme.textSecondary),
                  ),
                  if (listing['enhancementLevel'] != null) ...[
                    Text(
                      ' +${listing['enhancementLevel']}',
                      style: const TextStyle(fontSize: 13, color: AppTheme.primary, fontWeight: FontWeight.w600),
                    ),
                  ],
                ],
              ),
              const SizedBox(height: 8),
              Row(
                children: [
                  Text(
                    priceText,
                    style: const TextStyle(fontSize: 18, fontWeight: FontWeight.bold),
                  ),
                  if (listing['priceType'] == 'negotiable')
                    const Text(' (협상가능)', style: TextStyle(fontSize: 12, color: AppTheme.textSecondary)),
                  const Spacer(),
                  _metaChip(Icons.remove_red_eye_outlined, '${listing['viewCount']}'),
                  const SizedBox(width: 8),
                  _metaChip(Icons.favorite_border, '${listing['favoriteCount']}'),
                  const SizedBox(width: 8),
                  _metaChip(Icons.chat_bubble_outline, '${listing['chatCount']}'),
                ],
              ),
              const SizedBox(height: 6),
              Row(
                children: [
                  Text(
                    listing['author']?['nickname'] ?? '',
                    style: const TextStyle(fontSize: 12, color: AppTheme.textSecondary),
                  ),
                  const Spacer(),
                  Text(
                    utils.formatTimeAgo(listing['createdAt']),
                    style: const TextStyle(fontSize: 11, color: AppTheme.textSecondary),
                  ),
                ],
              ),
            ],
          ),
        ),
      ),
    );
  }

  Widget _metaChip(IconData icon, String count) {
    return Row(
      mainAxisSize: MainAxisSize.min,
      children: [
        Icon(icon, size: 14, color: AppTheme.textSecondary),
        const SizedBox(width: 2),
        Text(count, style: const TextStyle(fontSize: 12, color: AppTheme.textSecondary)),
      ],
    );
  }

  Widget _buildEmptyState() {
    return Center(
      child: Column(
        mainAxisSize: MainAxisSize.min,
        children: [
          Icon(Icons.search_off, size: 64, color: Colors.grey[300]),
          const SizedBox(height: 16),
          const Text('매물이 없습니다', style: TextStyle(fontSize: 16, color: AppTheme.textSecondary)),
          const SizedBox(height: 8),
          const Text('첫 매물을 등록해보세요!', style: TextStyle(fontSize: 14, color: AppTheme.textSecondary)),
        ],
      ),
    );
  }

  void _showSearchDialog() {
    showDialog(
      context: context,
      builder: (ctx) => AlertDialog(
        title: const Text('매물 검색'),
        content: TextField(
          controller: _searchController,
          decoration: const InputDecoration(hintText: '아이템명 또는 제목'),
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
            child: const Text('초기화'),
          ),
          TextButton(
            onPressed: () {
              _searchQuery = _searchController.text;
              Navigator.pop(ctx);
              _loadListings();
            },
            child: const Text('검색'),
          ),
        ],
      ),
    );
  }

}
