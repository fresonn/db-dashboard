import React from 'react'
import { cva, type VariantProps } from 'cva'

const typography = cva('antialiased', {
  variants: {
    variant: {
      h1: 'scroll-m-20 text-4xl font-semibold tracking-tight',
      h2: 'scroll-m-20 text-3xl font-semibold tracking-tight',
      h3: 'scroll-m-20 text-2xl font-semibold tracking-tight',
      h4: 'scroll-m-20 text-xl font-semibold tracking-tight',
      p: 'leading-7 [&:not(:first-child)]:mt-4',
      muted: 'text-muted-foreground text-sm',
      lead: 'text-muted-foreground text-xl',
      small: 'text-sm leading-normal',
      code: 'font-code text-sm leading-none font-medium tracking-tight text-balance',
      blockquote: 'border-muted text-muted-foreground mt-6 border-l-4 pl-4 italic'
    }
  },
  defaultVariants: {
    variant: 'p'
  }
})

export type TypographyTags =
  | 'p'
  | 'span'
  | 'h1'
  | 'h2'
  | 'h3'
  | 'h4'
  | 'small'
  | 'code'
  | 'blockquote'

export interface TypographyProps extends React.HTMLAttributes<HTMLElement> {
  children: React.ReactNode
  as?: TypographyTags
}

export function Typography({
  variant,
  className,
  children,
  as,
  ...props
}: VariantProps<typeof typography> & TypographyProps) {
  const Tag = as ?? 'p'
  return (
    <Tag className={typography({ variant, className })} {...props}>
      {children}
    </Tag>
  )
}
