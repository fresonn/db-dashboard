import type { RoleAccessLevel } from '@/lib/api/gen'

export type Role = {
  id: string
  name: string
  accessLevel: RoleAccessLevel
  memberOf: string[]
  flags: string[]
  capabilities: string[]
}
