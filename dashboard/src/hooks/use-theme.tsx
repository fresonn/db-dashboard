import { useContext } from 'react'
import { ThemeContext, type Theme } from '@/providers/theme-provider'

export const useTheme = (): [Theme, () => void] => {
  const context = useContext(ThemeContext)
  if (!context) {
    throw new Error('useTheme must be used within a ThemeProvider')
  }
  return context
}
