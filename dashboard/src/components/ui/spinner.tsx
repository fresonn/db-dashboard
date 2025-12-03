import * as React from 'react'
import { Loader2Icon } from 'lucide-react'

export function Spinner({ ...props }: React.ComponentProps<'svg'>) {
  return (
    <Loader2Icon
      role="status"
      aria-label="Loading"
      strokeWidth={3}
      className="size-5 animate-spin"
      {...props}
    />
  )
}
