import { useEffect, useMemo, useRef, useState } from 'react'
import { Chart as ChartJS, LinearScale, PointElement, LineElement, Tooltip, Legend } from 'chart.js'
import { Line } from 'react-chartjs-2'

ChartJS.register(LinearScale, PointElement, LineElement, Tooltip, Legend)

const WINDOW_MS = 60_000
const SAMPLE_INTERVAL = 1000
const DRAW_INTERVAL = 100

type Layer = 'all' | 'throughput' | 'flush' | 'latency' | 'pressure'

type Point = {
  x: number
  y: number
}

type MetricDefinition = {
  id: string
  label: string
  layer: Exclude<Layer, 'all'>
  color: string
  yAxisID: string
}

type MetricState = {
  points: Point[]

  previous: number
  lastDrawAt: number
  target: number

  transitionStartedAt: number
}

const metricDefinitions: MetricDefinition[] = [
  {
    id: 'records',
    label: 'Records/sec',
    layer: 'throughput',
    color: '#ef4444',
    yAxisID: 'throughput'
  },
  {
    id: 'fpi',
    label: 'FPI/sec',
    layer: 'throughput',
    color: '#3b82f6',
    yAxisID: 'throughput'
  },
  {
    id: 'writes',
    label: 'Writes/sec',
    layer: 'flush',
    color: '#22c55e',
    yAxisID: 'flush'
  },
  {
    id: 'syncs',
    label: 'Syncs/sec',
    layer: 'flush',
    color: '#f59e0b',
    yAxisID: 'flush'
  },
  {
    id: 'writeLatency',
    label: 'Write latency',
    layer: 'latency',
    color: '#8b5cf6',
    yAxisID: 'latency'
  },
  {
    id: 'syncLatency',
    label: 'Sync latency',
    layer: 'latency',
    color: '#ec4899',
    yAxisID: 'latency'
  },
  {
    id: 'bufferPressure',
    label: 'Buffer pressure',
    layer: 'pressure',
    color: '#06b6d4',
    yAxisID: 'pressure'
  }
]

function randomValue(id: string) {
  switch (id) {
    case 'records':
      return 10000 + Math.random() * 3000

    case 'fpi':
      return 300 + Math.random() * 200

    case 'writes':
      return 150 + Math.random() * 70

    case 'syncs':
      return 40 + Math.random() * 20

    case 'writeLatency':
      return 1 + Math.random() * 3

    case 'syncLatency':
      return 0.5 + Math.random()

    case 'bufferPressure':
      return 50 + Math.random() * 40

    default:
      return 0
  }
}

function lerp(a: number, b: number, t: number) {
  return a + (b - a) * t
}

export function WALAnalytics() {
  const chartRef = useRef<ChartJS<'line'>>(null)

  const metricsRef = useRef<Record<string, MetricState>>(
    Object.fromEntries(
      metricDefinitions.map((metric) => {
        const value = randomValue(metric.id)

        const now = Date.now()

        return [
          metric.id,
          {
            points: [],
            previous: value,
            target: value,

            transitionStartedAt: now,
            lastDrawAt: now
          }
        ]
      })
    )
  )

  const [layer, setLayer] = useState<Layer>('all')

  const [enabled, setEnabled] = useState<Record<string, boolean>>(
    Object.fromEntries(metricDefinitions.map((metric) => [metric.id, true]))
  )

  const visibleMetrics = useMemo(() => {
    return metricDefinitions.filter((metric) => {
      if (!enabled[metric.id]) {
        return false
      }

      if (layer === 'all') {
        return true
      }

      return metric.layer === layer
    })
  }, [enabled, layer])

  useEffect(() => {
    const interval = setInterval(() => {
      const now = Date.now()

      for (const metric of metricDefinitions) {
        const state = metricsRef.current[metric.id]

        state.previous = state.target

        state.target = randomValue(metric.id)

        state.transitionStartedAt = now
      }
    }, SAMPLE_INTERVAL)

    return () => clearInterval(interval)
  }, [])

  useEffect(() => {
    let frameId: number

    const animate = () => {
      const chart = chartRef.current

      if (!chart) {
        frameId = requestAnimationFrame(animate)
        return
      }

      const now = Date.now()
      const cutoff = now - WINDOW_MS

      for (const metric of metricDefinitions) {
        const state = metricsRef.current[metric.id]

        const progress = Math.min(1, (now - state.transitionStartedAt) / SAMPLE_INTERVAL)

        const current = lerp(state.previous, state.target, progress)

        const points = state.points

        if (now - state.lastDrawAt >= DRAW_INTERVAL) {
          points.push({
            x: now,
            y: current
          })

          state.lastDrawAt = now
        }

        while (points.length && points[0].x < cutoff) {
          points.shift()
        }
      }

      chart.data.datasets = visibleMetrics.map((metric) => ({
        label: metric.label,
        data: metricsRef.current[metric.id].points,
        borderColor: metric.color,
        backgroundColor: metric.color,
        yAxisID: metric.yAxisID,
        pointRadius: 0,
        tension: 0.3
      }))

      const x = chart.options.scales?.x

      if (x) {
        x.min = now - WINDOW_MS

        x.max = now
      }

      chart.update('none')

      frameId = requestAnimationFrame(animate)
    }

    animate()

    return () => cancelAnimationFrame(frameId)
  }, [visibleMetrics])

  return (
    <div className="space-y-6 p-6">
      <div className="flex flex-wrap gap-2">
        {['all', 'throughput', 'flush', 'latency', 'pressure'].map((item) => (
          <button key={item} onClick={() => setLayer(item as Layer)}>
            {item}
          </button>
        ))}
      </div>

      <div className="flex flex-wrap gap-4">
        {metricDefinitions.map((metric) => (
          <label key={metric.id} className="flex gap-2">
            <input
              type="checkbox"
              checked={enabled[metric.id]}
              onChange={() =>
                setEnabled((prev) => ({
                  ...prev,
                  [metric.id]: !prev[metric.id]
                }))
              }
            />

            {metric.label}
          </label>
        ))}
      </div>

      <div className="h-[500px]">
        <Line
          ref={chartRef}
          data={{
            datasets: []
          }}
          options={{
            responsive: true,
            maintainAspectRatio: false,

            animation: false,

            parsing: false,

            interaction: {
              mode: 'nearest',
              intersect: false
            },

            scales: {
              x: {
                type: 'linear',

                ticks: {
                  callback(value) {
                    return new Date(Number(value)).toLocaleTimeString()
                  }
                }
              },

              throughput: {
                type: 'linear',
                position: 'left'
              },

              flush: {
                type: 'linear',
                position: 'left'
              },

              latency: {
                type: 'linear',
                position: 'right',
                grid: {
                  drawOnChartArea: false
                }
              },

              pressure: {
                type: 'linear',
                position: 'right',
                grid: {
                  drawOnChartArea: false
                }
              }
            }
          }}
        />
      </div>
    </div>
  )
}
