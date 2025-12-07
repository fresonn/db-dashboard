import { type ReactNode } from 'react'
import { cva } from 'cva'

const msg = cva('h-5 pl-2.5 text-[13px] leading-none font-medium text-red-600', {
  variants: {
    show: {
      true: 'animate-in fade-in visible duration-200',
      false: 'invisible'
    },
    align: {
      start: 'text-start',
      center: 'text-center',
      end: 'text-end'
    }
  }
})

export function ErrorMessage({
  show,
  align = 'start',
  children
}: {
  show: boolean
  align?: 'start' | 'center' | 'end'
  children: ReactNode
}) {
  return <p className={msg({ show, align })}>{children}</p>
}
