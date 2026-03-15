export default function Home() {
  return (
    <main className="flex min-h-screen flex-col items-center justify-center gap-6 p-8">
      <h1 className="text-3xl font-bold text-primary">기란장터</h1>
      <p className="text-lg text-text-secondary">
        리니지 클래식 아이템 거래 중개 플랫폼
      </p>
      <div className="flex gap-3">
        <span className="rounded-lg bg-primary px-4 py-2 text-sm text-white">
          Primary
        </span>
        <span className="rounded-lg bg-secondary px-4 py-2 text-sm text-white">
          Secondary
        </span>
        <span className="rounded-lg bg-error px-4 py-2 text-sm text-white">
          Error
        </span>
        <span className="rounded-lg bg-warning px-4 py-2 text-sm text-white">
          Warning
        </span>
      </div>
    </main>
  );
}
