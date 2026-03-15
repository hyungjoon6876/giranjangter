export function Loading() {
  return (
    <div className="flex items-center justify-center py-20" role="status" aria-label="로딩 중">
      <div className="w-8 h-8 border-4 border-gold border-t-transparent rounded-full animate-spin" />
      <span className="sr-only">로딩 중</span>
    </div>
  );
}
