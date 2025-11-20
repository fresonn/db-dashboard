import { createContext, useState, useLayoutEffect, type ReactNode } from 'react'

export type Theme = 'light' | 'dark'

export const ThemeContext = createContext<[Theme, () => void] | null>(null)

const getInitialTheme = (): Theme => {
  if (typeof window === 'undefined') return 'light'

  const saved = localStorage.getItem('theme')
  if (saved === 'light' || saved === 'dark') return saved as Theme

  return window.matchMedia('(prefers-color-scheme: dark)').matches ? 'dark' : 'light'
}

export const ThemeProvider = ({ children }: { children: ReactNode }) => {
  const [theme, setTheme] = useState<Theme>(getInitialTheme)

  useLayoutEffect(() => {
    const root = document.documentElement
    root.classList.toggle('dark', theme === 'dark')
    root.classList.toggle('light', theme === 'light')
    localStorage.setItem('theme', theme)
  }, [theme])

  const toggleTheme = () => setTheme((t) => (t === 'light' ? 'dark' : 'light'))

  return <ThemeContext.Provider value={[theme, toggleTheme]}>{children}</ThemeContext.Provider>
}
