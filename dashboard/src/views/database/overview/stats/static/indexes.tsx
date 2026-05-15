import ContentLoader from 'react-content-loader'
import { useDatabaseStatsOverviewSuspense } from '@/lib/api/gen'
import { StatTrend, StatValue } from '../ui'

export function IndexesWidgetSkeleton() {
  return (
    <ContentLoader
      className="h-full"
      backgroundColor="var(--skeleton-bg)"
      foregroundColor="var(--skeleton-fg)"
      width="100%"
    >
      <rect x="0" y="0" width="40" rx="6" ry="6" height="34" />
      <rect x="0" y="68%" width="70%" rx="6" ry="6" height="20" />
    </ContentLoader>
  )
}

export function IndexesWidget({ databaseId }: { databaseId: string }) {
  const { data } = useDatabaseStatsOverviewSuspense(Number(databaseId))

  return (
    <div>
      <StatValue>{data.indexes.total}</StatValue>
      <StatTrend trend={data.indexes.trend} />
    </div>
  )
}
