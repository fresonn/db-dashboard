import { TriangleAlert } from 'lucide-react'
import { Suspense, type ReactNode } from 'react'
import { ErrorBoundary } from '@/components/error-boundary'
import { Typography } from '@/components/ui/typography'

export function WidgetErrorFallback() {
  return (
    <div className="flex h-full flex-col items-center justify-center text-red-700 dark:text-yellow-400">
      <div>
        <TriangleAlert className="size-8" strokeWidth={1.5} />
      </div>
      <Typography variant="small">Failed to fetch</Typography>
    </div>
  )
}

export function Widget({
  title,
  children,
  skeleton,
  className
}: {
  title: string
  children: ReactNode
  skeleton: ReactNode
  className?: string
}) {
  return (
    <div className={`dark:bg-section-box rounded-2xl bg-gray-100 p-4 ${className}`}>
      <div className="flex h-full flex-col">
        <Typography variant="h3" className="dark:text-theme-color">
          {title}
        </Typography>
        <div className="min-h-0 flex-1">
          <ErrorBoundary fallback={<WidgetErrorFallback />}>
            <Suspense fallback={skeleton}>{children}</Suspense>
          </ErrorBoundary>
        </div>
      </div>
    </div>
  )
}
