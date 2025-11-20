import * as React from 'react'
import * as LabelPrimitive from '@radix-ui/react-label'

export type LabelProps = Omit<React.ComponentProps<typeof LabelPrimitive.Root>, 'className'>

export function Label({ ...props }: LabelProps) {
  return (
    <LabelPrimitive.Root
      data-slot="label"
      className="flex items-center gap-2 text-sm leading-none font-medium select-none group-data-[disabled=true]:pointer-events-none group-data-[disabled=true]:opacity-50 peer-disabled:cursor-not-allowed peer-disabled:opacity-50"
      {...props}
    />
  )
}
