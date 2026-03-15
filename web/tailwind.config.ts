import type { Config } from "tailwindcss";
import tokens from "../shared/design-tokens.json";

const config: Config = {
  content: ["./app/**/*.{ts,tsx}", "./components/**/*.{ts,tsx}", "./lib/**/*.{ts,tsx}"],
  theme: {
    extend: {
      colors: {
        darkest: tokens.colors.darkest,
        dark: tokens.colors.dark,
        medium: tokens.colors.medium,
        light: tokens.colors.light,
        card: tokens.colors.card,
        border: tokens.colors.border,
        gold: { DEFAULT: tokens.colors.gold, bright: tokens.colors.goldBright, dim: tokens.colors.goldDim },
        blue: { DEFAULT: tokens.colors.blue, bright: tokens.colors.blueBright, dim: tokens.colors.blueDim },
        "text-primary": tokens.colors.textPrimary,
        "text-secondary": tokens.colors.textSecondary,
        "text-dim": tokens.colors.textDim,
        success: tokens.colors.success,
        danger: tokens.colors.danger,
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
      fontFamily: {
        display: ["'Black Han Sans'", "sans-serif"],
        serif: ["'Noto Serif KR'", "serif"],
        sans: ["'Noto Sans KR'", "-apple-system", "BlinkMacSystemFont", "sans-serif"],
      },
    },
  },
  plugins: [],
};
export default config;
