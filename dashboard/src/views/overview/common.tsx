import { TriangleAlert } from 'lucide-react'
import { Suspense, type ReactNode } from 'react'
import { ErrorBoundary } from '@/components/error-boundary'

export function WidgetErrorFallback() {
  return (
    <div className="flex h-full flex-col items-center justify-center text-red-700 dark:text-yellow-400">
      <div>
        <TriangleAlert className="size-8" strokeWidth={1.5} />
      </div>
      <p className="text-center">Failed to fetch</p>
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
    <div className={`rounded-2xl bg-gray-100 p-4 dark:bg-neutral-800/80 ${className}`}>
      <div className="flex h-full flex-col">
        <p className="text-xl font-semibold">{title}</p>
        <div className="min-h-0 flex-1">
          <ErrorBoundary fallback={<WidgetErrorFallback />}>
            <Suspense fallback={skeleton}>{children}</Suspense>
          </ErrorBoundary>
        </div>
      </div>
    </div>
  )
}
