import 'package:flutter/material.dart';
import 'package:flutter_riverpod/flutter_riverpod.dart';
import 'app/router.dart';
import 'shared/theme/app_theme.dart';

void main() {
  runApp(const ProviderScope(child: LincleApp()));
}

class LincleApp extends StatelessWidget {
  const LincleApp({super.key});

  @override
  Widget build(BuildContext context) {
    return MaterialApp.router(
      title: '기란장터',
      theme: AppTheme.dark,
      routerConfig: appRouter,
      debugShowCheckedModeBanner: false,
    );
  }
}
