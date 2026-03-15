import 'dart:async';
import 'package:cached_network_image/cached_network_image.dart';
import 'package:dio/dio.dart';
import 'package:flutter/material.dart';
import 'package:flutter_riverpod/flutter_riverpod.dart';
import 'package:go_router/go_router.dart';
import '../../shared/providers/app_providers.dart';
import '../../shared/theme/app_theme.dart';

class ListingCreateScreen extends ConsumerStatefulWidget {
  const ListingCreateScreen({super.key});

  @override
  ConsumerState<ListingCreateScreen> createState() => _ListingCreateScreenState();
}

class _ListingCreateScreenState extends ConsumerState<ListingCreateScreen> {
  final _formKey = GlobalKey<FormState>();
  String _listingType = 'sell';
  String? _serverId;
  String? _categoryId;
  final _titleCtrl = TextEditingController();
  final _itemNameCtrl = TextEditingController();
  final _descCtrl = TextEditingController();
  final _priceCtrl = TextEditingController();
  String _priceType = 'fixed';
  String _tradeMethod = 'either';
  final _enhancementCtrl = TextEditingController();
  bool _submitting = false;

  // Item autocomplete state
  Timer? _debounce;
  List<dynamic> _itemSuggestions = [];
  Map<String, dynamic>? _selectedItem;
  final _itemFocusNode = FocusNode();
  final _layerLink = LayerLink();
  OverlayEntry? _overlayEntry;

  @override
  void dispose() {
    _debounce?.cancel();
    _removeOverlay();
    _itemFocusNode.dispose();
    _titleCtrl.dispose();
    _itemNameCtrl.dispose();
    _descCtrl.dispose();
    _priceCtrl.dispose();
    _enhancementCtrl.dispose();
    super.dispose();
  }

  InputDecoration _darkInput(String label, {String? hint}) {
    return InputDecoration(
      labelText: label,
      hintText: hint,
      labelStyle: const TextStyle(color: AppColors.textSecondary),
      hintStyle: TextStyle(
        color: AppColors.textSecondary.withValues(alpha: 0.4),
      ),
      filled: true,
      fillColor: AppColors.bgSurface,
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
      errorBorder: OutlineInputBorder(
        borderRadius: BorderRadius.circular(10),
        borderSide: const BorderSide(color: AppColors.error),
      ),
      focusedErrorBorder: OutlineInputBorder(
        borderRadius: BorderRadius.circular(10),
        borderSide: const BorderSide(color: AppColors.error, width: 1.5),
      ),
      contentPadding: const EdgeInsets.symmetric(horizontal: 16, vertical: 14),
    );
  }

