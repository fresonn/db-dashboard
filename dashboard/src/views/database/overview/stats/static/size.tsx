import ContentLoader from 'react-content-loader'
import { useDatabaseStatsOverviewSuspense } from '@/lib/api/gen'
import { StatTrend, StatValue } from '../../ui/stat-trend'
import { CircleQuestionMark } from 'lucide-react'
import { HoverCard, HoverCardContent, HoverCardTrigger } from '@/components/ui/shadcn/hover-card'
import { Typography } from '@/components/ui/typography'

export function SizeWidgetSkeleton() {
  return (
    <ContentLoader
      className="h-full"
      backgroundColor="var(--skeleton-bg)"
      foregroundColor="var(--skeleton-fg)"
      width="100%"
    >
      <rect x="0" y="0" width="94" rx="6" ry="6" height="34" />
      <circle cx="108" cy="8" r="8" />
      <rect x="0" y="68%" width="70%" rx="6" ry="6" height="20" />
    </ContentLoader>
  )
}

export function SizeWidget({ databaseId }: { databaseId: string }) {
  const { data } = useDatabaseStatsOverviewSuspense(Number(databaseId))

  return (
    <div>
      <div className="flex">
        <StatValue>{data.size.sizePretty}</StatValue>
        <HoverCard openDelay={100}>
          <HoverCardTrigger asChild>
            <div className="pl-1">
              <CircleQuestionMark
                strokeWidth={1.5}
                size={16}
                className="cursor-pointer text-neutral-300"
              />
            </div>
          </HoverCardTrigger>
          <HoverCardContent side="right">
            <Typography variant="small">
              Precise byte size:{' '}
              <Typography as="span" variant="code">
                {data.size.sizeBytes}
              </Typography>
            </Typography>
          </HoverCardContent>
        </HoverCard>
      </div>
      <StatTrend trend={data.size.trend} />
    </div>
  )
}
