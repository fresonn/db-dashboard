import type { ReactNode } from 'react'

export function StaticStatsGrid({ children }: { children: ReactNode }) {
  return (
    <div className="grid grid-cols-[repeat(12,minmax(0,220px))] grid-rows-[repeat(12,15px)] gap-3">
      {children}
    </div>
  )
}