  @override
  Widget build(BuildContext context) {
    final servers = ref.watch(serversProvider);
    final categories = ref.watch(categoriesProvider);

    return Scaffold(
      backgroundColor: AppColors.bg,
      appBar: AppBar(
        backgroundColor: AppColors.bg,
        elevation: 0,
        scrolledUnderElevation: 0,
        foregroundColor: AppColors.textPrimary,
        title: const Text(
          '매물 등록',
          style: TextStyle(
            color: AppColors.gold,
            fontWeight: FontWeight.bold,
          ),
        ),
      ),
      body: Form(
        key: _formKey,
        child: ListView(
          padding: const EdgeInsets.all(16),
          children: [
            // Type toggle
            _buildTypeToggle(),
            const SizedBox(height: 20),

            // Server
            servers.when(
              data: (list) => DropdownButtonFormField<String>(
                decoration: _darkInput('서버 *'),
                dropdownColor: AppColors.bgSurface,
                style: const TextStyle(color: AppColors.textPrimary),
                items: list
                    .map((s) => DropdownMenuItem(
                          value: s['serverId'] as String,
                          child: Text(s['serverName']),
                        ))
                    .toList(),
                onChanged: (v) => _serverId = v,
                validator: (v) => v == null ? '서버를 선택해주세요' : null,
              ),
              loading: () => LinearProgressIndicator(
                color: AppColors.gold.withValues(alpha: 0.5),
                backgroundColor: AppColors.bgSurface,
              ),
              error: (_, __) => const Text(
                '서버 목록 로드 실패',
                style: TextStyle(color: AppColors.error),
              ),
            ),
            const SizedBox(height: 14),

            // Category
            categories.when(
              data: (list) {
                final topLevel =
                    list.where((c) => c['parentId'] == null).toList();
                return DropdownButtonFormField<String>(
                  key: ValueKey(_categoryId),
                  decoration: _darkInput('카테고리 *'),
                  dropdownColor: AppColors.bgSurface,
                  style: const TextStyle(color: AppColors.textPrimary),
                  value: _categoryId,
                  items: topLevel
                      .map((c) => DropdownMenuItem(
                            value: c['categoryId'] as String,
                            child: Text(c['categoryName']),
                          ))
                      .toList(),
                  onChanged: (v) => setState(() => _categoryId = v),
                  validator: (v) => v == null ? '카테고리를 선택해주세요' : null,
                );
              },
              loading: () => LinearProgressIndicator(
                color: AppColors.gold.withValues(alpha: 0.5),
                backgroundColor: AppColors.bgSurface,
              ),
              error: (_, __) => const Text(
                '카테고리 로드 실패',
                style: TextStyle(color: AppColors.error),
              ),
            ),
            const SizedBox(height: 14),

            CompositedTransformTarget(
              link: _layerLink,
              child: TextFormField(
                controller: _itemNameCtrl,
                focusNode: _itemFocusNode,
                style: const TextStyle(color: AppColors.textPrimary),
                cursorColor: AppColors.gold,
                decoration: _darkInput('아이템명 *', hint: '검색어를 입력하세요').copyWith(
                  prefixIcon: _selectedItem != null && _selectedItem!['iconUrl'] != null
                      ? Padding(
                          padding: const EdgeInsets.all(10),
                          child: CachedNetworkImage(
                            imageUrl: '${ref.read(apiClientProvider).staticBaseUrl}${_selectedItem!['iconUrl']}',
                            width: 24,
                            height: 24,
                          ),
                        )
                      : null,
                  suffixIcon: _selectedItem != null
                      ? IconButton(
                          icon: const Icon(Icons.clear, size: 18, color: AppColors.textMuted),
                          onPressed: () {
                            _itemNameCtrl.clear();
                            setState(() => _selectedItem = null);
                          },
                        )
                      : const Icon(Icons.search, size: 18, color: AppColors.textMuted),
                ),
                validator: (v) =>
                    (v?.isEmpty ?? true) ? '아이템명을 입력해주세요' : null,
                onChanged: _onItemSearchChanged,
              ),
            ),
            const SizedBox(height: 14),

            TextFormField(
              controller: _titleCtrl,
              style: const TextStyle(color: AppColors.textPrimary),
              cursorColor: AppColors.gold,
              decoration: _darkInput('제목 *', hint: '예: 집행검 +9 급처합니다'),
              validator: (v) =>
                  (v?.length ?? 0) < 2 ? '제목을 2자 이상 입력해주세요' : null,
            ),
            const SizedBox(height: 14),

            TextFormField(
              controller: _descCtrl,
              style: const TextStyle(color: AppColors.textPrimary),
              cursorColor: AppColors.gold,
              decoration: _darkInput('설명 *', hint: '아이템 상세 설명'),
              maxLines: 4,
              validator: (v) =>
                  (v?.length ?? 0) < 10 ? '설명을 10자 이상 입력해주세요' : null,
            ),
            const SizedBox(height: 14),

            // Price
            Row(
              crossAxisAlignment: CrossAxisAlignment.start,
              children: [
                Expanded(
                  flex: 2,
                  child: DropdownButtonFormField<String>(
                    decoration: _darkInput('가격 유형'),
                    dropdownColor: AppColors.bgSurface,
                    style: const TextStyle(color: AppColors.textPrimary),
                    value: _priceType,
                    items: const [
                      DropdownMenuItem(value: 'fixed', child: Text('고정가')),
                      DropdownMenuItem(
                          value: 'negotiable', child: Text('협상가능')),
                      DropdownMenuItem(value: 'offer', child: Text('제안받음')),
                    ],
                    onChanged: (v) =>
                        setState(() => _priceType = v ?? 'fixed'),
                  ),
                ),
                const SizedBox(width: 12),
                Expanded(
                  flex: 3,
                  child: TextFormField(
                    controller: _priceCtrl,
                    style: const TextStyle(color: AppColors.textPrimary),
                    cursorColor: AppColors.gold,
                    decoration: _darkInput('가격 (원)'),
                    keyboardType: TextInputType.number,
                    enabled: _priceType != 'offer',
                  ),
                ),
              ],
            ),
            const SizedBox(height: 14),

            TextFormField(
              controller: _enhancementCtrl,
              style: const TextStyle(color: AppColors.textPrimary),
              cursorColor: AppColors.gold,
              decoration: _darkInput('강화 수치 (선택)'),
              keyboardType: TextInputType.number,
            ),
            const SizedBox(height: 14),

            DropdownButtonFormField<String>(
              decoration: _darkInput('거래 방식'),
              dropdownColor: AppColors.bgSurface,
              style: const TextStyle(color: AppColors.textPrimary),
              value: _tradeMethod,
              items: const [
                DropdownMenuItem(value: 'in_game', child: Text('인게임')),
                DropdownMenuItem(
                    value: 'offline_pc_bang', child: Text('PC방/오프라인')),
                DropdownMenuItem(value: 'either', child: Text('무관')),
              ],
              onChanged: (v) => _tradeMethod = v ?? 'either',
            ),
            const SizedBox(height: 28),

            SizedBox(
              height: 50,
              child: ElevatedButton(
                onPressed: _submitting ? null : _submit,
                style: ElevatedButton.styleFrom(
                  backgroundColor: AppColors.gold,
                  foregroundColor: AppColors.bg,
                  disabledBackgroundColor:
                      AppColors.gold.withValues(alpha: 0.3),
                  shape: RoundedRectangleBorder(
                    borderRadius: BorderRadius.circular(10),
                  ),
                  elevation: 0,
                ),
                child: _submitting
                    ? SizedBox(
                        width: 20,
                        height: 20,
                        child: CircularProgressIndicator(
                          strokeWidth: 2,
                          color: AppColors.gold.withValues(alpha: 0.7),
                        ),
                      )
                    : const Text(
                        '등록하기',
                        style: TextStyle(
                          fontSize: 16,
                          fontWeight: FontWeight.bold,
                        ),
                      ),
              ),
            ),
            const SizedBox(height: 20),
          ],
        ),
      ),
    );
  }

