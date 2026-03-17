"use client";

import { useState, useRef, useCallback } from "react";
import { apiClient } from "@/lib/api-client";
import { useToast } from "@/lib/hooks/use-toast";
import type { UploadedImage } from "@/lib/types";

interface ImageUploadProps {
  images: UploadedImage[];
  onChange: (images: UploadedImage[]) => void;
  maxImages?: number;
}

const ALLOWED_TYPES = ["image/jpeg", "image/png", "image/webp"];
const MAX_SIZE = 10 * 1024 * 1024;

export function ImageUpload({
  images,
  onChange,
  maxImages = 5,
}: ImageUploadProps) {
  const [uploading, setUploading] = useState(false);
  const inputRef = useRef<HTMLInputElement>(null);
  const { addToast } = useToast();

  const upload = useCallback(
    async (files: FileList | File[]) => {
      const remaining = maxImages - images.length;
      if (remaining <= 0) {
        addToast("error", `최대 ${maxImages}장까지 업로드할 수 있습니다`);
        return;
      }

      const validFiles = Array.from(files)
        .slice(0, remaining)
        .filter((f) => {
          if (!ALLOWED_TYPES.includes(f.type)) {
            addToast("error", `${f.name}: JPG, PNG, WebP만 가능합니다`);
            return false;
          }
          if (f.size > MAX_SIZE) {
            addToast("error", `${f.name}: 10MB 이하만 가능합니다`);
            return false;
          }
          return true;
        });

      if (!validFiles.length) return;

      setUploading(true);
      try {
        const results: UploadedImage[] = [];
        for (const file of validFiles) {
          const result = await apiClient.uploadImage(file);
          results.push(result);
        }
        onChange([...images, ...results]);
      } catch {
        addToast("error", "이미지 업로드에 실패했습니다");
      } finally {
        setUploading(false);
      }
    },
    [images, maxImages, onChange, addToast],
  );

  const handleDrop = useCallback(
    (e: React.DragEvent) => {
      e.preventDefault();
      upload(e.dataTransfer.files);
    },
    [upload],
  );

  const handleRemove = (index: number) => {
    onChange(images.filter((_, i) => i !== index));
  };

  return (
    <div>
      <div
        onDrop={handleDrop}
        onDragOver={(e) => e.preventDefault()}
        onClick={() => inputRef.current?.click()}
        className="border-2 border-dashed border-border rounded-lg p-6 text-center cursor-pointer hover:border-gold/50 transition-colors"
      >
        <input
          ref={inputRef}
          type="file"
          accept={ALLOWED_TYPES.join(",")}
          multiple
          className="hidden"
          onChange={(e) => e.target.files && upload(e.target.files)}
        />
        {uploading ? (
          <p className="text-text-secondary text-sm">업로드 중...</p>
        ) : (
          <>
            <p className="text-text-secondary text-sm">
              클릭 또는 드래그하여 이미지 추가
            </p>
            <p className="text-text-dim text-xs mt-1">
              JPG, PNG, WebP · 최대 10MB · {images.length}/{maxImages}장
            </p>
          </>
        )}
      </div>

      {images.length > 0 && (
        <div className="grid grid-cols-5 gap-2 mt-3">
          {images.map((img, i) => (
            <div key={img.imageId} className="relative group">
              <img
                src={img.thumbnailUrl || img.url}
                alt={`업로드 이미지 ${i + 1}`}
                className="w-full aspect-square object-cover rounded-lg border border-border"
              />
              <button
                type="button"
                onClick={() => handleRemove(i)}
                className="absolute -top-1 -right-1 w-5 h-5 bg-danger text-white text-xs rounded-full opacity-0 group-hover:opacity-100 transition-opacity"
                aria-label={`이미지 ${i + 1} 삭제`}
              >
                ×
              </button>
              {i === 0 && (
                <span className="absolute bottom-1 left-1 bg-gold text-dark text-[10px] px-1.5 py-0.5 rounded font-medium">
                  대표
                </span>
              )}
            </div>
          ))}
        </div>
      )}
    </div>
  );
}
