import ContentLoader from 'react-content-loader'
import { useRolesSuspense } from '@/lib/api/gen'
import { RolesTable } from './table/table'

export function UptimeWidgetSkeleton() {
  return (
    <ContentLoader
      className="h-full"
      backgroundColor="var(--skeleton-bg)"
      foregroundColor="var(--skeleton-fg)"
      width="100%"
    >
      <rect x="0" y="0" width="100%" rx="6" ry="6" height="16" />
    </ContentLoader>
  )
}

export function RolesWidget() {
  const { data } = useRolesSuspense({
    query: {
      retry: false
    }
  })

  return (
    <div>
      <RolesTable data={data} />
    </div>
  )
}
