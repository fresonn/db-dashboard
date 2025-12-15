import { cva } from 'cva'
import { useTheme } from '@/hooks/use-theme'
import { Typography } from '@/components/ui/typography'
import { useSidebar } from '@/components/ui/shadcn/sidebar'
import { PostgresLogo } from '@/components/ui/pg-logo'

const head = cva('flex items-center transition-all duration-200', {
  variants: {
    open: {
      true: 'mb-4 py-4',
      false: 'mb-2 py-4'
    }
  }
})

const logo = cva('shrink-0 transition-all duration-200', {
  variants: {
    open: {
      true: 'size-10',
      false: 'size-8'
    }
  }
})

const title = cva(
  'pl-1 whitespace-nowrap transition-opacity duration-200 select-none dark:text-gray-200',
  {
    variants: {
      open: {
        true: 'opacity-100',
        false: 'pointer-events-none opacity-0'
      }
    }
  }
)

export function Header() {
  const { open } = useSidebar()
  const [theme] = useTheme()

  return (
    <div className={head({ open })}>
      <div className={logo({ open })}>
        <PostgresLogo className="h-full w-full" variant={theme} />
      </div>

      <Typography className={title({ open })} variant="h3" as="h3">
        Simple Dashboard
      </Typography>
    </div>
  )
}
