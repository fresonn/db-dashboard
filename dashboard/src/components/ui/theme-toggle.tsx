import { useTheme } from '@/hooks/use-theme'
import { Sun, Moon } from 'lucide-react'

export function ThemeToggle() {
  const [theme, toggleTheme] = useTheme()

  const isDark = theme === 'dark'

  return (
    <button
      type="button"
      onClick={toggleTheme}
      className={`relative inline-flex h-7 w-14 cursor-pointer items-center rounded-full transition-colors select-none ${
        isDark ? 'bg-neutral-700' : 'bg-gray-300'
      }`}
    >
      <div className="pointer-events-none absolute inset-0 z-40 flex items-center justify-between px-2">
        <Sun
          size={16}
          className={`text-black transition-opacity ${isDark ? 'opacity-0' : 'opacity-100'}`}
        />
        <Moon
          size={16}
          className={`text-gray-100 transition-opacity ${isDark ? 'opacity-100' : 'opacity-0'}`}
        />
      </div>

      <div
        className={`inline-block size-6 transform rounded-full shadow-lg transition-transform ${
          isDark ? 'translate-x-7 bg-neutral-800' : 'translate-x-1 bg-white'
        }`}
      />
    </button>
  )
}
