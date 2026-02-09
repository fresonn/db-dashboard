import React from 'react'
import { Tooltip, TooltipContent, TooltipTrigger, type TooltipContentProps } from './shadcn/tooltip'

export function MaybeTooltip({
  content,
  side,
  children
}: {
  content?: React.ReactNode
  side?: TooltipContentProps['side']
  children: React.ReactElement
}) {
  if (!content) return children

  return (
    <Tooltip>
      <TooltipTrigger asChild>{children}</TooltipTrigger>
      <TooltipContent side={side}>{content}</TooltipContent>
    </Tooltip>
  )
}
