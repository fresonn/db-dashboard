import { ThemeToggle } from '@/components/ui/theme-toggle'
import { TypographyH1 } from '@/components/ui/typography'
import { VersionWidget, VersionWidgetSkeleton } from './version'
import { Widget } from './common'
import { Button } from '@/components/ui/button'
import { useQueryClient } from '@tanstack/react-query'
import type { ReactNode } from 'react'

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
      <TypographyH1>Overview</TypographyH1>
      <ThemeToggle />

      <Refect />

      <div className="mt-20">
        <DashboardGrid>
          <Widget
            title="Version"
            className="col-span-3 row-span-2"
            skeleton={<VersionWidgetSkeleton />}
          >
            <VersionWidget />
          </Widget>
          <div className="col-span-3 row-span-2 dark:bg-neutral-800"></div>
          <div className="col-span-2 row-span-2 dark:bg-neutral-800"></div>
        </DashboardGrid>
      </div>
    </div>
  )
}

function Refect() {
  const queryClient = useQueryClient()

  const handleClick = () => {
    queryClient.refetchQueries({ queryKey: [{ url: '/cluster/version' }] })
  }

  return <Button onClick={handleClick}>Refetch</Button>
}
