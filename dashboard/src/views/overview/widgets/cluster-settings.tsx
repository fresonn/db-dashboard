import { usePostmasterSettingsSuspense, type PostgresSetting } from '@/lib/api/gen'
import { Typography } from '@/components/ui/typography'
import { MaybeTooltip } from '@/components/ui/maybe-tooltip'
import { CopyToClipboard } from '@/components/ui/copy-to-clipboard'
import ContentLoader from 'react-content-loader'
import React from 'react'

export function ClusterSettingsSkeleton() {
  const rowHeight = 16
  const gap = 14

  const leftMin = 15
  const leftMax = 30

  const rightMin = 20
  const rightMax = 50

  const rows = Array.from({ length: 8 }, () => ({
    left: Math.floor(Math.random() * (leftMax - leftMin + 1)) + leftMin,
    right: Math.floor(Math.random() * (rightMax - rightMin + 1)) + rightMin
  }))

  return (
    <ContentLoader
      className="h-full"
      backgroundColor="var(--skeleton-bg)"
      foregroundColor="var(--skeleton-fg)"
      width="100%"
    >
      {rows.map((row, ind) => {
        const y = ind * (rowHeight + gap)

        return (
          <React.Fragment key={ind}>
            <rect x="0" y={y} width={`${row.left}%`} height={rowHeight} rx="6" ry="6" />
            <rect
              x={`${100 - row.right}%`}
              y={y}
              width={`${row.right}%`}
              height={rowHeight}
              rx="6"
              ry="6"
            />
          </React.Fragment>
        )
      })}
    </ContentLoader>
  )
}

export function ClusterSettings() {
  const { data } = usePostmasterSettingsSuspense({
    query: {
      retry: false
    }
  })

  return (
    <div>
      <PropertyRow setting={data.walBuffers} />
      <PropertyRow setting={data.sharedBuffers} />
      <PropertyRow setting={data.maxConnections} />
      <PropertyRow setting={data.walLevel} />
      <PropertyRow setting={data.autovacuumMaxWorkers} />
      <PropertyRow setting={data.dataDirectory} withCopy />
      <PropertyRow setting={data.configFile} withCopy />
      <PropertyRow setting={data.hbaFile} withCopy />
    </div>
  )
}

export function PropertyRow({
  setting,
  withCopy = false
}: {
  setting: PostgresSetting
  withCopy?: boolean
}) {
  return (
    <div className="mb-4">
      <div className="flex justify-between">
        <MaybeTooltip side="bottom" content={setting.description && <p>{setting.description}</p>}>
          <Typography variant="code">{setting.name}</Typography>
        </MaybeTooltip>

        <div className="flex">
          <div>
            <MaybeTooltip side="bottom" content={setting.unit && <p>Unit: {setting.unit}</p>}>
              <Typography variant="code">{setting.value}</Typography>
            </MaybeTooltip>
          </div>
          {withCopy && (
            <div className="ml-2.5 dark:text-green-400">
              <CopyToClipboard what={setting.value} />
            </div>
          )}
        </div>
      </div>
    </div>
  )
}
