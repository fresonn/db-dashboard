import React from 'react'
import clsx from 'clsx'
import { cva } from 'cva'
import { Typography } from '@/components/ui/typography'
import { ArrowUp } from 'lucide-react'

export function StatValue({ children }: { children: React.ReactNode }) {
  return (
    <div className="mb-3">
      <p className="text-3xl">{children}</p>
    </div>
  )
}

const trendStyle = cva('flex items-center', {
  variants: {
    direction: {
      up: 'text-green-600',
      down: 'text-destructive'
    }
  }
})

export function StatTrend({
  trend
}: {
  trend: { diff: number; value: string; direction: string }
}) {
  if (trend.diff === 0) return null

  const sign = trend.direction === 'down' ? '-' : '+'

  return (
    <div className="flex items-center">
      <div className={trendStyle({ direction: trend.direction })}>
        <ArrowUp size={17} className={clsx('mr-1', trend.direction === 'down' && 'rotate-180')} />
        <Typography variant="small">
          {sign}
          {trend.value}
        </Typography>
      </div>
      <Typography variant="small" className="ml-1.5 text-neutral-400">
        in 15 minutes
      </Typography>
    </div>
  )
}
