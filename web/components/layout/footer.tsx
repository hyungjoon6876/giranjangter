import Link from "next/link";

export function Footer() {
  return (
    <footer className="hidden lg:block border-t border-border bg-dark mt-auto">
      <div className="max-w-6xl mx-auto px-6 py-8">
        <div className="flex flex-col lg:flex-row lg:items-start lg:justify-between gap-6">
          {/* Brand */}
          <div>
            <span className="text-gold font-display text-lg font-bold">
              기란JT
            </span>
            <p className="text-text-dim text-sm mt-1">
              리니지 클래식 아이템 거래 중개 플랫폼
            </p>
            <p className="text-text-dim text-xs mt-1">무료 · 커뮤니티 기반</p>
          </div>

          {/* Links */}
          <div className="flex gap-8">
            <div>
              <h3 className="text-text-secondary text-xs font-semibold uppercase tracking-wider mb-2">
                서비스
              </h3>
              <ul className="space-y-1.5">
                <li>
                  <Link
                    href="/"
                    className="text-text-dim text-sm hover:text-gold transition-colors"
                  >
                    매물 보기
                  </Link>
                </li>
                <li>
                  <Link
                    href="/create"
                    className="text-text-dim text-sm hover:text-gold transition-colors"
                  >
                    매물 등록
                  </Link>
                </li>
              </ul>
            </div>
            <div>
              <h3 className="text-text-secondary text-xs font-semibold uppercase tracking-wider mb-2">
                정보
              </h3>
              <ul className="space-y-1.5">
                <li>
                  <span className="text-text-dim text-sm">
                    문의: giranjt@gmail.com
                  </span>
                </li>
              </ul>
            </div>
          </div>
        </div>

        <div className="border-t border-border mt-6 pt-4 text-center">
          <p className="text-text-dim text-xs">
            &copy; {new Date().getFullYear()} 기란JT. 이 서비스는 리니지 클래식의
            공식 서비스가 아닙니다.
          </p>
        </div>
      </div>
    </footer>
  );
}
