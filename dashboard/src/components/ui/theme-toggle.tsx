import { useDarkMode } from '@/hooks/use-dark-mode'
import { Button } from './button'

export function ThemeToggle() {
  const [theme, toggleTheme] = useDarkMode()

  return (
    <div>
      <Button
        onClick={toggleTheme}
        className="rounded-full bg-gray-200 p-2 text-gray-800 shadow-md transition-colors duration-300 hover:bg-gray-300 focus:ring-2 focus:ring-gray-500 focus:outline-none dark:bg-gray-700 dark:text-gray-200 dark:hover:bg-gray-600"
        aria-label={`Switch to ${theme === 'dark' ? 'light' : 'dark'} mode`}
      >
        {theme === 'dark' ? (
          <span role="img" aria-label="Sun">
            ☀️ Light Mode
          </span>
        ) : (
          <span role="img" aria-label="Moon">
            🌙 Dark Mode
          </span>
        )}
      </Button>
    </div>
  )
}