  Widget _buildTypeToggle() {
    return Container(
      decoration: BoxDecoration(
        color: AppColors.bgSurface,
        borderRadius: BorderRadius.circular(10),
        border: Border.all(color: AppColors.border),
      ),
      child: Row(
        children: [
          Expanded(
            child: GestureDetector(
              onTap: () => setState(() => _listingType = 'sell'),
              child: Container(
                padding: const EdgeInsets.symmetric(vertical: 12),
                decoration: BoxDecoration(
                  color: _listingType == 'sell'
                      ? AppColors.gold.withValues(alpha: 0.15)
                      : Colors.transparent,
                  borderRadius: BorderRadius.circular(9),
                  border: _listingType == 'sell'
                      ? Border.all(color: AppColors.gold, width: 1.5)
                      : null,
                ),
                child: Text(
                  '판매',
                  textAlign: TextAlign.center,
                  style: TextStyle(
                    fontWeight: FontWeight.w600,
                    color: _listingType == 'sell'
                        ? AppColors.gold
                        : AppColors.textSecondary,
                  ),
                ),
              ),
            ),
          ),
          Expanded(
            child: GestureDetector(
              onTap: () => setState(() => _listingType = 'buy'),
              child: Container(
                padding: const EdgeInsets.symmetric(vertical: 12),
                decoration: BoxDecoration(
                  color: _listingType == 'buy'
                      ? AppColors.gold.withValues(alpha: 0.15)
                      : Colors.transparent,
                  borderRadius: BorderRadius.circular(9),
                  border: _listingType == 'buy'
                      ? Border.all(color: AppColors.gold, width: 1.5)
                      : null,
                ),
                child: Text(
                  '구매',
                  textAlign: TextAlign.center,
                  style: TextStyle(
                    fontWeight: FontWeight.w600,
                    color: _listingType == 'buy'
                        ? AppColors.gold
                        : AppColors.textSecondary,
                  ),
                ),
              ),
            ),
          ),
        ],
      ),
    );
  }

  // ── Item autocomplete methods ──

  void _onItemSearchChanged(String query) {
    _selectedItem = null;
    _debounce?.cancel();
    if (query.isEmpty) {
      _removeOverlay();
      return;
    }
    _debounce = Timer(const Duration(milliseconds: 300), () async {
      try {
        final api = ref.read(apiClientProvider);
        final results = await api.searchItems(query);
        if (!mounted) return;
        _itemSuggestions = results;
        if (results.isNotEmpty) {
          _showOverlay();
        } else {
          _removeOverlay();
        }
      } catch (_) {
        _removeOverlay();
      }
    });
  }

