import { useState, useLayoutEffect } from 'react'

const getInitialTheme = () => {
  const savedTheme = localStorage.getItem('theme')
  if (savedTheme === 'dark' || savedTheme === 'light') {
    return savedTheme
  }

  if (typeof window !== 'undefined' && window.matchMedia('(prefers-color-scheme: dark)').matches) {
    return 'dark'
  }
  return 'light'
}

export const useDarkMode = (): [string, () => void] => {
  const [theme, setTheme] = useState(getInitialTheme)

  useLayoutEffect(() => {
    const root = window.document.documentElement

    // Применяем класс
    if (theme === 'dark') {
      root.classList.add('dark')
      root.classList.remove('light')
    } else {
      root.classList.add('light')
      root.classList.remove('dark')
    }

    localStorage.setItem('theme', theme)
  }, [theme])

  const toggleTheme = () => {
    setTheme((currentTheme) => (currentTheme === 'light' ? 'dark' : 'light'))
  }

  return [theme, toggleTheme]
}
