import 'package:flutter_riverpod/flutter_riverpod.dart';
import '../api/api_client.dart';
import '../api/api_client_interface.dart';
import '../api/auth_service.dart';
import '../api/google_auth_service.dart';

final apiClientProvider = Provider<IApiClient>((ref) => ApiClient());

final authServiceProvider = Provider<AuthService>((ref) => GoogleAuthService());

final serversProvider = FutureProvider<List<dynamic>>((ref) async {
  final api = ref.watch(apiClientProvider);
  return api.getServers();
});

final categoriesProvider = FutureProvider<List<dynamic>>((ref) async {
  final api = ref.watch(apiClientProvider);
  return api.getCategories();
});

final currentUserProvider = NotifierProvider<CurrentUserNotifier, Map<String, dynamic>?>(CurrentUserNotifier.new);

class CurrentUserNotifier extends Notifier<Map<String, dynamic>?> {
  @override
  Map<String, dynamic>? build() => null;
  void set(Map<String, dynamic>? user) => state = user;
}

final isLoggedInProvider = Provider<bool>((ref) {
  return ref.watch(currentUserProvider) != null;
});
