import type { ReactNode } from 'react'

export type TypographyProps = {
  children: ReactNode
}

export function TypographyH1({ children }: TypographyProps) {
  return (
    <h1 className="scroll-m-20 text-center text-4xl font-extrabold tracking-tight text-balance select-none">
      {children}
    </h1>
  )
}

export function TypographySmall({ children }: TypographyProps) {
  return <small className="text-sm leading-none font-medium">{children}</small>
}

export function TypographyCode({ children }: TypographyProps) {
  return (
    <p className="font-code text-sm leading-none font-medium tracking-wide text-balance">
      {children}
    </p>
  )
}
