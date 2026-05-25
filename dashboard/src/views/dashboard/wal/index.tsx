import { useEffect, useMemo, useRef, useState } from 'react'
import {
  Chart as ChartJS,
  CategoryScale,
  LinearScale,
  PointElement,
  LineElement,
  Tooltip,
  Legend
} from 'chart.js'
import { Line } from 'react-chartjs-2'

ChartJS.register(CategoryScale, LinearScale, PointElement, LineElement, Tooltip, Legend)

type Layer = 'all' | 'throughput' | 'flush' | 'latency' | 'pressure'

const metrics = [
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

export function WALAnalytics() {
  const chartRef = useRef<ChartJS<'line'>>(null)

  const [layer, setLayer] = useState<Layer>('all')

  const [enabled, setEnabled] = useState<Record<string, boolean>>(
    Object.fromEntries(metrics.map((x) => [x.id, true]))
  )

  const visibleMetrics = useMemo(() => {
    return metrics.filter((metric) => {
      if (!enabled[metric.id]) return false

      if (layer === 'all') return true

      return metric.layer === layer
    })
  }, [enabled, layer])

  useEffect(() => {
    const interval = setInterval(() => {
      const chart = chartRef.current

      if (!chart) return

      const now = new Date().toLocaleTimeString()

      chart.data.labels?.push(now)

      visibleMetrics.forEach((metric) => {
        const dataset = chart.data.datasets.find((d) => d.label === metric.label)

        if (!dataset) return

        dataset.data.push(randomValue(metric.id))
      })

      // draft: with widnow
      if (chart.data.labels && chart.data.labels.length > 100) {
        chart.data.labels.shift()

        chart.data.datasets.forEach((d) => d.data.shift())
      }
    }, 1000)

    return () => clearInterval(interval)
  }, [visibleMetrics])

  useEffect(() => {
    let frameId: number

    const animate = () => {
      const chart = chartRef.current

      if (chart) {
        // chart.update('none')
        chart.update()
      }

      frameId = requestAnimationFrame(animate)
    }

    animate()

    return () => {
      cancelAnimationFrame(frameId)
    }
  }, [])

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
        {metrics.map((metric) => (
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

      <div className="h-[450px]">
        <Line
          ref={chartRef}
          options={{
            responsive: true,
            maintainAspectRatio: false,

            animation: {
              duration: 500,
              easing: 'linear'
            },

            interaction: {
              mode: 'index',
              intersect: false
            },

            elements: {
              point: {
                radius: 0
              }
            },

            scales: {
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
          data={{
            labels: [],

            datasets: visibleMetrics.map((metric) => ({
              label: metric.label,
              data: [],
              borderColor: metric.color,
              backgroundColor: metric.color,
              yAxisID: metric.yAxisID,
              tension: 0.3
            }))
          }}
        />
      </div>
    </div>
  )
}
