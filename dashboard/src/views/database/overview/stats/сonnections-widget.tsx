import React, { type ComponentType } from 'react'
import { Wrench } from 'lucide-react'
import ContentLoader from 'react-content-loader'

export function ConnectionsWidgetSkeleton() {
  return (
    <ContentLoader
      className="h-full"
      backgroundColor="var(--skeleton-bg)"
      foregroundColor="var(--skeleton-fg)"
      width="100%"
    >
      <rect x="0" y="0" width="100%" rx="6" ry="6" height="15" />
      <rect x="0" y="58%" width="45%" rx="6" ry="6" height="15" />
    </ContentLoader>
  )
}

export const ConnectionsWidget = React.lazy(() => {
  return new Promise<{ default: ComponentType }>((resolve) => {
    setTimeout(() => {
      resolve({
        default: () => (
          <div className="flex h-20 items-center justify-center">
            <Wrench className="mr-2" /> In dev
          </div>
        )
      })
    }, 200)
  })
})
