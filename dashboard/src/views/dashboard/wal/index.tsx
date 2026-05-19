import {
  Chart as ChartJS,
  CategoryScale,
  LinearScale,
  PointElement,
  LineElement,
  Title,
  Tooltip,
  Legend
} from 'chart.js'
import { Line } from 'react-chartjs-2'
import { faker } from '@faker-js/faker'

ChartJS.register(CategoryScale, LinearScale, PointElement, LineElement, Title, Tooltip, Legend)

export const options = {
  responsive: true,
  interaction: {
    mode: 'index' as const,
    intersect: false
  },
  stacked: false,
  plugins: {
    title: {
      display: true,
      text: 'Chart.js Line Chart - Multi Axis'
    }
  },
  scales: {
    y: {
      type: 'linear' as const,
      display: true,
      position: 'left' as const
    },
    y1: {
      type: 'linear' as const,
      display: true,
      position: 'right' as const,
      grid: {
        drawOnChartArea: false
      }
    }
  }
}

function getRandomDateInFormat() {
  const date = faker.date.birthdate()

  const day = String(date.getDate()).padStart(2, '0')
  const month = String(date.getMonth() + 1).padStart(2, '0') // 月份从0开始
  const year = date.getFullYear()

  return `${day}/${month}/${year}`
}

const labels = Array.from({ length: 10 }).map(() => getRandomDateInFormat())

export const data = {
  labels,
  datasets: [
    {
      label: 'Dataset 1',
      data: labels.map(() => faker.number.int({ min: 0, max: 1000 })),
      borderColor: 'rgb(255, 99, 132)',
      backgroundColor: 'rgba(255, 99, 132, 0.5)',
      yAxisID: 'y'
    },
    {
      label: 'Dataset 2',
      data: labels.map(() => faker.number.int({ min: 100, max: 1000 })),
      borderColor: '#00ff85',
      backgroundColor: 'rgba(53, 162, 235, 0.5)',
      yAxisID: 'y1'
    }
  ]
}

export function ChartWAL() {
  return (
    <div>
      <Line options={options} data={data} />
    </div>
  )
}
