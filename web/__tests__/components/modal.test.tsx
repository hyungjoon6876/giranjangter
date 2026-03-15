import { describe, it, expect, afterEach, vi } from "vitest";
import { render, screen, cleanup, fireEvent } from "@testing-library/react";
import { Modal } from "@/components/ui/modal";

afterEach(() => cleanup());

describe("Modal", () => {
  it("renders children when open", () => {
    render(
      <Modal open={true} onClose={vi.fn()} title="테스트 모달">
        <p>모달 내용입니다</p>
      </Modal>
    );
    expect(screen.getByText("테스트 모달")).toBeDefined();
    expect(screen.getByText("모달 내용입니다")).toBeDefined();
  });

  it("does not render when closed", () => {
    render(
      <Modal open={false} onClose={vi.fn()} title="테스트 모달">
        <p>모달 내용입니다</p>
      </Modal>
    );
    expect(screen.queryByText("테스트 모달")).toBeNull();
    expect(screen.queryByText("모달 내용입니다")).toBeNull();
  });

  it("close button calls onClose", () => {
    const onClose = vi.fn();
    render(
      <Modal open={true} onClose={onClose} title="테스트 모달">
        <p>내용</p>
      </Modal>
    );
    // The close button contains the × character
    const closeBtn = screen.getByText("\u00D7");
    fireEvent.click(closeBtn);
    expect(onClose).toHaveBeenCalledTimes(1);
  });
});