  void _showOverlay() {
    _removeOverlay();
    final api = ref.read(apiClientProvider);
    _overlayEntry = OverlayEntry(
      builder: (context) => Positioned(
        width: MediaQuery.of(context).size.width - 32,
        child: CompositedTransformFollower(
          link: _layerLink,
          offset: const Offset(0, 56),
          showWhenUnlinked: false,
          child: Material(
            elevation: 8,
            color: AppColors.bgElevated,
            borderRadius: BorderRadius.circular(10),
            child: Container(
              constraints: const BoxConstraints(maxHeight: 220),
              decoration: BoxDecoration(
                border: Border.all(color: AppColors.border),
                borderRadius: BorderRadius.circular(10),
              ),
              child: ListView.separated(
                shrinkWrap: true,
                padding: EdgeInsets.zero,
                itemCount: _itemSuggestions.length,
                separatorBuilder: (_, __) => const Divider(height: 1, color: AppColors.border),
                itemBuilder: (context, index) {
                  final item = _itemSuggestions[index] as Map<String, dynamic>;
                  final iconUrl = item['iconUrl'] as String?;
                  return ListTile(
                    dense: true,
                    leading: iconUrl != null
                        ? CachedNetworkImage(
                            imageUrl: '${api.staticBaseUrl}$iconUrl',
                            width: 32,
                            height: 32,
                            placeholder: (_, __) => const SizedBox(width: 32, height: 32),
                            errorWidget: (_, __, ___) => const Icon(Icons.inventory_2, size: 24, color: AppColors.textMuted),
                          )
                        : const Icon(Icons.inventory_2, size: 24, color: AppColors.textMuted),
                    title: Text(
                      item['name'] as String,
                      style: const TextStyle(color: AppColors.textPrimary, fontSize: 14),
                    ),
                    subtitle: Text(
                      item['categoryId'] as String? ?? '',
                      style: const TextStyle(color: AppColors.textMuted, fontSize: 12),
                    ),
                    onTap: () => _selectItem(item),
                  );
                },
              ),
            ),
          ),
        ),
      ),
    );
    Overlay.of(context).insert(_overlayEntry!);
  }

  void _removeOverlay() {
    _overlayEntry?.remove();
    _overlayEntry = null;
  }

  void _selectItem(Map<String, dynamic> item) {
    _itemNameCtrl.text = item['name'] as String;
    final categories = ref.read(categoriesProvider).valueOrNull ?? [];
    final itemCatId = item['categoryId'] as String?;

    // Resolve to top-level category for the dropdown
    String? resolvedCatId = itemCatId;
    if (itemCatId != null) {
      final cat = categories.cast<Map<String, dynamic>>().where((c) => c['categoryId'] == itemCatId).firstOrNull;
      if (cat != null && cat['parentId'] != null) {
        resolvedCatId = cat['parentId'] as String;
      }
    }

    setState(() {
      _selectedItem = item;
      if (resolvedCatId != null) _categoryId = resolvedCatId;
    });
    _removeOverlay();
    _itemFocusNode.unfocus();
  }

  Future<void> _submit() async {
    if (!_formKey.currentState!.validate()) return;

    setState(() => _submitting = true);
    try {
      final api = ref.read(apiClientProvider);
      final data = <String, dynamic>{
        'listingType': _listingType,
        'serverId': _serverId,
        'categoryId': _categoryId,
        'itemName': _itemNameCtrl.text,
        'title': _titleCtrl.text,
        'description': _descCtrl.text,
        'priceType': _priceType,
        'quantity': 1,
        'tradeMethod': _tradeMethod,
      };
      if (_priceType != 'offer' && _priceCtrl.text.isNotEmpty) {
        data['priceAmount'] = int.tryParse(_priceCtrl.text) ?? 0;
      }
      if (_enhancementCtrl.text.isNotEmpty) {
        data['enhancementLevel'] = int.tryParse(_enhancementCtrl.text);
      }

      await api.createListing(data);
      if (mounted) {
        ScaffoldMessenger.of(context).showSnackBar(
          const SnackBar(
            content: Text('매물이 등록되었습니다!'),
            backgroundColor: AppColors.success,
          ),
        );
        context.pop();
      }
    } catch (e) {
      if (mounted) {
        String msg = '$e';
        if (e is DioException) {
          msg = '${e.response?.statusCode}: ${e.response?.data ?? e.message}';
        }
        ScaffoldMessenger.of(context).showSnackBar(
          SnackBar(
            content: Text('등록 실패: $msg'),
            backgroundColor: AppColors.error,
            duration: const Duration(seconds: 5),
          ),
        );
      }
    } finally {
      if (mounted) setState(() => _submitting = false);
    }
  }
}
