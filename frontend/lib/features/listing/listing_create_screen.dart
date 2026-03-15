import 'package:dio/dio.dart';
import 'package:flutter/material.dart';
import 'package:flutter_riverpod/flutter_riverpod.dart';
import 'package:go_router/go_router.dart';
import '../../shared/providers/app_providers.dart';

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

  @override
  Widget build(BuildContext context) {
    final servers = ref.watch(serversProvider);
    final categories = ref.watch(categoriesProvider);

    return Scaffold(
      appBar: AppBar(title: const Text('매물 등록')),
      body: Form(
        key: _formKey,
        child: ListView(
          padding: const EdgeInsets.all(16),
          children: [
            // Type toggle
            SegmentedButton<String>(
              segments: const [
                ButtonSegment(value: 'sell', label: Text('판매')),
                ButtonSegment(value: 'buy', label: Text('구매')),
              ],
              selected: {_listingType},
              onSelectionChanged: (v) => setState(() => _listingType = v.first),
            ),
            const SizedBox(height: 16),

            // Server
            servers.when(
              data: (list) => DropdownButtonFormField<String>(
                decoration: const InputDecoration(labelText: '서버 *'),
                items: list.map((s) => DropdownMenuItem(value: s['serverId'] as String, child: Text(s['serverName']))).toList(),
                onChanged: (v) => _serverId = v,
                validator: (v) => v == null ? '서버를 선택해주세요' : null,
              ),
              loading: () => const LinearProgressIndicator(),
              error: (_, __) => const Text('서버 목록 로드 실패'),
            ),
            const SizedBox(height: 12),

            // Category
            categories.when(
              data: (list) {
                final topLevel = list.where((c) => c['parentId'] == null).toList();
                return DropdownButtonFormField<String>(
                  decoration: const InputDecoration(labelText: '카테고리 *'),
                  items: topLevel.map((c) => DropdownMenuItem(value: c['categoryId'] as String, child: Text(c['categoryName']))).toList(),
                  onChanged: (v) => _categoryId = v,
                  validator: (v) => v == null ? '카테고리를 선택해주세요' : null,
                );
              },
              loading: () => const LinearProgressIndicator(),
              error: (_, __) => const Text('카테고리 로드 실패'),
            ),
            const SizedBox(height: 12),

            TextFormField(
              controller: _itemNameCtrl,
              decoration: const InputDecoration(labelText: '아이템명 *'),
              validator: (v) => (v?.isEmpty ?? true) ? '아이템명을 입력해주세요' : null,
            ),
            const SizedBox(height: 12),

            TextFormField(
              controller: _titleCtrl,
              decoration: const InputDecoration(labelText: '제목 *', hintText: '예: 집행검 +9 급처합니다'),
              validator: (v) => (v?.length ?? 0) < 2 ? '제목을 2자 이상 입력해주세요' : null,
            ),
            const SizedBox(height: 12),

            TextFormField(
              controller: _descCtrl,
              decoration: const InputDecoration(labelText: '설명 *', hintText: '아이템 상세 설명'),
              maxLines: 4,
              validator: (v) => (v?.length ?? 0) < 10 ? '설명을 10자 이상 입력해주세요' : null,
            ),
            const SizedBox(height: 12),

            // Price
            Row(
              children: [
                Expanded(
                  flex: 2,
                  child: DropdownButtonFormField<String>(
                    decoration: const InputDecoration(labelText: '가격 유형'),
                    value: _priceType,
                    items: const [
                      DropdownMenuItem(value: 'fixed', child: Text('고정가')),
                      DropdownMenuItem(value: 'negotiable', child: Text('협상가능')),
                      DropdownMenuItem(value: 'offer', child: Text('제안받음')),
                    ],
                    onChanged: (v) => setState(() => _priceType = v ?? 'fixed'),
                  ),
                ),
                const SizedBox(width: 12),
                Expanded(
                  flex: 3,
                  child: TextFormField(
                    controller: _priceCtrl,
                    decoration: const InputDecoration(labelText: '가격 (원)'),
                    keyboardType: TextInputType.number,
                    enabled: _priceType != 'offer',
                  ),
                ),
              ],
            ),
            const SizedBox(height: 12),

            TextFormField(
              controller: _enhancementCtrl,
              decoration: const InputDecoration(labelText: '강화 수치 (선택)'),
              keyboardType: TextInputType.number,
            ),
            const SizedBox(height: 12),

            DropdownButtonFormField<String>(
              decoration: const InputDecoration(labelText: '거래 방식'),
              value: _tradeMethod,
              items: const [
                DropdownMenuItem(value: 'in_game', child: Text('인게임')),
                DropdownMenuItem(value: 'offline_pc_bang', child: Text('PC방/오프라인')),
                DropdownMenuItem(value: 'either', child: Text('무관')),
              ],
              onChanged: (v) => _tradeMethod = v ?? 'either',
            ),
            const SizedBox(height: 24),

            ElevatedButton(
              onPressed: _submitting ? null : _submit,
              child: _submitting
                  ? const SizedBox(width: 20, height: 20, child: CircularProgressIndicator(strokeWidth: 2))
                  : const Text('등록하기'),
            ),
          ],
        ),
      ),
    );
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
        ScaffoldMessenger.of(context).showSnackBar(const SnackBar(content: Text('매물이 등록되었습니다!')));
        context.pop();
      }
    } catch (e) {
      if (mounted) {
        String msg = '$e';
        if (e is DioException) {
          msg = '${e.response?.statusCode}: ${e.response?.data ?? e.message}';
        }
        ScaffoldMessenger.of(context).showSnackBar(SnackBar(content: Text('등록 실패: $msg'), duration: const Duration(seconds: 5)));
      }
    } finally {
      if (mounted) setState(() => _submitting = false);
    }
  }
}
