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
      title: '린클 - 리니지 클래식 거래',
      theme: AppTheme.light,
      routerConfig: appRouter,
      debugShowCheckedModeBanner: false,
    );
  }
}
