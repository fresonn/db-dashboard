import { usePostgresVersionSuspense } from '@/lib/api/gen'
import { TypographyCode } from '@/components/ui/typography'
import ContentLoader from 'react-content-loader'

export function VersionWidgetSkeleton() {
  return (
    <ContentLoader
      className="h-full pt-2"
      backgroundColor="var(--skeleton-bg)"
      foregroundColor="var(--skeleton-fg)"
      width="100%"
      viewBox="0 0 390 50"
    >
      <rect x="0" y="0" width="100%" rx="6" ry="6" height="18" />
      <rect x="0" y="30" width="25%" rx="6" ry="6" height="18" />
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
    <div className="pt-2">
      <TypographyCode>{data.version}</TypographyCode>
      {data.bitDepth !== '' && <TypographyCode>{data.bitDepth}</TypographyCode>}
    </div>
  )
}
