import type { Role } from './types'
import { RoleFlag } from '../ui/flag'
import { createColumnHelper } from '@tanstack/react-table'

const columnHelper = createColumnHelper<Role>()

export const columns = [
  columnHelper.accessor('name', {
    header: 'Role name',
    cell: (props) => props.getValue()
  }),
  columnHelper.accessor('accessLevel', {
    header: 'Access Level',
    cell: (props) => props.getValue()
  }),
  columnHelper.accessor('flags', {
    header: 'Flags',
    cell: (props) => {
      const flags = props.getValue()

      if (flags === null) {
        return
      }

      return (
        <ul className="flex items-center">
          {flags.map((flag, ind) => (
            <li key={ind} className="mr-1">
              <RoleFlag flag={flag} iconSize={20} />
            </li>
          ))}
        </ul>
      )
    }
  })
]
