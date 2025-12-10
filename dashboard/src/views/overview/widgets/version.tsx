import { usePostgresVersionSuspense } from '@/lib/api/gen'
import { Typography } from '@/components/ui/typography'
import ContentLoader from 'react-content-loader'

export function VersionWidgetSkeleton() {
  return (
    <ContentLoader
      className="h-full pt-2"
      backgroundColor="var(--skeleton-bg)"
      foregroundColor="var(--skeleton-fg)"
      width="100%"
    >
      <rect x="0" y="0" width="100%" rx="6" ry="6" height="16" />
      <rect x="0" y="60%" width="25%" rx="6" ry="6" height="16" />
    </ContentLoader>
  )
}

export function VersionWidget() {
  const { data } = usePostgresVersionSuspense({
    query: {
      retry: false
    }
  })

  return (
    <div className="animate-in fade-in slide-in-from-top-8 pt-2 duration-300">
      <Typography variant="code">{data.version}</Typography>
      {data.bitDepth !== '' && (
        <Typography className="mt-1.5" variant="code">
          {data.bitDepth}
        </Typography>
      )}
    </div>
  )
}
