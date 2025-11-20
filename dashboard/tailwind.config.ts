import type { Config } from 'tailwindcss'

const config: Config = {
  future: {
    hoverOnlyWhenSupported: true
  },
  content: ['./src/**/*.{js,ts,jsx,tsx,mdx}'],
  theme: {
    extend: {
      fontFamily: {
        sans: ['Roboto', 'Helvetica Neue', 'Arial', 'Menlo', 'Consolas', 'sans-serif'],
        code: ['Fira Code', 'Menlo', 'Consolas', 'monospace']
      }
    }
  }
}

export default config
