import type { Config } from "tailwindcss";
import tokens from "../shared/design-tokens.json";

const config: Config = {
  content: [
    "./app/**/*.{ts,tsx}",
    "./components/**/*.{ts,tsx}",
    "./lib/**/*.{ts,tsx}",
  ],
  theme: {
    extend: {
      colors: {
        primary: {
          DEFAULT: tokens.colors.primary,
          dark: tokens.colors.primaryDark,
        },
        secondary: tokens.colors.secondary,
        error: tokens.colors.error,
        warning: tokens.colors.warning,
        surface: tokens.colors.surface,
        "text-primary": tokens.colors.textPrimary,
        "text-secondary": tokens.colors.textSecondary,
        border: tokens.colors.border,
      },
      borderRadius: {
        sm: tokens.borderRadius.sm,
        md: tokens.borderRadius.md,
        lg: tokens.borderRadius.lg,
        xl: tokens.borderRadius.xl,
      },
      fontSize: {
        xs: tokens.fontSize.xs,
        sm: tokens.fontSize.sm,
        base: tokens.fontSize.base,
        lg: tokens.fontSize.lg,
        xl: tokens.fontSize.xl,
        "2xl": tokens.fontSize["2xl"],
        "3xl": tokens.fontSize["3xl"],
      },
    },
  },
  plugins: [],
};
export default config;
