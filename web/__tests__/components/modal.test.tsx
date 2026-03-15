import { describe, it, expect, afterEach, vi, beforeEach } from "vitest";
import {
  render,
  screen,
  cleanup,
  fireEvent,
} from "@testing-library/react";
import { Modal } from "@/components/ui/modal";

beforeEach(() => {
  // Ensure portal container exists for each test
  let portal = document.getElementById("modal-portal");
  if (!portal) {
    portal = document.createElement("div");
    portal.id = "modal-portal";
    document.body.appendChild(portal);
  }
});

afterEach(() => {
  cleanup();
  // Clean up portal container
  const portal = document.getElementById("modal-portal");
  if (portal) {
    portal.remove();
  }
  // Reset body overflow
  document.body.style.overflow = "";
});

describe("Modal", () => {
  it("renders children when open", () => {
    render(
      <Modal open={true} onClose={vi.fn()} title="테스트 모달">
        <p>모달 내용입니다</p>
      </Modal>,
    );
    expect(screen.getByText("테스트 모달")).toBeDefined();
    expect(screen.getByText("모달 내용입니다")).toBeDefined();
  });

  it("does not render when closed", () => {
    render(
      <Modal open={false} onClose={vi.fn()} title="테스트 모달">
        <p>모달 내용입니다</p>
      </Modal>,
    );
    expect(screen.queryByText("테스트 모달")).toBeNull();
    expect(screen.queryByText("모달 내용입니다")).toBeNull();
  });

  it("close button calls onClose", () => {
    const onClose = vi.fn();
    render(
      <Modal open={true} onClose={onClose} title="테스트 모달">
        <p>내용</p>
      </Modal>,
    );
    const closeBtn = screen.getByLabelText("닫기");
    fireEvent.click(closeBtn);
    expect(onClose).toHaveBeenCalledTimes(1);
  });

  it("has role=dialog and aria-modal", () => {
    render(
      <Modal open={true} onClose={vi.fn()} title="접근성 모달">
        <p>내용</p>
      </Modal>,
    );
    const dialog = screen.getByRole("dialog");
    expect(dialog).toBeDefined();
    expect(dialog.getAttribute("aria-modal")).toBe("true");
  });

  it("has aria-labelledby linked to title", () => {
    render(
      <Modal open={true} onClose={vi.fn()} title="제목 연결 테스트">
        <p>내용</p>
      </Modal>,
    );
    const dialog = screen.getByRole("dialog");
    const labelledBy = dialog.getAttribute("aria-labelledby");
    expect(labelledBy).toBeTruthy();

    // The h2 with this id should contain the title text
    const titleEl = document.getElementById(labelledBy!);
    expect(titleEl).not.toBeNull();
    expect(titleEl!.textContent).toBe("제목 연결 테스트");
  });

  it("close button has aria-label", () => {
    render(
      <Modal open={true} onClose={vi.fn()} title="접근성 모달">
        <p>내용</p>
      </Modal>,
    );
    const closeBtn = screen.getByLabelText("닫기");
    expect(closeBtn).toBeDefined();
    expect(closeBtn.getAttribute("aria-label")).toBe("닫기");
  });

  it("calls onClose on Escape key", () => {
    const onClose = vi.fn();
    render(
      <Modal open={true} onClose={onClose} title="ESC 테스트">
        <p>내용</p>
      </Modal>,
    );
    const dialog = screen.getByRole("dialog");
    fireEvent.keyDown(dialog, { key: "Escape" });
    expect(onClose).toHaveBeenCalledTimes(1);
  });
});
