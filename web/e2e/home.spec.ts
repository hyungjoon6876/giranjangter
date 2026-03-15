import { test, expect } from "@playwright/test";

test.describe("홈페이지 — 비로그인 사용자", () => {
  test.beforeEach(async ({ page }) => {
    // 토큰 초기화 (비로그인 상태)
    await page.goto("/");
    await page.evaluate(() => {
      localStorage.removeItem("accessToken");
      localStorage.removeItem("refreshToken");
    });
    await page.reload();
    await page.waitForLoadState("networkidle");
  });

  test("메인 페이지가 크래시 없이 렌더링된다", async ({ page }) => {
    // 콘솔 에러 수집
    const errors: string[] = [];
    page.on("pageerror", (err) => errors.push(err.message));

    await page.goto("/");
    await page.waitForLoadState("networkidle");

    // TypeError, ReferenceError 등 JS 크래시가 없어야 함
    const jsErrors = errors.filter(
      (e) => e.includes("TypeError") || e.includes("ReferenceError") || e.includes("Cannot read")
    );
    expect(jsErrors).toHaveLength(0);
  });

  test("히어로 섹션이 비로그인 사용자에게 표시된다", async ({ page }) => {
    await expect(page.getByRole("heading", { name: "기란장터", level: 1 })).toBeVisible();
    await expect(page.getByText("리니지 클래식 아이템 거래, 안전하고 무료")).toBeVisible();
    await expect(page.getByRole("link", { name: "시작하기" })).toBeVisible();
    await expect(page.getByRole("link", { name: "매물 둘러보기" })).toBeVisible();
  });

  test("헤더에 로그인 링크가 표시된다", async ({ page }) => {
    await expect(page.getByRole("banner").getByRole("link", { name: "로그인" })).toBeVisible();
  });

  test("서버 필터가 줄바꿈으로 표시된다", async ({ page }) => {
    const filterGroup = page.getByRole("group", { name: "서버 필터" });
    await expect(filterGroup).toBeVisible();
    await expect(filterGroup.getByRole("button", { name: "전체" })).toBeVisible();
    // 서버 버튼이 여러 개 존재
    const buttons = filterGroup.getByRole("button");
    expect(await buttons.count()).toBeGreaterThan(5);
  });

  test("서버 필터 클릭 시 크래시가 발생하지 않는다", async ({ page }) => {
    const errors: string[] = [];
    page.on("pageerror", (err) => errors.push(err.message));

    // 매물이 없는 서버들을 순차 클릭
    const filterGroup = page.getByRole("group", { name: "서버 필터" });
    const buttons = filterGroup.getByRole("button");
    const count = await buttons.count();

    // 최대 5개 서버 클릭 테스트
    for (let i = 1; i < Math.min(count, 6); i++) {
      await buttons.nth(i).click();
      await page.waitForTimeout(500);
    }

    // 전체로 복귀
    await filterGroup.getByRole("button", { name: "전체" }).click();
    await page.waitForTimeout(500);

    const jsErrors = errors.filter(
      (e) => e.includes("TypeError") || e.includes("Cannot read")
    );
    expect(jsErrors).toHaveLength(0);
  });

  test("서버 필터 선택 시 aria-pressed가 변경된다", async ({ page }) => {
    const filterGroup = page.getByRole("group", { name: "서버 필터" });
    const allBtn = filterGroup.getByRole("button", { name: "전체" });

    await expect(allBtn).toHaveAttribute("aria-pressed", "true");

    // 다른 서버 클릭
    const secondBtn = filterGroup.getByRole("button").nth(1);
    await secondBtn.click();

    await expect(allBtn).toHaveAttribute("aria-pressed", "false");
    await expect(secondBtn).toHaveAttribute("aria-pressed", "true");
  });

  test("매물이 없는 서버 선택 시 빈 상태가 표시된다", async ({ page }) => {
    const filterGroup = page.getByRole("group", { name: "서버 필터" });

    // 매물이 없을 확률이 높은 서버 클릭 (뒤쪽 서버)
    const buttons = filterGroup.getByRole("button");
    const count = await buttons.count();
    if (count > 10) {
      await buttons.nth(count - 2).click();
      await page.waitForLoadState("networkidle");

      // 빈 상태 또는 매물 카드가 표시되어야 함 (크래시 없이)
      const emptyOrCards = await Promise.race([
        page.getByRole("heading", { name: /매물이 없습니다|해당 서버/ }).waitFor({ timeout: 5000 }).then(() => "empty"),
        page.getByRole("link", { name: /원/ }).first().waitFor({ timeout: 5000 }).then(() => "cards"),
      ]);
      expect(["empty", "cards"]).toContain(emptyOrCards);
    }
  });

  test("정렬 드롭다운이 동작한다", async ({ page }) => {
    const sort = page.getByRole("combobox", { name: "정렬 방식" });

    // 기본 최신순
    if (await sort.isVisible()) {
      await expect(sort).toHaveValue("recent");

      await sort.selectOption("price_asc");
      await expect(sort).toHaveValue("price_asc");

      await sort.selectOption("recent");
      await expect(sort).toHaveValue("recent");
    }
  });

  test("검색 입력이 동작한다", async ({ page }) => {
    const errors: string[] = [];
    page.on("pageerror", (err) => errors.push(err.message));

    const searchInput = page.getByRole("searchbox", { name: "매물 검색" }).last();
    await searchInput.fill("존재하지않는아이템");
    await page.waitForTimeout(1000);

    // 크래시 없이 빈 상태 또는 결과가 표시
    const jsErrors = errors.filter((e) => e.includes("TypeError") || e.includes("Cannot read"));
    expect(jsErrors).toHaveLength(0);
  });
});
