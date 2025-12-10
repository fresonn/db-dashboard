import { ThemeToggle } from '@/components/ui/theme-toggle'
import { Typography } from '@/components/ui/typography'
import { VersionWidget, VersionWidgetSkeleton } from './version'
import { Widget } from './common'
import type { ReactNode } from 'react'
import { UptimeWidget, UptimeWidgetSkeleton } from './uptime'

export function DashboardGrid({ children }: { children: ReactNode }) {
  return (
    <div className="grid grid-cols-[repeat(8,minmax(0,220px))] grid-rows-[repeat(4,50px)] gap-3">
      {children}
    </div>
  )
}

export function Overview() {
  return (
    <div className="p-5">
      <Typography variant="h1" as="h1">
        Overview
      </Typography>
      <ThemeToggle />

      <div className="mt-20">
        <DashboardGrid>
          <Widget
            title="Version"
            className="col-span-3 row-span-2"
            skeleton={<VersionWidgetSkeleton />}
          >
            <VersionWidget />
          </Widget>
          <Widget
            title="Uptime"
            className="col-span-3 row-span-2"
            skeleton={<UptimeWidgetSkeleton />}
          >
            <UptimeWidget />
          </Widget>
          {/* <div className="animate-in fade-in in-from-top col-span-2 row-span-2 duration-300 dark:bg-neutral-800"></div> */}
        </DashboardGrid>
      </div>
    </div>
  )
}
